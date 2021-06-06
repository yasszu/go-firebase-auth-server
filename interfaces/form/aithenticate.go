package form

import "github.com/go-playground/validator/v10"

type Authenticate struct {
	IDToken string `form:"id_token" validate:"required"`
}

func (f *Authenticate) Validate() error {
	v := validator.New()
	err := v.Struct(f)
	return err
}
