// +build tinygo

package events

import "errors"

type receiverType int
const (
	receiverTypeUnknown receiverType = iota
	receiverTypeFuncNoArgs
)

type receiptHandler struct {
	receiver	interface{}
	rt			receiverType
}

func newReceiptHandler(receiver interface{}) (receiptHandler, error) {
	var rt receiverType = 0

	switch receiver.(type) {
	case func():
		rt = receiverTypeFuncNoArgs
	default:
		return receiptHandler{}, errors.New("unsupported receiver type")
	}

	return receiptHandler{receiver: receiver, rt: rt}, nil
}

func (rh *receiptHandler) invoke(values preparedArgs) {
	switch rh.rt {
	case receiverTypeFuncNoArgs:
		rh.receiver.(func())()
	}
}

type preparedArgs []interface{}

func prepareArgs(argIfacess []interface{}) preparedArgs {
	return argIfacess
}