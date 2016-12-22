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
	"crypto/tls"
	"log"
	"net"

	"gopkg.in/mgo.v2"

	"github.com/Kyperion/baldr/try"
	"github.com/spf13/cobra"
)

var mongoDB string

// mongodbCmd represents the mongodb command
var mongodbCmd = &cobra.Command{
	Use:   "mongodb",
	Short: "waits untill a mongodb server is ready to accept connections",
	Long: `this command will try to connect to a mongodb instance and retry a few times untill it succeeds or bails out. For example:

baldr mongodb -m mongodb://user:pass@db1:10013,db2:10014/auth?ssl=true`,

	PreRun: func(cmd *cobra.Command, args []string) {
		if mongoDB == "" {
			log.Fatalln("mongodb connection string is empty.")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		tlsConfig := &tls.Config{}
		tlsConfig.InsecureSkipVerify = true

		log.Println("Connecting to ", mongoDB)

		err := try.Do(func(attempt int) (bool, error) {
			var err error
			dialInfo, err := mgo.ParseURL(mongoDB)
			if err != nil {
				log.Println("Connection error: ", err.Error())
				return attempt < retry, err
			}
			dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
				conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
				return conn, err
			}
			session, err := mgo.DialWithInfo(dialInfo)
			if err == nil {
				session.Close()
			} else {
				log.Println("Connection error: ", err.Error())
			}

			return attempt < retry, err
		}, wait)

		if err != nil {
			log.Fatalln("error:", err)
		} else {
			log.Println("Connected to mongo successful")
		}
	},
}

func init() {
	RootCmd.AddCommand(mongodbCmd)
	mongodbCmd.Flags().StringVarP(&mongoDB, "mongodb", "m", "", "mongo instances to connect to.")
}
