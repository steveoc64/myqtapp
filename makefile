all: run

test:
	time qtdeploy -fast test desktop .

build:
	time qtdeploy -fast build desktop .

run: build
	#./deploy/darwin/myqtapp.app/Contents/MacOS/myqtapp -cpuprofile out.cpu -memprofile out.mem
	open deploy/darwin/myqtapp.app
