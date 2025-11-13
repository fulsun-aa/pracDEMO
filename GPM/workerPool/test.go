package main

//改进版：任务不应该知道waitgroup的存在，而是由线程池集中管理waitGroup
import (
	"fmt"
	"sync"
)

type task interface {
	run()
	getResult()
}

type doctorTask struct{ docID int }

func (d *doctorTask) run()       { fmt.Println("doctor", d.docID, "is working") }
func (d *doctorTask) getResult() { fmt.Println("doctor", d.docID, "is done") }

type teacherTask struct{ teacherID int }

func (t *teacherTask) run()       { fmt.Println("teacher", t.teacherID, "is working") }
func (t *teacherTask) getResult() { fmt.Println("teacher", t.teacherID, "is done") }

type programmerTask struct{ programmerID int }

func (p *programmerTask) run()       { fmt.Println("programmer", p.programmerID, "is working") }
func (p *programmerTask) getResult() { fmt.Println("programmer", p.programmerID, "is done") }

// -------------------- Thread Pool --------------------

type threadPool struct {
	coreNum   int
	taskQueue chan task
	wg        sync.WaitGroup
}

func NewPool(num int, queueSize int) *threadPool {
	pool := &threadPool{
		coreNum:   num,
		taskQueue: make(chan task, queueSize),
	}

	for i := 0; i < pool.coreNum; i++ {
		go pool.worker()
	}
	return pool
}

func (p *threadPool) worker() {
	for task := range p.taskQueue {
		task.run()
		task.getResult()
		p.wg.Done()
	}
}

func (p *threadPool) AddTask(t task) {
	p.wg.Add(1)
	p.taskQueue <- t
}

func (p *threadPool) Wait() {
	p.wg.Wait()
}

func (p *threadPool) Close() {
	close(p.taskQueue)
}

// -------------------- Main --------------------

func main() {
	pool := NewPool(3, 10)
	for i := 0; i < 7; i++ {
		pool.AddTask(&doctorTask{docID: i})
	}
	for i := 0; i < 8; i++ {
		pool.AddTask(&teacherTask{teacherID: i})
	}
	for i := 0; i < 9; i++ {
		pool.AddTask(&programmerTask{programmerID: i})
	}

	pool.Wait()
	pool.Close()
}
