package models

import (
	"api/src/service"

	"github.com/asaskevich/govalidator"
	"gopkg.in/validator.v2"
)

type State struct {
	Uf   string `json:"tb001_sigla_uf,omitempty" gorm:"type:varchar(255)" validate:"notzz,nonzero"`
	Name string `json:"tb001_nome_estado,omitempty" gorm:"type:varchar(255)" validate:"notzz,nonzero"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	validator.SetValidationFunc("notzz", service.WhiteSpaces)

}

func (state *State) Prepare() error {
	if err := validator.Validate(state); err != nil {
		return err
	}
	return nil
}

func (state *State) validate() error {

	_, err := govalidator.ValidateStruct(state)
	if err != nil {
		return err
	}

	return nil
}
