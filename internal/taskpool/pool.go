package taskpool

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// TaskFunc 是协程池执行的任务函数类型
type TaskFunc func(ctx context.Context, params interface{}) error

// Task 是任务结构体
type Task struct {
	Func   TaskFunc
	Params interface{}
	Ctx    context.Context
	Cancel context.CancelFunc
}

// TaskPool 是协程池结构体
type TaskPool struct {
	jobs    chan Task // 任务通道
	workers int       // 工人数量
	wg      sync.WaitGroup
	mu      sync.Mutex
	running map[context.Context]struct{} // 当前运行的任务
}

// NewTaskPool 创建新的协程池
func NewTaskPool(workers int) *TaskPool {
	return &TaskPool{
		jobs:    make(chan Task, workers),
		workers: workers,
		running: make(map[context.Context]struct{}),
	}
}

// Run 启动协程池的工作goroutine
func (p *TaskPool) Run() {
	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go p.worker()
	}
}

// worker 是协程池中的工作goroutine
func (p *TaskPool) worker() {
	defer p.wg.Done()
	for task := range p.jobs {
		p.mu.Lock()
		p.running[task.Ctx] = struct{}{}
		p.mu.Unlock()

		err := task.Func(task.Ctx, task.Params)
		if err != nil {
			fmt.Printf("Task failed: %v\n", err)
		}

		p.mu.Lock()
		delete(p.running, task.Ctx)
		p.mu.Unlock()
	}
}

// AddTask 添加一个任务到协程池
func (p *TaskPool) AddTask(taskFunc TaskFunc, params interface{}) {
	ctx, cancel := context.WithCancel(context.Background())
	task := Task{
		Func:   taskFunc,
		Params: params,
		Ctx:    ctx,
		Cancel: cancel,
	}
	select {
	case p.jobs <- task:
		// 任务成功添加到队列
	default:
		// 队列已满，处理失败情况（可选）
		cancel()
		fmt.Println("Task queue is full, task was not added.")
	}
}

// CancelTask 取消一个任务（如果它正在运行）
func (p *TaskPool) CancelTask(taskCtx context.Context) {
	p.mu.Lock()
	if _, ok := p.running[taskCtx]; ok {
		// 如果任务正在运行，则取消它
		taskCtx.Done() // 假设任务在检查这个任务上下文是否被取消
	}
	p.mu.Unlock()
}

// Stop 关闭协程池，停止接受新任务，并等待所有任务完成
func (p *TaskPool) Stop() {
	close(p.jobs) // 关闭任务通道，停止接收新任务
	p.wg.Wait()   // 等待所有工作goroutine完成
}

// 示例任务函数
func exampleTask(ctx context.Context, params interface{}) error {
	select {
	case <-ctx.Done():
		// 如果任务被取消，则退出
		fmt.Println("Task cancelled:", params)
		return ctx.Err()
	case <-time.After(time.Duration(params.(int)) * time.Second):
		// 模拟任务完成
		fmt.Println("Task completed:", params)
		return nil
	}
}

func demoTest() {
	pool := NewTaskPool(3) // 创建一个包含3个工作goroutine的协程池
	pool.Run()             // 启动协程池

	// 添加任务
	pool.AddTask(exampleTask, 2) // 模拟运行2秒的任务
	pool.AddTask(exampleTask, 1) // 模拟运行1秒的任务

	// 取消任务（假设我们想要取消运行2秒的任务）
	time.Sleep(1 * time.Second)           // 等待1秒，确保1秒的任务已经开始运行
	pool.CancelTask(context.Background()) // 注意：这里只是一个示例，实际应该使用AddTask返回的任务上下文

	// 等待一段时间以确保任务执行完成
	time.Sleep(3 * time.Second)

	// 关闭协程池
	pool.Stop()
}
