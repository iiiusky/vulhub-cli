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
	"github.com/briandowns/spinner"
	"github.com/go-resty/resty/v2"
	"github.com/iiiusky/vulhub-cli/utils"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"time"
)

// CacheDB Cache vulhub data information to the local machine.
func CacheDB() {
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Start()
	s.Suffix = fmt.Sprintf("Download Db file for %s..", utils.GetConfig().Mirror)
	s.Color("green")
	s.FinalMSG = "\r\n Cache Db file Over! \r\n"

	vulhubDBResp, err := resty.New().R().Get(utils.VulhubAppDBUrl)

	if err != nil {
		fmt.Printf("\r\nGet vulhub db data error,try --mirror raw.fastgit.org again\n")
		utils.Error("Get vulhub db data error", zap.Error(err))
		os.Exit(-2)
	}

	err = ioutil.WriteFile(utils.VulhubDBName, vulhubDBResp.Body(), 0666)

	if err != nil {
		fmt.Printf("\r\nWriteFile vulhub meta data error \n")
		utils.Error("WriteFile vulhub meta data error", zap.Error(err))
		os.Exit(-2)
	}

	s.Stop()
}
