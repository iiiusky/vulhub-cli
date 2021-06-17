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

package core

import (
	"fmt"
	"github.com/alexeyco/simpletable"
	"github.com/briandowns/spinner"
	"github.com/go-resty/resty/v2"
	"github.com/iiiusky/vulhub-cli/utils"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// Deploy Deploy app to dest path
func Deploy(appPath, destPath string) {
	appInfo := GetAppInfoForAppPath(appPath)
	if appInfo != nil {
		AppNormalOutput([]utils.VulhubAppDBStruct{*appInfo})

		if !utils.IsDir(path.Join(destPath, appInfo.Path)) || utils.ForceFlag {
			err := os.MkdirAll(destPath, 0700)
			if err != nil {
				fmt.Println(err)
			}
			download(*appInfo, destPath)
		}

		compose := utils.ParseDockerCompose(path.Join(destPath, appInfo.Path, "docker-compose.yml"))

		commands := utils.GetExecCommands(destPath, *appInfo)

		fmt.Println("Start Commands is: \r\n")
		fmt.Println("cd " + path.Join(destPath, appInfo.Path))

		for _, command := range commands {
			if strings.TrimSpace(command) == "" {
				continue
			}
			fmt.Println(command)
		}
		fmt.Println("\r\nApp Readme Link: https://github.com/vulhub/vulhub/tree/master/" + appInfo.Path)

		displayPortInfo(compose)
	}
}

// GetAppInfoForAppPath get App Info for App Path
func GetAppInfoForAppPath(appPath string) (files *utils.VulhubAppDBStruct) {
	for _, datum := range utils.VulhubDBs {
		if datum.Path == appPath {
			return &datum
		}
	}

	return nil
}

// download app files to destPath
func download(appInfo utils.VulhubAppDBStruct, destPath string) {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()
	s.Suffix = "Download ing.."
	s.Color("green")
	s.FinalMSG = "\r\n Download Over!\n"

	for _, file := range appInfo.Files {
		resp, err := resty.New().R().Get(
			fmt.Sprintf("https://%s/vulhub/vulhub/master/%s%s", utils.GetConfig().Mirror, appInfo.Path, file))
		if err != nil {
			fmt.Printf("\r\n Download %s error: %v ,please try --mirror raw.fastgit.org --force again\n", file, err)
			utils.Error(fmt.Sprintf("Download %s error", file), zap.Error(err))
			os.Exit(-2)
		}

		content := resp.Body()
		mkdirPath := path.Join(destPath, appInfo.Path,
			strings.Join(strings.Split(file, "/")[0:len(strings.Split(file, "/"))-1], "/"))
		err = os.MkdirAll(mkdirPath, 0700)

		if err != nil {
			fmt.Printf("\r\nMkdirAll Path %s error: %v \n", mkdirPath, err)
			utils.Error(fmt.Sprintf("MkdirAll Path %s error", file), zap.Error(err))
			os.Exit(-2)
		}

		err = ioutil.WriteFile(path.Join(destPath, appInfo.Path, file), content, 0666)

		if err != nil {
			fmt.Printf("\r\nWriteFile %s to %s error: %v \n", file, path.Join(destPath, appInfo.Path, file), err)
			utils.Error(fmt.Sprintf("WriteFile %s to %s error:", file, path.Join(destPath, appInfo.Path, file)),
				zap.Error(err))
			continue
		}
	}

	s.Stop()
}

// GetAppInfoForId Get app info for id
func GetAppInfoForId(id int) (info *utils.VulhubAppDBStruct) {
	for _, info := range utils.VulhubDBs {
		if info.Id == id {
			return &info
		}
	}

	return nil
}

// displayPortInfo display port info
func displayPortInfo(compose utils.DockerComposeStruct) {
	fmt.Println("The current docker-compose.yml port information of this app is as follows:")

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Service Name"},
			{Align: simpletable.AlignCenter, Text: "Public Port"},
			{Align: simpletable.AlignCenter, Text: "Internal Port"},
		},
	}
	i := 0

	for serviceName, serviceStruct := range compose.Services {
		for _, port := range serviceStruct.Ports {
			portInfo := strings.Split(port, ":")
			r := []*simpletable.Cell{
				{Align: simpletable.AlignCenter, Text: strconv.Itoa(i)},
				{Align: simpletable.AlignCenter, Text: serviceName},
				{Align: simpletable.AlignCenter, Text: portInfo[0]},
				{Align: simpletable.AlignCenter, Text: portInfo[1]},
			}
			table.Body.Cells = append(table.Body.Cells, r)
			i += 1
		}
	}

	fmt.Println(table.String())
}
