package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type infoCmdFlags struct {
	Gender         *string
	CountryOfBirth *string
}

var infoFlags infoCmdFlags = infoCmdFlags{
	Gender:         new(string),
	CountryOfBirth: new(string),
}

func printInfo(cmd *cobra.Command, args []string) {
	var info string = fmt.Sprintf("gender: %v, country of birth: %v", *infoFlags.Gender, *infoFlags.CountryOfBirth)
	fmt.Println(info)
}

var infoCmd *cobra.Command = &cobra.Command{
	Use:   "info",
	Short: "displays info",
	Long:  "will print user info",
	Run:   printInfo,
}

func init() {
	infoCmd.Flags().StringVar(infoFlags.Gender, "gender", "choose one", "it will set the gender to be printed")
	infoCmd.Flags().StringVar(infoFlags.CountryOfBirth, "country", "choose one", "it will set the value for the country to be printed")
	infoCmd.MarkFlagRequired("gender")
}
