pc = require("pc")

	function Init()
		pc.SetSize(250, 280)
		pc.SetResizable(false)
		pc.SetTitle("Calculator")
		pc.Run()
	end
	
pc.Init()