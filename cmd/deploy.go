package cmd

import (
	"github.com/iiiusky/vulhub-cli/core"
	"github.com/iiiusky/vulhub-cli/utils"
	"github.com/spf13/cobra"
)

var appId int
var appPath string
var force bool
var destPath string

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy vulhub files to local",
	Long:  `Deploy vulhub files to local.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.ForceFlag = force
		if appId != -1 {
			var appInfo *utils.VulhubAppDBStruct
			appInfo = core.GetAppInfoForId(appId)
			appPath = appInfo.Path
		}

		if appPath == "" {
			return
		}

		core.Deploy(appPath, destPath)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.Flags().IntVarP(&appId, "appId", "i", -1, "App ID")
	deployCmd.Flags().StringVarP(&appPath, "appPath", "a", "", "App Path")
	deployCmd.Flags().StringVarP(&destPath, "destPath", "d", "vulhub-apps", "Deployment path")
	deployCmd.Flags().BoolVar(&force, "force", false, "Re-download anyway")
}
