package cmd

import (
	"github.com/iiiusky/vulhub-cli/core"

	"github.com/spf13/cobra"
)

// cacheCmd represents the cache command
var cacheCmd = &cobra.Command{
	Use:   "cache",
	Short: "Download cached data file information",
	Long:  `Download cached data file information.`,
	Run: func(cmd *cobra.Command, args []string) {
		core.CacheDB()
	},
}

func init() {
	rootCmd.AddCommand(cacheCmd)
}
