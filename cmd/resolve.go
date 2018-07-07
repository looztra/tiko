// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/looztra/tiko/logger"
	"github.com/looztra/tiko/pkg"
	"github.com/spf13/cobra"
)

// resolveCmd represents the resolve command
var resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "resolve the provided URL",
	Long: `resolve the provided URL:
	tiko resolve https://t.co/SCjrdEBmyx`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				logger.Log.Fatal("Something went wrong when trying to print the help stuff")
			}
		} else if len(args) > 1 {
			logger.Log.Fatal("Too many arguments for the 'resolve' subcommand")
		} else {
			logger.Log.Debug("'resolve' command called with arg <", args[0], ">")
			resolved, err := tiko.Resolve(args[0])
			if err != nil {
				logger.Log.Error("Something went wrong when trying to resolve url <" + args[0] + ">")
				logger.Log.Fatal(err)
			}
			fmt.Printf("url <%s> resolved to <%s>\n", args[0], resolved)
		}
	},
}

func init() {
	RootCmd.AddCommand(resolveCmd)
}
