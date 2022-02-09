package models

import (
	"api/src/service"
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"github.com/klassmann/cpfcnpj"
	"gopkg.in/validator.v2"
)

type User struct {
	Id        uuid.UUID `json:"tb011_logins,omitempty" gorm:"type:varchar(255)"`
	Cpf       string    `json:"tb010_cpf,omitempty" gorm:"type:varchar(255)" validate:"notzz,nonzero,min=1,max=16"`
	Password  string    `json:"tb011_senha,omitempty" gorm:"type:varchar(255)" validate:"notzz,nonzero"`
	Create_at time.Time `json:"tb011_data_cadastro,omitempty"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	validator.SetValidationFunc("notzz", service.WhiteSpaces)

}

func (user *User) Prepare() error {
	if err := validator.Validate(user); err != nil {
		return err
	}
	return nil
}

func (user *User) validate() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	cpf := cpfcnpj.NewCNPJ(user.Cpf)
	if !cpf.IsValid() {
		return errors.New("bad formatted cpf")
	}

	return nil

}
