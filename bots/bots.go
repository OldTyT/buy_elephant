package bots

import (
	"os"
	"sync"

	"github.com/OldTyT/buy_elephant/utils/log"
	tb "gopkg.in/telebot.v3"
)

type Arg struct {
	Name     string
	Desc     string
	Required bool
}

type Command struct {
	Cmd     string
	Desc    string
	Args    []Arg
	Handler tb.HandlerFunc
}

type my_bot struct {
	Bot *tb.Bot

	// sql.DB for disabling etc
	mut      sync.Mutex
	commands map[string]*Command
}

var instance *my_bot

func Get() *my_bot {
	return instance
}

func Run(m ...Module) {
	instance = new(my_bot)

	s := tb.Settings{
		Token:     os.Getenv("TOKEN"),
		ParseMode: "HTML",
		OnError: func(e error, c tb.Context) {
			log.Warn.Println(e)
		},
	}

	var err error
	Get().Bot, err = tb.NewBot(s)
	if err != nil {
		log.Error.Fatalln(err)
	}
	Get().commands = make(map[string]*Command)

	loadMiddlewares()
	loadBuiltins()
	loadModules(m)
	Get().setCommands()

	log.Info.Println("starting as: " + Get().Bot.Me.Username)
	Get().Bot.Start()
}

func (u *my_bot) AddCommand(c *Command) {
	u.mut.Lock()
	defer u.mut.Unlock()

	u.commands[c.Cmd[1:]] = c
	u.Bot.Handle(c.Cmd, c.Handler)
}

func (u *my_bot) setCommands() {
	var cmds []tb.Command

	for _, c := range u.commands {
		cmds = append(cmds, tb.Command{
			Text:        c.Cmd[1:],
			Description: c.Desc,
		})
	}

	if err := u.Bot.SetCommands(cmds); err != nil {
		log.Error.Println(err)
	}
}

// add builtin middleware that rejects all disabled commands
// func (u *my_bot) disableLocal(cmd string, chatID int) {
// }

// func (u *my_bot) disableGlobal(cmd string) {
// }
