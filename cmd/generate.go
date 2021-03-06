/*
Copyright © 2021 mass8

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"io/ioutil"
	"log"

	"github.com/goccy/go-yaml"
	"github.com/h64m1/gridtodo/todo"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a yaml file (default is todo.yaml) for markdown table",
	Long: `Generate a yaml file:
	Use subcommand 'generate' to create template yaml file for gridtodo
	$ gridtodo generate
	where default yaml file name is 'todo.yaml'.

	Write filename with 'generate' command if you want to specify the output yaml file
	$ gridtodo generate test.yaml
	which displays markdown table in standard output based on the test.yaml.`,
	Run: func(cmd *cobra.Command, args []string) {
		var yamlFile string
		if len(args) == 0 {
			yamlFile = "todo.yml"
		} else {
			yamlFile = args[0]
		}

		todo := todo.Todo{
			Goal:  "",
			Panel: [9]todo.Panel{},
		}

		data, err := yaml.Marshal(todo)
		if err != nil {
			panic(err)
		}

		err = ioutil.WriteFile(yamlFile, data, 0644)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
