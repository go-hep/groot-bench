// Copyright ©2020 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <iostream>
#include <stdint.h>
#include <vector>

#include "TFile.h"
#include "TTree.h"
#include "TTreeReader.h"
#include "TTreeReaderArray.h"

int main(int argc, char **argv) {
	auto fname = "./testdata/Run2012B_DoubleElectron.root";
	auto tname = "Events";

	if (argc > 1) {
		fname = argv[1];
	}
	if (argc > 2) {
		tname = argv[2];
	}

	auto f = TFile::Open(fname);
	auto r = TTreeReader(tname, f);

	// Readers to access the data (delete the ones you do not need).
	TTreeReaderValue<Int_t> run = {r, "run"};
	TTreeReaderValue<UInt_t> luminosityBlock = {r, "luminosityBlock"};
	TTreeReaderValue<ULong64_t> event = {r, "event"};
	TTreeReaderValue<Int_t> PV_npvs = {r, "PV_npvs"};
	TTreeReaderValue<Float_t> PV_x = {r, "PV_x"};
	TTreeReaderValue<Float_t> PV_y = {r, "PV_y"};
	TTreeReaderValue<Float_t> PV_z = {r, "PV_z"};
	TTreeReaderValue<UInt_t> nMuon = {r, "nMuon"};
	TTreeReaderArray<Float_t> Muon_pt = {r, "Muon_pt"};
	TTreeReaderArray<Float_t> Muon_eta = {r, "Muon_eta"};
	TTreeReaderArray<Float_t> Muon_phi = {r, "Muon_phi"};
	TTreeReaderArray<Float_t> Muon_mass = {r, "Muon_mass"};
	TTreeReaderArray<Int_t> Muon_charge = {r, "Muon_charge"};
	TTreeReaderArray<Float_t> Muon_pfRelIso03_all = {r, "Muon_pfRelIso03_all"};
	TTreeReaderArray<Float_t> Muon_pfRelIso04_all = {r, "Muon_pfRelIso04_all"};
	TTreeReaderArray<Bool_t> Muon_tightId = {r, "Muon_tightId"};
	TTreeReaderArray<Bool_t> Muon_softId = {r, "Muon_softId"};
	TTreeReaderArray<Float_t> Muon_dxy = {r, "Muon_dxy"};
	TTreeReaderArray<Float_t> Muon_dxyErr = {r, "Muon_dxyErr"};
	TTreeReaderArray<Float_t> Muon_dz = {r, "Muon_dz"};
	TTreeReaderArray<Float_t> Muon_dzErr = {r, "Muon_dzErr"};
	TTreeReaderValue<UInt_t> nElectron = {r, "nElectron"};
	TTreeReaderArray<Float_t> Electron_pt = {r, "Electron_pt"};
	TTreeReaderArray<Float_t> Electron_eta = {r, "Electron_eta"};
	TTreeReaderArray<Float_t> Electron_phi = {r, "Electron_phi"};
	TTreeReaderArray<Float_t> Electron_mass = {r, "Electron_mass"};
	TTreeReaderArray<Int_t> Electron_charge = {r, "Electron_charge"};
	TTreeReaderArray<Float_t> Electron_pfRelIso03_all = {r, "Electron_pfRelIso03_all"};
	TTreeReaderArray<Float_t> Electron_dxy = {r, "Electron_dxy"};
	TTreeReaderArray<Float_t> Electron_dxyErr = {r, "Electron_dxyErr"};
	TTreeReaderArray<Float_t> Electron_dz = {r, "Electron_dz"};
	TTreeReaderArray<Float_t> Electron_dzErr = {r, "Electron_dzErr"};

	int n = r.GetEntries();
	auto freq = n/10;
	auto i = 0;
	auto sum = 0;

	while (r.Next()) {
		if (i%freq==0) {
			std::cout << "Processing event " << i << "\n";
		}
		i++;
		sum += *event;
	}
	std::cout << "sum=" << sum << "\n";
}

