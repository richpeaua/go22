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

	"github.com/richpeaua/go22/pkg/backend"
	"github.com/richpeaua/go22/pkg/log"
	"github.com/richpeaua/go22/pkg/ssh"
	"github.com/spf13/cobra"

)


// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an SSH connection",
	Long: `Save a connection for a target host by defining the following items:
 • Target hostname AND/OR 
 • Target IP address
 • Auth type (key or password)
 • Username
 • Password
 	
Two modes: interactive with prompts or non-interactive using flags

Examples: 

[interactive]
$ go22 add 

[non-interactive]
$ go22 add -c mycomp -n localhost -i 127.0.0.1 -a key -u admin -d $HOME/.ssh/id_rsa.pub`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := backend.NewDB(DataFile)
		if err != nil {
			log.Error(err.Error())
			return
		}
		
		connName, _  := cmd.Flags().GetString("conn-name") 
		hostName, _  := cmd.Flags().GetString("host-name")
		ipAddress, _ := cmd.Flags().GetString("ip-address")
		authType, _  := cmd.Flags().GetString("auth-type")
		userName, _  := cmd.Flags().GetString("username")
		password, _  := cmd.Flags().GetString("password")
		privKey, pubKey, _ := ssh.GenSSHKeyPair()


		connection := backend.Connection{
			ConnName: connName,
			HostName: hostName,
			IPAddress: ipAddress,
			AuthType: authType,
			Username: userName,
			Password: password,
			PrivKey: privKey,
			PubKey: pubKey,
		}
		// fmt.Println(connection, db)
		err = db.AddConn(connection)
		if err != nil {
			log.Error(err.Error())
			return
		}
		log.Info("Connection \"%s\" successfully created", connName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("conn-name", "c", "", "connection name \t\t*required" )
	addCmd.Flags().StringP("host-name", "n", "", "hostname of target machine \t^optional if ip address is set" )
	addCmd.Flags().StringP("ip-address", "i", "", "ip of target machine \t^optional if hostname is set" )
	addCmd.Flags().StringP("auth-type", "a", "password", "ssh authentication type \t[credentials | key] *required" )
	addCmd.Flags().StringP("username", "u", "", "connection username \t*required")
	addCmd.Flags().StringP("password", "p", "", "connection password \t*required if --auth-type=password")
	addCmd.Flags().StringP("key", "k", "", "connection key path \t^optional")
}

