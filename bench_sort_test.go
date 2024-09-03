/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2018-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * bench_sort_test.go
 *
 *  Created on: Jun 09 2018
 *      Author: Massimiliano Ghilardi
 */
package main

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/truthtracer/gomacro/classic"
	"github.com/truthtracer/gomacro/fast"
)

var verbosesort = verbose

func quicksortInts(v []int) {
	for len(v) >= 30 {
		n := len(v)

		left, right := 0, n-1

		pivot := rand.Intn(n)

		v[pivot], v[right] = v[right], v[pivot]

		for i, _ := range v {
			if v[i] < v[right] {
				v[left], v[i] = v[i], v[left]
				left++
			}
		}

		v[left], v[right] = v[right], v[left]

		if left < n-left {
			quicksortInts(v[:left])
			v = v[left+1:]
		} else {
			quicksortInts(v[left+1:])
			v = v[:left]
		}
	}
	if n := len(v); n >= 3 {
		shellsortInts(v)
	} else if n == 2 {
		if v[0] > v[1] {
			v[0], v[1] = v[1], v[0]
		}
	}
}

// ---------------------- arrays: shellsort ------------------------

// array indexing is faster that slice indexing,
// provided the array is *not* copied. so use a pointer to array
var shellshort_gaps = &[...]int{8839, 3797, 1631, 701, 301, 132, 57, 23, 10, 4, 1}

func shellsortInts(v []int) {
	var i, j, n, gap, temp int
	n = len(v)
	for _, gap = range shellshort_gaps {
		for i = gap; i < n; i++ {
			temp = v[i]
			for j = i; j >= gap && v[j-gap] > temp; j -= gap {
				v[j] = v[j-gap]
			}
			v[j] = temp
		}
	}
}

func shellsortIntSlice(ints []int) {
	var v sort.IntSlice = ints
	var i, j, n, gap int
	n = v.Len()
	for _, gap = range shellshort_gaps {
		for i = gap; i < n; i++ {
			for j = i; j >= gap && v.Less(j, j-gap); j -= gap {
				v.Swap(j, j-gap)
			}
		}
	}
}

func shellsortInterfaces(ints []int) {
	// if v is declared with type sort.IntSlice, performance reaches shellsortInts() above
	var v sort.Interface = sort.IntSlice(ints)
	var i, j, n, gap int
	n = v.Len()
	for _, gap = range shellshort_gaps {
		for i = gap; i < n; i++ {
			for j = i; j >= gap && v.Less(j, j-gap); j -= gap {
				v.Swap(j, j-gap)
			}
		}
	}
}

const shellsort_ints_source_string = `
var shellshort_gaps = [...]int{8839, 3797, 1631, 701, 301, 132, 57, 23, 10, 4, 1}

func shellsort(v []int) {
	var i, j, n, temp int
	n = len(v)
	for _, gap := range shellshort_gaps {
		for i = gap; i < n; i++ {
			temp = v[i]
			for j = i; j >= gap && v[j-gap] > temp; j -= gap {
				v[j] = v[j-gap]
			}
			v[j] = temp
		}
	}
}`

const shellsort_intslice_source_string = `
import "sort"

var shellshort_gaps = [...]int{8839, 3797, 1631, 701, 301, 132, 57, 23, 10, 4, 1}

func shellsort(ints []int) {
	var v sort.IntSlice = ints
	var i, j, n, gap int
	n = v.Len()
	for _, gap = range shellshort_gaps {
		for i = gap; i < n; i++ {
			for j = i; j >= gap && v.Less(j, j-gap); j -= gap {
				v.Swap(j, j-gap)
			}
		}
	}
}`

const shellsort_interfaces_source_string = `
import "sort"

var shellshort_gaps = [...]int{8839, 3797, 1631, 701, 301, 132, 57, 23, 10, 4, 1}

func shellsort(ints []int) {
	var v sort.Interface = sort.IntSlice(ints)
	var i, j, n, gap int
	n = v.Len()
	for _, gap = range shellshort_gaps {
		for i = gap; i < n; i++ {
			for j = i; j >= gap && v.Less(j, j-gap); j -= gap {
				v.Swap(j, j-gap)
			}
		}
	}
}`

var shellsort_generic_source_string = `
var shellshort_gaps = [...]int{8839, 3797, 1631, 701, 301, 132, 57, 23, 10, 4, 1}

` + generic_func("shellsort", "T") + ` (v []T) {
	var i, j, n int
	var temp T
	n = len(v)
	for _, gap := range shellshort_gaps {
		for i = gap; i < n; i++ {
			temp = v[i]
			for j = i; j >= gap && v[j-gap] > temp; j -= gap {
				v[j] = v[j-gap]
			}
			v[j] = temp
		}
	}
}`

func BenchmarkQuickSortCompilerInts(b *testing.B) {
	benchmark_sort(b, quicksortInts)
}

func BenchmarkShellSortCompilerInts(b *testing.B) {
	benchmark_sort(b, shellsortInts)
}

func BenchmarkShellSortCompilerIntSlice(b *testing.B) {
	benchmark_sort(b, shellsortIntSlice)
}

func BenchmarkShellSortCompilerInterfaces(b *testing.B) {
	benchmark_sort(b, shellsortInterfaces)
}

func BenchmarkStdSortCompilerInterfaces(b *testing.B) {
	// use standard library sort
	benchmark_sort(b, sort.Ints)
}

func BenchmarkShellSortFastInts(b *testing.B) {
	ir := fast.New()
	ir.Eval(shellsort_ints_source_string)

	// extract the function shellsort()
	sort := ir.ValueOf("shellsort").Interface().(func([]int))

	benchmark_sort(b, sort)
}

func BenchmarkShellSortFastGeneric(b *testing.B) {
	ir := fast.New()
	ir.Eval(shellsort_generic_source_string)

	// extract the function shellsort#[int]()
	vs, _ := ir.Eval("shellsort#[int]")
	sort := vs[0].Interface().(func([]int))

	benchmark_sort(b, sort)
}

func BenchmarkShellSortFastInterfaces(b *testing.B) {
	ir := fast.New()
	ir.Eval(shellsort_interfaces_source_string)

	// extract the function shellsort()
	sort := ir.ValueOf("shellsort").Interface().(func([]int))

	benchmark_sort(b, sort)
}

func BenchmarkShellSortFastIntSlice(b *testing.B) {
	ir := fast.New()
	ir.Eval(shellsort_intslice_source_string)

	// extract the function shellsort()
	sort := ir.ValueOf("shellsort").Interface().(func([]int))

	benchmark_sort(b, sort)
}

func BenchmarkShellSortFastCompileLoop(b *testing.B) {
	ir := fast.New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ir.Comp.Binds = make(map[string]*fast.Bind)
		ir.Comp.BindNum = fast.NoIndex
		ir.Compile(shellsort_ints_source_string)
	}
}

func BenchmarkShellSortClassicInts(b *testing.B) {
	ir := classic.New()
	ir.Eval(shellsort_ints_source_string)

	// extract the function shellsort()
	sort := ir.ValueOf("shellsort").Interface().(func([]int))

	benchmark_sort(b, sort)
}

var sort_data = make_sort_data()

func make_sort_data() []int {
	const n = 10000
	v := make([]int, n)
	for i := 0; i < n; i++ {
		v[i] = rand.Int()
	}
	return v
}

func benchmark_sort(b *testing.B, sort func([]int)) {
	// call sort once for warm-up
	v := make([]int, len(sort_data))
	copy(v, sort_data)
	sort(v)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		copy(v, sort_data)
		sort(v)
	}
	if verbosesort {
		fmt.Println(v)
	}
}
