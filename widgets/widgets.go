package widgets

import (
	"./button"
	"./label"
	"./layout"
	"./edittext"
	lua "github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	label.RegisterLabelType(L)
	button.RegisterButtonType(L)
	layout.RegisterVerticalType(L)
	layout.RegisterHorizontalType(L)
	edittext.RegisterLineEditType(L)
	edittext.RegisterMultilineEditType(L)
	return 1
}
