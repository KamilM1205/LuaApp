pc = require("pc")

	function Init()
		pc.SetSize(300, 400)
		pc.SetResizable(false)
		pc.SetTitle("Calculator")
		pc.Run()
	end
	
pc.Init()