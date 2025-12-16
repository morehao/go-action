package renderer

import "github.com/morehao/go-action/bizCase/llmproxy/types"

// Renderer 定义工具定义渲染器接口
type Renderer interface {
	// RenderTools 将工具定义转换为提示词并添加到请求中
	RenderTools(req *types.ChatRequest) *types.ChatRequest
}

// NewRenderer 根据格式创建渲染器
func NewRenderer(format string) Renderer {
	switch format {
	case "qwen":
		return &QwenRenderer{}
	case "llama":
		return &LlamaRenderer{}
	default:
		return &GenericRenderer{}
	}
}
