// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/fyuan1316/proxyserver/cmd/option"
	"github.com/fyuan1316/proxyserver/pkg/server"

	"github.com/spf13/cobra"
)

var conf option.Configuration

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "proxyserver",
	Long:  `proxyserver`,
	Run: func(cmd *cobra.Command, args []string) {
		s := server.NewMyServer(
			conf.Addr, option.ModelName(conf.ModelName),
			option.TargetURL(conf.TargetURL),
			option.ProxyPrefixURL(conf.ProxyPrefixURLCondition))
		s.Start()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	serverCmd.Flags().StringVarP(&conf.Addr, "address", "", ":9999", "server listen on")
	serverCmd.Flags().StringVarP(&conf.ModelName, "modelname", "", "", "model name")
	serverCmd.Flags().StringVarP(&conf.ProxyPrefixURLCondition, "prefix", "", "/alauda/", "watch for url that has this prefix")
	serverCmd.Flags().StringVarP(&conf.TargetURL, "target", "", "", "proxy to this address")
	serverCmd.MarkFlagRequired("modelname")
}
