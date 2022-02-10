package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type greetingCmdFlags struct {
	Msg         *string
	Repetitions *int
}

//flags for greetingCmd
var greetingFlags greetingCmdFlags = greetingCmdFlags{
	Msg:         new(string),
	Repetitions: new(int),
}

//function to be run by greetingCmd
func greeting(cmd *cobra.Command, args []string) {
	for i := 0; i < *greetingFlags.Repetitions; i++ {
		fmt.Printf("%v\n", *greetingFlags.Msg)
		fmt.Printf("my name is %v\n", *personFlags.Name)
	}
}

var greetingCmd *cobra.Command = &cobra.Command{
	Use:   "greet",
	Run:   greeting,
	Short: "will print a greeting",
	Long:  "will take argument and print it as many times as user indicates",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("yahallo")
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("this is goodbye")
	},
}

func init() {
	greetingCmd.Flags().StringVar(greetingFlags.Msg, "msg", "default", "use it to run greeting")
	greetingCmd.Flags().IntVar(greetingFlags.Repetitions, "reps", 0, "indicate how many times msg should be printed")
	greetingCmd.MarkFlagRequired("reps")
}
