package main

import (
	"log"

	"github.com/geekymv/go-mq/server"
	"github.com/judwhite/go-svc"
)

type program struct {
	server *server.MQServer
}

func main() {
	prg := &program{}
	if err := svc.Run(prg); err != nil {
		log.Printf("run err:%v\n", err)
	}
}

func (p *program) Init(env svc.Environment) error {
	s, err := server.New()
	if err != nil {
		log.Printf("init err:%v", err)
	}
	p.server = s

	return nil
}

func (p *program) Start() error {
	go func() {
		p.server.Main()
	}()
	return nil
}

func (p *program) Stop() error {

	return nil
}
