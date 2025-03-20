package main

import (
	cmdMain "github.com/GreekMilkBot/Core/cmd"

	// plug in modules here
	_ "github.com/GreekMilkBot/Core/modules/standard"
)

func main() {
	cmdMain.Main()
}
