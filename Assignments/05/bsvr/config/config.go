package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/addr"
	"github.com/Univ-Wyo-Education/S20-4010/Assignments/05/bsvr/lib"
)

type InitialAccountType struct {
	AcctStr string
	Value   int
	Acct    addr.AddressType `json:"-"` // Converted from AcctStr above with validation.
}

// GlobalConfigData is the gloal configuration data.
// It holds all the data from the cfg.json file.
type GlobalConfigData struct {
	DataDir          string
	MiningReward     int
	MiningDifficulty string // "MiningDifficulty":"0"
	InitialAccounts  []InitialAccountType
	MiningRewardAcct string
	AcctCoinbase     addr.AddressType `json:"-"`
	ControlAcct      string
	AcctPin          map[string]string // map from 0x00000 -> PIN for auth to use.
}

var gCfg GlobalConfigData // global configuration data.

// ReadConfig will read a configuration file into the global congiguration structure.
func ReadConfig(filename string) (err error) {
	// Create Defaults
	gCfg.MiningDifficulty = "0000"

	var buf []byte
	buf, err = ioutil.ReadFile(filename)
	if err != nil {
		// TODO - xyzzy
	}
	err = json.Unmarshal(buf, &gCfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid initialization - Unable to parse JSON file, %s\n", err)
		os.Exit(1)
	}
	for ii, aa := range gCfg.InitialAccounts {
		if !lib.IsValidAddress(aa.AcctStr) {
			fmt.Fprintf(os.Stderr, "Invalid initialization - invalid account at [%d] value [%s] in list of initial account values\n", ii, aa.AcctStr)
			os.Exit(1)
		}
		tmp, _ := lib.ConvAddrStrToAddressType(aa.AcctStr)
		aa.Acct = tmp
		gCfg.InitialAccounts[ii] = aa
		if aa.Value < 0 {
			fmt.Fprintf(os.Stderr, "Invalid initialization - invalid alue at [%d] value [%d] in list of initial account values\n", ii, aa.Value)
			os.Exit(1)
		}
	}
	tmp, err := lib.ConvAddrStrToAddressType(gCfg.MiningRewardAcct)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid initialization - invalid account for MiningRewardAcct value [%s]\n", gCfg.MiningRewardAcct)
		os.Exit(1)
	}
	gCfg.AcctCoinbase = tmp
	return err
}

// GetGlobalConfig returns a copy of the global config structure.
func GetGlobalConfig() (rv GlobalConfigData) {
	return gCfg
}
