package atomic

import (
	"fmt"
	"sync"
	"testing"
)

func TestPerson_Update(t *testing.T) {
	person = &Person{}
	wg := sync.WaitGroup{}
	wg.Add(10)
	// 10 个协程并发更新
	for i := 0; i < 10; i++ {
		name, age := fmt.Sprintf("nobody:%v", i), i
		go func() {
			defer wg.Done()
			//person.Update(name, age)
			UpdateV2(name, age)
		}()
	}
	wg.Wait()
	p := ato.Load().(*Person)
	fmt.Printf("p.name=%s\np.age=%v\n", p.name, p.age)
}


