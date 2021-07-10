package entity

import "github.com/go-playground/validator/v10"

type IDToken string

type UID string

type IDTokenForm struct {
	IDToken string `validate:"required"`
}

func (f *IDTokenForm) Validate() error {
	v := validator.New()
	err := v.Struct(f)
	return err
}

func (f *IDTokenForm) Entity() IDToken {
	return IDToken(f.IDToken)
}
