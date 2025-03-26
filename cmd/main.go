package cmd

import (
	"context"

	"github.com/GreekMilkBot/Core"
)

func Main() {
	bot := Core.NewBot(context.Background())
	bot.Start()
}
