gui = require("GUI")
widgets = require("Widgets")
	
function Init() 
	screen = Screen.new()

	v = Vertical.new()
	v:setHAlign(Align.Fill)
	v:setVAlign(Align.Fill)

	menu = Menu.new()
	menu:addMenuItem("About")

	text = LineEdit.new()
	--text:setHeight()
	text:setEditable(false)
	text:setHAlign(Align.Fill)
	--text:setVAlign(Align.Fill)
	text:setMarginTop(15)
	text:setMarginStart(10)
	text:setMarginEnd(10)

	btn_vert = Vertical.new()
	btn_vert:setMarginTop(5)

	btn_group1 = Horizontal.new()

	btn1 = Button.new("1")
	btn_group1:addWidget(btn1:getParent())

	btn2 = Button.new("2")
	btn_group1:addWidget(btn1:getParent())

	btn3 = Button.new("3")
	btn_group1:addWidget(btn1:getParent())

	btn_vert:addWidget(btn_group1:getParent())

	v:addWidget(text:getParent())
	v:addWidget(btn_vert:getParent())

	screen:addWidget(menu:getParent())
	screen:addWidget(v:getParent())

	gui.setScreen(screen)
end
gui.Init()