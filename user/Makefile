run:
	GOPROXY=direct GOSUMDB=off go run main.go

check_swag:
	which swag || GO111MODULE=off go get -u github.com/swaggo/swag/cmd/swag
.PHONY: swag
swag: check_swag
	swag init -parseDependency=true