.PHONY: all common micro grpc build test run deploy local clean echo monitor

GOPATH:=$(shell go env GOPATH)

all: micro grpc monitor
# To do: there should be a script to loop through all the micro services.
# To do: there should be only one copy of each proto file.
common:
	@echo "Generating common proto"
	@cd common && protoc \
	--go_out=. \
	*.proto

micro: common
	@echo "Generating micro proto"
	@for type in email usermgr appmgr dcmgr rolemgr; do \
		cp $$type/v1/*.proto $$type/v1/micro/; \
		protoc -I. -Icommon \
			--micro_out=Mcommon/common.proto=github.com/Ankr-network/dccn-common/protos/common:. \
			--go_out=Mcommon/common.proto=github.com/Ankr-network/dccn-common/protos/common:. \
			$$type/v1/micro/*.proto; \
		rm $$type/v1/micro/*.proto; \
	done

grpc: common
	@echo "Generating gPRC proto"
	@for type in email usermgr appmgr dcmgr rolemgr; do \
		cp $$type/v1/*.proto $$type/v1/grpc/; \
		protoc -I. -Icommon \
			--go_out=plugins=grpc,Mcommon/common.proto=github.com/Ankr-network/dccn-common/protos/common:. \
			$$type/v1/grpc/*.proto; \
		rm $$type/v1/grpc/*.proto; \
	done

monitor:
	@echo "Generating monitor proto"
	@cd monitor && make

