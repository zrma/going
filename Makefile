VERSION = $(shell gobump show -r)
CURRENT_REVISION = $(shell git rev-parse --short HEAD)
BUILD_LDFLAGS = "-X github.com/zrma/going.revision=$(CURRENT_REVISION)"
ifdef update
  u=-u
endif

GO ?= GO111MODULE=on go

deps:
	env GO111MODULE=on go mod download

devel-deps: deps
	$(GO) get ${u} \
	  golang.org/x/lint/golint             \
	  github.com/mattn/goveralls           \
	  github.com/x-motemen/gobump/cmd/gobump \
	  github.com/Songmu/goxz/cmd/goxz      \
	  github.com/Songmu/ghch/cmd/ghch
	$(GO) install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow

test: deps
	./cover.sh

lint: devel-deps
	$(GO) vet ./...
	shadow ./...
	golint -set_exit_status ./...

cover: devel-deps
	COVERALLS_TOKEN=${COVERALLS_TOKEN} goveralls -coverprofile=coverage.out --service=travis-ci

bump: devel-deps
	_tools/releng

upload:
	ghr v$(VERSION) dist/v$(VERSION)

release: bump crossbuild upload

.PHONY: deps devel-deps test lint cover build crossbuild build upload release
