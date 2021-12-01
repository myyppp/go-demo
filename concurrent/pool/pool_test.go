package pool_test

import (
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/myyppp/go-demo/concurrent/pool"
)

const (
	maxGroutines    = 20
	pooledResources = 5 // 池中资源的数量
)

// 模拟要共享的资源
type dbConnection struct {
	ID int32 // 保存每个连接的唯一标识
}

// Close 实现了io.Closer接口
// Close 关闭资源
func (dbConn *dbConnection) Close() error {
	log.Println("关闭资源：资源id为", dbConn.ID)
	return nil
}

// 给每一个连接分配一个独一无二的id
var idCounter int32

// createConnection 是一个工厂函数
// 创建资源
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("创建资源：资源id为", id)
	return &dbConnection{id}, nil
}

func TestPool(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(maxGroutines) // 并发数量

	log.Println("----------------开始 创建资源池----------------")
	// 创建连接池，返回一个连接池p
	p, err := pool.New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
	}

	// 模拟请求
	for i := 0; i < maxGroutines; i++ {
		time.Sleep(time.Duration(rand.Intn(10+i)) * time.Millisecond)
		go func(i int) {
			performQueries(i, p)
			wg.Done()
		}(i)
	}

	wg.Wait()
	// 关闭池
	log.Println("----------------结束 关闭资源池----------------")
	p.Close()
}

// 模拟请求
// i gorotine的ID
// p 资源池
func performQueries(i int, p *pool.Pool) {
	// 从池中请求一个资源
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}
	// 将资源释放回池中
	defer p.Release(conn)

	// 用等待模拟查询响应
	log.Printf("模拟执行：gorotine的id为%d，资源的id为%d\n", i, conn.(*dbConnection).ID)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
}
