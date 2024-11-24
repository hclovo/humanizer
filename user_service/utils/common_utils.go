package utils

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/bwmarrin/snowflake"
	"gopkg.in/gomail.v2"
)

// 随机生成字符串的函数
func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLZXCVBNMQWERTYUIOP")
	result := make([]byte, n)
 
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
 
	return string(result)
}

// 发送验证码邮件
func SendEmailCode(to string, code string) error {
	log.Println("Sending email to", to)
	config := AppConfig.SMTP

	// SMTP服务器配置
	smtpHost := config.Host
	smtpPortStr := config.Port
	sender := config.Sender
	password := config.Password

	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		return fmt.Errorf("failed to convert smtpPort to int: %w", err)
	}

	// 构建邮件内容
	subject := "Your Verification Code"
	body := fmt.Sprintf("Your verification code is: %s. It is valid for 5 minutes.", code)
	message := fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body)

	m := gomail.NewMessage()
	m.SetHeader("From", sender)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", message)

	d := gomail.NewDialer(smtpHost, smtpPort, sender, password)
	return d.DialAndSend(m)
}


// 根据雪花算法生成uid
func GetUID() (string, error) {
    // 设置自定义起始时间戳（Twitter Snowflake 默认起始时间为 1970 年）
	startTime := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	snowflake.Epoch = startTime.UnixNano() / 1e6 // 转换为毫秒

	// 创建一个节点，节点 ID 范围为 [0, 1023]
	node, err := snowflake.NewNode(1) // 这里设置节点 ID 为 1
	if err != nil {
		return "", fmt.Errorf("failed to create snowflake node: %w", err)
	}

	// 生成唯一 ID
	id := node.Generate()

	return id.String(), nil
}