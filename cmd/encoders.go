package cmd

import (
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"github.com/mr-menno/resi-cli/helper"
	"github.com/mr-menno/resi-cli/resi"
	"fmt"
	"errors"
	"encoding/json"
)

var strEncoderUUID string
var strEventUuid string
var strPresetUuid string

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

		if(jsonOutput) {
			if strEncoderUUID != "" {
				for _, v := range encoders {
					if strEncoderUUID != "" {
						if v.UUID == strEncoderUUID {
							jsonOutput, _ := json.Marshal(v)
							fmt.Println(string(jsonOutput))
							break
						}
					}
				}
			} else {
				jsonOutput, _ := json.Marshal(encoders)
				fmt.Println(string(jsonOutput))
			}	
		} else {
			fmt.Println(" # Encoder Name                   UUID                                 Serial           Status   OperationalState")
			fmt.Println("-- ------------------------------ ------------------------------------ ---------------- -------- ----------------")
			for i, v := range encoders {
				if strEncoderUUID != "" {
					if v.UUID != strEncoderUUID {
						continue
					}
				}
				fmt.Printf("%2d %30s %36s %16s %8s %s\n", i, v.Name, v.UUID, v.SerialNumber, v.Status, v.OperationalState)
			}
		}
	},
}

var encoderStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "The 'stop' command stops a specific encoder.",
	Long: "The 'stop' command stops a specific encoder.",
	Run: func(cmd *cobra.Command, args []string) {
		if strEncoderUUID == "" {
			helper.HandleError(errors.New("ERROR: missing --encoder-uuid <uuid>"))
		}
		token, err := resi.Authenticate(viper.GetString("username"), viper.GetString("password"))
		if err != nil {
			helper.HandleError(err)
		}
		result, err := resi.StopEncoder(token, strEncoderUUID)
		if err != nil {
			helper.HandleError(err)
		}
		if result {
			fmt.Println("successfully stopped encoder")
		} else {
			fmt.Println("failed to stop encoder")
		}
	},
}

var encoderStartCmd = &cobra.Command{
	Use:   "start",
	Short: "The 'start' command starts a specific encoder.",
	Long: "The 'start' command starts a specific encoder.",
	Run: func(cmd *cobra.Command, args []string) {
		if strEncoderUUID == "" {
			helper.HandleError(errors.New("ERROR: missing --encoder-uuid <uuid>"))
		}
		if strEventUuid == "" {
			helper.HandleError(errors.New("ERROR: missing --event-uuid <uuid>"))
		}
		if strPresetUuid == "" {
			helper.HandleError(errors.New("ERROR: missing --webevent-uuid <uuid>"))
		}
		token, err := resi.Authenticate(viper.GetString("username"), viper.GetString("password"))
		if err != nil {
			helper.HandleError(err)
		}
		result, err := resi.StartEncoder(token,strEncoderUUID,strEventUuid,strPresetUuid)
		if err != nil {
			helper.HandleError(err)
		}
		if result {
			fmt.Println("successfully started encoder")
		} else {
			fmt.Println("failed to start encoder")
		}
	},
}

func init() {
	encodersCmd.PersistentFlags().StringVarP((&strEncoderUUID), "encoder-uuid", "", "", "encoder-uuid")
	encodersCmd.PersistentFlags().StringVarP((&strEventUuid), "event-uuid", "", "", "event-uuid")
	encodersCmd.PersistentFlags().StringVarP((&strPresetUuid), "preset-uuid", "", "", "preset-uuid")
	encodersCmd.AddCommand(encoderStopCmd)
	encodersCmd.AddCommand(encoderStartCmd)
	rootCmd.AddCommand(encodersCmd)
}