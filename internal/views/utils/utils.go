package utils

import "errors"

func RequiredValidation(v string) error {
	if len(v) <= 1 {
		return errors.New("Campo obrigatório, deve ser preenchido.")
	}
	return nil
}
