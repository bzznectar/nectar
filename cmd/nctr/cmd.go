// Copyright 2020 The Swarm Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nctr

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethersphere/bee/pkg/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	optionNameDataDir      = "data-dir"
	optionNameBootnodes    = "bootnode"
	optionNamePassword     = "password"
	optionNamePasswordFile = "password-file"
	//optionNameAPIAddr      = "api-addr"
	optionNameMainNet         = "mainnet"
	optionNameDebugAPIAddr    = "debug-api-addr"
	optionNameBeeDebugAPIAddr = "bee-debug-api"

	optionNameNetworkID = "network-id"

	optionNameVerbosity = "verbosity"

	optionNameSwapEndpoint = "swap-endpoint"
	optionNameSwapInitialDeposit="swap-initial-deposit"
	optionNameHarvestGas="harvest-gas"
)

func init() {
	cobra.EnableCommandSorting = false
}

type command struct {
	root           *cobra.Command
	config         *viper.Viper
	passwordReader passwordReader
	cfgFile        string
	homeDir        string
}

type option func(*command)

func newCommand(opts ...option) (c *command, err error) {
	c = &command{
		root: &cobra.Command{
			Use:           "nctr",
			Short:         "nctr client",
			SilenceErrors: true,
			SilenceUsage:  true,
			PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
				return c.initConfig()
			},
		},
	}

	for _, o := range opts {
		o(c)
	}
	if c.passwordReader == nil {
		c.passwordReader = new(stdInPasswordReader)
	}

	// Find home directory.
	if err := c.setHomeDir(); err != nil {
		return nil, err
	}

	c.initGlobalFlags()

	if err := c.initStartCmd(); err != nil {
		return nil, err
	}

	// if err := c.initInitCmd(); err != nil {
	// 	return nil, err
	// }

	// if err := c.initDeployCmd(); err != nil {
	// 	return nil, err
	// }

	c.initVersionCmd()

	if err := c.initConfigurateOptionsCmd(); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *command) Execute() (err error) {
	return c.root.Execute()
}

// Execute parses command line arguments and runs appropriate functions.
func Execute() (err error) {
	c, err := newCommand()
	if err != nil {
		return err
	}
	return c.Execute()
}

func (c *command) initGlobalFlags() {
	globalFlags := c.root.PersistentFlags()
	globalFlags.StringVar(&c.cfgFile, "config", "", "config file (default is $HOME/.nctr.yaml)")
}

func (c *command) initConfig() (err error) {
	config := viper.New()
	configName := ".nctr"
	if c.cfgFile != "" {
		// Use config file from the flag.
		config.SetConfigFile(c.cfgFile)
	} else {
		// Search config in home directory with name ".bee" (without extension).
		config.AddConfigPath(c.homeDir)
		config.SetConfigName(configName)
	}

	// Environment
	config.SetEnvPrefix("nctr")
	config.AutomaticEnv() // read in environment variables that match
	config.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	if c.homeDir != "" && c.cfgFile == "" {
		c.cfgFile = filepath.Join(c.homeDir, configName+".yaml")
	}

	// If a config file is found, read it in.
	if err := config.ReadInConfig(); err != nil {
		var e viper.ConfigFileNotFoundError
		if !errors.As(err, &e) {
			return err
		}
	}
	c.config = config
	return nil
}

func (c *command) setHomeDir() (err error) {
	if c.homeDir != "" {
		return
	}
	dir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	c.homeDir = dir
	return nil
}

func (c *command) setAllFlags(cmd *cobra.Command) {
	cmd.Flags().String(optionNameDataDir, filepath.Join(c.homeDir, ".nctr"), "data directory")
	cmd.Flags().String(optionNamePassword, "", "password for decrypting keys")
	// cmd.Flags().String(optionNamePasswordFile, "", "path to a file that contains password for decrypting keys")
	//cmd.Flags().String(optionNameAPIAddr, ":1633", "Bee HTTP API listen address")
	cmd.Flags().String(optionNameBeeDebugAPIAddr, ":1635", "Bee DEBUG API listen address")
	cmd.Flags().String(optionNameDebugAPIAddr, ":9000", "HCTR HTTP API listen address")
	// cmd.Flags().String(optionNameBeeDebugAPIAddr, ":1635", "debug HTTP API listen address")
	cmd.Flags().Uint64(optionNameNetworkID, 10, "ID of the nctr network")
	// cmd.Flags().String(optionNameVerbosity, "info", "log verbosity level 0=silent, 1=error, 2=warn, 3=info, 4=debug, 5=trace")
	cmd.Flags().String(optionNameSwapEndpoint, "ws://localhost:8546", "swap ethereum blockchain endpoint")
	cmd.Flags().Uint64(optionNameSwapInitialDeposit,7000000,"swap-initial-deposit")
	cmd.Flags().Uint64(optionNameHarvestGas,7000000,"harvest gas")
	cmd.Flags().Bool(optionNameMainNet, false, "Is it the MainNet")
	//cmd.Flags().String(optionNameBootnodes,"","initial nodes to connect to")
}

func newLogger(cmd *cobra.Command, verbosity string) (logging.Logger, error) {
	var logger logging.Logger
	switch verbosity {
	case "0", "silent":
		logger = logging.New(ioutil.Discard, 0)
	case "1", "error":
		logger = logging.New(cmd.OutOrStdout(), logrus.ErrorLevel)
	case "2", "warn":
		logger = logging.New(cmd.OutOrStdout(), logrus.WarnLevel)
	case "3", "info":
		logger = logging.New(cmd.OutOrStdout(), logrus.InfoLevel)
	case "4", "debug":
		logger = logging.New(cmd.OutOrStdout(), logrus.DebugLevel)
	case "5", "trace":
		logger = logging.New(cmd.OutOrStdout(), logrus.TraceLevel)
	default:
		// return nil, fmt.Errorf("unknown verbosity level %q", verbosity)
		logger = logging.New(cmd.OutOrStdout(), logrus.InfoLevel)
	}
	return logger, nil
}
