package utils

import (
	"sync"
	"time"
)

type AsyncTaskPool struct {
	taskChan      chan *AsyncTask // queue used in non-runtime  tasks
	threadCnt     int             // thread num
	quitFlagChan  chan bool       // quit signal for the watcher to quit
	quitWaitGroup *sync.WaitGroup // quit waitgroup
}

func NewAsyncTaskPool(threads int, taskChanSize int) *AsyncTaskPool {
	p := &AsyncTaskPool{
		quitFlagChan:  make(chan bool, threads),
		taskChan:      make(chan *AsyncTask, taskChanSize),
		threadCnt:     threads,
		quitWaitGroup: &sync.WaitGroup{},
	}

	for i := 0; i < threads; i++ {
		go func() {
			for {
				select {
				case task := <-p.taskChan:
					task.Do()
				case waitTaskDone := <-p.quitFlagChan:
					if waitTaskDone { //等待未执行的任务执行结束
						for {
							select {
							case <-time.After(time.Second / 10):
								p.quitWaitGroup.Done()
								return
							case task := <-p.taskChan:
								task.Do()
							}
						}
					} else { //不等待未执行的任务执行完成，直接退出
						p.quitWaitGroup.Done()
						return
					}
				}
			}
		}()
		p.quitWaitGroup.Add(1)
	}

	return p
}

//注意，添加的异步任务不能抛异常
func (p *AsyncTaskPool) Do(handler interface{}, params ...interface{}) {
	//这里创建一个task放到channel
	t := NewAsyncTask(handler, params...)
	p.taskChan <- t
}

func (p *AsyncTaskPool) Close(waitTaskDone bool, waitThreadStop bool) {
	for i := 0; i < p.threadCnt; i++ {
		p.quitFlagChan <- waitTaskDone
	}
	if waitThreadStop {
		p.quitWaitGroup.Wait()
	}
}
