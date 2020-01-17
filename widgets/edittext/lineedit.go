package edittext

import (
	"../../utils"
	"github.com/gotk3/gotk3/gtk"
	lua "github.com/yuin/gopher-lua"
)

//LineEdit type of LuaApp
type LineEdit struct {
	edit *gtk.Entry
}

//RegisterLineEditType is func for register edittext type
func RegisterLineEditType(L *lua.LState) int {
	mt := L.NewTypeMetatable("LineEdit")
	L.SetGlobal("LineEdit", mt)

	L.SetField(mt, "new", L.NewFunction(newLEditText))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), lEditMethods))
	return 1
}

var lEditMethods = map[string]lua.LGFunction{
	"setEditable": luaLSetEditable,
	"setText": luaLSetText,
	"setHAlign": luaLSetHAlign,
	"setVAlign": luaLSetVAlign,
	"setMarginTop": luaLSetMarginTop,
	"setMarginBottom": luaLSetMarginBottom,
	"setMarginStart": luaLSetMarginStart,
	"setMarginEnd": luaLSetMarginEnd,
	"setWidth": luaLSetWidth,
	"setHeight": luaLSetHeight,
	"getText": luaLGetText,
	"setMaxLength": luaLSetMaxLength,
	"getParent": luaLGetParent,
}

func newLEditText(L *lua.LState) int{
	gtkEdit, err := gtk.EntryNew()

	if err != nil {
		utils.FatalMessage(err.Error())
	}

	gtkEdit.SetPlaceholderText(L.ToString(2))
	editText := &LineEdit{gtkEdit}
	ud := L.NewUserData()
	ud.Value = editText
	L.SetMetatable(ud, L.GetTypeMetatable("LineEdit"))
	L.Push(ud)
	return 1
}

func checkLineEdit(L *lua.LState) *LineEdit {
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*LineEdit); ok {
		return v
	}
	L.ArgError(1, "LineEdit expected")
	return nil
}

func luaLSetEditable(L *lua.LState) int {
	e := checkLineEdit(L)
	e.edit.SetEditable(L.ToBool(2))
	return 1
}

func luaLSetText(L *lua.LState) int{
	e := checkLineEdit(L)
	e.edit.SetText(L.ToString(2))
	return 1
}

func luaLSetHAlign(L *lua.LState) int {
	e := checkLineEdit(L)
	ud := L.ToUserData(2)
	e.edit.SetHAlign(ud.Value.(gtk.Align))
	e.edit.SetHExpand(true)
	return 1
}

func luaLSetVAlign(L *lua.LState) int {
	e := checkLineEdit(L)
	ud := L.ToUserData(2)
	e.edit.SetVAlign(ud.Value.(gtk.Align))
	e.edit.SetVExpand(true)
	return 1
}

func luaLSetMarginTop(L *lua.LState) int {
	e := checkLineEdit(L)
	e.edit.SetMarginTop(L.ToInt(2))
	return 1
}

func luaLSetMarginBottom(L *lua.LState) int {
	e := checkLineEdit(L)
	e.edit.SetMarginBottom(L.ToInt(2))
	return 1
}

func luaLSetMarginStart(L *lua.LState) int {
	e := checkLineEdit(L)
	e.edit.SetMarginStart(L.ToInt(2))
	return 1
}

func luaLSetMarginEnd(L *lua.LState) int {
	e := checkLineEdit(L)
	e.edit.SetMarginEnd(L.ToInt(2))
	return 1
}

func luaLSetWidth(L *lua.LState) int {
	e := checkLineEdit(L)
	e.edit.GetAllocation().SetWidth(L.ToInt(2))
	return 1
}

func luaLSetHeight(L *lua.LState) int {
	e := checkLineEdit(L)
	e.edit.GetAllocation().SetHeight(L.ToInt(2))
	return 1
}

func luaLSetMaxLength(L *lua.LState) int {
	e := checkLineEdit(L)
	e.edit.SetMaxLength(L.ToInt(2))
	return 1
}

func luaLGetText(L *lua.LState) int{
	e := checkLineEdit(L)
	text, _ := e.edit.GetText()
	L.Push(lua.LString(text))
	return 1
}

func luaLGetParent(L *lua.LState) int {
	e := checkLineEdit(L)
	ud := L.NewUserData()
	ud.Value = e.edit
	L.Push(ud)
	return 1
}