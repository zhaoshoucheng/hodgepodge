package atomic

import (
	"sync/atomic"
	"time"
)

/*
原子操作

atomic.Value 没有内存数据拷贝，所以存储的是临时变量，内部是维护一个指针
写入数据时设置协程不可抢占，存储完毕时释放，切换临时指针。因为每一个临时变量的指
不可能互相影响，所以保证原子性
*/

var person *Person
var ato atomic.Value

type Person struct {
	name string
	age int
}

func (p *Person)Update(name string, age int) {
	p.name = name
	time.Sleep(time.Millisecond * 200)
	p.age = age
}

func UpdateV2(name string, age int) {
	p := &Person{}
	p.name = name
	time.Sleep(time.Millisecond * 200)
	p.age = age
	ato.Store(p)
}