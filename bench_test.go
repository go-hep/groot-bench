// Copyright ©2020 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bench_test

import (
	"io/ioutil"
	"os/exec"
	"testing"
)

func BenchmarkReadScalar(b *testing.B) {
	for _, lc := range []struct {
		kind string
		cmd  string
	}{
		{
			kind: "GoHEP",
			cmd:  "./bin/read-scalar",
		},
		{
			kind: "ROOT-TreeBranch",
			cmd:  "./bin/cxx-read-scalar-br",
		},
		{
			kind: "ROOT-TreeReader",
			cmd:  "./bin/cxx-read-scalar-rd",
		},
	} {
		b.Run(lc.kind, func(b *testing.B) {
			for _, bc := range []struct {
				name  string
				fname string
			}{
				{
					name:  "None",
					fname: "./testdata/scalar-none.root",
				},
				{
					name:  "LZ4",
					fname: "./testdata/scalar-lz4.root",
				},
				{
					name:  "Zlib-Lvl1",
					fname: "./testdata/scalar-zlib-1.root",
				},
				{
					name:  "Zlib-Lvl6",
					fname: "./testdata/scalar-zlib-6.root",
				},
				{
					name:  "Zlib-Lvl9",
					fname: "./testdata/scalar-zlib-9.root",
				},
			} {
				b.Run(bc.name, func(b *testing.B) {
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						b.StopTimer()
						cmd := exec.Command(lc.cmd, bc.fname)
						cmd.Stdout = ioutil.Discard
						cmd.Stderr = ioutil.Discard
						b.StartTimer()
						err := cmd.Run()
						if err != nil {
							b.Fatal(err)
						}
					}
				})
			}
		})
	}
}

func BenchmarkReadSlices(b *testing.B) {
	for _, lc := range []struct {
		kind string
		cmd  string
	}{
		{
			kind: "GoHEP",
			cmd:  "./bin/read-slices",
		},
		{
			kind: "ROOT-TreeBranch",
			cmd:  "./bin/cxx-read-slices-br",
		},
		{
			kind: "ROOT-TreeReader",
			cmd:  "./bin/cxx-read-slices-rd",
		},
	} {
		b.Run(lc.kind, func(b *testing.B) {
			for _, bc := range []struct {
				name  string
				fname string
			}{
				{
					name:  "None",
					fname: "./testdata/f64s-none.root",
				},
				{
					name:  "LZ4",
					fname: "./testdata/f64s-lz4.root",
				},
				{
					name:  "Zlib-Lvl1",
					fname: "./testdata/f64s-zlib-1.root",
				},
				{
					name:  "Zlib-Lvl6",
					fname: "./testdata/f64s-zlib-6.root",
				},
				{
					name:  "Zlib-Lvl9",
					fname: "./testdata/f64s-zlib-9.root",
				},
			} {
				b.Run(bc.name, func(b *testing.B) {
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						b.StopTimer()
						cmd := exec.Command(lc.cmd, bc.fname)
						cmd.Stdout = ioutil.Discard
						cmd.Stderr = ioutil.Discard
						b.StartTimer()
						err := cmd.Run()
						if err != nil {
							b.Fatal(err)
						}
					}
				})
			}
		})
	}
}

func BenchmarkReadCMSAll(b *testing.B) {
	for _, lc := range []struct {
		kind string
		cmd  string
		skip bool
	}{
		{
			kind: "GoHEP",
			cmd:  "./bin/read-cms-all",
		},
		{
			kind: "ROOT-TreeBranch",
			cmd:  "./bin/cxx-read-cms-all-br",
		},
	} {
		if lc.skip {
			b.Skip()
		}

		b.Run(lc.kind, func(b *testing.B) {
			for _, bc := range []struct {
				name  string
				fname string
			}{
				{
					name:  "Zlib",
					fname: "./testdata/Run2012B_DoubleElectron.root",
				},
				// {
				// 	name:  "Zlib-XRD",
				// 	fname: "root://eospublic.cern.ch//eos/root-eos/cms_opendata_2012_nanoaod/Run2012B_DoubleElectron.root",
				// },
			} {
				b.Run(bc.name, func(b *testing.B) {
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						b.StopTimer()
						cmd := exec.Command(lc.cmd, bc.fname)
						cmd.Stdout = ioutil.Discard
						cmd.Stderr = ioutil.Discard
						b.StartTimer()
						err := cmd.Run()
						if err != nil {
							b.Fatal(err)
						}
					}
				})
			}
		})
	}
}

func BenchmarkReadCMSScalar(b *testing.B) {
	for _, lc := range []struct {
		kind string
		cmd  string
		skip bool
	}{
		{
			kind: "GoHEP",
			cmd:  "./bin/read-cms-sca",
		},
		{
			kind: "ROOT-TreeBranch",
			cmd:  "./bin/cxx-read-cms-sca-br",
		},
		{
			kind: "ROOT-TreeReader",
			cmd:  "./bin/cxx-read-cms-sca-rd",
		},
	} {
		if lc.skip {
			b.Skip()
		}

		b.Run(lc.kind, func(b *testing.B) {
			for _, bc := range []struct {
				name  string
				fname string
			}{
				{
					name:  "Zlib",
					fname: "./testdata/Run2012B_DoubleElectron.root",
				},
				// {
				// 	name:  "Zlib-XRD",
				// 	fname: "root://eospublic.cern.ch//eos/root-eos/cms_opendata_2012_nanoaod/Run2012B_DoubleElectron.root",
				// },
			} {
				b.Run(bc.name, func(b *testing.B) {
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						b.StopTimer()
						cmd := exec.Command(lc.cmd, bc.fname)
						cmd.Stdout = ioutil.Discard
						cmd.Stderr = ioutil.Discard
						b.StartTimer()
						err := cmd.Run()
						if err != nil {
							b.Fatal(err)
						}
					}
				})
			}
		})
	}
}

func BenchmarkReadCMSAna(b *testing.B) {
	b.Skip() // not ready yet

	for _, lc := range []struct {
		kind string
		cmd  string
		skip bool
	}{
		{
			kind: "GoHEP",
			cmd:  "./bin/read-cms-ana",
		},
		{
			kind: "ROOT-TreeBranch",
			cmd:  "./bin/cxx-read-cms-ana-br",
		},
		{
			kind: "ROOT-TreeReader",
			cmd:  "./bin/cxx-read-cms-ana-rd",
		},
	} {
		if lc.skip {
			b.Skip()
		}

		b.Run(lc.kind, func(b *testing.B) {
			for _, bc := range []struct {
				name  string
				fname string
			}{
				{
					name:  "Zlib",
					fname: "./testdata/Run2012B_DoubleElectron.root",
				},
				// {
				// 	name:  "Zlib-XRD",
				// 	fname: "root://eospublic.cern.ch//eos/root-eos/cms_opendata_2012_nanoaod/Run2012B_DoubleElectron.root",
				// },
			} {
				b.Run(bc.name, func(b *testing.B) {
					b.ResetTimer()
					for i := 0; i < b.N; i++ {
						b.StopTimer()
						cmd := exec.Command(lc.cmd, bc.fname)
						cmd.Stdout = ioutil.Discard
						cmd.Stderr = ioutil.Discard
						b.StartTimer()
						err := cmd.Run()
						if err != nil {
							b.Fatal(err)
						}
					}
				})
			}
		})
	}
}
