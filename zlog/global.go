package zlog

import "sync/atomic"

var (
	gLevel          = new(uint32)
	disableSampling = new(uint32)
)
// GlobalLevel returns the current global log level
func GlobalLevel() Level {
	return Level(atomic.LoadUint32(gLevel))
}

func samplingDisabled() bool {
	return atomic.LoadUint32(disableSampling) == 1
}