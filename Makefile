APPPATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
MUSES_SYSTEM:=github.com/mygomod/muses/pkg/system
NAME:=webhook
APPOUT:=$(APPPATH)/bin/$(NAME)

# 执行go指令
prod:
	@cd $(APPPATH) && $(APPPATH)/tool/build.sh $(NAME) $(APPOUT) $(MUSES_SYSTEM) && $(APPOUT) start --conf=conf/conf.toml

go:
	@cd $(APPPATH) && go run main.go start --conf=conf/conf.toml