package base

import (
	"errors"
	"fmt"
	"github.com/tango-contrib/binding"
)

type BindRouter struct {
	binding.Binder
}

func (br *BindRouter) BindAndValidate(v interface{}) error {
	errs := br.Bind(v)
	if errs.Len() > 0 {
		// todo : friendly message
		return errors.New(fmt.Sprintf("'%s' has %s", errs[0].Fields()[0], errs[0].Classification))
	}
	return nil
}
