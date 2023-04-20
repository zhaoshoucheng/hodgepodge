package time

import (
	"fmt"
	"time"
)

func Location() {
	fmt.Println(time.Now())
	loc, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(time.ParseInLocation("2006-01-02 15:04:05", "2021-10-14 21:19:22",loc))

}
