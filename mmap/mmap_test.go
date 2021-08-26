package mmap

import (
	"testing"
)

func TestName(t *testing.T) {
	mmap := NewMMap("my.db")
	mmap.Write([]byte("hello world"))
	mmap.Write([]byte("hello go"))
	t.Log(string(mmap.Read()))
	t.Log(string(mmap.Read()))
	//syscall.Munmap(mmap.b)
}
