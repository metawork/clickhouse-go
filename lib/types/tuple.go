package types

import (
	"database/sql"
	"database/sql/driver"
	"errors"
)

var (
	_ driver.Valuer = &Tuple{}
	_ sql.Scanner   = &Tuple{}
)

type Tuple []interface{}

func (t *Tuple) Value() (driver.Value, error) {
	return driver.Value(t), nil
}

func (t *Tuple) Scan(src interface{}) error {
	switch v := src.(type) {
	case Tuple:
		*t = v
	case []interface{}:
		*t = Tuple(v)
	default:
		return errors.New("incompatible type for Tuple")
	}
	return nil
}
