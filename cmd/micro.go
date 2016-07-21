// Copyright © 2016 Kyperion S.L. <info@kyperion.com>
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
	"log"

	"github.com/spf13/cobra"
)

var microService string

// microCmd represents the micro command
var microCmd = &cobra.Command{
	Use:   "micro",
	Short: "waits untill a micro service is registered",
	Long: `this command will check micro registry and look for a spesific micro service. For example:

baldr micro -s foo.bar`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if microService == "" {
			log.Fatalln("Micro service name is empty.")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("micro called")
	},
}

func init() {
	RootCmd.AddCommand(microCmd)
	microCmd.Flags().StringVarP(&microService, "servicename", "s", "", "Micro service name to watch for.")
}
