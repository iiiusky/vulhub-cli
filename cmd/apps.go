package cmd

import (
	"github.com/iiiusky/vulhub-cli/core"

	"github.com/spf13/cobra"
)

var filter string

// appsCmd represents the apps command
var appsCmd = &cobra.Command{
	Use:   "apps",
	Short: "Get all supported vulnerability environment information",
	Long:  `Get all supported vulnerability environment information,And support search.`,
	Run: func(cmd *cobra.Command, args []string) {
		core.GetApps(filter)
	},
}

func init() {
	rootCmd.AddCommand(appsCmd)
	appsCmd.Flags().StringVarP(&filter, "filter", "f", "", "Filter condition(Only support path and name filtering)")
}
