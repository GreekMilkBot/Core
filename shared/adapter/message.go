package adapter

import (
	"github.com/GreekMilkBot/Core/shared/common"
)

type Message struct {
	// Adapter Name
	Platform string
	BotID    string
	UserID   string

	Content common.Content // must
	Quote   *Message       // optional
}

type GroupMessage struct {
	GroupID string
	Message
}
