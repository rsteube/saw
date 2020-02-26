package cmd

import (
	"fmt"
	"github.com/rsteube/saw/config"
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

// TODO: colorize based on logGroup prefix (/aws/lambda, /aws/kinesisfirehose, etc...)
var groupsConfig config.Configuration

var groupsCommand = &cobra.Command{
	Use:   "groups",
	Short: "List log groups",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		for _, name := range groupNames() {
			fmt.Println(name)
		}
	},
}

func init() {
	groupsCommand.Flags().StringVar(&groupsConfig.Prefix, "prefix", "", "log group prefix filter")
	SawCommand.AddCommand(groupsCommand)

	carapace.Gen(groupsCommand).FlagCompletion(carapace.ActionMap{
		"prefix": carapace.ActionCallback(func(args []string) carapace.Action {
			return carapace.ActionMultiParts('/', groupNames()...)
		}),
	})
}
