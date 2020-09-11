package task

import (
	"github.com/google/uuid"
	"sync"
)

type Pool struct {
	tasks sync.Map
}

func NewPool() *Pool {
	return &Pool{}
}

func (p *Pool) Add(task Interface) {
	task.SetId(genToken())
	task.SetPool(p)
	p.tasks.Store(task.GetId(), task)
}

func (p *Pool) Get(id string) (Interface, bool) {
	tmp, ok := p.tasks.Load(id)
	if !ok {
		return nil, false
	}
	return tmp.(Interface), true
}

func (p *Pool) Remove(task Interface) {
	p.tasks.Delete(task.GetId())
}

func genToken() string {
	return uuid.New().String()
}
