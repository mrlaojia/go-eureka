package v1

import (
	"fmt"
	"net/http"
	"time"
)

/**
* @Author: jack.walker
* @File: heartbeat.go
* @CreateDate: 2025/11/21 16:27
* @Version: 1.0.0
* @Description:
 */

func (e *EurekaClientV1) SendHeartbeat(instance *Instance) error {
	url := fmt.Sprintf("%s/eureka/apps/%s/%s", e.eurekaHost, instance.App, instance.InstanceID)
	req, _ := http.NewRequest("PUT", url, nil)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if e.Debug {
		fmt.Println("Heartbeat Response:", resp.Status)
	}
	return nil
}
