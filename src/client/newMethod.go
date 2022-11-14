package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	client := &http.Client{}
	request, _ := http.NewRequest("DELETE", "http://localhost:18888", nil)
	resp, _ := client.Do(request)
	dump, _ := httputil.DumpResponse(resp, true)
	log.Println(string(dump))
}
