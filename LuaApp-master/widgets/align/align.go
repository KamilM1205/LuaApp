package align

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/yuin/gopher-lua"
)

//RegisterAlignType функция регистрирующая Align
func RegisterAlignType(L *lua.LState) int {
	mt := L.NewTypeMetatable("Align")
	L.SetGlobal("Align", mt)

	udCenter := L.NewUserData()
	udCenter.Value = gtk.ALIGN_CENTER
	L.SetField(mt, "Center", udCenter)

	udFill := L.NewUserData()
	udFill.Value = gtk.ALIGN_FILL
	L.SetField(mt, "Fill", udFill)

	udEnd := L.NewUserData()
	udEnd.Value = gtk.ALIGN_FILL
	L.SetField(mt, "End", udEnd)

	udStart := L.NewUserData()
	udStart.Value = gtk.ALIGN_FILL
	L.SetField(mt, "Start", udStart)
	return 1
}