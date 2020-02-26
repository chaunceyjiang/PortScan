package main

import (
	"flag"
	"os"
)

var (
	port = flag.String("p", "80", "端口号范围 例如:-p 80,81,88-1000")
	help = flag.Bool("h", false, "帮助信息")
	ip = flag.String("ip","127.0.0.1","ip地址或域名  192.168.0.1-255,192.168.2.1,localhost")
	process = flag.Int("n",32,"并发数")
	timeout = flag.Int("t",250,"超时时间(毫秒)")
)


func main() {
	flag.Parse()
	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}
	//wg:=NewSizeWG(*process)

}
