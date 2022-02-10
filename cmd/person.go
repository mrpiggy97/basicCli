package cmd

import "github.com/spf13/cobra"

type personCmdFlags struct {
	Name *string
	Age  *int
}

var personFlags personCmdFlags = personCmdFlags{
	Name: new(string),
	Age:  new(int),
}

var personCmd *cobra.Command = &cobra.Command{
	Use:   "person",
	Short: "it has a bunch of sub commands to use",
	Long:  "use its subcommands to print greetings and stuff",
}

func init() {
	personCmd.AddCommand(greetingCmd, infoCmd)
	personCmd.PersistentFlags().StringVar(personFlags.Name, "name", "default", "sets name value")
	personCmd.PersistentFlags().IntVar(personFlags.Age, "age", 0, "sets age value")
}
