PKG  = github.com/DevMine/srcanlzr
EXEC = srcanlzr

all: check test build

install:
	go install ${PKG}

build:
	go build -o ${EXEC} ${PKG}

test:
	go test ${PKG}/...

deps:
	 go get -u -f github.com/DevMine/repotool/model

dev-deps:
	 go get -u github.com/golang/lint/golint

check:
	go vet ${PKG}/...
	golint ${PKG}/...

cover:
	go test -cover ${PKG}/...

clean:
	rm -f ./${EXEC}
