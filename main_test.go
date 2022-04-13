package main

import (
    "net/http"
	"fmt"
	"testing"
)

func TestRootHandler(t *testing.T){
	resp, _ := http.Get("http://localhost:2345/")
	if (resp.StatusCode != 200) {
		t.Error("Expected status 200, got", resp.StatusCode)
	}
	fmt.Println("/hello test passed");
}

func TestHelloHandler(t *testing.T){
	resp, _ := http.Get("http://localhost:2345/hello")
	if (resp.StatusCode != 200) {
		t.Error("Expected status 200, got", resp.StatusCode)
	}
	fmt.Println("/hello_there test passed");
}

func TestHelloThereHandler(t *testing.T){
	resp, _ := http.Get("http://localhost:2345/hello_there")
	if (resp.StatusCode != 200) {
		t.Error("Expected status 200, got", resp.StatusCode)
	}
	fmt.Println("/ test passed");
}