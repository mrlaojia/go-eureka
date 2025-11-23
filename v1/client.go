package v1

import "strings"

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
