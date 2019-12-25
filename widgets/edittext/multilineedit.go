package edittext

import (
	//"../../utils"
	"github.com/gotk3/gotk3/gtk"
	lua "github.com/yuin/gopher-lua"
)

//MultilineEdit type of LuaApp
type MultilineEdit struct {
	edit *gtk.TextView
}

//RegisterMultilineEditType is func for register edittext type
func RegisterMultilineEditType(L *lua.LState) int {
	mt := L.NewTypeMetatable("MultilineEdit")
	L.SetGlobal("MultilineEdit", mt)

	L.SetField(mt, "new", L.NewFunction(newMEditText))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), mEditMethods))
	return 1
}

var mEditMethods = map[string]lua.LGFunction{
	"setEditable": luaMSetEditable,
	"setText": luaMSetText,
	"getText": luaMGetText,
	"getParent": luaMGetParent,
}

func newMEditText(L *lua.LState) int{
	gtkTextView, _ := gtk.TextViewNew()
	print(gtkTextView)

	//if err != nil {
	//	utils.FatalMessage(err.Error())
	//}

	//editText := &MultilineEdit{gtkTextView}
	ud := L.NewUserData()
	//ud.Value = editText
	L.SetMetatable(ud, L.GetTypeMetatable("MultilineEdit"))
	L.Push(ud)
	return 1
}

func checkMultilineEdit(L *lua.LState) *MultilineEdit {
	ud := L.CheckUserData(1)
	if m, ok := ud.Value.(*MultilineEdit); ok {
		return m
	}
	L.ArgError(1, "MultilineEdit expected")
	return nil
}

func luaMSetEditable(L *lua.LState) int {
	e := checkMultilineEdit(L)
	e.edit.SetEditable(L.ToBool(2))
	return 1
}

func luaMSetText(L *lua.LState) int{
	e := checkMultilineEdit(L)
	buff, _ := e.edit.GetBuffer()
	buff.SetText(L.ToString(2))
	return 1
}

func luaMGetText(L *lua.LState) int{
	e := checkMultilineEdit(L)
	buff, _ := e.edit.GetBuffer()
	text, _ := buff.GetText(buff.GetStartIter(), buff.GetEndIter(), false)
	L.Push(lua.LString(text))
	return 1
}

func luaMGetParent(L *lua.LState) int {
	e := checkMultilineEdit(L)
	ud := L.NewUserData()
	ud.Value = e.edit
	L.Push(ud)
	return 1
}