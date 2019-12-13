/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"strconv"

	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http [status code]",
	Short: "Search status codes",
	Long:  `Search status codes such as for 301 or 422 or ye old 404`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		code, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		if !(code >= 100 && code < 600) {
			return fmt.Errorf("invalid status code: %d", code)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		code, _ := strconv.Atoi(args[0])
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408
		fmt.Printf("https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/%d\n", code)
		fmt.Printf("https://httpstatuses.com/%d\n", code)
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
