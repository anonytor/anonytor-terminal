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
	p.tasks.Store(task.GetId(), task)
}

func (p *Pool) Get(id string) (Interface, bool) {
	tmp, ok := p.tasks.Load(id)
	if !ok {
		return nil, false
	}
	return tmp.(Interface), true
}

func genToken() string {
	return uuid.New().String()
}
