package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(greetingCmd)
	greetingCmd.Flags().StringVar(Msg, "msg", "default", "use it to run greeting")
	greetingCmd.Flags().IntVar(Repetitions, "reps", 0, "indicate how many times msg should be printed")
}

var Msg *string = new(string)
var Repetitions *int = new(int)

func greeting(cmd *cobra.Command, args []string) {
	for i := 0; i < *Repetitions; i++ {
		fmt.Printf("%v\n", *Msg)
	}
}

var greetingCmd *cobra.Command = &cobra.Command{
	Use:   "greet",
	Run:   greeting,
	Short: "will print a greeting",
	Long:  "will take argument and print it as many times as user indicates",
}
