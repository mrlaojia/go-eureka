package v1

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
	return &EurekaClientV1{
		eurekaHost: eurekaHost,
		Debug:      false,
	}
}
