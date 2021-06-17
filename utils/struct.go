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

// VulhubAppDBStruct vulhub app database struct
type VulhubAppDBStruct struct {
	Id    int      `json:"id"`
	Cve   string   `json:"cve"`
	Name  string   `json:"name"`
	App   string   `json:"app"`
	Path  string   `json:"path"`
	Files []string `json:"files"`
}

// Config config
type Config struct {
	Mirror string `json:"mirrors"`
}

// DockerComposeServiceStruct docker-compose service struct
type DockerComposeServiceStruct struct {
	Image     string   `yaml:"image"`
	Ports     []string `yaml:"ports"`
	Build     string   `yaml:"build"`
	Networks  []string `yaml:"networks"`
	Volumes   []string `yaml:"volumes"`
	DependsOn []string `yaml:"depends_on"`
	Deploy    struct {
		Mode          string   `yaml:"mode"`
		Replicas      int      `yaml:"replicas"`
		Labels        []string `yaml:"labels"`
		RestartPolicy struct {
			Condition   string `yaml:"condition"`
			Delay       string `yaml:"delay"`
			MaxAttempts int    `yaml:"max_attempts"`
			Window      string `yaml:"window"`
		} `json:"restart_policy"`
		Placement struct {
			Constraints []string `yaml:"constraints"`
		} `json:"placement"`
		UpdateConfig struct {
			Parallelism int    `yaml:"parallelism"`
			Delay       string `yaml:"delay"`
		} `yaml:"update_config"`
	} `yaml:"deploy"`
	StopGracePeriod string `yaml:"stop_grace_period"`
}

// DockerComposeStruct docker-compose struct
type DockerComposeStruct struct {
	Version  string                                `yaml:"version"`
	Services map[string]DockerComposeServiceStruct `yaml:"services"`
}
