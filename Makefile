.PHONY: all binaries clean bench data-scalar data-slices

ROOT_FLAGS=`root-config --cflags --libs`
OPT=-O2

all: binaries

clean:
	/bin/rm -fr ./bin

binaries:
	mkdir -p ./bin
	## bench
	go get ./...
	go build -o ./bin/run-bench ./cmd/run-bench
	## generation
	go build -o ./bin/gen-data-scalar ./cmd/gen-data-scalar
	go build -o ./bin/gen-data-slices ./cmd/gen-data-slices
	## scalar
	go build -o ./bin/read-scalar ./cmd/read-scalar
	$(CXX) $(OPT) $(ROOT_FLAGS) -o bin/cxx-read-scalar-br ./cxx-root/read-scalar-br.cxx
	$(CXX) $(OPT) $(ROOT_FLAGS) -o bin/cxx-read-scalar-rd ./cxx-root/read-scalar-rd.cxx
	## slices
	go build -o ./bin/read-slices ./cmd/read-slices
	$(CXX) $(OPT) $(ROOT_FLAGS) -o bin/cxx-read-slices-br ./cxx-root/read-slices-br.cxx
	$(CXX) $(OPT) $(ROOT_FLAGS) -o bin/cxx-read-slices-rd ./cxx-root/read-slices-rd.cxx
	## CMS OpenData 2012
	go build -o ./bin/read-cms-all ./cmd/read-cms-all
	go build -o ./bin/read-cms-sca ./cmd/read-cms-sca
	go build -o ./bin/read-cms-ana ./cmd/read-cms-ana
	$(CXX) $(OPT) $(ROOT_FLAGS) -o bin/cxx-read-cms-all-br ./cxx-root/read-cms-all-br.cxx
	$(CXX) $(OPT) $(ROOT_FLAGS) -o bin/cxx-read-cms-sca-br ./cxx-root/read-cms-sca-br.cxx
	$(CXX) $(OPT) $(ROOT_FLAGS) -o bin/cxx-read-cms-sca-rd ./cxx-root/read-cms-sca-rd.cxx
	$(CXX) $(OPT) $(ROOT_FLAGS) -o bin/cxx-read-cms-ana-br ./cxx-root/read-cms-ana-br.cxx
	$(CXX) $(OPT) $(ROOT_FLAGS) -o bin/cxx-read-cms-ana-rd ./cxx-root/read-cms-ana-rd.cxx

bench:
	./bin/run-bench -count=20

data-scalar: binaries
	./bin/gen-data-scalar -zip=none -o ./testdata/scalar-none.root
	./bin/gen-data-scalar -zip=lz4  -o ./testdata/scalar-lz4.root
	./bin/gen-data-scalar -zip=zlib -lvl=1 -o ./testdata/scalar-zlib-1.root
	./bin/gen-data-scalar -zip=zlib -lvl=6 -o ./testdata/scalar-zlib-6.root
	./bin/gen-data-scalar -zip=zlib -lvl=9 -o ./testdata/scalar-zlib-9.root

data-slices: binaries
	./bin/gen-data-slices -zip=none -o ./testdata/f64s-none.root
	./bin/gen-data-slices -zip=lz4  -o ./testdata/f64s-lz4.root
	./bin/gen-data-slices -zip=zlib -lvl=1 -o ./testdata/f64s-zlib-1.root
	./bin/gen-data-slices -zip=zlib -lvl=6 -o ./testdata/f64s-zlib-6.root
	./bin/gen-data-slices -zip=zlib -lvl=9 -o ./testdata/f64s-zlib-9.root
