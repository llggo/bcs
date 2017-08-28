package ctn

import "errors"

type MapString map[string]interface{}

func NewMapString() *MapString {
	var res MapString
	res = make(map[string]interface{})
	return &res
}

//Add : Add a key
func (ms *MapString) Add(key string, value interface{}) error {
	if ms.Get(key) != nil {
		return errors.New("Key is exist")
	}
	(*ms)[key] = value
	return nil
}

//Set : If not exist, add
func (ms *MapString) Set(key string, val interface{}) error {
	oldVal := ms.Get(key)
	if oldVal == nil {
		return ms.Add(key, val)
	}
	(*ms)[key] = val
	return nil
}

//Get :
func (ms *MapString) Get(key string) interface{} {
	val, ok := (*ms)[key]
	if ok {
		return val
	}
	return nil
}

//Loop :
func (ms *MapString) Loop(loopIter func(interface{}), breakPoint func() bool) {
	for _, m := range *ms {
		if breakPoint() {
			break
		}
		loopIter(m)
	}
}

//LoopError :
func (ms *MapString) LoopError(loopIter func(interface{}) error, breakPoint func() bool) error {
	for _, m := range *ms {
		if breakPoint() {
			break
		}
		err := loopIter(m)
		if err != nil {
			return err
		}
	}
	return nil
}
