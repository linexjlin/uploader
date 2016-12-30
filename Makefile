export GOPATH=$(shell pwd)

BUILDTAGS=release
default: linux 

linux:
	GOOS=linux GOARCH=amd64 go install -tags '$(BUILDTAGS)' uploader

windows:
	GOOS=windows GOARCH=amd64 go install -tags '$(BUILDTAGS)' uploader

all: linux windows
