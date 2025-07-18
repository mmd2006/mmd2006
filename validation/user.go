package validation

import "errors"

type SignupInput struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

func (s *SignupInput) Validate() error {
	if s.Username == "" {
		return errors.New("username is required")
	}
	if s.Password == "" {
		return errors.New("password is required")
	}
	return nil
}
