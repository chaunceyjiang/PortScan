package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	port    = flag.String("p", "1-65535", "端口号范围 例如:-p 80,81,88-1000")
	help    = flag.Bool("h", false, "帮助信息")
	ip      = flag.String("ip", "127.0.0.1", "ip地址或域名  192.168.0.1-255,192.168.2.1,localhost")
	process = flag.Int("n", 32, "并发数")
	timeout = flag.Int("t", 250, "超时时间(毫秒)")
)

func init()  {
	log.SetFlags(log.Llongfile|log.LUTC|log.Ldate)
}


func main() {
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	sc := NewScanIpPort(*ip, *port, *timeout, *process)
	t := time.Now()
	fmt.Printf("扫描目标: ip %s port %s\n", *ip, *port)
	fmt.Printf("开始扫描: %s \n", t.Format("2006-01-02 15:04:05"))
	sc.Scan()
	fmt.Printf("总共耗时: %s\n", time.Now().Sub(t))
}
