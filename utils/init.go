/*
Copyright © 2020 iiusky sky@03sec.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path"
)

// InitConfigPath init global config path
func InitConfigPath(configPath string) {
	var dir string
	var err error

	if configPath == "$HOME/.vulhub" {
		dir, err = homedir.Dir()
		if err != nil {
			Error("Get $HOME dir error", zap.Error(err))
			fmt.Println("Get $HOME dir error")
			os.Exit(-1)
		}
		HomeDir = path.Join(dir, ".vulhub")
	} else {
		dir = configPath
		HomeDir = path.Join(dir)
	}

	VulhubConfigDBPath = path.Join(HomeDir, "db")

	err = os.MkdirAll(VulhubConfigDBPath, 0700)
	if err != nil {
		Error(fmt.Sprintf("Mkdir config path %s error", VulhubConfigDBPath), zap.Error(err))
		fmt.Printf("Mkdir config path %s error \n", VulhubConfigDBPath)
		os.Exit(-1)
	}

	VulhubDBName = path.Join(VulhubConfigDBPath, "vulhub.db")
	VulhubConfig = path.Join(HomeDir, "config.json")
}

// InitConfig init config data and set VulhubAppDBUrl
func InitConfig(config Config) {
	configBytes, _ := json.Marshal(config)
	_ = ioutil.WriteFile(VulhubConfig, configBytes, 0655)

	VulhubAppDBUrl = fmt.Sprintf("https://%s/iiiusky/gen-db-tools/main/vulhub.db", config.Mirror)
}

// InitDBFile init db file
func InitDBFile() bool {
	dbBytes, err := ioutil.ReadFile(VulhubDBName)
	if err != nil {
		fmt.Printf("Read db file error %v \r", err)
		Error("Read db file error", zap.Error(err))
		return false
	}
	err = json.Unmarshal(dbBytes, &VulhubDBs)

	if err != nil {
		fmt.Printf("Unmarshal db data error %v \r", err)
		Error("Unmarshal db data error", zap.Error(err))
		return false
	}
	return true
}
