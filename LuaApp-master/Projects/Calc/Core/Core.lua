gui = require("GUI")
widgets = require("Widgets")
mainScreen = require("Projects/Calc/Core/Screens/MainScreen")
	
function Init()
	gui.setScreen(mainScreen:getScreen())
end
gui.Init()