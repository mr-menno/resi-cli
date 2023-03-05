package cmd

import (
	"github.com/mr-menno/resi-cli/helper"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/mr-menno/resi-cli/resi"
)

var profilesCmd = &cobra.Command{
	Use:   "profiles",
	Short: "The 'profiles' command fetches the list of profiles.",
	Long: "The 'profiles' command fetches the list of profiles.",
}

var eventProfilesCmd = &cobra.Command{
	Use:   "event",
	Short: "The 'event' command will show fetch event profiles.",
	Long: `The 'event' command will show fetch event profiles.`,
	Run: func(cmd *cobra.Command, args []string) {
		token, err := resi.Authenticate(viper.GetString("username"), viper.GetString("password"))
		if err != nil {
			helper.HandleError(err)
		}

		me, err := resi.Whoami(token)
		if err != nil {
			helper.HandleError(err)
		}

		eventProfiles, err := resi.EventProfiles(token, me.CustomerId)
		fmt.Println(" # UUID                                 Name")
		fmt.Println("-- ------------------------------------ ----")
		for i, v := range eventProfiles {
			fmt.Printf("%2d %36s %s\n", i, v.UUID, v.Name)
		}

	},
}

var encoderProfilesCmd = &cobra.Command{
	Use:   "encoder",
	Short: "The 'encoder' command will show fetch encoder profiles.",
	Long: `The 'encoder' command will show fetch encoder profiles.`,
	Run: func(cmd *cobra.Command, args []string) {
		token, err := resi.Authenticate(viper.GetString("username"), viper.GetString("password"))
		if err != nil {
			helper.HandleError(err)
		}

		me, err := resi.Whoami(token)
		if err != nil {
			helper.HandleError(err)
		}

		encoderProfiles, err := resi.EncoderProfiles(token, me.CustomerId)
		fmt.Println(" # UUID                                 Name")
		fmt.Println("-- ------------------------------------ ----")
		for i, v := range encoderProfiles {
			fmt.Printf("%2d %36s %s\n", i, v.UUID, v.Name)
		}

	},
}

var webEventProfilesCmd = &cobra.Command{
	Use:   "webevent",
	Short: "The 'webevent' command will show fetch webevent profiles.",
	Long: `The 'webevent' command will show fetch webevent profiles.`,
	Run: func(cmd *cobra.Command, args []string) {
		token, err := resi.Authenticate(viper.GetString("username"), viper.GetString("password"))
		if err != nil {
			helper.HandleError(err)
		}

		me, err := resi.Whoami(token)
		if err != nil {
			helper.HandleError(err)
		}

		webEventProfiles, err := resi.WebEventProfiles(token, me.CustomerId)
		fmt.Println(" # UUID                                 Active\tName")
		fmt.Println("-- ------------------------------------ ------\t----")
		for i, v := range webEventProfiles {
			fmt.Printf("%2d %36s %s\t%s\n", i, v.UUID, v.Active, v.Name)
		}

	},
}

func init() {
	rootCmd.AddCommand(profilesCmd)
	profilesCmd.AddCommand(eventProfilesCmd)
	profilesCmd.AddCommand(encoderProfilesCmd)
	profilesCmd.AddCommand(webEventProfilesCmd)
}
