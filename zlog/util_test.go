package zlog

import (
	"testing"
	"time"
)

func TestGetNow(t *testing.T) {
	t.Log(getNow(time.Now()))

	t.Log(getCurHour(time.Now()))
}
