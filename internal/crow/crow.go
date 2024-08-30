package crow

import (
    "os"
    "os/signal"
    "syscall"
    "github.com/spf13/cobra"
)

func NewCrowCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use: "crow",
        SilenceUsage: true,
        Short: "crow a mini command line tools.",
        Long: "crow a mini powerful command line tools.",
        RunE: func(cmd *cobra.Command, args []string) error {
            return runService()
        },
    }
    return cmd
}

func runService() error {
    quit := make(chan os.Signal, 1)

    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

    <-quit
    return nil
}

