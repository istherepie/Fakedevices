all: testing build

# Get current commit hash
commit_hash := $(shell git rev-parse HEAD)

# Get current directory
current_dir := $(shell pwd)

testing:
	@echo "Running all tests"

build:
	@echo "Building binaries"

	mkdir $(current_dir)/build
	go build -o $(current_dir)/build/$(commit_hash) $(current_dir)/cmd/fakedevices/fakedevices.go

	ln -s $(current_dir)/build/$(commit_hash) $(current_dir)/build/fakedevices

clean:
	@echo "Cleaning up..."
	rm -rf $(current_dir)/build