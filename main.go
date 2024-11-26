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
	fuzzel := flag.Bool("fuzzel", false, "Run fuzzel mode")
	resume := flag.Bool("resume", false, "Launch last savegame. Rofi mode only.")
	flag.Parse()

	base.Config()

	if *rofi {
		rofimode.RunRofiMode("rofi", *resume)
		return
	}

	if *wofi {
		rofimode.RunRofiMode("wofi", *resume)
		return
	}

	if *dmenu {
		rofimode.RunRofiMode("dmenu", *resume)
		return
	}

	if *tofi {
		rofimode.RunRofiMode("tofi", *resume)
		return
	}

	if *fuzzel {
		rofimode.RunRofiMode("fuzzel", *resume)
		return
	}

	//cfg.GetInstance().Configured = false
	tui.Draw()
}
