GO_VER = $(shell go version | sed -n -r 's/.*go([0-9.]+).*/\1/p')
ifeq ($(strip $(GO_VER)), 1.10)
  $(error Please install and config GoLang 1.13+ environment)
endif

export GO111MODULE=on

GOFMT=gofmt
GOIMPORTS=goimports
DEBUG=Y

# module & pkg name
MODULE = $(shell env GO111MODULE=on go list -m)
#PKGS = $(or $(PKG),$(shell env GO111MODULE=on $(GO) list ./...))

# get GOPATH
GOPATH = $(shell go env GOPATH)

# these for --version cmd
BUILD_VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || cat $(CURDIR)/.version 2 > /dev/null || echo v0)
BUILD_TIME ?= $(shell date +%FT%T%z)
BUILD_NAME ?= $(MODULE)/$@_$(shell date "+%Y%m%d%H" )
COMMIT_SHA1 ?= $(shell git rev-parse HEAD )

# extract branch or tag func
BRANCH_NAME ?= $(shell git -C . rev-parse --abbrev-ref HEAD | grep -v HEAD || git -C . describe --exact-match HEAD || git -C . rev-parse HEAD)

# build dir
BIN_DIR = $(CURDIR)/build

# printf config
V = 0
Q = $(if $(filter 1,$V),,@)
M = $(shell printf "\033[34;1m▶\033[0m")


# build flag
COMMON_PKG = $(MODULE)/pkg
DEFINE_VARIABLE = -X '$(COMMON_PKG)/util.BuildVersion=$(BUILD_VERSION)' \
				  -X '$(COMMON_PKG)/util.BuildTime=$(BUILD_TIME)'       \
				  -X '$(COMMON_PKG)/util.BuildName=$(BUILD_NAME)'       \
				  -X '$(COMMON_PKG)/util.CommitID=$(COMMIT_SHA1)'       \
				  -X '$(COMMON_PKG)/util.BranchName=$(BRANCH_NAME)'     \
				  -X 'google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=warn' \

ifeq ($(MAKECMDGOALS),unittest)
	# BLD_FLAGS += -gcflags=all=-l
	DEFINE_VARIABLE += -X '$(COMMON_PKG)/util.UnittestCompile=true'
endif
BLD_FLAGS += -ldflags "$(DEFINE_VARIABLE)"

ifeq ($(DEBUG),Y)
	BLD_FLAGS +=  -gcflags=all="-N -l"
endif

# module list
MODS = second
# unittest files
UNITTEST_FILES = coverprofile.cov report.out cover.html test_result*

.PHONY: all
all: fmt $(MODS)
	@echo "Build all modules successfully."

.PHONY: unittest
unittest: create_errors
	$Q go test $(BLD_FLAGS) -v -covermode=count -coverprofile=coverprofile.cov ./... | tee report.out
	$Q go tool cover -html=coverprofile.cov  -o cover.html
	$Q mkdir -p test_result/
	$Q cp coverprofile.cov report.out cover.html test_result/
	$Q zip -q test_result.zip  test_result/*

.PHONY: $(MODS)
$(MODS): $(BIN); $(info $(M) building executable $@) @ ## Build program binary
	$Q go build $(BLD_FLAGS) -o $(BIN_DIR)/$@/$@ cmd/$@/*.go


$(BIN):
	@mkdir -p $@


.PHONY: fmt
fmt:
	$(GOFMT) -s -w .
	@echo "Go files have been formatted."

.PHONY: imports
goimports:
	$(GOIMPORTS) -l -w ./...
	@echo "Go files have been goimports fix."

.PHONY: clean
clean: ; $(info $(M) cleaning)	@ ## Cleanup everything
	-rm -f $(PROTO_META_GO)
	-rm -f $(PROTO_RPC_GO)
	-rm -rf $(BIN_DIR)/*
	-rm -rf ${UNITTEST_FILES}
	-rm -rf ${CURDIR}/pkg/common/errors/errors.go

.PHONY: env
env:
	@go env

.PHONY: version
version:
	@echo $(BUILD_VERSION)

.PHONY: help
help:
	@echo "--------------- ---------------"
	@echo "modules: $(MODS)"
	@echo "--------------- ---------------"
	@echo "Build all protos: make protos"
	@echo "--------------- ---------------"
	@echo "Build all modules: make all"
	@echo "Build certain module: make <modules>"
	@echo "--------------- ---------------"
	@echo "Clean all build files: make clean"
	@echo "--------------- ---------------"
	@echo "Print golang environment: make env"
	@echo "--------------- ---------------"
	@echo "Print build version: make version"
	@echo "--------------- ---------------"
