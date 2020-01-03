gui = require("GUI")
widgets = require("Widgets")

function btn_click(event)
    print(event)
end

function Init() 
    screen = Screen.new()
    h = Vertical.new()
    edt = LineEdit.new("gfhg")
    btn1 = Button.new("1")
    btn2 = Button.new("2")
    btn2:setEnable(false)
    h:addWidget(btn1:getParent())
    h:addWidget(btn2:getParent())
    screen:addWidget(edt:getParent())
    screen:addWidget(h:getParent())
    gui.setScreen(screen)
end
gui.Init()