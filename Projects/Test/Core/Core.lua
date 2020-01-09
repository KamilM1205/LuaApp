gui = require("GUI")
widgets = require("Widgets")

function btn_click(event)
    print(event)
end

function Init() 
    screen = Screen.new()
    v = Vertical.new()
    v:setHAlign(Align.Fill)
    v:setVAlign(Align.Fill)
    label = Label.new("Hello, world")
    label:setHAlign(Align.Center)
    label:setVAlign(Align.Center)
    v:addWidget(label:getParent())
    screen:addWidget(v:getParent())
    gui.setScreen(screen)
end
gui.Init()