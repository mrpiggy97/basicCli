/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import "github.com/mrpiggy97/basicCli/cmd"

func main() {
	var err error = cmd.RootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
