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

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize go22 program ",
	Long: `creates the ".go22/" and ".go22/data" directories in your home directory
and creates the "connections.db" sqlite3 database with initalized "connections" table`,
	Run: func(cmd *cobra.Command, args []string) {
		createAppDir(AppDir, AppDataDir)
		fmt.Println("initialized; ready to save connections")

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

var (
	home 		string		= os.Getenv("HOME")
	AppDir		string 		= home + "/.go22"
	AppDataDir 	string  	= AppDir + "/data"
	DataFile	string		= AppDataDir + "/" + "connections.db"

)

// Creates go22 user application directory in user home dir
func createAppDir(appDir string, appDataDir string) {

	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		os.Mkdir(appDir, 0750)
		os.Mkdir(appDataDir, 0750)
		fmt.Println("go22 user app directory created.")
	} else {
		fmt.Println("go22 user app directory already exists")
	}
}
