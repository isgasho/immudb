/*
Copyright 2019 vChain, Inc.

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

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/codenotary/immudb/pkg/client"
	"github.com/codenotary/immudb/pkg/server"
)

func main() {
	cmd := &cobra.Command{
		Use: "immu",
	}
	getCommand := &cobra.Command{
		Use:     "get",
		Aliases: []string{"g"},
		RunE: func(cmd *cobra.Command, args []string) error {
			options, err := options(cmd)
			if err != nil {
				return err
			}
			key := args[0]
			response, err := client.Get(options, key)
			if err != nil {
				return err
			}
			fmt.Println("Get", key, "=", string(response))
			return nil
		},
		Args: cobra.ExactArgs(1),
	}
	setCommand := &cobra.Command{
		Use:     "set",
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			options, err := options(cmd)
			if err != nil {
				return err
			}
			key, value := args[0], args[1]
			if err := client.Set(options, key, []byte(value)); err != nil {
				return err
			}
			fmt.Println("Set", key, "=", value)
			return nil
		},
		Args: cobra.ExactArgs(2),
	}
	configureOptions(getCommand)
	configureOptions(setCommand)
	cmd.AddCommand(getCommand)
	cmd.AddCommand(setCommand)
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func configureOptions(cmd *cobra.Command) {
	cmd.Flags().IntP("port", "p", server.DefaultOptions().Port, "port number")
	cmd.Flags().StringP("address", "a", server.DefaultOptions().Address, "bind address")
}

func options(cmd *cobra.Command) (*client.Options, error) {
	port, err := cmd.Flags().GetInt("port")
	if err != nil {
		return nil, err
	}
	address, err := cmd.Flags().GetString("address")
	if err != nil {
		return nil, err
	}
	options := client.DefaultOptions().
		WithPort(port).
		WithAddress(address)
	return &options, nil
}