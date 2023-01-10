package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"github.com/mr-menno/resi-cli/helper"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:	"resi",
	Short: "A simple CLI to control resi.io",
	Long: "A simple CLI to control resi.io",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var strUsername string
var strPassword string

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP((&strUsername), "username", "u", "", "username required")
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("resi.username"))
	rootCmd.PersistentFlags().StringVarP((&strPassword), "password", "p", "", "password required")
}

func initConfig() {
	configFile = ".resi-io.yml"
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)

	
	viper.AutomaticEnv()
	viper.SetEnvPrefix("RESI")
	helper.HandleError(viper.BindEnv("USERNAME"))
	helper.HandleError(viper.BindEnv("PASSWORD"))

	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using configuration file: ", viper.ConfigFileUsed())
	}

	if viper.GetString("username") == "" {
		username, err := helper.PromptText("resi.io username")
		if err != nil {
			helper.HandleError(err)
		}
		viper.Set("username",username)
		viper.WriteConfig()
	}
	if viper.GetString("password") == "" {
		fmt.Println("Logging in with username: "+viper.GetString("username"))
		password, err := helper.PromptPassword("resi.io password")
		if err != nil {
			helper.HandleError(err)
		}
		viper.Set("password",password)
		viper.WriteConfig()
	}

}
