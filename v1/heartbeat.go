package v1

import (
	"fmt"
	"log"
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
	if e.Debug {
		log.Printf("SendHeartbeat request Url: %v", url)
	}

	resp, err := e.put(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("eureka server returned status %s", resp.Status)
	}

	if e.Debug {
		log.Println("Heartbeat Response:", resp.Status)
		log.Printf("%v send Heartbeat to %v sucess", instance.App, e.eurekaHost)
	}
	return nil
}
