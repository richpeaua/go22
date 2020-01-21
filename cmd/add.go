/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

	"github.com/richpeaua/go22/pkg/backend"
	"github.com/spf13/cobra"

)


// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an SSH connection",
	Long: `Add the connection information for a target host, such as:
 • Hostname AND/OR Host IP Address
 • Auth type (key or credentials)
 • Auth payload (key/key path/username & password)
 • Metadata (host OS type, group label, server/device type [router, database, webserver])	
two modes: interactive with prompts or non-interactive using flags

examples: 

[interactive]
$ go22 add 

[non-interactive]
$ go add -c mycomp -n localhost -i 127.0.0.1 -a key -d $HOME/.ssh/id_rsa.pub`,
	Run: func(cmd *cobra.Command, args []string) {

		db, err := backend.NewDB(DataFile)
		if err != nil {
			panic(err)
		}
		connName, _  := cmd.Flags().GetString("conn-name") 
		hostName, _  := cmd.Flags().GetString("host-name")
		ipAddress, _ := cmd.Flags().GetString("ip-address")
		authType, _  := cmd.Flags().GetString("auth-type")
		userName, _  := cmd.Flags().GetString("username")
		password, _  := cmd.Flags().GetString("password")
		privkey, _ 	 := cmd.Flags().GetString("key")
		pubkey  	 := "bca"


		connection := backend.Connection{
			ConnName: connName,
			HostName: hostName,
			IPAddress: ipAddress,
			AuthType: authType,
			Username: userName,
			Password: password,
			PrivKey: privkey,
			PubKey: pubkey,
		}
		fmt.Println(connection, db)
		// err = db.AddConn(connection)
		// if err != nil {
		// 	panic(err)
		// }
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("conn-name", "c", "", "connection name \t\t*required" )
	addCmd.Flags().StringP("host-name", "n", "", "hostname of target machine \t*optional if ip address is set" )
	addCmd.Flags().StringP("ip-address", "i", "", "ip of target machine \t*optional if hostname is set" )
	addCmd.Flags().StringP("auth-type", "a", "password", "ssh authentication type \t[credentials | key] *required" )
	addCmd.Flags().StringP("username", "u", "", "connection username \t\t*required")
	addCmd.Flags().StringP("password", "p", "", "connection password \t\t*required if --auth-type=password")
	addCmd.Flags().StringP("key", "k", "", "connection key path \t\t *optional")
}

