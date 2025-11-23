package info

import (
	v1 "github.com/mrlaojia/go-eureka/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

/**
* @Author: jack.walker
* @File: cmd.go
* @CreateDate: 2025/11/22 08:26
* @Version: 1.0.0
* @Description:
 */

// InfoCmd 创建一个命令
var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "获取 eureka 服务器的信息",
	Long: `获取 eureka 服务器的信息

示例：
# go-eureka.exe status --app=jack2 --ip=1.1.1.1 --port=2025 -e=10.0.199.177:8080
2025/11/23 09:27:03 JACK2 current status: UP
`,

	// 位置参数，这里接受
	// 位置参数，这里接受
	RunE: run,
}

// 参数设置
func init() {

}

// 执行逻辑
func run(cmd *cobra.Command, args []string) error {

	client := v1.CreateEurekaClientV1(viper.GetString("eureka"))
	client.Debug = viper.GetBool("verbose")

	status, err := client.EurekaStatus()
	if err != nil {
		return err
	}

	log.Println("eureka info:")
	log.Printf("  Environment：%v\n", status.GeneralStats.Environment)
	log.Printf("  ServerUptime: %v\n", status.GeneralStats.ServerUptime)
	log.Printf("  Cpus: %v\n", status.GeneralStats.Cpus)
	log.Printf("  TotalMemory: %v\n", status.GeneralStats.TotalMemory)
	log.Printf("  CurrentMemoryUsage: %v\n", status.GeneralStats.CurrentMemoryUsage)
	log.Printf("  Status: %v\n", status.InstanceInfo.Status)
	log.Printf("  HostName: %v\n", status.InstanceInfo.HostName)
	log.Printf("  App: %v\n", status.InstanceInfo.App)
	log.Printf("  IpAddr: %v\n", status.InstanceInfo.IpAddr)
	log.Printf("  Port: %v\n", status.InstanceInfo.Port.Value)

	return nil
}
