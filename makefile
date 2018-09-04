all: run

test:
	time qtdeploy -fast test desktop .

build:
	time qtdeploy -fast build desktop .

run: build
	open deploy/darwin/myqtapp.app
