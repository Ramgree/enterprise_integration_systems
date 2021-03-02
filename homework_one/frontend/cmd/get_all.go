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
	"fmt"
	"todocli/model"
	"todocli/service"

	"github.com/spf13/cobra"
)

// getallCmd represents the get_all command
var getallCmd = &cobra.Command{
	Use:   "get-all",
	Short: "See all the TODOs.",
	Long: "See all the TODOs.",
	Run: func(cmd *cobra.Command, args []string) {
		getAllTodos()
	},
}

func getAllTodos() {
	todos := service.GetAllTodos()
	edges := service.GetAllEdges()

	for _, todo := range todos {

		completeFirst := filterEdges(todo.Id, edges)

		fmt.Println()

		fmt.Println("Id:", todo.Id)
		fmt.Println("Name:",todo.Title)
		fmt.Println("Status:", todo.Status)
		fmt.Println("TODO before this: ", completeFirst)

		fmt.Println()
	}
}

func filterEdges(todoID string, edges model.EdgeList) []string {
	output := []string{}

	for _, edge := range edges {
		if edge.To == todoID {
			output = append(output, edge.From)
		}
	}

	return output
}

func init() {
	rootCmd.AddCommand(getallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
