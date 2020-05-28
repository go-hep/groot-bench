// Copyright ©2020 The go-hep Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <iostream>
#include <stdint.h>
#include <vector>

#include "TROOT.h"
#include "TFile.h"
#include "TTree.h"

struct Event {
   Int_t     run;
   UInt_t    luminosityBlock;
   ULong64_t event;
   Int_t     PV_npvs;
   Float_t   PV_x;
   Float_t   PV_y;
   Float_t   PV_z;
   UInt_t    nMuon;
   UInt_t    nElectron;
};

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
	auto t = f->Get<TTree>(tname);

	t->SetBranchStatus("*", 0);
	t->SetBranchStatus("run", 1);
	t->SetBranchStatus("luminosityBlock", 1);
	t->SetBranchStatus("event", 1);
	t->SetBranchStatus("PV_npvs", 1);
	t->SetBranchStatus("PV_x", 1);
	t->SetBranchStatus("PV_y", 1);
	t->SetBranchStatus("PV_z", 1);
	t->SetBranchStatus("nMuon", 1);
	t->SetBranchStatus("nElectron", 1);


	Event evt;

	t->SetBranchAddress("run", &evt.run);
	t->SetBranchAddress("luminosityBlock", &evt.luminosityBlock);
	t->SetBranchAddress("event", &evt.event);
	t->SetBranchAddress("PV_npvs", &evt.PV_npvs);
	t->SetBranchAddress("PV_x", &evt.PV_x);
	t->SetBranchAddress("PV_y", &evt.PV_y);
	t->SetBranchAddress("PV_z", &evt.PV_z);
	t->SetBranchAddress("nMuon", &evt.nMuon);
	t->SetBranchAddress("nElectron", &evt.nElectron);

	int n = t->GetEntries();
	auto freq = n/10;
	auto sum = 0;

	for (int i=0; i<n; i++) {
		if (i%freq==0) {
			std::cout << "Processing event " << i << "\n";
		}
		t->GetEntry(i);
		sum += evt.event;
	}
	std::cout << "sum=" << sum << "\n";
}

