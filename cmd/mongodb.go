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
	"log"
	"os"

	"gopkg.in/mgo.v2"

	"github.com/Kyperion/baldr/try"
	"github.com/spf13/cobra"
)

var mongoDB string
var wait int
var retry int

// mongodbCmd represents the mongodb command
var mongodbCmd = &cobra.Command{
	Use:   "mongodb",
	Short: "waits untill a mongodb is ready to accept connections",
	Long: `this command will try to connect to a mongodb instance and retry a few times untill it succeeds or bails out. For example:

baldr mongodb -m mongodb://user:pass@db1:10013,db2:10014/auth?ssl=true`,

	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Connecting to ", mongoDB)

		err := try.Do(func(attempt int) (bool, error) {
			var err error
			session, err := mgo.Dial(mongoDB)

			if err == nil {
				session.Close()
			}
			return attempt < retry, err
		}, wait)

		if err != nil {
			log.Fatalln("error:", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(mongodbCmd)
	mongodbCmd.Flags().StringVarP(&mongoDB, "mongodb", "m", os.Getenv("MONGODB"), "mongo instances to connect to.")
	mongodbCmd.Flags().IntVarP(&wait, "timeout", "t", 5000, "Timeout in ms to wait before retrys.")
	mongodbCmd.Flags().IntVarP(&retry, "retry", "r", 5, "number of times to retry before bailing out.")
}
