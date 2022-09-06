package _domain

import "fmt"

type Boolean bool

func (b *Boolean) Scan(src interface{}) error {
	str, ok := src.(int64)
	if !ok {
		return fmt.Errorf("Unexpected type for Boolean: %T", src)
	}
	switch str {
	case 0:
		v := false
		*b = Boolean(v)
	case 1:
		v := true
		*b = Boolean(v)
	}
	return nil
}
