package adapter

import "github.com/GreekMilkBot/Core/share"

type Message struct {
	// Adapter Name
	Platform string
	BotID    string
	UserID   string

	Content share.Content // must
	Quote   *Message      // optional
}

type GroupMessage struct {
	GroupID string
	Message
}
