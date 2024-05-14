package workpool

import (
	"fmt"
	"sync"

	"github.com/robfig/cron/v3"
)

type CronTask struct {
	EntryID cron.EntryID
	Spec    string
	Job     func()
}

type CronPool struct {
	cron  *cron.Cron
	tasks map[cron.EntryID]*CronTask
	mu    sync.RWMutex
}

func NewCronPool() *CronPool {
	return &CronPool{
		cron:  cron.New(),
		tasks: make(map[cron.EntryID]*CronTask),
	}
}

func (tp *CronPool) AddTask(spec string, job func()) (*CronTask, error) {
	tp.mu.Lock()
	defer tp.mu.Unlock()

	entryID, err := tp.cron.AddFunc(spec, job)
	if err != nil {
		return nil, err
	}

	task := &CronTask{
		EntryID: entryID,
		Spec:    spec,
		Job:     job,
	}
	tp.tasks[entryID] = task
	return task, nil
}

func (tp *CronPool) ViewTasks() {
	tp.mu.RLock()
	defer tp.mu.RUnlock()

	for _, task := range tp.tasks {
		fmt.Printf("Task ID: %v, Spec: %s\n", task.EntryID, task.Spec)
	}
}

func (tp *CronPool) RemoveTask(entryID cron.EntryID) error {
	tp.mu.Lock()
	defer tp.mu.Unlock()

	task, ok := tp.tasks[entryID]
	if !ok {
		return fmt.Errorf("task with ID %v not found", entryID)
	}

	tp.cron.Remove(task.EntryID)
	delete(tp.tasks, entryID)
	return nil
}

func (tp *CronPool) Start() {
	tp.cron.Start()
}

func (tp *CronPool) Stop() {
	tp.cron.Stop()
}

// func demoTest() {
// 	taskPool := NewCronPool()

// 	// 添加一个每5秒执行一次的定时任务
// 	task, err := taskPool.AddTask("@every 5s", func() {
// 		fmt.Println("Task is running...")
// 	})
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Added task with ID: %v\n", task.EntryID)

// 	// 查看所有任务
// 	taskPool.ViewTasks()

// 	// 移除刚才添加的任务
// 	err = taskPool.RemoveTask(task.EntryID)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Removed task")

// 	// 再次查看任务，确认任务已被移除
// 	taskPool.ViewTasks()

// 	// 启动任务调度器
// 	taskPool.Start()

// 	// 等待一段时间让任务有机会执行
// 	time.Sleep(10 * time.Second)

// 	// 停止任务调度器
// 	taskPool.Stop()
// }
