package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/penguinn/pgo/jsonrpc2"
	"time"
	"sync"
	"github.com/mitchellh/mapstructure"
)
type Echo int

type AddParams struct {
	AddStart 	int 	`json:"addStart"`
	AddEnd 	  	int 	`json:"addEnd"`
}

type AddResult struct {
	Sum 	int    `json:"sum"`
}

func(p *Echo)Add(c context.Context, params *json.RawMessage) (interface{}, *jsonrpc2.Error) {
	var param AddParams
	if err := jsonrpc2.Unmarshal(params, &param); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return AddResult{
		Sum: param.AddStart + param.AddEnd,
	}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go serClient(wg)
	go serServer(wg)
	wg.Wait()
}

func serClient(wg sync.WaitGroup) {
	time.Sleep(1)
	client := jsonrpc2.NewClient("http://127.0.0.1:8080/v1/jrpc")
	params := AddParams{AddStart:1, AddEnd:2}
	resp, err := client.Call("Echo.Add", params, 1)
	if err != nil && resp.Error != nil{
		log.Fatal(err)
	}else {
		var sum AddResult
		err := mapstructure.WeakDecode(resp.Result, &sum)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(sum.Sum)
	}
	wg.Done()
}

func serServer(wg sync.WaitGroup) {
	log.Println(1)
	p := new(Echo)
	jsonrpc2.RegisterMethod("Echo.Add", p.Add, AddParams{}, AddResult{})
	http.HandleFunc("/v1/jrpc", jsonrpc2.Handler)
	http.HandleFunc("/v1/jrpc/debug", jsonrpc2.DebugHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
	wg.Done()
}
