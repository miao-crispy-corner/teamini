package main

import (
	"os"

	_ "go.uber.org/automaxprocs"

	"github.com/miao-crispy-corner/teamini/internal/teamini"
)

// Go 程序的默认入口函数(主函数).
func main() {
	command := teamini.NewTeaMiniCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
