package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type ScanIpPort struct {
	timeout time.Duration
	ips     string
	ports   string
	process int
}

func (sc *ScanIpPort) Scan() {
	ips := getAllIP(sc.ips)
	ports := getAllPort(sc.ports)
	wg := NewSizeWG(sc.process)
	if len(ips) != 0 && len(ports) != 0 {
		for _, ip := range ips {
			for _, port := range ports {
				wg.Add()
				go func(ip string, port int) {
					defer wg.Done()
					if sc.isOpen(ip, port) {
						fmt.Printf("ip地址 %s 端口%d  开放\n", ip, port)
					} //else {
					//	fmt.Printf("ip地址 %s 端口%d  未开放\n", ip, port)
					//}
				}(ip, port)
			}
		}
	}
	wg.Wait()
}

func (sc *ScanIpPort) isOpen(ip string, port int) bool {
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, port), sc.timeout*time.Millisecond)
	if err != nil {
		if !strings.Contains(err.Error(), "connection refused") &&
			!strings.Contains(err.Error(), "i/o timeout") {
			log.Println(err)
		}
		return false
	}
	_ = conn.Close()
	return true

}

func NewScanIpPort(ips string, ports string, timeout int, process int) *ScanIpPort {
	return &ScanIpPort{
		timeout: time.Duration(timeout),
		ips:     ips,
		ports:   ports,
		process: process,
	}

}
