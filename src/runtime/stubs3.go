// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !plan9
// +build !solaris
// +build !freebsd
// +build !darwin
// +build !aix
// +build !js
// +build !wasi

package runtime

//go:wasmimport nanotime1 go runtime.nanotime1 abi0
func nanotime1() int64
