package zlog

import (
	"sync"
)



var eventPool = &sync.Pool{
	New: func() interface{} {
		return &Event{
			buf: make([]byte, 0, 1 << 14),
		}
	},
}

// Event represents a log event. It is instanced by one of the level method of
// Logger and finalized by the Msg or Msgf method.
type Event struct {
	buf   []byte
	w     LevelWriter
	level Level
	done  func(msg string)
	ch    []Hook // hooks from context
}

func newEvent(w LevelWriter, level Level) *Event {
	e := eventPool.Get().(*Event)
	e.buf = e.buf[:0]
	e.ch = nil
	e.w = w
	e.level = level
	return e
}
