package domain

import (
	value_object "github.com/axel-andrade/finance_planner_api/internal/application/domain/value_objects"
)

/**
* As tags json e bson indicam que um User pode ser serializado nestes formatos
**/

type User struct {
	Base
	Name     value_object.Name     `json:"name"`
	Email    value_object.Email    `json:"email"`
	Password value_object.Password `json:"-"`
}

func BuildUser(name string, email string, password string) (*User, error) {
	user := &User{
		Name:     value_object.Name{Value: name},
		Email:    value_object.Email{Value: email},
		Password: value_object.Password{Value: password},
	}

	if err := user.validate(); err != nil {
		return nil, err
	}

	return user, nil
}

/*
1 - O trecho user *User antes do nome da função representa uma amarração entre a função
Prepare e a struc User, ou seja, é como se User fosse uma classe e a função Prepare
fosse um método público.
2 - É usado o * no User, pois todas as vezes que o user for alterado ele será atualizado
em todos os objetos pois esta utilizando o mesmo local na memória
3 - O retorno da função é um error que pode ter valor nil, ou seja, se o erro for nil quer dizer
que a função funcionou corretamente. Esta é uma forma de validação.
4 - A função Prepare começa com letra maiuscula pois é um método publico. Funções que começam com
letra minuscula são funcões privadas.
*/

func (u *User) validate() error {
	if err := u.Name.Validate(); err != nil {
		return err
	}

	if err := u.Email.Validate(); err != nil {
		return err
	}

	if err := u.Password.Validate(); err != nil {
		return err
	}

	return nil
}
