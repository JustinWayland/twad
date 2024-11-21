package rofimode

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"

	"github.com/zmnpl/twad/base"
	"github.com/zmnpl/twad/games"
)

// RunRofiMode starts rofi (or any other dmenu-like program) to select and run a already created game.
// It pipes all games as a list of names to the external program
func RunRofiMode(command string) {
	base.EnableBasePath()
	var params []string
	prompt := "Rip & Tear: "
	if command == "rofi" && commandExists("rofi") {
		params = []string{"--dmenu", "-p", prompt}
	} else if command == "wofi" && commandExists("wofi") {
		params = []string{"--dmenu", "-p", prompt}
	} else if command == "dmenu" && commandExists("dmenu") {
		params = []string{"-p", prompt}
	} else if command == "tofi" && commandExists("tofi") {
		// TODO: Add placeholder text so that Rip & Tear can live on
		params = []string{"--prompt-text", prompt}
	} else if command == "fuzzel" && commandExists("fuzzel") {
		// TODO: Add placeholder text so that Rip & Tear can live on
		params = []string{"-d", "-p", prompt}
	} else {
		return
	}

	rofi := exec.Command(command, params...)
	r, w := io.Pipe()
	rofi.Stdin = r
	var stdout bytes.Buffer
	rofi.Stdout = &stdout
	err := rofi.Start()
	if err != nil {
		//return err
	}

	// create map with games / indices to later on select and start one from rofi
	rofiToGame := make(map[string]int)
	for i, v := range games.Games() {
		displayName := fmt.Sprintf("%v: %s\n", i+1, v.Name)
		rofiToGame[displayName] = i
		w.Write([]byte(displayName)) // pipe game name to rofi
	}
	w.Close()

	rofi.Wait()

	result := stdout.String()
	fmt.Println(result)

	// run selected game
	if i, exists := rofiToGame[result]; exists {
		games.Games()[i].Run()
	}
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
