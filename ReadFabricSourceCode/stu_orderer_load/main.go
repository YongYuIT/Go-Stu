package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const Prefix = "ORDERER"

func main() {
	config := viper.New()
	InitViper(config, "orderer")
	config.SetEnvPrefix(Prefix)
	config.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	config.SetEnvKeyReplacer(replacer)

	if err := config.ReadInConfig(); err != nil {
		_ = fmt.Errorf("Error reading configuration: %s", err)
	}

	fmt.Println(config.ConfigFileUsed())
	fmt.Println("success")
}

const OfficialPath = "/home/yong/Go-Stu20191008001/ReadFabricSourceCode/stu_orderer_load/tmp_fabric"

func InitViper(v *viper.Viper, configName string) error {
	var altPath = os.Getenv("FABRIC_CFG_PATH")
	//参照/mnt/hgfs/fabric-env/fabric-samples/first-network/byfn.sh
	altPath = "/home/yong/Go-Stu20191008001/ReadFabricSourceCode/stu_orderer_load/tmp_fabric"
	if altPath != "" {
		// If the user has overridden the path with an envvar, its the only path
		// we will consider

		if !dirExists(altPath) {
			return fmt.Errorf("FABRIC_CFG_PATH %s does not exist", altPath)
		}

		AddConfigPath(v, altPath)
	} else {
		// If we get here, we should use the default paths in priority order:
		//
		// *) CWD
		// *) /etc/hyperledger/fabric

		// CWD
		AddConfigPath(v, "./")

		// And finally, the official path
		if dirExists(OfficialPath) {
			AddConfigPath(v, OfficialPath)
		}
	}

	// Now set the configuration file.
	if v != nil {
		v.SetConfigName(configName)
	} else {
		viper.SetConfigName(configName)
	}

	return nil
}

func dirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func AddConfigPath(v *viper.Viper, p string) {
	if v != nil {
		v.AddConfigPath(p)
	} else {
		viper.AddConfigPath(p)
	}
}
