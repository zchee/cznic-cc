# Copyright 2019 The CC Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	all bench clean cover cpu editor internalError later mem nuke todo edit devbench

grep=--include=*.go --include=*.l --include=*.y --include=*.yy
ngrep='internalError\|TODOOK\|lexer\.go\|ast.go\|trigraphs\.go\|.*_string\.go\|stringer\.go\|testdata\/gcc'

all: lexer.go
	date
	go version 2>&1 | tee log
	./unconvert.sh
	gofmt -l -s -w *.go
	go test -i
	go test 2>&1 -csmith 10s | tee -a log
	go install
	GOOS=linux GOARCH=arm go build
	GOOS=linux GOARCH=386 go build
	GOOS=linux GOARCH=amd64 go build
	GOOS=windows GOARCH=386 go build
	GOOS=windows GOARCH=amd64 go build
	go vet 2>&1 | grep -v $(ngrep) || true
	golint 2>&1 | grep -v $(ngrep) || true
	make todo
	misspell *.go
	staticcheck | grep -v 'lexer\.go' || true
	maligned || true
	pcregrep -nM 'FAIL|false|<nil>|:\n}' ast_test.go || true
	grep -n 'FAIL\|PASS' log 
	go version
	date 2>&1 | tee -a log

devbench:
	date 2>&1 | tee log-devbench
	go test -timeout 24h -dev -run @ -bench . 2>&1 | tee -a log-devbench
	grep -n 'FAIL\|SKIP' log-devbench || true

bench:
	date 2>&1 | tee log-bench
	go test -timeout 24h -v -run '^[^E]' -bench . 2>&1 | tee -a log-bench
	grep -n 'FAIL\|SKIP' log-bench || true

clean:
	go clean
	rm -f *~ *.test *.out

cover:
	t=$(shell mktemp) ; go test -coverprofile $$t && go tool cover -html $$t && unlink $$t

cpu: clean
	go test -run @ -bench . -cpuprofile cpu.out
	go tool pprof -lines *.test cpu.out

edit:
	touch log
	gvim -p Makefile *.l *.yy *.go &

editor: lexer.go
	gofmt -l -s -w *.go
	go test -i
	go test -short 2>&1 | tee log
	go install

ast.go lexer.go stringer.go: lexer.l parser.yy enum.go
	go generate

later:
	@grep -n $(grep) LATER * || true
	@grep -n $(grep) MAYBE * || true

mem: clean
	go test -run @ -dev -bench BenchmarkDevParse -memprofile mem.out -timeout 24h
	go tool pprof -lines -web -alloc_space *.test mem.out

nuke: clean
	go clean -i

todo:
	@grep -nr $(grep) ^[[:space:]]*_[[:space:]]*=[[:space:]][[:alpha:]][[:alnum:]]* * | grep -v $(ngrep) || true
	@grep -nr $(grep) 'TODO\|panic' * | grep -v $(ngrep) || true
	@grep -nr $(grep) BUG * | grep -v $(ngrep) || true
	@grep -nr $(grep) [^[:alpha:]]println * | grep -v $(ngrep) || true
