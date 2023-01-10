package cmd

import (
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"github.com/mr-menno/resi-cli/helper"
	"github.com/mr-menno/resi-cli/resi"
)

// configCmd represents the config command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "The 'auth' command validated username and password.",
	Long: "The 'auth' command validated username and password.",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := resi.Authenticate(viper.GetString("username"), viper.GetString("password"))

		if err != nil {
			helper.HandleError(err)
		}
		helper.JsonResult("ok","successfully authenticated")
	},
}

var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "The 'token' fetches the session token.",
	Long: "ThThe 'token' fetches the session token.",
	Run: func(cmd *cobra.Command, args []string) {
		token, err := resi.Authenticate(viper.GetString("username"), viper.GetString("password"))

		if err != nil {
			helper.HandleError(err)
		}
		helper.JsonResult("ok",token)
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.AddCommand(tokenCmd)
}