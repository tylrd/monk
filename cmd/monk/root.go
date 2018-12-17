package main

import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tylrd/monk/cmd/monk/version"
	"github.com/tylrd/monk/pkg/config"
	"github.com/tylrd/monk/pkg/util"
	"os"
	"path"
)

var rootCmd = &cobra.Command{
	Use:   "monk",
	Short: "Run and test HTTP collections",
	Long:  `Monk is a CLI for executing, managing, and testing HTTP collections.`,
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.AddCommand(
		version.NewCommand(),
	)
}

func initConfig() {
	configDir := GetConfigDir()
	viper.AddConfigPath(configDir)
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("MONK")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error reading config file: %s \n", err))
	}
}

func GetConfigDir() string {
	baseDir, err := BaseDir()

	if err != nil {
		panic(fmt.Errorf("Fatal error getting home directory: %s \n", err))
	}

	configDir := path.Join(baseDir, ".monk")

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err2 := os.MkdirAll(configDir, 0755)
		if err2 != nil {
			panic(fmt.Errorf("Could not create configuration directory: %s \n", err2))
		}
	}

	writeDefaultConfig(configDir)
	return configDir
}

func BaseDir() (string, error) {
	baseDir, exists := os.LookupEnv("MONK_BASE_DIR")
	var err error

	if !exists {
		baseDir, err = homedir.Dir()
		return baseDir, err
	}

	return baseDir, err
}

func writeDefaultConfig(configDir string) {
	configPath := path.Join(configDir, "config.toml")
	f, err := os.Create(configPath)
	util.CheckError(err)

	c := config.Default()
	var firstBuffer bytes.Buffer
	e := toml.NewEncoder(&firstBuffer)
	err2 := e.Encode(c)
	util.CheckError(err2)

	_, err3 := f.Write(firstBuffer.Bytes())
	util.CheckError(err3)
}
