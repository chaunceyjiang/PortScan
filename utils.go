package main

import (
	"errors"
	"log"
	"net"
	"strconv"
	"strings"
)

func getAllIP(ip string) []string {
	// 192.168.1.12-31,localhost
	var ips []string
	ipList := strings.Split(ip, ",")
	for _, i := range ipList {
		if strings.Contains(i, "-") {
			exIP, err := extendIP(i)
			if err != nil {
				log.Printf("%s %s", i, err)
				continue
			}
			ips = append(ips, exIP...)
		} else {
			ipAddr, err := net.ResolveIPAddr("ip", i)
			if err != nil {
				log.Printf("%s %s", i, err)
				continue
			}
			ips = append(ips, ipAddr.String())
		}
	}
	return ips
}

func extendIP(ip string) ([]string, error) {
	ipList := strings.Split(strings.Trim(ip, " "), "-")
	end, err := strconv.Atoi(ipList[1])
	if err != nil {
		return nil, errors.New("ip地址解析错误")
	}
	firstIP := net.ParseIP(ipList[0])
	if firstIP == nil {
		return nil, errors.New("ip地址解析错误")
	}
	startIP := strings.Split(firstIP.String(), ".")
	if len(startIP) != 4 {
		return nil, errors.New("ip地址解析错误")
	}
	start, err := strconv.Atoi(startIP[3])
	if err != nil {
		return nil, errors.New("ip地址解析错误")
	}
	var ips []string
	for s := start; s <= end; s++ {
		startIP[3] = strconv.Itoa(s)
		ips = append(ips,strings.Join(startIP,"."))
	}
	return ips, nil
}
