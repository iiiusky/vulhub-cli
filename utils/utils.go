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
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

// FileExists Determine whether the given path file exists
func FileExists(path string) bool {
	x, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) && !x.IsDir() {
			return true
		}
		return false
	}
	return true
}

// IsDir Determine whether the given path is a folder
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// GetExecCommands get app start commands
func GetExecCommands(destPath string, appInfo VulhubAppDBStruct) []string {
	var commands []string
	var result []string
	readmeByte, err := ioutil.ReadFile(path.Join(destPath, appInfo.Path, "README.md"))

	if err != nil {
		readmeByte, err = ioutil.ReadFile(path.Join(destPath, appInfo.Path, "README.zh-cn.md"))
		if err != nil {
			fmt.Printf("Open %s README.md file faild.\n", appInfo.Path)
			Error(fmt.Sprintf("Open %s README.md file faild.", appInfo.Path), zap.Error(err))
			os.Exit(-2)
		}
	}

	reg1 := regexp.MustCompile("```\\s?(.+\\s)+```")
	result1 := reg1.FindAllString(string(readmeByte), -1)

	for _, result := range result1 {
		result = strings.Trim(result, "```")
		tmpCommands := strings.Split(result, "\n")

		if strings.HasPrefix(tmpCommands[1], "docker-compose") {
			commands = append(commands, tmpCommands...)
		}
	}

	for _, command := range commands {
		if command != "" {
			result = append(result, command)
		}
	}

	if len(commands) == 0 {
		fmt.Printf("Did not get any startup commands, please check the README.md file \n")
		Error(fmt.Sprintf("Did not get any startup commands, please check the README.md file"))
		os.Exit(-2)
	}

	return result
}

// ParseDockerCompose Parse docker-compose.yml to struct
func ParseDockerCompose(filePath string) DockerComposeStruct {
	var compose DockerComposeStruct

	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Open docker-compose.yml file fail.")
		Error("Open docker-compose.yml file fail.", zap.Error(err))
		return compose
	}

	if err = yaml.Unmarshal(bs, &compose); err != nil {
		fmt.Println("Unmarshal docker-compose.yml file fail.")
		Error("Unmarshal docker-compose.yml file fail.", zap.Error(err))
		return compose
	}

	return compose
}

// GetConfig read config file and Unmarshal config data
func GetConfig() (config Config) {
	configBytes, _ := ioutil.ReadFile(VulhubConfig)
	err := json.Unmarshal(configBytes, &config)

	if err != nil {
		Error("Unmarshal Config error", zap.Error(err))
		fmt.Printf("Unmarshal Config  error \n")
		os.Exit(-2)
	}

	return config
}
