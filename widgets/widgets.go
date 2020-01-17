package widgets

import (
	"./button"
	"./menu"
	"./label"
	"./layout"
	"./edittext"
	"./align"
	lua "github.com/yuin/gopher-lua"
)

func Loader(L *lua.LState) int {
	label.RegisterLabelType(L)
	button.RegisterButtonType(L)
	layout.RegisterVerticalType(L)
	layout.RegisterHorizontalType(L)
	edittext.RegisterLineEditType(L)
	edittext.RegisterMultilineEditType(L)
	align.RegisterAlignType(L)
	menu.RegisterMenuType(L)
	return 1
}
