package utils

import (
	"sync"
)

type Set struct {
	mapInst map[interface{}]bool
	sync.RWMutex
}

func New() *Set {
	return &Set{
		mapInst: map[interface{}]bool{},
	}
}

func (p *Set) Add(item interface{}) {
	p.Lock()
	defer p.Unlock()
	p.mapInst[item] = true
}

func (p *Set) Remove(item interface{}) {
	p.Lock()
	defer p.Unlock()
	delete(p.mapInst, item)
}

func (p *Set) PopOne() (item interface{}, ok bool) {
	p.Lock()
	defer p.Unlock()
	if len(p.mapInst) > 0 {
		for k := range p.mapInst {
			item = k
			delete(p.mapInst, k)
			break
		}
		ok = true
	} else {
		ok = false
	}
	return
}

func (p *Set) Has(item interface{}) bool {
	p.RLock()
	defer p.RUnlock()
	_, ok := p.mapInst[item]
	return ok
}

func (p *Set) Len() int {
	return int(len(p.mapInst))
}

func (p *Set) Clear() {
	p.Lock()
	defer p.Unlock()
	p.mapInst = map[interface{}]bool{}
}

func (p *Set) IsEmpty() bool {
	if p.Len() == 0 {
		return true
	}
	return false
}

func (p *Set) List() []interface{} {
	p.RLock()
	defer p.RUnlock()
	list := []interface{}{}

	for item := range p.mapInst {
		list = append(list, item)
	}
	return list
}
