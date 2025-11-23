package v1

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/**
* @Author: jack.walker
* @File: delete_instance.go
* @CreateDate: 2025/11/21 16:46
* @Version: 1.0.0
* @Description:
 */

func (e *EurekaClientV1) DeRegister(instance *Instance) error {
	url := fmt.Sprintf("%s/eureka/apps/%s/%s", e.eurekaHost, instance.App, instance.InstanceID)
	if e.Debug {
		log.Printf("Deregister request Url: %v", url)
	}
	req, _ := http.NewRequest("DELETE", url, nil)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("eureka server returned status %s", resp.Status)
	}

	if e.Debug {
		log.Println("Deregister Response:", resp.Status)
		log.Printf("Deregister %v to %v sucess", instance.App, e.eurekaHost)
	}
	return nil
}
