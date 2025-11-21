package v1

import "testing"

/**
* @Author: jack.walker
* @File: eurekaStatus_test.go
* @CreateDate: 2025/11/21 16:16
* @Version: 1.0.0
* @Description:
 */

func TestEurekaClientV1_Status(t *testing.T) {
	
	client := CreateEurekaClientV1("http://10.0.199.177:8080")
	status, err := client.EurekaStatus()
	if err != nil {
		t.Error(err.Error())
	}
	
	t.Log(status)
}
