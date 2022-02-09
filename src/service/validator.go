package service

import (
	"errors"
	"reflect"

	"github.com/asaskevich/govalidator"
)

func WhiteSpaces(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	if govalidator.HasWhitespaceOnly(st.String()) {
		return errors.New("value cannot be only white spaces")
	}
	return nil
}
