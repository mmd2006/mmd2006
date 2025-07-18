package validation

import "errors"

type TaskInput struct {
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
}

func (t *TaskInput) Validate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	if t.Description == "" {
		return errors.New("description is required")
	}
	return nil
}
