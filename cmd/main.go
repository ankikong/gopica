package main

import (
	"flag"
	"fmt"

	"github.com/ankikong/gopica/pkg/session"
)

func main() {
	ipAddr := flag.String("addr", "", "设置直连ip")
	action := flag.String("action", "category", "想要获取什么信息")
	email := flag.String("email", "", "登录的账号")
	password := flag.String("password", "", "登录的密码")
	sess := flag.String("session", "./session.txt", "保存或加载session的文件")
	proxy := flag.String("proxy", "", "http代理")

	flag.Parse()

	s := session.NewPicaSession(*proxy, *ipAddr)
	err := s.Load(*sess)
	if err != nil {
		res, err := s.Login(*email, *password)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(res.Content))
		s.Save(*sess)
	}
	if *action == "category" {
		res, err := s.GetCategory()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(res.Content))
		}
	}
}
