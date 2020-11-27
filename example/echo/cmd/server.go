package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/sqjian/bingo/acceptor/entrance"
	"log"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "example server",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		viperInst := func() *viper.Viper {
			viperInst := viper.New()
			viperInst.SetConfigFile(cfgFile)
			err := viperInst.ReadInConfig()
			if err == nil {
				log.Printf("using config file:%v", cfgFile)
			} else {
				log.Fatalf("read config failed,err:%v", err)
			}
			return viperInst
		}()
		if err := entrance.Serve(viperInst); err != nil {
			log.Fatalf("serve failed,err:%v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
