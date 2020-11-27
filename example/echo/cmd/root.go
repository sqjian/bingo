package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "bingo",
	Short: "basic skeleton",
}
var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	_ = rootCmd.MarkPersistentFlagRequired("config")
}

func initConfig() {
	if cfgFile != "" {
		_, err := os.Lstat(cfgFile)
		if os.IsNotExist(err) {
			panic(fmt.Sprintf("%v IsNotExist", cfgFile))
		}
	}
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
