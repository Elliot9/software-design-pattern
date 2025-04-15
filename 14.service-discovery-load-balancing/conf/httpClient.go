package conf

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	ServiceDiscovery = make(map[string][]string)
	BlackList        = make(map[string]bool)
)

func init() {
	loadConfig("config.ini")
}

func loadConfig(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentSection string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// 跳過空行或註解
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// 檢查是否是區塊標題
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentSection = line[1 : len(line)-1]
			continue
		}

		// BlackList
		if currentSection == "BlackList" {
			entries := strings.Split(line, ",")
			for _, entry := range entries {
				BlackList[strings.TrimSpace(entry)] = true
			}
			continue
		}

		//  ServiceDiscovery (key: value)
		if currentSection == "ServiceDiscovery" {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) != 2 {
				fmt.Println("Invalid line format:", line)
				continue
			}

			host := strings.TrimSpace(parts[0])
			ips := strings.Split(parts[1], ",")
			for i, ip := range ips {
				ips[i] = strings.TrimSpace(ip)
			}

			ServiceDiscovery[host] = ips
		}
	}

	return scanner.Err()
}

func GetServiceDiscovery() map[string][]string {
	return ServiceDiscovery
}

func GetBlackList() map[string]bool {
	return BlackList
}
