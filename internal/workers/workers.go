package workers

import (
	"github.com/vsrtferrum/OzonIntro/internal/model"
	"github.com/vsrtferrum/OzonIntro/internal/module"
)

type WorkerPool struct {
	tasks chan func()
}

func NewWorkerPool(workerCount, queueSize int) *WorkerPool {
	wp := &WorkerPool{
		tasks: make(chan func(), queueSize),
	}

	for i := 0; i < workerCount; i++ {
		go wp.worker()
	}
	return wp
}

func (wp *WorkerPool) worker() {
	for task := range wp.tasks {
		task()
	}
}

func (wp *WorkerPool) Submit(task func()) {
	wp.tasks <- task
}

type ConcurrentModule struct {
	Module *module.Module
	wp     *WorkerPool
}

func NewConcurrentModule(module *module.Module, workerCount, queueSize int) *ConcurrentModule {
	return &ConcurrentModule{
		Module: module,
		wp:     NewWorkerPool(workerCount, queueSize),
	}
}

func (cm *ConcurrentModule) GetPosts() (*[]model.PostList, error) {
	type result struct {
		post *[]model.PostList
		err  error
	}

	resChan := make(chan result, 1)

	cm.wp.Submit(func() {
		info, err := cm.Module.GetPosts()
		resChan <- result{info, err}
	})

	res := <-resChan
	return res.post, res.err
}

func (cm *ConcurrentModule) GetPost(id uint64) (*model.Post, *[]model.Comments, error) {
	type result struct {
		post     *model.Post
		comments *[]model.Comments
		err      error
	}

	resChan := make(chan result, 1)

	cm.wp.Submit(func() {
		post, comments, err := cm.Module.GetPost(id)
		resChan <- result{post, comments, err}
	})

	res := <-resChan
	return res.post, res.comments, res.err
}

func (cm *ConcurrentModule) AddComment(data *model.WriteComment) (uint64, error) {
	type result struct {
		id  uint64
		err error
	}
	resChan := make(chan result, 1)
	cm.wp.Submit(func() {
		id, err := cm.Module.AddComment(data)
		resChan <- result{id, err}
	})
	res := <-resChan
	return res.id, res.err
}

func (cm *ConcurrentModule) AddPost(data *model.WritePost) (uint64, error) {
	type result struct {
		id  uint64
		err error
	}
	resChan := make(chan result, 1)
	cm.wp.Submit(func() {
		id, err := cm.Module.AddPost(data)
		resChan <- result{id, err}
	})
	res := <-resChan
	return res.id, res.err
}
