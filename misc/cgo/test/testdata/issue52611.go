// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Issue 52611: inconsistent compiler behaviour when compiling a C.struct.
// No runtime test; just make sure it compiles.

package cgotest

import (
	_ "cgotest/issue52611a"
	_ "cgotest/issue52611b"
)
