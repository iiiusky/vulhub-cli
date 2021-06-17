package cmd

import (
	"fmt"
	"github.com/iiiusky/cos-upgrade"
	"github.com/iiiusky/vulhub-cli/utils"
	"github.com/spf13/cobra"
)

// upgradeCmd represents the upgrade command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Online upgrade vulhub cli tools",
	Long:  `Online upgrade vulhub cli tools.`,
	Run: func(cmd *cobra.Command, args []string) {
		u := cos_upgrade.CosUpgrade{
			AppName:        "vulhub-cli",
			CurrentVersion: utils.AppVersion,
			CosBucket:      "builder-1253375937",
			CosLocation:    "ap-beijing",
			CosFolder:      "builder",
			TmpBinFile:     "vulhub-cli" + "_TMP",
			Debug:          false,
		}
		fmt.Println(u.Upgrade())
	},
}

func init() {
	rootCmd.AddCommand(upgradeCmd)
}
