package models

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/asaskevich/govalidator"
	"gopkg.in/validator.v2"
)

type State struct {
	Uf   string `json:"tb001_sigla_uf,omitempty" gorm:"type:varchar(255)" validate:"notzz,nonzero"`
	Name string `json:"tb001_nome_estado,omitempty" gorm:"type:varchar(255)" validate:"notzz,nonzero"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	validator.SetValidationFunc("notzz", WhiteSpaces)

}

func (state *State) Prepare() error {
	if err := validator.Validate(state); err != nil {
		return err
	}
	return nil
}

func (state *State) validate() error {
	fmt.Printf(strconv.Itoa(len(state.Uf)))

	_, err := govalidator.ValidateStruct(state)
	if err != nil {
		return err
	}

	return nil
}

func WhiteSpaces(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	if govalidator.HasWhitespaceOnly(st.String()) {
		return errors.New("value cannot be only white spaces")
	}
	return nil
}
