package v1

import (
	"io"
	"net/http"
	"strings"
	"time"
)

/**
* @Author: jack.walker
* @File: client.go
* @CreateDate: 2025/11/21 15:43
* @Version: 1.0.0
* @Description:
 */

type EurekaClientV1 struct {
	eurekaHost string
	Debug      bool
}

func CreateEurekaClientV1(eurekaHost string) *EurekaClientV1 {
	var eureka string

	// 如果已包含 http:// 或 https://，直接返回
	if strings.HasPrefix(eurekaHost, "http://") || strings.HasPrefix(eurekaHost, "https://") {
		eureka = eurekaHost
	} else {
		// 否则自动补 http://
		eureka = "http://" + eurekaHost
	}

	return &EurekaClientV1{
		eurekaHost: eureka,
		Debug:      false,
	}
}

func (e *EurekaClientV1) get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/xml")
	client := &http.Client{Timeout: 5 * time.Second}

	return client.Do(req)
}

func (e *EurekaClientV1) put(url string) (*http.Response, error) {
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 5 * time.Second}

	return client.Do(req)
}

func (e *EurekaClientV1) delete(url string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: 5 * time.Second}

	return client.Do(req)
}

func (e *EurekaClientV1) post(url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/xml")
	client := &http.Client{Timeout: 10 * time.Second}

	return client.Do(req)
}
