package fsnotify

import (
	"sync"
	"testing"
)

func TestWatch(t *testing.T) {
	path := "./"
	wg := sync.WaitGroup{}
	wg.Add(1)
	Watch(path)
	wg.Wait()
}