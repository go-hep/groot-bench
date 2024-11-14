// Copyright ©2020 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"time"

	bench "github.com/go-hep/groot-bench"
	"go-hep.org/x/hep/groot"
	"go-hep.org/x/hep/groot/rtree"
)

func main() {
	var (
		nevts = flag.Int64("nevts", -1, "number of events to read")
		tname = flag.String("t", "Events", "name of the ROOT tree to read")

		cpuProf = flag.String("cpu-profile", "", "path to the output CPU profile")
	)

	log.SetPrefix("groot: ")
	log.SetFlags(0)

	flag.Parse()

	bench.Version()

	if *cpuProf != "" {
		prof, err := os.Create(*cpuProf)
		if err != nil {
			log.Fatalf("could not create CPU profile: %+v", err)
		}
		defer prof.Close()
		err = pprof.StartCPUProfile(prof)
		if err != nil {
			log.Fatalf("could not start CPU profile: %+v", err)
		}
		defer pprof.StopCPUProfile()
	}

	read(flag.Arg(0), *tname, *nevts)
}

func read(fname, tname string, nevts int64) {
	start := time.Now()
	defer func() {
		log.Printf("read time: %v", time.Since(start))
	}()

	log.Printf("reading ROOT file: %s", fname)
	f, err := groot.Open(fname)
	if err != nil {
		log.Fatalf("could not open ROOT file: %+v", err)
	}
	defer f.Close()

	obj, err := f.Get(tname)
	if err != nil {
		log.Fatalf("could not get ROOT tree: %+v", err)
	}
	t := obj.(rtree.Tree)

	var (
		evt   Event
		rvars = rtree.ReadVarsFromStruct(&evt)
	)

	if nevts < 0 {
		nevts = t.Entries()
	}

	r, err := rtree.NewReader(t, rvars, rtree.WithRange(0, nevts))
	if err != nil {
		log.Fatalf("could not create tree reader: %+v", err)
	}
	defer r.Close()

	log.Printf("-- created reader: evts=%d", nevts)
	for i, rv := range rvars {
		log.Printf("branch[%d]: name=%q", i, rv.Name)
	}

	var (
		n    = 0
		freq = nevts / 10
	)
	err = r.Read(func(rctx rtree.RCtx) error {
		if rctx.Entry%freq == 0 {
			log.Printf("event %d...", rctx.Entry)
		}
		n++
		return nil
	})
	if err != nil {
		log.Fatalf("could not read tree: %+v", err)
	}
	log.Printf("read %d entries", n)
}

// Event is the data contained in a rtree.Tree.
type Event struct {
	ROOT_run                     int32     `groot:"run"`
	ROOT_luminosityBlock         uint32    `groot:"luminosityBlock"`
	ROOT_event                   uint64    `groot:"event"`
	ROOT_PV_npvs                 int32     `groot:"PV_npvs"`
	ROOT_PV_x                    float32   `groot:"PV_x"`
	ROOT_PV_y                    float32   `groot:"PV_y"`
	ROOT_PV_z                    float32   `groot:"PV_z"`
	ROOT_nMuon                   uint32    `groot:"nMuon"`
	ROOT_Muon_pt                 []float32 `groot:"Muon_pt"`
	ROOT_Muon_eta                []float32 `groot:"Muon_eta"`
	ROOT_Muon_phi                []float32 `groot:"Muon_phi"`
	ROOT_Muon_mass               []float32 `groot:"Muon_mass"`
	ROOT_Muon_charge             []int32   `groot:"Muon_charge"`
	ROOT_Muon_pfRelIso03_all     []float32 `groot:"Muon_pfRelIso03_all"`
	ROOT_Muon_pfRelIso04_all     []float32 `groot:"Muon_pfRelIso04_all"`
	ROOT_Muon_dxy                []float32 `groot:"Muon_dxy"`
	ROOT_Muon_dxyErr             []float32 `groot:"Muon_dxyErr"`
	ROOT_Muon_dz                 []float32 `groot:"Muon_dz"`
	ROOT_Muon_dzErr              []float32 `groot:"Muon_dzErr"`
	ROOT_nElectron               uint32    `groot:"nElectron"`
	ROOT_Electron_pt             []float32 `groot:"Electron_pt"`
	ROOT_Electron_eta            []float32 `groot:"Electron_eta"`
	ROOT_Electron_phi            []float32 `groot:"Electron_phi"`
	ROOT_Electron_mass           []float32 `groot:"Electron_mass"`
	ROOT_Electron_charge         []int32   `groot:"Electron_charge"`
	ROOT_Electron_pfRelIso03_all []float32 `groot:"Electron_pfRelIso03_all"`
	ROOT_Electron_dxy            []float32 `groot:"Electron_dxy"`
	ROOT_Electron_dxyErr         []float32 `groot:"Electron_dxyErr"`
	ROOT_Electron_dz             []float32 `groot:"Electron_dz"`
	ROOT_Electron_dzErr          []float32 `groot:"Electron_dzErr"`
}
