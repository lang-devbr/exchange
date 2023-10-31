package main

import (
	"github.com/lang-dev/exchange/cmd"
	"github.com/lang-dev/exchange/config"
)

func init() {
	config.Load()
}

func main() {
	cmd.Execute()
}
