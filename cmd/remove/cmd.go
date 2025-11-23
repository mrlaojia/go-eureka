package remove

import (
	v1 "github.com/mrlaojia/go-eureka/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

/**
* @Author: jack.walker
* @File: cmd.go
* @CreateDate: 2025/11/22 08:26
* @Version: 1.0.0
* @Description:
 */

// RemoveCmd 创建一个命令
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "从 eureka 移除应用信息",
	Long: `从 eureka 移除应用信息，instanceID=ip:port

示例：
go-eureka.exe remove --app=jack2 --ip=1.1.1.1 --port=2025 -e=10.0.199.177:8080
`,

	// 位置参数，这里接受
	// 位置参数，这里接受
	RunE: run,
}

var (
	app, ip string
	port    int
)

// 参数设置
func init() {
	RemoveCmd.Flags().StringVar(&app, "app", "", "必须，app名字")
	RemoveCmd.Flags().StringVar(&ip, "ip", "", "必须，IP地址")
	RemoveCmd.Flags().IntVar(&port, "port", 8080, "必须，服务端口")

	RemoveCmd.PersistentFlags().SortFlags = false
	RemoveCmd.Flags().SortFlags = false

	RemoveCmd.MarkFlagRequired("app")
	RemoveCmd.MarkFlagRequired("ip")
	RemoveCmd.MarkFlagRequired("port")

}

// 执行逻辑
func run(cmd *cobra.Command, args []string) error {

	client := v1.CreateEurekaClientV1(viper.GetString("eureka"))
	client.Debug = viper.GetBool("verbose")

	ins := client.BuildInstance(app, ip, port)

	return client.DeRegister(ins)
}
