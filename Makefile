# Vars

# Get current commit hash
commit_hash := $(shell git rev-parse --short=7 HEAD)

# Get current directory
current_dir := $(shell pwd)

# Targets
all: testing build

testing:
	@echo "Running all tests"

build:
	@echo "Building binaries"

	mkdir $(current_dir)/build
	go build -o $(current_dir)/build/$(commit_hash) $(current_dir)/cmd/main.go

	ln -s $(current_dir)/build/$(commit_hash) $(current_dir)/build/fakedevices

clean:
	@echo "Cleaning up..."
	rm -rf $(current_dir)/build