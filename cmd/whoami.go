package cmd

import (
	"github.com/mr-menno/resi-cli/helper"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/mr-menno/resi-cli/resi"
)

// viewCmd represents the view command
var whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "The 'whoami' command will show details about my login.",
	Long: `The 'whoami' command will show details about my login.`,
	Run: func(cmd *cobra.Command, args []string) {
		token, err := resi.Authenticate(viper.GetString("username"), viper.GetString("password"))
		if err != nil {
			helper.HandleError(err)
		}

		me, err := resi.Whoami(token)
		if err != nil {
			helper.HandleError(err)
		}
		fmt.Println("userName:   "+me.UserName)
		fmt.Println("userId:     "+me.UserId)
		fmt.Println("customerId: "+me.CustomerId)
	},
}

func init() {
	authCmd.AddCommand(whoamiCmd)
}