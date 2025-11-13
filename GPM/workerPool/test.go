// 测试一下写一个线程池
package main

import (
	"fmt"
	"sync"
	_ "time"
)

type task interface {
	run()
	getResult()
}
type doctorTask struct {
	wg    *sync.WaitGroup
	docID int
	done  string
}

func (d *doctorTask) run() {
	d.wg.Done()
	fmt.Println("doctor", d.docID, "is working")
}
func (d *doctorTask) getResult() {
	d.done = fmt.Sprintf("doctor %d is done", d.docID)
	fmt.Println(d.done)
}

type teacherTask struct {
	wg        *sync.WaitGroup
	teacherID int
	done      string
}

func (t *teacherTask) run() {
	t.wg.Done()
	fmt.Println("teacher", t.teacherID, "is working")
}
func (t *teacherTask) getResult() {
	t.done = fmt.Sprintf("teacher %d is done", t.teacherID)
	fmt.Println(t.done)
}

type programmerTask struct {
	wg           *sync.WaitGroup
	programmerID int
	done         string
}

func (p *programmerTask) run() {
	p.wg.Done()
	fmt.Println("programmer", p.programmerID, "is working")
}

func (p *programmerTask) getResult() {
	p.done = fmt.Sprintf("programmer %d is done", p.programmerID)
	fmt.Println(p.done)
}

// 创建一个线程池对象
type threadPool struct {
	//核心数量
	coreNum int
	//创建一个任务队列
	taskQueue chan task
}

func NewPool(num int) *threadPool {
	pool := &threadPool{
		coreNum:   num,
		taskQueue: make(chan task),
	}

	//创建核心数量的goroutine
	for i := 0; i < pool.coreNum; i++ {
		go func() {
			for task := range pool.taskQueue {
				task.run()
				task.getResult()
			}
		}()
	}
	return pool
}
func (p *threadPool) AddTask(t task, wg *sync.WaitGroup) {
	wg.Add(1)
	p.taskQueue <- t
}

func main() {
	var w = sync.WaitGroup{}
	pool := NewPool(3)
	defer close(pool.taskQueue)

	//创建任务
	for i := 0; i < 7; i++ {
		obj := &doctorTask{docID: i}
		obj.wg = &w
		pool.AddTask(obj, &w)
	}
	for i := 0; i < 8; i++ {
		obj := &teacherTask{teacherID: i}
		obj.wg = &w
		pool.AddTask(obj, &w)
	}
	for i := 0; i < 9; i++ {
		obj := &programmerTask{programmerID: i}
		obj.wg = &w
		pool.AddTask(obj, &w)
	}

	w.Wait()

}
