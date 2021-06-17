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
	"github.com/iiiusky/vulhub-cli/utils"
	"strconv"
	"strings"
)

// GetApps Get all supported app lists and show them
func GetApps(filter string) {
	if filter == "" {
		AppNormalOutput(utils.VulhubDBs)
	} else {
		var filterData []utils.VulhubAppDBStruct
		for _, app := range utils.VulhubDBs {
			if filter != "" {
				if strings.Contains(strings.Replace(strings.ToLower(app.Path), " ", "", -1), strings.ToLower(filter)) {
					filterData = append(filterData, app)
				}
			} else {
				filterData = append(filterData, app)
			}
		}

		AppNormalOutput(filterData)
	}

}

// AppNormalOutput Output table format
func AppNormalOutput(apps []utils.VulhubAppDBStruct) {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "ID"},
			{Align: simpletable.AlignCenter, Text: "APP"},
			{Align: simpletable.AlignCenter, Text: "Name"},
			{Align: simpletable.AlignCenter, Text: "CVE"},
			{Align: simpletable.AlignCenter, Text: "Path"},
		},
	}

	for _, app := range apps {
		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: strconv.Itoa(app.Id)},
			{Align: simpletable.AlignCenter, Text: app.App},
			{Align: simpletable.AlignCenter, Text: app.Name},
			{Align: simpletable.AlignCenter, Text: app.Cve},
			{Align: simpletable.AlignCenter, Text: app.Path},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}

	fmt.Println(table.String())
}
