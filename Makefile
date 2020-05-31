run:
	GOPROXY=direct GOSUMDB=off go run ./user/main.go
check_swag:
	which swag || GO111MODULE=off go get -u github.com/swaggo/swag/cmd/swag
.PHONY: doc
doc: check_swag
	swag init -d=./user/ -o=./user/docs -parseDependency=true