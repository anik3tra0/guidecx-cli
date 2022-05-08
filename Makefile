# Overcome m2c: Command not found
# https://lists.mcs.anl.gov/pipermail/petsc-users/2010-November/007299.html
 %.o: %.mod
.PHONY: build

mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))
abs_dir := $(patsubst %/,%,$(dir $(mkfile_path)))
current_dir := $(notdir $(abs_dir))

lint:
	bash stern_test.sh
	go mod tidy -compat=1.17

format:
	gofmt -s -w .
	goimports -w -l .

gotests:
	[ "`goimports -l .`" = "" ] # Checks if you have formatted your file
	go test -v ./... --reset

test: gotests lint

build:
	./go-executable-build.sh github.com/anik3tra0/guidecx-cli

build_gobin_linux:
	env GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -o gcx -ldflags="-X 'cmd.VersionLabel=$(versionLabel)'" .

build_gobin_osx:
	env GOOS=darwin GARCH=amd64 CGO_ENABLED=0 go build -o gcx -ldflags="-X 'cmd.VersionLabel=$(versionLabel)'" .
