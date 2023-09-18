package ctrUser

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"go-homework/dto"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	var ip string
	// 获取请求ip
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		ip = forwarded
	}
	ip = r.RemoteAddr
	header := map[string]interface{}{
		"ip":      ip,
		"version": os.Getenv("GOVERSION"),
	}
	for k, v := range r.Header {
		header[k] = v
	}
	res, _ := json.Marshal(&(dto.HttpResponse{
		ErrNo: 0,
		Data: map[string]interface{}{
			"msg":    "I am health!",
			"header": header,
		},
	}))
	log.Printf("ip:%s", ip)
	log.Printf("res:%s", string(res))
	w.Header().Set("content-type", "text/json")
	w.WriteHeader(200)
	w.Write(res)
}
