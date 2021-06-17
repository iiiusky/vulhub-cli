package cmd

import (
	"fmt"
	"github.com/iiiusky/vulhub-cli/core"
	"github.com/iiiusky/vulhub-cli/utils"
	"github.com/spf13/cobra"
	"os"
)

var configPath string
var mirror string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vulhub-cli",
	Short: "A vulhub cli tools",
	Long:  `A vulhub cli tools.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		utils.InitConfigPath(configPath)
		config := utils.Config{Mirror: mirror}
		utils.InitConfig(config)
		if !utils.InitDBFile() {
			core.CacheDB()
			if !utils.InitDBFile() {
				os.Exit(-1)
			}
		}
	},
	ValidArgs: []string{"apps", "deploy", "version", "upgrade", "cache"},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&configPath, "configPath", "p", "$HOME/.vulhub", "")
	rootCmd.PersistentFlags().StringVar(&mirror, "mirror", "raw.githubusercontent.com", "")
}
