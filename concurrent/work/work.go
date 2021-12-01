// 使用无缓冲的通道来创建完成工作的 goroutine 池
// 使用无缓冲的通道的方法 允许使用者知道什么时候 groutine 池正在执行工作
package work

import "sync"

// Worker 必须满足接口类型，才能使用工作池
type Worker interface {
	Task()
}

// Pool 提供一个 groutine 池，这个池可以完成任何已提交的 Worker 任务
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// New 新建一个工作池
func New(maxGroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}
	p.wg.Add(maxGroutines)
	for i := 0; i < maxGroutines; i++ {
		go func() {
			for w := range p.work {
				// 阻塞，直到收到一个 Worker 接口值
				w.Task()
			}
			p.wg.Done()
		}()
	}
	return &p
}

// Run 提交工作到工作池
// work 是一个无缓冲的通道，调用者必须等待工作池里的某个 groutine 接收到这个值才会返回
// 保证调用的 Run() 方法返回时，提交的工作已经开始执行
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown 等待所有 goroutine 停止工作，关闭工作池
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
