package main

import (
	"crow/internal/dictagraph"
	"os"
)

func main() {
    command := dictagraph.NewDictagraphCommand()
    if err := command.Execute(); err != nil {
        os.Exit(1)
    }
}

