package label

import (
	"../../utils"
	"github.com/gotk3/gotk3/gtk"
	"github.com/yuin/gopher-lua"
)


//Label структура luaLabel
type Label struct {
	Label *gtk.Label
}

//RegisterLabelType функция регистрации Label
func RegisterLabelType(L *lua.LState) int {
	mt := L.NewTypeMetatable("Label")
	L.SetGlobal("Label", mt)
	
	L.SetField(mt, "new", L.NewFunction(newLabel))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), labelMethods))
	return 1
}

var labelMethods = map[string]lua.LGFunction{
	"setText": luaSetText,
	"setHAlign": luaSetHAlign,
	"setVAlign": luaSetVAlign,
	"setMarginTop": luaSetMarginTop,
	"setMarginBottom": luaSetMarginBottom,
	"setMarginStart": luaSetMarginStart,
	"setMarginEnd": luaSetMarginEnd,
	"setWidth": luaSetWidth,
	"setHeight": luaSetHeight,
	"getParent": luaGetParent,
}

func newLabel(L *lua.LState) int{
	gtkLabel, err := gtk.LabelNew(L.ToString(1))
	if err != nil {
		utils.FatalMessage(err.Error())
	}

	label := &Label{gtkLabel}
	ud := L.NewUserData()
	ud.Value = label
	L.SetMetatable(ud, L.GetTypeMetatable("Label"))
	L.Push(ud)
	return 1
}

func checkLabel(L *lua.LState) *Label{
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Label); ok {
		return v
	}
	L.ArgError(1, "Label expected")
	return nil
}

func luaSetText(L *lua.LState) int {
	l := checkLabel(L)
	l.Label.SetText(L.ToString(1))
	return 1
}

func luaSetHAlign(L *lua.LState) int {
	l := checkLabel(L)
	ud := L.ToUserData(2)
	l.Label.SetHAlign(ud.Value.(gtk.Align))
	l.Label.SetHExpand(true)
	return 1
}

func luaSetVAlign(L *lua.LState) int {
	l := checkLabel(L)
	ud := L.ToUserData(2)
	l.Label.SetVAlign(ud.Value.(gtk.Align))
	l.Label.SetVExpand(true)
	return 1
}

func luaSetMarginTop(L *lua.LState) int {
	l := checkLabel(L)
	l.Label.SetMarginTop(L.ToInt(2))
	return 1
}

func luaSetMarginBottom(L *lua.LState) int {
	l := checkLabel(L)
	l.Label.SetMarginBottom(L.ToInt(2))
	return 1
}

func luaSetMarginStart(L *lua.LState) int {
	l := checkLabel(L)
	l.Label.SetMarginStart(L.ToInt(2))
	return 1
}

func luaSetMarginEnd(L *lua.LState) int {
	l := checkLabel(L)
	l.Label.SetMarginEnd(L.ToInt(2))
	return 1
}

func luaSetWidth(L *lua.LState) int {
	l := checkLabel(L)
	l.Label.GetAllocation().SetWidth(L.ToInt(2))
	return 1
}

func luaSetHeight(L *lua.LState) int {
	l := checkLabel(L)
	l.Label.GetAllocation().SetHeight(L.ToInt(2))
	return 1
}

func luaGetParent(L *lua.LState) int {
	l := checkLabel(L)
	ud := L.NewUserData()
	ud.Value =  l.Label
	L.Push(ud)
	return 1
}