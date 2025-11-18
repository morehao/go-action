package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://httpbingo.org/cookies", nil)
	if err != nil {
		panic(err)
	}

	// 方式1：使用 AddCookie 方法（推荐）
	cookie1 := &http.Cookie{
		Name:  "foo",
		Value: "bar",
	}
	cookie2 := &http.Cookie{
		Name:  "token",
		Value: "123456",
	}

	req.AddCookie(cookie1)
	req.AddCookie(cookie2)

	// 可选：添加更多 Cookie 属性
	cookie3 := &http.Cookie{
		Name:     "session",
		Value:    "abc789",
		Path:     "/",
		Domain:   "httpbingo.org",
		Secure:   true,
		HttpOnly: true,
	}
	req.AddCookie(cookie3)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("响应内容:")
	fmt.Println(string(body))

	// 打印请求中实际发送的 Cookie 头
	fmt.Println("\n发送的 Cookie 头:")
	fmt.Println(req.Header.Get("Cookie"))
}
