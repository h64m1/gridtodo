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
	"fmt"
	"io/ioutil"

	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
)

// 3*3 panel = (3*3 item) * 9
type Todo struct {
	Goal  string
	Panel [9]Panel
}

// 1 panel = 3*3 item
type Panel struct {
	Item [9]string
}

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
		fmt.Println("generate called")
		// var yamlName string
		// if len(args) == 0 {
		// 	yamlName = "todo.yaml"
		// } else {
		// 	yamlName = args[0]
		// }

		todo := Todo{
			Goal:  "",
			Panel: [9]Panel{},
		}

		bytes, err := yaml.Marshal(todo)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(bytes))

		buf, err := ioutil.ReadFile("test.yml")
		if err != nil {
			return
		}

		err = yaml.Unmarshal([]byte(buf), &todo)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%v\n", todo)
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
