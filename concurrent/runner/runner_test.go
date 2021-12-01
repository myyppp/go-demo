package runner_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/myyppp/go-demo/concurrent/runner"
)

const timeout = 3 * time.Second

func TestRunner(t *testing.T) {

	log.Println("开始工作！")
	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("超时了!")
			os.Exit(1)

		case runner.ErrInterrupt:
			log.Println("错误中断")
			os.Exit(2)
		}
	}

	log.Println("结束！")
}

func createTask() func(int) {
	return func(i int) {
		log.Printf("第%d个任务", i)
		time.Sleep(time.Duration(i) * time.Second)
	}
}
