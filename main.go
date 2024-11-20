package main

import (
	"flag"

	"github.com/zmnpl/twad/base"
	"github.com/zmnpl/twad/rofimode"
	"github.com/zmnpl/twad/tui"
)

func main() {
	rofi := flag.Bool("rofi", false, "Run rofi mode.")
	wofi := flag.Bool("wofi", false, "Run wofi mode.")
	dmenu := flag.Bool("dmenu", false, "Run dmenu mode.")
	tofi := flag.Bool("tofi", false, "Run tofi mode.")
	flag.Parse()

	base.Config()

	if *rofi {
		rofimode.RunRofiMode("rofi")
		return
	}

	if *wofi {
		rofimode.RunRofiMode("wofi")
		return
	}

	if *dmenu {
		rofimode.RunRofiMode("dmenu")
		return
	}

	if *tofi {
		rofimode.RunRofiMode("tofi")
		return
	}

	//cfg.GetInstance().Configured = false
	tui.Draw()
}
