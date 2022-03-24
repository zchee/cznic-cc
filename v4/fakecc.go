// Copyright 2022 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fn := os.Getenv("FAKE_CC_LOG")
	if fn == "" {
		fmt.Fprintf(os.Stderr, "FAKE_CC_LOG not set\n")
		os.Exit(1)
	}

	cc := os.Getenv("FAKE_CC_CC")
	if cc == "" {
		fmt.Fprintf(os.Stderr, "FAKE_CC_CC not set\n")
		os.Exit(1)
	}

	f, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_RDWR|os.O_SYNC, 0660)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	cmd := exec.Command(cc, os.Args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		os.Exit(err.(*exec.ExitError).ExitCode())
	}

	a := append([]string{wd}, os.Args[1:]...)
	var b bytes.Buffer
	if err := json.NewEncoder(&b).Encode(a); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	if _, err := f.Write(b.Bytes()); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
