package label

import (
	"../../utils"
	"github.com/gotk3/gotk3/gtk"
	"github.com/yuin/gopher-lua"
)

type Label struct {
	Label *gtk.Label
}

func RegisterLabelType(L *lua.LState) int {
	mt := L.NewTypeMetatable("Label")
	L.SetGlobal("Label", mt)
	
	L.SetField(mt, "new", L.NewFunction(newLabel))
	L.SetField(mt, "__index", L.SetFuncs(L.NewTable(), labelMethods))
	return 1
}

var labelMethods = map[string]lua.LGFunction{
	"setText": luaSetText,
	"getParent": luaGetParent,
}

func newLabel(L *lua.LState) int{
	gtkLabel, err := gtk.LabelNew(L.ToString(1))
	if err != nil {
		utils.FatalMessage(err.Error())
	}

	screen := &Label{gtkLabel}
	ud := L.NewUserData()
	ud.Value = screen
	L.SetMetatable(ud, L.GetTypeMetatable("Label"))
	L.Push(ud)
	return 1
}

func checkLabel(L *lua.LState) *Label{
	ud := L.CheckUserData(1)
	if v, ok := ud.Value.(*Label); ok {
		return v
	}
	L.ArgError(1, "Screen expected")
	return nil
}

func luaSetText(L *lua.LState) int {
	l := checkLabel(L)
	l.Label.SetText(L.ToString(1))
	return 1
}

func luaGetParent(L *lua.LState) int {
	l := checkLabel(L)
	ud := L.NewUserData()
	ud.Value =  l.Label
	L.Push(ud)
	return 1
}