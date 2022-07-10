package message

import (
	"strings"

	u "github.com/OldTyT/buy_elephant/bots"
	"github.com/OldTyT/buy_elephant/utils/log"
	"github.com/OldTyT/buy_elephant/utils/style"
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
			Cmd:  "/me",
			Desc: "command from IRC",
			Args: []u.Arg{
				{
					Name:     "action",
					Desc:     "the action you want to perform",
					Required: true,
				},
			},
			Handler: me,
		},
	}
}

func me(c telebot.Context) error {
	if len(c.Args()) == 0 {
		return c.Reply("please provide an action")
	}

	return c.Send(
		style.Bold(c.Sender().FirstName + " " + c.Sender().LastName + " " + strings.Join(c.Args(), " ")),
	)
}
