package validators

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func OneOf(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	if st.Kind() != reflect.String {
		return fmt.Errorf("valor inválido")
	}

	options := strings.Split(param, " ")
	for _, opt := range options {
		if st.String() == opt {
			return nil
		}
	}
	return errors.New("valor inválido")
}
