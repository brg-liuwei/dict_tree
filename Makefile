ENV = env.tmp
$(shell ./env.sh ${ENV})
include ${ENV}

.PHONY: all benchmark clean

all:
	@go build -o bin/dict src/main.go

bench:
	@cd src/dict_tree && go test -bench="." && cd ../..

test:
	@go test dict_tree

clean:
	@rm -rf ${ENV}
	@rm -rf bin/*
