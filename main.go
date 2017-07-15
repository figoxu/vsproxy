package main

import (
	"github.com/astaxie/beego/config"
	"github.com/quexer/utee"
	"log"
	"github.com/figoxu/Figo"
)

var (
	G_CFG config.Configer
)

func main() {
	cfg, err := config.NewConfig("ini", "conf.ini")
	utee.Chk(err)
	G_CFG = cfg

	source_addr := cfg.String("net::source_addr")
	target_addr:=cfg.String("net::target_addr")
	cocurrent,err := cfg.Int("net::cocurrent")
	utee.Chk(err)

	log.Println("@source_addr:",source_addr," @target_addr:",target_addr)
	Figo.NewTcpProxy(source_addr,target_addr,cocurrent).Listen()



}
