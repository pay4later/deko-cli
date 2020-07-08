package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "deko-cli",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(configCmd)
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
	home, err := homedir.Dir()
	if err != nil {
		er(err)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(".deko-cli")
	viper.SetConfigType("yaml")

	if _, err := os.Stat(viper.ConfigFileUsed()); err != nil {
		viper.SafeWriteConfig()
	}

	if err := viper.ReadInConfig(); err != nil {
		er(err)
	}
}