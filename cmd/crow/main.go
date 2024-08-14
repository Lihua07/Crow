package main

import (
	"crow/internal/crow"
	"os"
)

func main() {
    command := crow.NewCrowCommand()
    if err := command.Execute(); err != nil {
        os.Exit(1)
    }
}
