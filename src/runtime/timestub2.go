// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !darwin
// +build !windows
// +build !freebsd
// +build !aix
// +build !solaris

package runtime

//go:wasmimport walltime1 go runtime.walltime1 abi0
func walltime1() (sec int64, nsec int32)
