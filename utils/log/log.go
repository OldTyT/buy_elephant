package log

import (
	"log"
	"os"
)

var (
	reset   = "\033[0m"
	warning = "\033[33m"
	error   = "\033[31m"
	info    = "\033[32m"

	Warn  = log.New(os.Stderr, warning+"WARN : "+reset, log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, error+"ERROR: "+reset, log.Ldate|log.Ltime|log.Lshortfile)
	Info  = log.New(os.Stderr, info+"INFO : "+reset, log.Ltime|log.Lshortfile)
)
