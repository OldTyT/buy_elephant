package message

import (
	u "github.com/OldTyT/buy_elephant/bots"
	"github.com/OldTyT/buy_elephant/utils/log"
	"gopkg.in/telebot.v3"
)

type Message struct{}

func (m *Message) Init() error {
	log.Info.Println("initialized module `Message`")
	return nil
}

func (m *Message) Commands() []*u.Command {
	return []*u.Command{
		{
			Text:    true,
			Handler: message,
		},
	}
}

func message(c telebot.Context) error {
	return c.Reply(c.Text())
}
