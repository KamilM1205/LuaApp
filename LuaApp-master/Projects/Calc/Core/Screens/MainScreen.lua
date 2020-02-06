--
-- Created by IntelliJ IDEA.
-- User: Камиль
-- Date: 27.01.2020
-- Time: 17:03
-- To change this template use File | Settings | File Templates.
--
Parser = require("Projects/Calc/Core/Parser")
local MainScreen = {}

local screen
local v
local menu
local text
local btn_vert
local btn_group1
local btn1
local btn2
local btn3
local btnC
local btn_group2
local btn4
local btn5
local btn6
local btnR
local btn_group3
local btn7
local btn8
local btn9
local btnP
local btn_group4
local btnM
local btn0
local btnS
local btnD

function btn_click(event, btn)
    text:setText(text:getText()..tostring(btn))
end

function btn_clear(event)
    text:setText("")
end

function btn_result(event)
    parser = Parser:new(text:getText())
    text:setText(parser:Parse())
end

function MainScreen.getScreen()
    screen = Screen.new()

    v = Vertical.new()
    v:setHAlign(Align.Fill)
    v:setVAlign(Align.Fill)

    menu = Menu.new()
    menu:addMenuItem("About")

    text = LineEdit.new("Test")
    text:setEditable(false)
    text:setTextAlign(1)
    text:setHAlign(Align.Fill)
    text:setMarginTop(15)
    text:setMarginStart(10)
    text:setMarginEnd(10)

    btn_vert = Vertical.new()
    btn_vert:setMarginTop(20)
    --btn_vert:setMarginStart(5)

    --------------------------------
    btn_group1 = Horizontal.new()

    btn1 = Button.new("1")
    btn1:setMarginStart(10)
    btn1:setClick("btn_click", 1)
    btn_group1:addWidget(btn1:getParent())

    btn2 = Button.new("2")
    btn2:setMarginStart(10)
    btn2:setClick("btn_click", 2)
    btn_group1:addWidget(btn2:getParent())

    btn3 = Button.new("3")
    btn3:setMarginStart(10)
    btn3:setClick("btn_click", 3)
    btn_group1:addWidget(btn3:getParent())

    btnC = Button.new("C")
    btnC:setMarginStart(10)
    btnC:setClick("btn_clear")
    btn_group1:addWidget(btnC:getParent())

    btn_vert:addWidget(btn_group1:getParent())

    ------------------------------------------
    btn_group2 = Horizontal.new()
    btn_group2:setMarginTop(10)

    btn4 = Button.new("4")
    btn4:setMarginStart(10)
    btn4:setClick("btn_click", 4)
    btn_group2:addWidget(btn4:getParent())

    btn5 = Button.new("5")
    btn5:setMarginStart(10)
    btn5:setClick("btn_click", 5)
    btn_group2:addWidget(btn5:getParent())

    btn6 = Button.new("6")
    btn6:setClick("btn_click", 6)
    btn6:setMarginStart(10)
    btn_group2:addWidget(btn6:getParent())

    btnR = Button.new("=")
    btnR:setMarginStart(10)
    btnR:setClick("btn_result")
    btn_group2:addWidget(btnR:getParent())

    btn_vert:addWidget(btn_group2:getParent())

    ------------------------------------
    btn_group3 = Horizontal.new()
    btn_group3:setMarginTop(10)

    btn7 = Button.new("7")
    btn7:setMarginStart(10)
    btn7:setClick("btn_click", 7)
    btn_group3:addWidget(btn7:getParent())

    btn8 = Button.new("8")
    btn8:setMarginStart(10)
    btn8:setClick("btn_click", 8)
    btn_group3:addWidget(btn8:getParent())

    btn9 = Button.new("9")
    btn9:setMarginStart(10)
    btn9:setClick("btn_click", 9)
    btn_group3:addWidget(btn9:getParent())

    btnP = Button.new("+")
    btnP:setMarginStart(10)
    btnP:setClick("btn_click", "+")
    btn_group3:addWidget(btnP:getParent())

    btn_vert:addWidget(btn_group3:getParent())

    btn_group4 = Horizontal.new()
    btn_group4:setMarginTop(10)

    btnM = Button.new("-")
    btnM:setMarginStart(10)
    btnM:setClick("btn_click", "-")
    btn_group4:addWidget(btnM:getParent())

    btn0 = Button.new("0")
    btn0:setMarginStart(10)
    btn0:setClick("btn_click", 0)
    btn_group4:addWidget(btn0:getParent())

    btnS = Button.new("*")
    btnS:setMarginStart(10)
    btnS:setClick("btn_click", "*")
    btn_group4:addWidget(btnS:getParent())

    btnD = Button.new("/")
    btnD:setMarginStart(10)
    btnD:setClick("btn_click", "/")
    btn_group4:addWidget(btnD:getParent())

    btn_vert:addWidget(btn_group4:getParent())

    v:addWidget(text:getParent())
    v:addWidget(btn_vert:getParent())

    screen:addWidget(menu:getParent())
    screen:addWidget(v:getParent())
    return screen
end

return MainScreen
