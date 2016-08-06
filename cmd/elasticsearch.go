// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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

	elastic "gopkg.in/olivere/elastic.v3"

	"github.com/Kyperion/baldr/try"
	"github.com/spf13/cobra"
)

var elasticsearch string

// elasticsearchCmd represents the elasticsearch command
var elasticsearchCmd = &cobra.Command{
	Use:   "elasticsearch",
	Short: "waits untill a elasticsearch server is ready to accept connections",
	Long: `this command will try to connect to a elasticsearch instance and retry a few times untill it succeeds or bails out. For example:

baldr elasticsearch -s http://localhost:9200/`,

	PreRun: func(cmd *cobra.Command, args []string) {
		if elasticsearch == "" {
			log.Fatalln("elasticsearch connection string is empty.")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Connecting to ", elasticsearch)

		err := try.Do(func(attempt int) (bool, error) {
			var err error
			_, err = elastic.NewClient(
				elastic.SetURL(elasticsearch),
				elastic.SetMaxRetries(3),
			)

			return attempt < retry, err
		}, wait)

		if err != nil {
			log.Fatalln("error:", err)
		}
	},
}

func init() {
	RootCmd.AddCommand(elasticsearchCmd)
	elasticsearchCmd.Flags().StringVarP(&elasticsearch, "elasticsearch", "s", "", "elasticsearch instances to connect to.")
}
