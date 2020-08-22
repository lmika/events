// +build !tinygo

package events

import (
	"errors"
	"reflect"
)

type receiptHandler struct {
	receiverFunc reflect.Value
	funcType     reflect.Type
}

func newReceiptHandler(receiver interface{}) (receiptHandler, error) {
	val := reflect.ValueOf(receiver)
	if val.Type().Kind() != reflect.Func {
		return receiptHandler{}, errors.New("not a function")
	}

	return receiptHandler{receiverFunc: val, funcType: val.Type()}, nil
}

func (rh *receiptHandler) invoke(values preparedArgs) {
	args := make([]reflect.Value, rh.funcType.NumIn())
	for i := range args {
		args[i] = reflect.Zero(rh.funcType.In(i))
		if i < len(values) {
			if rh.funcType.In(i).AssignableTo(values[i].Type()) {
				args[i] = values[i]
			}
		}
	}

	rh.receiverFunc.Call(args)
}

type preparedArgs []reflect.Value

func prepareArgs(argIfacess []interface{}) preparedArgs {
	values := make([]reflect.Value, len(argIfacess))
	for i, a := range argIfacess {
		values[i] = reflect.ValueOf(a)
	}
	return values
}