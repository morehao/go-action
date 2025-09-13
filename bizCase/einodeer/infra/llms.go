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

package infra

import (
	"context"

	"github.com/cloudwego/eino-ext/components/model/deepseek"
	"github.com/morehao/go-action/bizCase/einodeer/config"
)

var (
	ChatModel *deepseek.ChatModel
	PlanModel *deepseek.ChatModel
)

func InitModel() {
	chatModelConfig := &deepseek.ChatModelConfig{
		BaseURL: config.Config.Model.BaseURL,
		APIKey:  config.Config.Model.APIKey,
		Model:   config.Config.Model.DefaultModel,
	}
	ChatModel, _ = deepseek.NewChatModel(context.Background(), chatModelConfig)

	planModelConfig := &deepseek.ChatModelConfig{
		BaseURL: config.Config.Model.BaseURL,
		APIKey:  config.Config.Model.APIKey,
		Model:   config.Config.Model.DefaultModel,
	}
	PlanModel, _ = deepseek.NewChatModel(context.Background(), planModelConfig)
}
