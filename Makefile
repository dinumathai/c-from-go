
all: clean build

clean:
	rm main 2> /dev/null || true
	cd c-lib; make clean

build:
	cd c-lib; make static
	go run main.go
	# go build -o main main.go && ./main
