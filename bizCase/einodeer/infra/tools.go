package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"os/exec"
	"strings"
	"time"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
	"github.com/morehao/go-action/bizCase/einodeer/config"
	"github.com/morehao/golib/glog"
)

func init() {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())
}

// 模拟搜索结果结构
type MockSearchResult struct {
	Results []SearchResult `json:"results"`
	Query   string         `json:"query"`
}

// 搜索结果项
type SearchResult struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Published   string `json:"published"`
}

// 模拟爬取结果结构
type MockCrawlResult struct {
	URL     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Links   []Link `json:"links"`
}

// 链接结构
type Link struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

// generateMockSearchResults 生成模拟搜索结果
func generateMockSearchResults(query string) MockSearchResult {
	// 随机生成2-5个搜索结果
	resultCount := rand.Intn(4) + 2
	results := make([]SearchResult, resultCount)

	// 常见的域名列表
	domains := []string{"example.com", "mockdata.org", "testsite.net", "demoinfo.io", "sampledata.com"}

	// 生成随机的搜索结果
	for i := 0; i < resultCount; i++ {
		// 随机选择一个域名
		domain := domains[rand.Intn(len(domains))]

		// 生成随机的标题
		title := fmt.Sprintf("关于 %s 的信息 #%d", query, i+1)

		// 生成随机的URL
		url := fmt.Sprintf("https://%s/article-%d-%s", domain, rand.Intn(1000), strings.ReplaceAll(query, " ", "-"))

		// 生成随机的描述
		description := fmt.Sprintf("这是关于 %s 的详细信息。包含了相关的背景、历史和最新发展。这是一个模拟的搜索结果，用于测试目的。", query)

		// 生成随机的发布日期（过去一年内）
		daysAgo := rand.Intn(365)
		publishDate := time.Now().AddDate(0, 0, -daysAgo).Format("2006-01-02")

		results[i] = SearchResult{
			Title:       title,
			URL:         url,
			Description: description,
			Published:   publishDate,
		}
	}

	return MockSearchResult{
		Results: results,
		Query:   query,
	}
}

// generateMockCrawlResults 生成模拟爬取结果
func generateMockCrawlResults(url string) MockCrawlResult {
	// 从URL中提取可能的主题
	parts := strings.Split(url, "/")
	lastPart := ""
	if len(parts) > 0 {
		lastPart = parts[len(parts)-1]
	}
	if lastPart == "" {
		lastPart = "homepage"
	}

	// 生成随机的标题
	title := fmt.Sprintf("%s - 网页内容", strings.ReplaceAll(lastPart, "-", " "))

	// 生成随机的内容段落
	paragraphs := []string{
		"这是一个模拟的网页内容，用于测试目的。",
		"该页面包含了关于特定主题的详细信息和数据。",
		"以下是一些相关的要点和分析：",
		"1. 这是第一个要点，包含了基本信息。",
		"2. 这是第二个要点，提供了更多细节。",
		"3. 这是第三个要点，包含了一些结论。",
		"总结：这是一个模拟的网页爬取结果，仅用于测试和开发目的。",
	}

	// 将段落组合成内容
	content := strings.Join(paragraphs, "\n\n")

	// 生成随机的链接（1-3个）
	linkCount := rand.Intn(3) + 1
	links := make([]Link, linkCount)

	// 常见的域名列表
	domains := []string{"example.com", "mockdata.org", "testsite.net", "demoinfo.io", "sampledata.com"}

	// 生成随机的链接
	for i := 0; i < linkCount; i++ {
		// 随机选择一个域名
		domain := domains[rand.Intn(len(domains))]

		// 生成随机的链接标题
		linkTitle := fmt.Sprintf("相关链接 #%d", i+1)

		// 生成随机的链接URL
		linkURL := fmt.Sprintf("https://%s/related-%d", domain, rand.Intn(100))

		links[i] = Link{
			URL:   linkURL,
			Title: linkTitle,
		}
	}

	return MockCrawlResult{
		URL:     url,
		Title:   title,
		Content: content,
		Links:   links,
	}
}

// generateMockGolangCode 生成模拟Golang代码
func generateMockGolangCode(input string) string {
	// 根据输入生成不同类型的代码
	if strings.Contains(input, "分析") || strings.Contains(input, "analyze") {
		return generateDataAnalysisCode()
	} else if strings.Contains(input, "爬取") || strings.Contains(input, "crawl") {
		return generateWebCrawlerCode()
	} else if strings.Contains(input, "API") || strings.Contains(input, "服务") || strings.Contains(input, "service") {
		return generateAPIServiceCode()
	} else {
		return generateBasicCode(input)
	}
}

// generateDataAnalysisCode 生成数据分析相关的Golang代码
func generateDataAnalysisCode() string {
	return `package main

// 这里不需要import语句，使用文件顶部的import块

// 数据点结构
type DataPoint struct {
	ID    int
	Value float64
	Label string
}

// 分析结果结构
type AnalysisResult struct {
	Count       int
	Sum         float64
	Average     float64
	Min         float64
	Max         float64
	LabelCounts map[string]int
}

// 分析数据
func analyzeData(data []DataPoint) AnalysisResult {
	result := AnalysisResult{
		Count:       len(data),
		Sum:         0,
		Min:         0,
		Max:         0,
		LabelCounts: make(map[string]int),
	}

	if len(data) == 0 {
		return result
	}

	// 初始化最小值和最大值
	result.Min = data[0].Value
	result.Max = data[0].Value

	// 计算总和、最小值、最大值和标签计数
	for _, point := range data {
		result.Sum += point.Value

		if point.Value < result.Min {
			result.Min = point.Value
		}
		if point.Value > result.Max {
			result.Max = point.Value
		}

		result.LabelCounts[point.Label]++
	}

	// 计算平均值
	result.Average = result.Sum / float64(result.Count)

	return result
}

// 加载CSV数据
func loadCSVData(filename string) ([]DataPoint, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []DataPoint
	for i, record := range records {
		if i == 0 { // 跳过标题行
			continue
		}

		if len(record) < 3 {
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			continue
		}

		value, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			continue
		}

		data = append(data, DataPoint{
			ID:    id,
			Value: value,
			Label: record[2],
		})
	}

	return data, nil
}

func main() {
	// 模拟数据
	data := []DataPoint{
		{ID: 1, Value: 10.5, Label: "A"},
		{ID: 2, Value: 15.2, Label: "B"},
		{ID: 3, Value: 8.7, Label: "A"},
		{ID: 4, Value: 12.3, Label: "C"},
		{ID: 5, Value: 9.8, Label: "B"},
	}

	// 分析数据
	result := analyzeData(data)

	// 输出结果
	fmt.Println("数据分析结果:")
	fmt.Printf("数据点数量: %d\n", result.Count)
	fmt.Printf("总和: %.2f\n", result.Sum)
	fmt.Printf("平均值: %.2f\n", result.Average)
	fmt.Printf("最小值: %.2f\n", result.Min)
	fmt.Printf("最大值: %.2f\n", result.Max)

	fmt.Println("\n标签统计:")
	for label, count := range result.LabelCounts {
		fmt.Printf("%s: %d\n", label, count)
	}
}
`
}

// generateWebCrawlerCode 生成网页爬虫相关的Golang代码
func generateWebCrawlerCode() string {
	return `package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/html"
)

// 页面信息结构
type PageInfo struct {
	URL     string
	Title   string
	Links   []string
	Content string
}

// 爬取网页
func crawlPage(url string) (*PageInfo, error) {
	// 创建HTTP客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// 发送GET请求
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP请求失败，状态码: %d", resp.StatusCode)
	}

	// 解析HTML
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	// 提取信息
	pageInfo := &PageInfo{
		URL:   url,
		Links: []string{},
	}

	// 提取标题和链接
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			pageInfo.Title = n.FirstChild.Data
		}

		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					pageInfo.Links = append(pageInfo.Links, a.Val)
					break
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	// 提取页面内容（简化版，仅提取文本）
	var contentBuilder strings.Builder
	extractText(doc, &contentBuilder)
	pageInfo.Content = contentBuilder.String()

	return pageInfo, nil
}

// 提取文本内容
func extractText(n *html.Node, builder *strings.Builder) {
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text != "" {
			builder.WriteString(text)
			builder.WriteString(" ")
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractText(c, builder)
	}
}

func main() {
	// 模拟爬取网页
	fmt.Println("模拟网页爬虫")
	fmt.Println("URL: https://example.com")

	// 在实际应用中，这里会调用crawlPage函数
	// 但为了演示，我们使用模拟数据
	pageInfo := &PageInfo{
		URL:   "https://example.com",
		Title: "Example Domain",
		Links: []string{
			"https://www.iana.org/domains/example",
			"https://www.example.org",
			"https://www.example.net",
		},
		Content: "This domain is for use in illustrative examples in documents. You may use this domain in literature without prior coordination or asking for permission.",
	}

	// 输出结果
	fmt.Println("\n爬取结果:")
	fmt.Printf("标题: %s\n", pageInfo.Title)
	fmt.Printf("内容: %s\n", pageInfo.Content)

	fmt.Println("\n发现的链接:")
	for i, link := range pageInfo.Links {
		fmt.Printf("%d. %s\n", i+1, link)
	}
}
`
}

// generateAPIServiceCode 生成API服务相关的Golang代码
func generateAPIServiceCode() string {
	return `package main

// 这里不需要import语句，使用文件顶部的import块

// 用户结构
type User struct {
	ID       int    "json:\"id\""
	Username string "json:\"username\""
	Email    string "json:\"email\""
	Age      int    "json:\"age\""
}

// 用户存储
type UserStore struct {
	users map[int]User
	mutex sync.RWMutex
	nextID int
}

// 创建新的用户存储
func NewUserStore() *UserStore {
	return &UserStore{
		users:  make(map[int]User),
		nextID: 1,
	}
}

// 添加用户
func (s *UserStore) AddUser(user User) User {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	user.ID = s.nextID
	s.users[user.ID] = user
	s.nextID++

	return user
}

// 获取用户
func (s *UserStore) GetUser(id int) (User, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	user, exists := s.users[id]
	return user, exists
}

// 获取所有用户
func (s *UserStore) GetAllUsers() []User {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	users := make([]User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}

	return users
}

// 更新用户
func (s *UserStore) UpdateUser(id int, user User) (User, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.users[id]; !exists {
		return User{}, false
	}

	user.ID = id
	s.users[id] = user

	return user, true
}

// 删除用户
func (s *UserStore) DeleteUser(id int) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.users[id]; !exists {
		return false
	}

	delete(s.users, id)
	return true
}

func main() {
	// 创建用户存储
	userStore := NewUserStore()

	// 添加一些示例用户
	userStore.AddUser(User{Username: "user1", Email: "user1@example.com", Age: 25})
	userStore.AddUser(User{Username: "user2", Email: "user2@example.com", Age: 30})
	userStore.AddUser(User{Username: "user3", Email: "user3@example.com", Age: 35})

	// 创建HTTP服务器
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			// 获取所有用户或单个用户
			idStr := r.URL.Query().Get("id")
			if idStr != "" {
				id, err := strconv.Atoi(idStr)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, "{\"error\":\"Invalid user ID\"}")
					return
				}

				user, exists := userStore.GetUser(id)
				if !exists {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprintf(w, "{\"error\":\"User not found\"}")
					return
				}

				json.NewEncoder(w).Encode(user)
			} else {
				// 获取所有用户
				users := userStore.GetAllUsers()
				json.NewEncoder(w).Encode(users)
			}

		case http.MethodPost:
			// 添加新用户
			var user User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "{\"error\":\"Invalid request body\"}")
				return
			}

			createdUser := userStore.AddUser(user)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(createdUser)

		case http.MethodPut:
			// 更新用户
			idStr := r.URL.Query().Get("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "{\"error\":\"Invalid user ID\"}")
				return
			}

			var user User
			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "{\"error\":\"Invalid request body\"}")
				return
			}

			updatedUser, exists := userStore.UpdateUser(id, user)
			if !exists {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "{\"error\":\"User not found\"}")
				return
			}

			json.NewEncoder(w).Encode(updatedUser)

		case http.MethodDelete:
			// 删除用户
			idStr := r.URL.Query().Get("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "{\"error\":\"Invalid user ID\"}")
				return
			}

			if success := userStore.DeleteUser(id); !success {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "{\"error\":\"User not found\"}")
				return
			}

			w.WriteHeader(http.StatusNoContent)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "{\"error\":\"Method not allowed\"}")
		}
	})

	// 启动服务器
	fmt.Println("启动API服务器在 http://localhost:8080")
	fmt.Println("可用的端点:")
	fmt.Println("- GET /users - 获取所有用户")
	fmt.Println("- GET /users?id=1 - 获取指定ID的用户")
	fmt.Println("- POST /users - 创建新用户")
	fmt.Println("- PUT /users?id=1 - 更新指定ID的用户")
	fmt.Println("- DELETE /users?id=1 - 删除指定ID的用户")

	// 在实际应用中，这里会启动HTTP服务器
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// 为了演示，我们只显示一些模拟的API调用结果
	fmt.Println("\n模拟API调用结果:")

	// 获取所有用户
	fmt.Println("\nGET /users:")
	users := userStore.GetAllUsers()
	usersJSON, _ := json.MarshalIndent(users, "", "  ")
	fmt.Println(string(usersJSON))

	// 获取单个用户
	fmt.Println("\nGET /users?id=1:")
	user, _ := userStore.GetUser(1)
	userJSON, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(userJSON))
}
`
}

// generateBasicCode 生成基本的Golang代码
func generateBasicCode(input string) string {
	return fmt.Sprintf(`package main

import (
	"fmt"
	"time"
)

// 模拟数据结构
type MockData struct {
	ID        int
	Name      string
	Value     float64
	Timestamp time.Time
}

// 生成模拟数据
func generateMockData(count int) []MockData {
	data := make([]MockData, count)
	
	for i := 0; i < count; i++ {
		data[i] = MockData{
			ID:        i + 1,
			Name:      fmt.Sprintf("Item-%d", i+1),
			Value:     float64(i*10) + 0.5,
			Timestamp: time.Now().Add(time.Duration(i) * time.Hour),
		}
	}
	
	return data
}

// 处理模拟数据
func processMockData(data []MockData) {
	fmt.Println("处理模拟数据:")
	
	for _, item := range data {
		fmt.Printf("ID: %d, 名称: %s, 值: %.2f, 时间戳: %s\n", 
			item.ID, item.Name, item.Value, item.Timestamp.Format("2006-01-02 15:04:05"))
	}
}

func main() {
	fmt.Println("开始执行: %s")
	
	// 生成模拟数据
	data := generateMockData(5)
	
	// 处理数据
	processMockData(data)
	
	fmt.Println("\n执行完成!")
}
`, input)
}

// ToolManager 管理Eino工具
type ToolManager struct {
	tools map[string]tool.InvokableTool
}

// NewToolManager 创建工具管理器
func NewToolManager() *ToolManager {
	return &ToolManager{
		tools: make(map[string]tool.InvokableTool),
	}
}

// RegisterTool 注册工具
func (tm *ToolManager) RegisterTool(name string, t tool.InvokableTool) {
	tm.tools[name] = t
}

// GetTool 获取工具
func (tm *ToolManager) GetTool(name string) tool.InvokableTool {
	return tm.tools[name]
}

// GetToolByNameSuffix 通过名称后缀获取工具
func (tm *ToolManager) GetToolByNameSuffix(suffix string) tool.InvokableTool {
	for name, t := range tm.tools {
		if strings.HasSuffix(name, suffix) {
			return t
		}
	}
	return nil
}

// GetAllTools 获取所有工具
func (tm *ToolManager) GetAllTools() map[string]tool.InvokableTool {
	return tm.tools
}

// 全局工具管理器实例
var DefaultToolManager = NewToolManager()

// InitTools 初始化所有工具
func InitTools() error {
	// 初始化搜索工具
	for name, server := range config.Config.Tools.Servers {
		if err := initToolFromConfig(name, server); err != nil {
			return err
		}
	}

	return nil
}

// initToolFromConfig 从配置初始化工具
func initToolFromConfig(name string, server struct {
	APIKey string `yaml:"api_key"`
}) error {
	// 这里简化实现，根据配置创建相应的工具
	// 在实际应用中，可能需要更复杂的逻辑来创建不同类型的工具
	glog.Infof(context.Background(), "initializing tool: %s", name)

	// 使用Mock工具替代外部进程工具
	var tool tool.InvokableTool

	// 根据工具名称决定使用哪种Mock实现
	switch {
	case strings.Contains(name, "browser_search"):
		tool = &MockSearchTool{
			name: name,
		}
	case strings.Contains(name, "web_crawler"):
		tool = &MockCrawlTool{
			name: name,
		}
	default:
		// 默认使用简单的Mock工具
		tool = &MockSearchTool{
			name: name,
		}
	}

	// 注册工具
	DefaultToolManager.RegisterTool(name, tool)
	return nil
}

// ExternalProcessTool 外部进程工具实现
type ExternalProcessTool struct {
	name    string
	command string
	args    []string
	env     map[string]string
}

// Info 返回工具信息
func (t *ExternalProcessTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: t.name,
		Desc: fmt.Sprintf("External process tool: %s", t.command),
	}, nil
}

// InvokableRun 执行工具
func (t *ExternalProcessTool) InvokableRun(ctx context.Context, input string, opts ...tool.Option) (string, error) {
	// 创建命令
	cmd := exec.CommandContext(ctx, t.command, append(t.args, input)...)

	// 设置环境变量
	if t.env != nil {
		env := cmd.Environ()
		for k, v := range t.env {
			env = append(env, fmt.Sprintf("%s=%s", k, v))
		}
		cmd.Env = env
	}

	// 执行命令并获取输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		glog.Errorf(ctx, "tool execution error: %v, output: %s", err, string(output))
		return "", err
	}

	return string(output), nil
}

// MockSearchTool 模拟搜索工具实现
type MockSearchTool struct {
	name string
}

// Info 返回工具信息
func (t *MockSearchTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: t.name,
		Desc: "Mock search tool for tavily",
	}, nil
}

// InvokableRun 执行模拟搜索
func (t *MockSearchTool) InvokableRun(ctx context.Context, input string, opts ...tool.Option) (string, error) {
	glog.Infof(ctx, "Mock search tool running with input: %s", input)

	// 生成随机搜索结果
	result := generateMockSearchResults(input)

	// 转换为JSON
	outputJSON, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(outputJSON), nil
}

// MockCrawlTool 模拟网页爬取工具实现
type MockCrawlTool struct {
	name string
}

// Info 返回工具信息
func (t *MockCrawlTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: t.name,
		Desc: "Mock crawl tool for firecrawl",
	}, nil
}

// InvokableRun 执行模拟爬取
func (t *MockCrawlTool) InvokableRun(ctx context.Context, input string, opts ...tool.Option) (string, error) {
	glog.Infof(ctx, "Mock crawl tool running with input: %s", input)

	// 生成随机爬取结果
	result := generateMockCrawlResults(input)

	// 转换为JSON
	outputJSON, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(outputJSON), nil
}

// MockGolangTool 模拟Golang工具实现
type MockGolangTool struct {
	name string
}

// Info 返回工具信息
func (t *MockGolangTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: t.name,
		Desc: "Mock Golang tool",
	}, nil
}

// InvokableRun 执行模拟Golang代码
func (t *MockGolangTool) InvokableRun(ctx context.Context, input string, opts ...tool.Option) (string, error) {
	glog.Infof(ctx, "Mock Golang tool running with input: %s", input)

	// 生成随机Golang代码结果
	result := generateMockGolangCode(input)

	return result, nil
}
