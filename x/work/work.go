package work

import (
	"sync"
)

type Work struct {
	Number    int64
	TaskLimit int64
	done      sync.WaitGroup
	Work      func(index int64)
	Done      func()
	sync.Mutex
}

func (sw *Work) Run() {
	sw.done.Add(sw.Count())
	for i := 0; i < sw.Count(); i++ {
		var start = sw.TaskLimit * int64(i)
		var stop = sw.TaskLimit*int64(i) + sw.TaskLimit
		if stop > sw.Number {
			stop = sw.Number
		}
		go sw.Loop(start, stop)
	}
	sw.done.Wait()
	sw.Done()
}

func (sw *Work) Loop(start, end int64) {
	for i := start; i < end; i++ {
		sw.Work(i)
	}
	sw.done.Done()
}

func (sw *Work) Count() int {
	var num int64
	num = sw.Number / sw.TaskLimit
	if !(sw.Number%sw.TaskLimit == 0) {
		num++
	}
	return int(num)
}
