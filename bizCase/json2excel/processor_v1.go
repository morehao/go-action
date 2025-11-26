package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// ProcessorV1 处理 v1 版本的 JSON 格式
// 对应文件: es_data_v1.json
type ProcessorV1 struct{}

// ESResponseV1 表示 v1 版本的 Elasticsearch 查询响应结构
type ESResponseV1 struct {
	Hits struct {
		Hits []struct {
			Source struct {
				ID              int64  `json:"id"`
				Type            int    `json:"type"`
				Status          int    `json:"status"`
				Name            string `json:"name"`
				SpaceID         int64  `json:"space_id"`
				OwnerID         int64  `json:"owner_id"`
				CreateTime      int64  `json:"create_time"`
				UpdateTime      int64  `json:"update_time"`
				FavTime         int64  `json:"fav_time"`
				IsFav           int    `json:"is_fav"`
				IsRecentlyOpen  int    `json:"is_recently_open"`
				RecentlyOpenTime int64 `json:"recently_open_time"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func (p *ProcessorV1) Version() string {
	return "v1"
}

func (p *ProcessorV1) Parse(jsonData []byte) ([]DataRow, error) {
	var esResp ESResponseV1
	if err := json.Unmarshal(jsonData, &esResp); err != nil {
		return nil, fmt.Errorf("解析 v1 版本 JSON 失败: %w", err)
	}

	rows := make([]DataRow, 0, len(esResp.Hits.Hits))
	for _, hit := range esResp.Hits.Hits {
		// 格式化时间戳为可读的日期时间格式（毫秒时间戳）
		createTimeStr := formatTimestamp(hit.Source.CreateTime)
		
		rows = append(rows, DataRow{
			CreatedAt: createTimeStr,
			Question:  hit.Source.Name,
			Answer:    fmt.Sprintf("ID: %d | 类型: %d | 状态: %d | 空间ID: %d | 所有者ID: %d", 
				hit.Source.ID, hit.Source.Type, hit.Source.Status, hit.Source.SpaceID, hit.Source.OwnerID),
		})
	}

	return rows, nil
}

// formatTimestamp 将毫秒时间戳格式化为可读的日期时间字符串
func formatTimestamp(ms int64) string {
	if ms == 0 {
		return ""
	}
	// 将毫秒时间戳转换为秒时间戳
	t := time.Unix(ms/1000, (ms%1000)*1000000)
	return t.Format("2006-01-02 15:04:05")
}

