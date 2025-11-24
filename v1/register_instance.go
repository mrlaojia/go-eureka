package v1

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"time"
)

/**
* @Author: jack.walker
* @File: build_instance.go
* @CreateDate: 2025/11/21 16:20
* @Version: 1.0.0
* @Description:
 */

func (e *EurekaClientV1) BuildInstance(appName, appIp string, appPort int) *Instance {

	instanceID := fmt.Sprintf("%s:%v", appIp, appPort)
	instanceUrl := fmt.Sprintf("http://%v", instanceID)

	now := time.Now().UnixMilli()
	return &Instance{
		InstanceID: instanceID,
		HostName:   appIp,
		App:        appName,
		IPAddr:     appIp,
		Status:     "UP",
		Port:       Port{Enabled: "true", Value: appPort},
		SecurePort: Port{Enabled: "false", Value: 443},
		DataCenterInfo: DataCenter{
			Class: "com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo",
			Name:  "MyOwn",
		},
		HomePageUrl:             instanceUrl,
		StatusPageUrl:           instanceUrl + "/actuator/info",
		HealthCheckUrl:          instanceUrl + "/actuator/health",
		VipAddress:              appName,
		SecureVipAddress:        appName,
		IsCoordinatingDiscovery: false,
		LastUpdatedTimestamp:    now,
		LastDirtyTimestamp:      now,
		ActionType:              "ADDED",
		LeaseInfo: LeaseInfo{
			RenewalIntervalInSecs: 30,
			DurationInSecs:        90,
			RegistrationTimestamp: now,
			LastRenewalTimestamp:  now,
		},
		Metadata: Metadata{
			ManagementPort: appPort,
		},
	}
}

func (e *EurekaClientV1) RegisterInstance(instance *Instance) error {
	xmlData, err := xml.MarshalIndent(instance, "", "  ")
	if err != nil {
		return err
	}
	xmlData = append([]byte(xml.Header), xmlData...)

	reqUrl := fmt.Sprintf("%s/eureka/apps/%s", e.eurekaHost, instance.App)
	if e.Debug {
		log.Printf("RegisterInstance request Url: %v", reqUrl)
	}

	resp, err := e.post(reqUrl, bytes.NewBuffer(xmlData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		return fmt.Errorf("eureka server returned status %s", resp.Status)
	}

	if e.Debug {
		log.Println("Register Response:", resp.Status)
		log.Printf("Register %v to %v sucess", instance.App, e.eurekaHost)
	}
	return nil
}
