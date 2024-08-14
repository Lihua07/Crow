package crow

import "github.com/spf13/cobra"

func NewCrowCommand() *cobra.Command {
    cmd := &cobra.Command{
        Use: "crow",
    }
    return cmd
}
