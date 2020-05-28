// Copyright ©2020 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <iostream>
#include <stdint.h>
#include <vector>

#include "TROOT.h"
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

#ifdef GROOT_ENABLE_IMT
	ROOT::EnableImplicitMT();
#endif // GROOT_ENABLE_IMT

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
	TTreeReaderValue<UInt_t> nElectron = {r, "nElectron"};

	int n = r.GetEntries();
	auto freq = n/10;
	auto i = 0;
	auto sum = 0;

	while (r.Next()) {
		if (i%freq==0) {
			std::cout << "Processing event " << i << "\n";
		}
		i++;
		sum += 0
		+ *run
		+ *luminosityBlock
		+ *event
		+ *PV_npvs
		+ *nMuon
		+ *PV_npvs
		+ *PV_x
		+ *PV_y
		+ *PV_z
		+ *nMuon
		+ *nElectron
		;
	}
	std::cout << "sum=" << sum << "\n";
}

