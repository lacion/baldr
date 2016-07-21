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
	"strings"
	"time"

	"github.com/Kyperion/baldr/try"
	"github.com/coreos/etcd/clientv3"
	"github.com/spf13/cobra"
)

var endpoints string

// etcd3Cmd represents the etcd3 command
var etcd3Cmd = &cobra.Command{
	Use:   "etcd3",
	Short: "waits untill an etcd3 is ready to accept connections",
	Long: `this command will try to connect to a etcd3 instance and retry a few times untill it succeeds or bails out. For example:

baldr etcd3 -e etcd1:2379,etcd2:2379,etcd3:2379`,

	PreRun: func(cmd *cobra.Command, args []string) {
		if endpoints == "" {
			log.Fatalln("etcd3 connection string is empty.")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Connecting to", endpoints)

		err := try.Do(func(attempt int) (bool, error) {
			var err error
			cli, err := clientv3.New(clientv3.Config{
				Endpoints:   strings.Split(endpoints, ","),
				DialTimeout: 5 * time.Second,
			})
			if err == nil {
				cli.Close()
			}
			return attempt < retry, err
		}, wait)

		if err != nil {
			log.Fatalln("error:", err)
		}

	},
}

func init() {
	RootCmd.AddCommand(etcd3Cmd)

	etcd3Cmd.Flags().StringVarP(&endpoints, "endpoints", "e", "", "etcd3 instances to connect to")
}
