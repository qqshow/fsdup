package fsdup

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

type fileChunkStore struct {
	root     string
	chunkMap *sync.Map
}

func NewFileChunkStore(root string) *fileChunkStore {
	return &fileChunkStore{
		root:     root,
		chunkMap: &sync.Map{},
	}
}

func (idx *fileChunkStore) Stat(checksum []byte) error {
	checksumStr := fmt.Sprintf("%x", checksum)

	if _, ok := idx.chunkMap.Load(checksumStr); ok {
		return nil
	}

	dir := fmt.Sprintf("%s/%s/%s", idx.root, checksumStr[0:3], checksumStr[3:6])
	file := fmt.Sprintf("%s/%s", dir, checksumStr)

	_, err := os.Stat(file)
	return err
}

func (idx *fileChunkStore) Write(checksum []byte, buffer []byte) error {
	checksumStr := fmt.Sprintf("%x", checksum)

	if _, ok := idx.chunkMap.Load(checksumStr); !ok {
		dir := fmt.Sprintf("%s/%s/%s", idx.root, checksumStr[0:3], checksumStr[3:6])
		file := fmt.Sprintf("%s/%s", dir, checksumStr)

		if _, err := os.Stat(file); err != nil {
			if err := os.MkdirAll(dir, 0770); err != nil {
				return err
			}

			err = ioutil.WriteFile(file, buffer, 0666)
			if err != nil {
				return err
			}
		}

		idx.chunkMap.Store(checksumStr, true)
	}

	return nil
}

func (idx *fileChunkStore) ReadAt(checksum []byte, buffer []byte, offset int64) (int, error) {
	checksumStr := fmt.Sprintf("%x", checksum)
	dir := fmt.Sprintf("%s/%s/%s", idx.root, checksumStr[0:3], checksumStr[3:6])
	file := fmt.Sprintf("%s/%s", dir, checksumStr)

	if _, err := os.Stat(file); err != nil {
		return 0, err
	}

	chunk, err := os.OpenFile(file, os.O_RDONLY, 0666)
	if err != nil {
		return 0, err
	}
	defer chunk.Close()

	read, err := chunk.ReadAt(buffer, offset)
	if err != nil {
		return 0, err
	} else if read != len(buffer) {
		return 0, errors.New("cannot read full section")
	}

	return read, nil
}

func (idx *fileChunkStore) Remove(checksum []byte) error {
	checksumStr := fmt.Sprintf("%x", checksum)
	dir1 := fmt.Sprintf("%s/%s", idx.root, checksumStr[0:3])
	dir2 := fmt.Sprintf("%s/%s/%s", idx.root, checksumStr[0:3], checksumStr[3:6])
	file := fmt.Sprintf("%s/%s", dir2, checksumStr)

	if err := os.Remove(file); err != nil {
		return err
	}

	os.Remove(dir2)
	os.Remove(dir1)

	idx.chunkMap.Delete(checksumStr)

	return nil
}