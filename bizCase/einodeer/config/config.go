/*
 * Copyright 2025 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"context"

	"github.com/morehao/golib/conf"
	"github.com/morehao/golib/glog"
)

// 定义一个结构体来解析 YAML 文件中的 mcp.servers 部分
type DeerConfig struct {
	MCP     MCPConfig     `yaml:"mcp"`
	Model   ModelConfig   `yaml:"model"`
	Setting SettingConfig `yaml:"setting"`
}

type ModelConfig struct {
	DefaultModel string `yaml:"default_model"`
	APIKey       string `yaml:"api_key"`
	BaseURL      string `yaml:"base_url"`
}

type SettingConfig struct {
	MaxPlanIterations int `yaml:"max_plan_iterations"`
	MaxStepNum        int `yaml:"max_step_num"`
}

type MCPConfig struct {
	Servers map[string]struct {
		Command string            `yaml:"command"`
		Args    []string          `yaml:"args"`
		Env     map[string]string `yaml:"env,omitempty"`
	} `yaml:"servers"`
}

var (
	Config *DeerConfig = &DeerConfig{}
)

func LoadDeerConfig() {

	configPath := conf.GetAppRootDir() + "/config/config.yaml"

	// 读取 YAML 文件内容
	var deerConfig DeerConfig
	conf.LoadConfig(configPath, &deerConfig)

	glog.Infof(context.Background(), "load_config: %s", deerConfig)

	Config = &deerConfig
}
