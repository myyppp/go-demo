// 有缓冲的通道实现资源池
// 场景：共享数据库连接和内存缓冲区
package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type closerFunc func() (io.Closer, error)

// Pool 管理一组可以安全在多个 goroutine 间共享的资源，被管理的资源必须实现 io.Closer 接口
type Pool struct {
	m         sync.Mutex     // 保证池内值的安全
	resources chan io.Closer // 缓冲通道，保存共享的资源
	factory   closerFunc     // 创建池中的新资源，由包的使用者提供
	closed    bool
}

var ErrPoolClosed = errors.New("Pool 已经关闭！")

// New 创建一个资源池
// fn：分配新资源的函数
// size：池的大小
func New(fn closerFunc, size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("池太小了")
	}

	return &Pool{
		resources: make(chan io.Closer, size),
		factory:   fn,
	}, nil
}

// Acquire 请求资源
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// 检查是否有空闲资源
	case r, ok := <-p.resources:
		log.Println("请求资源：池中有共享资源！")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil

	// 没用空闲的资源，创建一个新资源
	default:
		log.Println("请求资源：池中没有空闲资源，需要新建！")
		return p.factory()
	}
}

// Release 释放资源
func (p *Pool) Release(r io.Closer) {
	// 保证安全
	p.m.Lock()
	defer p.m.Unlock()
	// 如果池关闭，则销毁资源
	// 对closed标志的读写必须进行同步，否则可能会误导其他的goroutine，认为资源池是打开的
	if p.closed {
		r.Close()
		return
	}
	select {
	// 试图将资源放入队列中
	case p.resources <- r:
		log.Println("释放资源：将资源放入池中！")

	// 如果队列已满，则关闭资源
	default:
		log.Println("释放资源：资源池已满，关闭资源！")
		r.Close()
	}
}

// Close 关闭池和资源
func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}
	// 将池关闭
	p.closed = true

	// 先关闭chan，再关闭资源
	// 如果不这样做，会发生死锁
	close(p.resources)

	// 关闭资源
	for r := range p.resources {
		r.Close()
	}
}
