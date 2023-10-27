package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"project-layout/internal/config"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "project-layout",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello project-layout CLI")
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra-demo.yaml)")
}

func initConfig() {
	conf := config.GetConfig()
	conf.Viper.SetConfigType("toml")

	if cfgFile != "" {
		conf.Viper.SetConfigFile(cfgFile)
	} else {
		conf.Viper.AddConfigPath("./configs")
		conf.Viper.SetConfigName("config.toml")
	}

	conf.Viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := conf.Viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", conf.Viper.ConfigFileUsed())
	}

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
