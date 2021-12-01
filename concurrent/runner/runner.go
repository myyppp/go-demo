// runner 包管理处理任务的运行和生命周期
// 场景：调度后台任务的程序
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner 在给定的时间内执行一组任务，并在操作系统发送中断信号时结束任务
// 三个通道辅助管理程序的生命周期
type Runner struct {
	interrupt chan os.Signal   // 从操作系统发出的信号
	complete  chan error       // 报告完成
	timeout   <-chan time.Time // 报告超时
	tasks     []func(int)      // 顺序执行不同任务的函数切片
}

var ErrTimeout = errors.New("received timeout")

var ErrInterrupt = errors.New("received interrupt")

// New 返回一个Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add 添加任务
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start 执行任务，并监视通道
func (r *Runner) Start() error {
	// 接受到所有的中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	// 执行任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 任务处理完成
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

// run 执行每一个已注册的任务
func (r *Runner) run() error {
	for id, task := range r.tasks {
		// 检测到中断
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		// 执行任务
		task(id)
	}
	return nil
}

// gotInterrupt 验收是否接受到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	// 当中断事件被触发时发出信号
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true

	default:
		return false
	}
}
