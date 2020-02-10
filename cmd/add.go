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
	"os"
	"os/exec"
	"bufio"
	"strings"

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
 • Connection Name
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
$ go22 add -c myconn01 -n localhost -i 127.0.0.1 -a key -u test`,
	Run: func(cmd *cobra.Command, args []string) {
		// New database connection
		db, err := backend.NewDB(DataFile)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// Create new SSH connection
		newConn := backend.Connection{}

		switch flagSet := cmd.Flags().NFlag(); flagSet {
		case 0:
			promptAdd(&newConn)
		default:
		       	cliAdd(&newConn, cmd)
		}

		// Save new SSH connection
		err = db.AddConn(newConn)
		if err != nil {
			log.Error(err.Error())
			return
		}
		log.Info("Connection \"%s\" successfully created", newConn.ConnName)

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

// Add connection from flags
func cliAdd(c *backend.Connection, cmd *cobra.Command) {
	//Grab connection attribute values from CLI flag input
	c.ConnName, _  = cmd.Flags().GetString("conn-name") 
	c.HostName, _  = cmd.Flags().GetString("host-name")
	c.IPAddress, _ = cmd.Flags().GetString("ip-address")
	c.AuthType, _  = cmd.Flags().GetString("auth-type")
	c.Username, _  = cmd.Flags().GetString("username")
	c.Password, _  = cmd.Flags().GetString("password")
	c.PrivKey, c.PubKey, _ = ssh.GenSSHKeyPair()

}

// Add connection from interactive prompt
func promptAdd(c *backend.Connection) {
	reader := bufio.NewReader(os.Stdin)

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println("New Connection Prompt \n\nEnter the following:")

	fmt.Print("\tConnection Name: ")	
	connName, _ :=  reader.ReadString('\n')
	c.ConnName = strings.TrimSpace(connName)

	fmt.Print("\tHost Name: ")
	hostName, _  := reader.ReadString('\n')
	c.HostName = strings.TrimSpace(hostName)

	fmt.Print("\tIP Address: ")
	ipAddress, _ := reader.ReadString('\n')
	c.IPAddress = strings.TrimSpace(ipAddress)

	fmt.Print("\tAuth Method [key | password]: ")
	authType, _  := reader.ReadString('\n')
	c.AuthType = strings.TrimSpace(authType)

	fmt.Print("\tUsername: ")
	userName, _  := reader.ReadString('\n')
	c.Username = strings.TrimSpace(userName)

	fmt.Print("\tPassword: ")
	password, _  := reader.ReadString('\n')
	c.Password = strings.TrimSpace(password)

	c.PrivKey, c.PubKey, _ = ssh.GenSSHKeyPair()
	fmt.Println()
}


