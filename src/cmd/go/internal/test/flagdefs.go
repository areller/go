// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by genflags.go — DO NOT EDIT.

package test

// passFlagToTest contains the flags that should be forwarded to
// the test binary with the prefix "test.".
var passFlagToTest = map[string]bool{
	"bench":                true,
	"benchmem":             true,
	"benchtime":            true,
	"blockprofile":         true,
	"blockprofilerate":     true,
	"count":                true,
	"coverprofile":         true,
	"cpu":                  true,
	"cpuprofile":           true,
	"failfast":             true,
	"fullpath":             true,
	"fuzz":                 true,
	"fuzzminimizetime":     true,
	"fuzztime":             true,
	"list":                 true,
	"memprofile":           true,
	"memprofilerate":       true,
	"mutexprofile":         true,
	"mutexprofilefraction": true,
	"outputdir":            true,
	"parallel":             true,
	"run":                  true,
	"short":                true,
	"shuffle":              true,
	"skip":                 true,
	"timeout":              true,
	"trace":                true,
	"v":                    true,
}

var passAnalyzersToVet = map[string]bool{
	"appends":          true,
	"asmdecl":          true,
	"assign":           true,
	"atomic":           true,
	"bool":             true,
	"bools":            true,
	"buildtag":         true,
	"buildtags":        true,
	"cgocall":          true,
	"composites":       true,
	"copylocks":        true,
	"defers":           true,
	"directive":        true,
	"errorsas":         true,
	"framepointer":     true,
	"hostport":         true,
	"httpresponse":     true,
	"ifaceassert":      true,
	"loopclosure":      true,
	"lostcancel":       true,
	"methods":          true,
	"nilfunc":          true,
	"printf":           true,
	"rangeloops":       true,
	"shift":            true,
	"sigchanyzer":      true,
	"slog":             true,
	"stdmethods":       true,
	"stdversion":       true,
	"stringintconv":    true,
	"structtag":        true,
	"testinggoroutine": true,
	"tests":            true,
	"timeformat":       true,
	"unmarshal":        true,
	"unreachable":      true,
	"unsafeptr":        true,
	"unusedresult":     true,
	"waitgroup":        true,
}
