package db

import "github.com/izsal/go-todo-cli/model"

type IMemory interface {
	Add(*model.ToDo)
	Update(*model.ToDo)
	Delete(*model.ToDo)
	GetById(id int) *model.ToDo
	GetAll() []*model.ToDo
}

type memory struct {
	mem     map[int]*model.ToDo
	counter int
}

func New() IMemory {
	return &memory{mem: make(map[int]*model.ToDo), counter: 0}
}

func (m *memory) Add(t *model.ToDo) {
	m.counter++
	t.Id = m.counter
	m.mem[t.Id] = t
}

func (m *memory) Update(t *model.ToDo) {
	m.mem[t.Id] = t
}

func (m *memory) Delete(t *model.ToDo) {
	delete(m.mem, t.Id)
}

func (m *memory) GetById(id int) *model.ToDo {
	val, ok := m.mem[id]
	if !ok {
		return nil
	}
	return val
}

func (m *memory) GetAll() []*model.ToDo {
	var todos []*model.ToDo
	for _, v := range m.mem {
		todos = append(todos, v)
	}
	return todos
}
