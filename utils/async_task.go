package utils

import (
	"reflect"
	"runtime/debug"

	log "github.com/cihub/seelog"
)

type AsyncTask struct {
	handler reflect.Value
	params  []reflect.Value
}

func NewAsyncTask(handler interface{}, params ...interface{}) *AsyncTask {
	handlerValue := reflect.ValueOf(handler)

	if handlerValue.Kind() == reflect.Func {
		task := AsyncTask{
			handler: handlerValue,
			params:  make([]reflect.Value, 0),
		}
		if paramNum := len(params); paramNum > 0 {
			task.params = make([]reflect.Value, paramNum)
			for index, v := range params {
				task.params[index] = reflect.ValueOf(v)
			}
		}
		return &task
	}
	panic("handler not func")
}

//吃掉异步任务异常
func (p *AsyncTask) Do() []reflect.Value {
	defer func() { //增加异步捕获
		if errErr := recover(); errErr != nil {
			log.Error("defer err:", errErr, "\nsatck:", string(debug.Stack()))
		}
		return
	}()
	return p.handler.Call(p.params)
}
