package bootstrap

import (
	"github.com/oxyrinchus/goilerplate/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "goilerplate",
	Short:            "For the perfect architecture.",
	Long:             `For the perfect architecture.`,
	TraverseChildren: true,
}

type App struct {
	*cobra.Command
}

func NewApp() App {
	cmd := App{
		Command: rootCmd,
	}

	cmd.AddCommand(commands.GetSubCommands(CommonModules)...)

	return cmd
}

var RootApp = NewApp()
