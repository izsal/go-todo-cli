package service

import (
	"errors"

	"github.com/izsal/go-todo-cli/db"
	"github.com/izsal/go-todo-cli/model"
)

type ITodoService interface {
	Add(*model.ToDo) (*model.ToDo, error)
	Update(*model.ToDo) (*model.ToDo, error)
	Delete(id int) error
	GetById(id int) *model.ToDo
	GetAll() []*model.ToDo
}

func NewService(db db.IMemory) ITodoService {
	return &toDoService{db: db}
}

type toDoService struct {
	db db.IMemory
}

func (s *toDoService) Add(t *model.ToDo) (*model.ToDo, error) {
	s.db.Add(t)
	return t, nil
}

func (s *toDoService) Update(t *model.ToDo) (*model.ToDo, error) {
	found := s.db.GetById(t.Id)
	if found != nil {
		found.Label = t.Label
		s.db.Update(found)
		return found, nil
	}
	return nil, errors.New("Todo not found")
}

func (s *toDoService) Delete(id int) error {
	found := s.db.GetById(id)
	if found == nil {
		return errors.New("Todo not found")
	}
	s.db.Delete(s.db.GetById(id))
	return nil
}

func (s *toDoService) GetById(id int) *model.ToDo {
	return s.db.GetById(id)
}

func (s *toDoService) GetAll() []*model.ToDo {
	return s.db.GetAll()
}
