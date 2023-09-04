/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "godo",
	Short: "A CLI-based todo app to manage your tasks.",
	Long: `A simple and efficient CLI-based todo app that
allows you to manage your tasks.`,
}

var cfgFile string
var dataDir string

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	viper.AddConfigPath("$HOME/.godo")
	viper.AddConfigPath("/opt/godo")
	viper.SetConfigName(".godo")
	viper.SetEnvPrefix("GODO")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No configuration file found. Using default values.")
		} else {
			fmt.Println("Error reading the configuration file:", err)
			os.Exit(1)
		}
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "godo config file")

	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(`Unable to detect godo directory. Please set a data directory using --datadir`)
	}

	rootCmd.PersistentFlags().StringVarP(&dataDir, "datadir", "d", home, "the directory to store godo data")
	viper.BindPFlag("datadir", rootCmd.PersistentFlags().Lookup("datadir"))
}
