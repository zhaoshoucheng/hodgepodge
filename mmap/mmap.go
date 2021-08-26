package mmap

import (
	"fmt"
	"os"
	"syscall"
)

/*
利用Go's syscall package 使用GO的内存映射文件
*/
const maxMapSize = 0x8000000000
const maxMmapStep = 1 << 30 // 1GB

type MMap struct {
	file *os.File
	b []byte
	index int
}

func (m *MMap) Write(data []byte) {
	defer func() {
		//err := syscall.Munmap(m.b)
		var err error
		if err != nil {
			panic(err)
		}
	}()
	for index, bb := range data {
		m.b[m.index + index] = bb
	}
	m.index = m.index + len(data)
}

func (m *MMap) Read() []byte {
	return m.b
}

func (m *MMap) Close() {
	m.file.Close()
}
func NewMMap(dbFile string) *MMap {
	file, err := os.OpenFile(dbFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	stat, err := os.Stat(dbFile)
	if err != nil {
		panic(err)
	}

	size, err := mmapSize(int(stat.Size()))
	if err != nil {
		panic(err)
	}

	err = syscall.Ftruncate(int(file.Fd()), int64(size))
	if err != nil {
		panic(err)
	}

	b, err := syscall.Mmap(int(file.Fd()), 0, size, syscall.PROT_WRITE|syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		panic(err)
	}
	return &MMap{
		file: file,
		b: b,
	}
}


func mmapSize(size int) (int, error) {
	// Double the size from 32KB until 1GB.
	for i := uint(15); i <= 30; i++ {
		if size <= 1<<i {
			return 1 << i, nil
		}
	}

	// Verify the requested size is not above the maximum allowed.
	if size > maxMapSize {
		return 0, fmt.Errorf("mmap too large")
	}

	// If larger than 1GB then grow by 1GB at a time.
	sz := int64(size)
	if remainder := sz % int64(maxMmapStep); remainder > 0 {
		sz += int64(maxMmapStep) - remainder
	}

	// Ensure that the mmap size is a multiple of the page size.
	// This should always be true since we're incrementing in MBs.
	pageSize := int64(os.Getpagesize())
	if (sz % pageSize) != 0 {
		sz = ((sz / pageSize) + 1) * pageSize
	}

	// If we've exceeded the max size then only grow up to the max size.
	if sz > maxMapSize {
		sz = maxMapSize
	}

	return int(sz), nil
}
