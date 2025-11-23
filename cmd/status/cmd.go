package status

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

// StatusCmd 创建一个命令
var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "从 eureka 获取应用状态",
	Long: `从 eureka 获取应用状态，instanceID=ip:port

示例：
# go-eureka.exe status --app=jack2 --ip=1.1.1.1 --port=2025 -e=10.0.199.177:8080
2025/11/23 09:27:03 JACK2 current status: UP
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
	StatusCmd.Flags().StringVar(&app, "app", "", "必须，app名字")
	StatusCmd.Flags().StringVar(&ip, "ip", "", "必须，IP地址")
	StatusCmd.Flags().IntVar(&port, "port", 8080, "必须，服务端口")

	StatusCmd.PersistentFlags().SortFlags = false
	StatusCmd.Flags().SortFlags = false

	StatusCmd.MarkFlagRequired("app")
	StatusCmd.MarkFlagRequired("ip")
	StatusCmd.MarkFlagRequired("port")

}

// 执行逻辑
func run(cmd *cobra.Command, args []string) error {

	client := v1.CreateEurekaClientV1(viper.GetString("eureka"))
	client.Debug = viper.GetBool("verbose")

	ins := client.BuildInstance(app, ip, port)

	ins, err := client.GetInstance(ins)
	if err != nil {
		return err
	}

	log.Printf("%v current status: %v", ins.App, ins.Status)

	return nil
}
