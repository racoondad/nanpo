package workpool

import "github.com/panjf2000/ants/v2"

type WorkerPool struct {
	pool *ants.Pool
}

// size 		池大小
// isBlockage 	是否阻塞
// true=不阻塞，直接报错
// false=阻塞，直至可以提交
func New(size int, isBlockage bool) *WorkerPool {
	if size <= 0 {
		size = 1
	}
	pool, err := ants.NewPool(size, ants.WithNonblocking(isBlockage))
	if err != nil {
		panic("worker_pool init error:" + err.Error())
	}
	return &WorkerPool{pool: pool}
}

func (w *WorkerPool) Submit(f func()) error {
	return w.pool.Submit(f)
}

func (w *WorkerPool) Release() {
	w.pool.Release()
}
