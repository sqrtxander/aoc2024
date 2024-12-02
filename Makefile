.PHONY: all
all: getinput submit

getinput:
	go build -o ./bin/aocinput ./support/getinput/main.go

submit:
	go build -o ./bin/aocsubmit ./support/submit/main.go

.PHONY: clean
clean:
	rm ./bin/aocinput
	rm ./bin/aocsubmit

.PHONY: hs
hs:
	mkdir -p ./day${d}/part${p}/bin/
	mkdir -p ./day${d}/part${p}/build/
	ghc -o ./day${d}/part${p}/bin/solve -outputdir ./day${d}/part${p}/build/ ./day${d}/part${p}/solve.hs
	rm -r ./day${d}/part${p}/build
