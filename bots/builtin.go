package bots

import (
	"gopkg.in/telebot.v3"
)

func loadMiddlewares() {
	ignoreFwdCommands()
}

func loadBuiltins() {
	ping()
	help()
	start()
}

func ping() {
	Get().AddCommand(&Command{
		Cmd:  "/ping",
		Desc: "check if bot is active",
		Args: nil,
		Handler: func(c telebot.Context) error {
			return c.Reply("pong")
		},
	},
	)
}

func help() {
	Get().AddCommand(&Command{
		Cmd:  "/help",
		Desc: "usage for the bot",
		Args: nil,
		Handler: func(c telebot.Context) error {
			return c.Reply("help")
		},
	})
}

func start() {
	Get().AddCommand(&Command{
		Cmd:  "/start",
		Desc: "start message",
		Args: nil,
		Handler: func(c telebot.Context) error {
			if c.Message().Private() {
				msg := "start_"

				return c.Send(msg, &telebot.SendOptions{
					DisableWebPagePreview: true,
				})
			}
			return nil
		},
	})
}

func ignoreFwdCommands() {
	Get().Bot.Use(func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			if msg := c.Message(); msg != nil && msg.IsForwarded() {
				return nil
			}

			return next(c)
		}
	})
}
