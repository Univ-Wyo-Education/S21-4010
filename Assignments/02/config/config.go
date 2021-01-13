package config

import (
	"encoding/json"
	"io/ioutil"
)

// GlobalConfigData is the gloal configuration data.
// It holds all the data from the cfg.json file.
type GlobalConfigData struct {
	DataDir          string
	MiningDifficulty string
}

var gCfg GlobalConfigData // global configuration data.

// ReadConfig will read a configuration file into the global congiguration structure.
func ReadConfig(filename string) (err error) {
	// Create Defaults
	gCfg.MiningDifficulty = "0000"

	var buf []byte
	buf, err = ioutil.ReadFile(filename)
	if err != nil {
	}
	err = json.Unmarshal(buf, &gCfg)
	if err != nil {
	}
	return err
}

// GetGlobalConfig returns a copy of the global config structure.
func GetGlobalConfig() (rv GlobalConfigData) {
	return gCfg
}
