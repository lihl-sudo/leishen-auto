package main

import (
	"fmt"
	"os"

	"leishen-auto/api"
	"leishen-auto/config"
)

func main() {
	fmt.Println("⌛️开始运行")

	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("❌错误: %v\n", err)
		os.Exit(1)
	}

	client := api.NewClient()

	resp, err := client.Pause(cfg.AccountToken, cfg.Lang)
	if err != nil {
		fmt.Printf("❌暂停失败: %v\n", err)
		os.Exit(1)
	}

	if resp.Code != 0 {
		if resp.Code == 400803 { // 400803 - 账号已经停止加速，请不要重复操作
			fmt.Printf("👌已经暂停: %d - %s\n", resp.Code, resp.Msg)
			fmt.Println("⌛️结束运行")
			return
		}
		fmt.Printf("❌暂停失败: %d - %s\n", resp.Code, resp.Msg)
		os.Exit(1)
	}

	fmt.Printf("%d:%s\n", resp.Code, resp.Msg)
	fmt.Println("✔️暂停成功")
	fmt.Println("⌛️结束运行")
}
