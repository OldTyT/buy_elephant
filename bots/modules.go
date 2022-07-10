package bots

import "github.com/OldTyT/buy_elephant/utils/log"

type Module interface {
	Init() error
	Commands() []*Command
}

func loadModules(modules []Module) {
	for _, m := range modules {
		err := m.Init()
		if err != nil {
			log.Warn.Println(err)
			break
		}

		for _, c := range m.Commands() {
			Get().AddCommand(c)
		}
	}
}
