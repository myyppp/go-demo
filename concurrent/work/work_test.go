package work_test

import (
	"log"
	"sync"
	"testing"
	"time"

	"github.com/myyppp/go-demo/concurrent/work"
)

// namePrinter 使用特定方式打印名字
type namePrinter struct {
	name string
}

// Task 实现 Worker 接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

// names 提供一组用来显示的名字
var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

func TestWork(t *testing.T) {
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 10; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}
			go func() {
				// 将任务提交执行，当 Run() 返回时我们就知道任务已处理完成
				// 调用工作池的 Run() 方法，将 namePrinter 提交到池中
				// 一旦工作池的 goroutine 接收到这个值，Run() 方法就会返回
				p.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	p.Shutdown()
}
