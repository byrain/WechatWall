CURRENT_GIT_GROUP := myGo
CURRENT_GIT_REPO  := httpServer
COMMONENVVAR      = GOOS=linux GOARCH=amd64
BUILDENVVAR       = CGO_ENABLED=0

export GOPATH := $(CURDIR)/_project
export GOBIN := $(CURDIR)/bin

.PHONY: deps install test add_dep clean docker

folder_dep:
	mkdir -p $(CURDIR)/_project/src/$(CURRENT_GIT_GROUP)
	test -d $(CURDIR)/_project/src/$(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO) || ln -s $(CURDIR) $(CURDIR)/_project/src/$(CURRENT_GIT_GROUP)

deps: folder_dep
	mkdir -p $(CURDIR)/vendor
	glide install

build: folder_dep
	go install $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)

linux_build: deps
	$(COMMONENVVAR) $(BUILDENVVAR) make build

all: deps linux_build

test:
	go test $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/test

test-v:
	go test -v $(CURRENT_GIT_GROUP)/$(CURRENT_GIT_REPO)/test

clean:
	@rm -rf vendor bin _project
