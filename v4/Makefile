# Copyright 2021 The CC Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

.PHONY:	edit editor ast.go

edit:
	@touch log
	@if [ -f "Session.vim" ]; then gvim -S & else gvim -p Makefile *.go & fi

editor: ast.go stringer.go
	gofmt -l -s -w *.go
	go test -o /dev/null -c
	go install 2>&1 | tee log

ast.go:
	go generate
