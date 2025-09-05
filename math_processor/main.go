package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("数学计算处理器启动...")

	// 检查输入文件
	if _, err := os.Stat("data/problems.txt"); os.IsNotExist(err) {
		log.Fatal("输入文件 data/problems.txt 不存在")
	}

	// 创建输出目录
	if err := os.MkdirAll("output", 0755); err != nil {
		log.Fatal("创建输出目录失败:", err)
	}

	// 记录开始时间
	startTime := time.Now()

	// TODO: 在这里实现生产者-消费者模式
	// 1. 生产者：读取problems.txt，发送计算任务到channel
	// 2. 消费者：从channel接收任务，计算结果，写入results.txt
	// 3. 协调：确保所有任务完成后正常退出

	fmt.Println("所有计算任务处理完成")

	// 计算总耗时
	duration := time.Since(startTime)
	fmt.Printf("总耗时: %v\n", duration)

	// TODO: 验证输出文件
	if _, err := os.Stat("output/results.txt"); err != nil {
		log.Printf("警告: 输出文件不存在或无法访问")
	} else {
		fmt.Println("结果已保存到 output/results.txt")
	}
}
