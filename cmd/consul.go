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

	"github.com/Kyperion/baldr/try"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/cobra"
)

var consul string

// consulCmd represents the consul command
var consulCmd = &cobra.Command{
	Use:   "consul",
	Short: "waits untill an consul is ready to accept connections",
	Long: `this command will try to connect to a consul instance and retry a few times untill it succeeds or bails out. For example:

baldr consul -e consul:8500`,

	PreRun: func(cmd *cobra.Command, args []string) {
		if consul == "" {
			log.Fatalln("consul connection string is empty.")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Connecting to", consul)

		err := try.Do(func(attempt int) (bool, error) {
			var err error

			config := api.DefaultConfig()
			config.Address = consul

			client, err := api.NewClient(config)

			kv := client.KV()
			p := &api.KVPair{Key: "foo", Value: []byte("test")}
			_, err = kv.Put(p, nil)

			if err != nil {
				log.Println("error connecting:", err)
				return attempt < retry, err
			}

			if err == nil {
				log.Println("connected to consul")
			}
			return attempt < retry, err
		}, wait)

		if err != nil {
			log.Fatalln("error:", err)
		}

	},
}

func init() {
	RootCmd.AddCommand(consulCmd)

	consulCmd.Flags().StringVarP(&consul, "consul", "c", "", "consul instance to connect to")
}
