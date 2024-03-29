/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"todocli/service"

	"github.com/spf13/cobra"
)

// todoCmd represents the todo command
var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "A command to add a new item to todo list",
	Long: `To use this command, type 'todocli todo "Insert your task here!"'.
	 Optionally, you may want to add dependencies (tasks which are to be completed before this task),
	 to do this simply add the ids with whitespaces like:
	  'todocli todo "Insert your task here!" 1 2 3' - to make the program depend on task 1,2 and 3
	  'todocli todo "Insert your task here!" 1 2 3 4 5' - to make the program depend on task 1,2,3, 4 and 5s
	  `,
	Run: func(cmd *cobra.Command, args []string) {
		var id = service.AddTodo(args[0])
		if id != "" {
			service.AddDependencies(id, args[1:])
		}
	},
}

func init() {
	rootCmd.AddCommand(todoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
