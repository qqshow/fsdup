package fsdup

import (
	"fmt"
	"io"
)

type fixedChunker struct {
	reader      io.ReaderAt
	store       ChunkStore
	start       int64
	sizeInBytes int64
	skip        *manifest
}

func NewFixedChunker(reader io.ReaderAt, index ChunkStore, offset int64, size int64) *fixedChunker {
	skip := NewManifest()
	return NewFixedChunkerWithSkip(reader, index, offset, size, skip)
}

func NewFixedChunkerWithSkip(reader io.ReaderAt, store ChunkStore, offset int64, size int64, skip *manifest) *fixedChunker {
	return &fixedChunker{
		reader:      reader,
		store:       store,
		start:       offset,
		sizeInBytes: size,
		skip:        skip,
	}
}

func (d *fixedChunker) Dedup() (*manifest, error) {
	out := NewManifest()

	sliceOffsets := d.skip.Offsets()

	currentOffset := int64(0)
	breakpointIndex := 0
	breakpoint := int64(0)

	chunk := NewChunk()
	buffer := make([]byte, chunkSizeMaxBytes)

	statusf("Creating gap chunks ...")
	chunkBytes := int64(0)
	chunkCount := int64(0)

	for currentOffset < d.sizeInBytes {
		hasNextBreakpoint := breakpointIndex < len(sliceOffsets)

		if hasNextBreakpoint {
			// At this point, we figure out if the space from the current offset to the
			// next breakpoint will fit in a full chunk.

			breakpoint = sliceOffsets[breakpointIndex]
			bytesToBreakpoint := breakpoint - currentOffset

			if bytesToBreakpoint > chunkSizeMaxBytes {
				// We can fill an entire chunk, because there are enough bytes to the next breakpoint

				chunkEndOffset := minInt64(currentOffset + chunkSizeMaxBytes, d.sizeInBytes)

				bytesRead, err := d.reader.ReadAt(buffer, d.start + currentOffset)
				if err != nil {
					return nil, err
				} else if bytesRead != chunkSizeMaxBytes {
					return nil, fmt.Errorf("cannot read all bytes from disk, %d read\n", bytesRead)
				}

				chunk.Reset()
				chunk.Write(buffer[:bytesRead])

				if err := d.store.Write(chunk.Checksum(), chunk.Data()); err != nil {
					return nil, err
				}

				debugf("offset %d - %d, NEW chunk %x, size %d\n",
					currentOffset, chunkEndOffset, chunk.Checksum(), chunk.Size())

				out.Add(&chunkSlice{
					checksum:  chunk.Checksum(),
					kind:      kindGap,
					diskfrom:  currentOffset,
					diskto:    currentOffset + chunk.Size(),
					chunkfrom: 0,
					chunkto:   chunk.Size(),
					length:    chunk.Size(),
				})

				chunkBytes += chunk.Size()
				chunkCount++
				statusf("Creating gap chunk(s) (%d chunk(s), %s) ...", chunkCount, convertBytesToHumanReadable(chunkBytes))

				currentOffset = chunkEndOffset
			} else {
				// There are NOT enough bytes to the next breakpoint to fill an entire chunk

				if bytesToBreakpoint > 0 {
					// Create and emit a chunk from the current position to the breakpoint.
					// This may create small chunks and is inefficient.
					// FIXME this should just buffer the current chunk and not emit is right away. It should FILL UP a chunk later!

					bytesRead, err := d.reader.ReadAt(buffer[:bytesToBreakpoint], d.start + currentOffset)
					if err != nil {
						return nil, err
					} else if int64(bytesRead) != bytesToBreakpoint {
						return nil, fmt.Errorf("cannot read all bytes from disk, %d read\n", bytesRead)
					}

					chunk.Reset()
					chunk.Write(buffer[:bytesRead])

					if err := d.store.Write(chunk.Checksum(), chunk.Data()); err != nil {
						return nil, err
					}

					out.Add(&chunkSlice{
						checksum:  chunk.Checksum(),
						kind:      kindGap,
						diskfrom:  currentOffset,
						diskto:    currentOffset + chunk.Size(),
						chunkfrom: 0,
						chunkto:   chunk.Size(),
						length:    chunk.Size(),
					})

					chunkBytes += chunk.Size()
					chunkCount++
					statusf("Creating gap chunk(s) (%d chunk(s), %s) ...", chunkCount, convertBytesToHumanReadable(chunkBytes))

					debugf("offset %d - %d, NEW2 chunk %x, size %d\n",
						currentOffset, currentOffset + bytesToBreakpoint, chunk.Checksum(), chunk.Size())

					currentOffset += bytesToBreakpoint
				}

				// Now we are AT the breakpoint.
				// Simply add this entry to the manifest.

				part := d.skip.Get(breakpoint)
				partSize := part.chunkto - part.chunkfrom

				debugf("offset %d - %d, size %d  -> FILE chunk %x, offset %d - %d\n",
					currentOffset, currentOffset + partSize, partSize, part.checksum, part.chunkfrom, part.chunkto)

				currentOffset += partSize
				breakpointIndex++
			}
		} else {
			chunkEndOffset := minInt64(currentOffset + chunkSizeMaxBytes, d.sizeInBytes)
			chunkSize := chunkEndOffset - currentOffset

			bytesRead, err := d.reader.ReadAt(buffer[:chunkSize], d.start + currentOffset)
			if err != nil {
				panic(err)
			} else if int64(bytesRead) != chunkSize {
				panic(fmt.Errorf("cannot read bytes from disk, %d read\n", bytesRead))
			}

			chunk.Reset()
			chunk.Write(buffer[:bytesRead])

			if err := d.store.Write(chunk.Checksum(), chunk.Data()); err != nil {
				return nil, err
			}

			debugf("offset %d - %d, NEW3 chunk %x, size %d\n",
				currentOffset, chunkEndOffset, chunk.Checksum(), chunk.Size())

			out.Add(&chunkSlice{
				checksum:  chunk.Checksum(),
				kind:      kindGap,
				diskfrom:  currentOffset,
				diskto:    currentOffset + chunk.Size(),
				chunkfrom: 0,
				chunkto:   chunk.Size(),
				length:    chunk.Size(),
			})

			chunkBytes += chunk.Size()
			chunkCount++
			statusf("Indexing gap chunks (%d chunk(s), %s) ...", chunkCount, convertBytesToHumanReadable(chunkBytes))

			currentOffset = chunkEndOffset
		}
	}

	statusf("Indexed %s of gaps (%d chunk(s))\n", convertBytesToHumanReadable(chunkBytes), chunkCount)
	return out, nil
}
