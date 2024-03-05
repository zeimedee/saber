package services

import "sync"

type Value struct {
	Value int
	mutex sync.RWMutex
}

func NewValue() *Value {
	return &Value{
		Value: 0,
		mutex: sync.RWMutex{},
	}
}

func (v *Value) AddTotal(value int) {
	v.mutex.RLock()
	defer v.mutex.RUnlock()
	v.Value += value
}
