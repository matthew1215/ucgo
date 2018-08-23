package demo

import (
	"ucgo/library/ires"
)

type GetDemoInModel struct {
	Test string
}

func (this GetDemoInModel) GetDemo() ires.ResponseStruct {
	return ires.Ok(nil)
}
