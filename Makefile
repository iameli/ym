
all: .
	go build -o bin/ym github.com/iameli/ym/cmd/ym

install:
	sudo cp bin/ym /usr/local/bin/ym
