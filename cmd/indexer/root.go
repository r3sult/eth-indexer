// Copyright © 2018 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/getamis/sirius/log"
	"github.com/maichain/eth-indexer/indexer"
	manager "github.com/maichain/eth-indexer/store/store_manager"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "indexer",
	Short: "blockchain data indexer",
	Long:  `blockchain data indexer`,
	Run: func(cmd *cobra.Command, args []string) {
		vp := viper.New()
		vp.BindPFlags(cmd.Flags())
		vp.AutomaticEnv() // read in environment variables that match
		if configFile := vp.GetString(configFileFlag); configFile != "" {
			if err := loadConfigUsingViper(vp, configFile); err != nil {
				log.Error("Failed to load config file", "err", err)
				return
			}
			loadFlagToVar(vp)
		}

		// eth-client
		ethClient := MustEthConn(fmt.Sprintf("%s://%s:%d", ethProtocol, ethHost, ethPort))
		// log.Info("eth=client" + ethClient)

		// database
		db := MustNewDatabase()
		defer db.Close()

		store := manager.NewStoreManager(db)
		indexer := indexer.NewIndexer(ethClient, store)

		if listen {
			ch := make(chan *types.Header)
			indexer.Listen(context.Background(), ch)
		} else {
			indexer.Start(start, end)
		}

		return
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	// indexer flags
	RootCmd.Flags().Int64Var(&start, startFlag, 0, "The start block height")
	RootCmd.Flags().Int64Var(&end, endFlag, 0, "The end block height")
	RootCmd.Flags().BoolVar(&listen, "listen", false, "listen mode to recent block")

	// eth-client flags
	RootCmd.Flags().StringVar(&ethProtocol, ethProtocolFlag, "ws", "The eth-client protocol")
	RootCmd.Flags().StringVar(&ethHost, ethHostFlag, "127.0.0.1", "The eth-client host")
	RootCmd.Flags().IntVar(&ethPort, ethPortFlag, 8546, "The eth-client port")

	// Database flags
	RootCmd.Flags().StringVar(&dbDriver, dbDriverFlag, "mysql", "The database driver")
	RootCmd.Flags().StringVar(&dbHost, dbHostFlag, "", "The database host")
	RootCmd.Flags().IntVar(&dbPort, dbPortFlag, 3306, "The database port")
	RootCmd.Flags().StringVar(&dbName, dbNameFlag, "eth-db", "The database name")
	RootCmd.Flags().StringVar(&dbUser, dbUserFlag, "root", "The database username to login")
	RootCmd.Flags().StringVar(&dbPassword, dbPasswordFlag, "my-secret-pw", "The database password to login")
}