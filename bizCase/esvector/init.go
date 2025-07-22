package main

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/morehao/golib/storages/dbes"
)

var ESClient *elasticsearch.Client

func init() {
	cfg := &dbes.ESConfig{
		Service: "es",
		Addr:    "http://localhost:9200",
	}
	simpleClient, _, err := dbes.InitES(cfg)
	if err != nil {
		panic(err)
	}
	ESClient = simpleClient
}
