package model

import "errors"

type ToDo struct {
	Id    int    `json:"id"`
	Label string `json:"label"`
}

func (t *ToDo) validate() error {
	if t.Label == "" {
		return errors.New("Labels cannot be empty")
	}
	return nil
}
