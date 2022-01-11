# Copyright 2021 The CC Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	ast.go build_all_targets clean cpu edit editor mem 

ast.go:
	go generate

build_all_targets:
	GOOS=darwin GOARCH=amd64 go build -v ./...
	GOOS=darwin GOARCH=arm64 go build -v ./...
	GOOS=freebsd GOARCH=amd64 go build -v ./...
	GOOS=linux GOARCH=386 go build -v ./...
	GOOS=linux GOARCH=amd64 go build -v ./...
	GOOS=linux GOARCH=arm go build -v ./...
	GOOS=linux GOARCH=arm64 go build -v ./...
	GOOS=netbsd GOARCH=amd64 go build -v ./...
	GOOS=windows GOARCH=386 go build -v ./...
	GOOS=windows GOARCH=amd64 go build -v ./...

clean:
	rm -f log-* cpu.test mem.test *.out
	go clean

cpu:
	go test -c -o cpu.test
	./cpu.test -test.run TestTranslationPhase4 -test.cpuprofile cpu.out
	go tool pprof --lines cpu.test cpu.out

edit:
	@touch log
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile *.go & fi

editor: ast.go
	gofmt -l -s -w *.go
	go test -o /dev/null -c
	go install 2>&1 | tee log

mem:
	go test -c -o mem.test
	./mem.test -test.run TestTranslationPhase4 -test.memprofile mem.out
	go tool pprof --lines --web --alloc_space mem.test mem.out
