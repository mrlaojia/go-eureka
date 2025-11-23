package register

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

// RegisterCmd 创建一个命令
var RegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "注册应用信息到 eureka",
	Long: `注册应用信息到 eureka，instanceID=ip:port

示例：
go-eureka.exe register --app=jack2 --ip=1.1.1.1 --port=2025 --eureka=10.0.199.177:8080
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
	RegisterCmd.Flags().StringVar(&app, "app", "", "必须，app名字")
	RegisterCmd.Flags().StringVar(&ip, "ip", "", "必须，IP地址")
	RegisterCmd.Flags().IntVar(&port, "port", 8080, "必须，服务端口")

	RegisterCmd.PersistentFlags().SortFlags = false
	RegisterCmd.Flags().SortFlags = false

	RegisterCmd.MarkFlagRequired("app")
	RegisterCmd.MarkFlagRequired("ip")
	RegisterCmd.MarkFlagRequired("port")

}

// 执行逻辑
func run(cmd *cobra.Command, args []string) error {

	client := v1.CreateEurekaClientV1(viper.GetString("eureka"))
	client.Debug = viper.GetBool("verbose")

	ins := client.BuildInstance(app, ip, port)

	return client.RegisterInstance(ins)
}
