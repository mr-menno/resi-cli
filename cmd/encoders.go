package cmd

import (
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"github.com/mr-menno/resi-cli/helper"
	"github.com/mr-menno/resi-cli/resi"
	"fmt"
)

var encodersCmd = &cobra.Command{
	Use:   "encoders",
	Short: "The 'encoders' command fetches the list of encoders.",
	Long: "The 'encoders' command fetches the list of encoders.",
	Run: func(cmd *cobra.Command, args []string) {
		token, err := resi.Authenticate(viper.GetString("username"), viper.GetString("password"))
		if err != nil {
			helper.HandleError(err)
		}
		encoders, err := resi.Encoders(token)
		fmt.Println(" # Encoder Name                   UUID                                 Serial           Status   OperationalState")
		fmt.Println("-- ------------------------------ ------------------------------------ ---------------- -------- ----------------")
		for i, v := range encoders {
			fmt.Printf("%2d %30s %36s %16s %8s %s\n", i, v.Name, v.UUID, v.SerialNumber, v.Status, v.OperationalState)
		}
	},
}

func init() {
	rootCmd.AddCommand(encodersCmd)
}