package parser

// InternalToolCall 表示内部使用的工具调用（解析后）
type InternalToolCall struct {
	Name      string                 `json:"name"`
	Arguments map[string]interface{} `json:"arguments"`
}

// Parser 定义工具调用解析器接口
type Parser interface {
	// Parse 从内容中解析工具调用
	// 返回: 工具调用列表, 剩余内容, 错误
	Parse(content string) ([]InternalToolCall, string, error)
}

// StreamParser 定义流式解析器接口
type StreamParser interface {
	// Add 添加内容块并尝试解析
	// 返回: 工具调用列表, 可输出的内容, 是否完成
	Add(chunk string) ([]InternalToolCall, string, bool)

	// Flush 刷新缓冲区，返回所有剩余内容
	Flush() ([]InternalToolCall, string)
}

// NewParser 根据格式创建解析器
func NewParser(format string) Parser {
	switch format {
	case "json":
		return NewJSONParser()
	default:
		return NewXMLParser()
	}
}

// NewStreamParser 根据格式创建流式解析器
func NewStreamParser(format string) StreamParser {
	switch format {
	case "json":
		return NewStreamJSONParser()
	default:
		return NewStreamXMLParser()
	}
}
