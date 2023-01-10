package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "The 'view' subcommand will provide a list of keys and a map of the values.",
	Long: `The 'view' subcommand will provide a list of keys and a map of the values. For example:
'<cmd> config view'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("** All keys including environment variables for CLI.\n")
		fmt.Printf("%s\n\n", viper.AllKeys())
		fmt.Printf("** Bind vars.\n")
		bind := rootCmd.PersistentFlags().Lookup("username").Value.String()
		fmt.Printf("Binding definition provided: %s\n", bind)
		fmt.Printf("Binding definition provided from viper: %s\n", viper.GetString("USERNAME"))
	
		settings := viper.AllSettings()
		fmt.Printf("** Configuration file keys and values.\n")
		for i, v := range settings {
			fmt.Printf("%v: %v\n", i, v)
		}
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}