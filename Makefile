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
