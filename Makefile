ACCOUNT="grengojbo"
TAG="1.4.0"
NAME="ivy"
PUBLIC_PORT=4900
PORT=${PUBLIC_PORT}
IMAGE_NAME = "${ACCOUNT}/${NAME}"
SITE="img.uatv.me"


# Program version
VERSION := $(shell grep "const Version " ivy/version.go | sed -E 's/.*"(.+)"$$/\1/')

# Binary name for bintray
BIN_NAME=$(shell basename $(abspath ./))

# Project owner for bintray
OWNER=${ACCOUNT}

# Project name for bintray
PROJECT_NAME=$(shell basename $(abspath ./))

# Project url used for builds
# examples: github.com, bitbucket.org
REPO_HOST_URL=bitbucket.org

# Grab the current commit
GIT_COMMIT="$(shell git rev-parse HEAD)"

# Check if there are uncommited changes
GIT_DIRTY="$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)"

# Add the godep path to the GOPATH
#GOPATH=$(shell godep path):$(shell echo $$GOPATH)

push:
	git push deis master

destroy:
	deis apps:destroy --app=${NAME} --confirm=${NAME}

create:
	deis create ${NAME}
	deis domains:add ${SITE} -a ${NAME}
	deis limits:set -m cmd=64M -a ${NAME}
	deis tags:set cluster=yes -a ${NAME}
	deis config:set NAME_APP=${NAME} -a ${NAME}
	# deis config:set NEW_RELIC_LICENSE_KEY=<key> -a ${NAME}
	# deis config:set NEW_RELIC_APP_NAME=${NAME} -a ${NAME}
	# deis config:set NEW_RELIC_APDEX=<0.010>

init:
	go get github.com/tools/godep

save:
	godep save

install:
	go get -v -u
	go get -v -u github.com/bradhe/stopwatch
	go get -v -u github.com/plimble/ace
	go get -v -u github.com/stretchr/testify

oldinstall:
	# go get -v  -u github.com/astaxie/beego
	# go get -v -u github.com/beego/bee
	go get -v -u github.com/astaxie/beego/orm
	go get -v -u github.com/beego/i18n
	go get -v -u github.com/go-sql-driver/mysql
	go get -v -u github.com/astaxie/beego/validation
	go get -v -u github.com/grengojbo/beego/modules/utils
	go get -v -u github.com/xyproto/permissions2
	go get -v -u github.com/yvasiyarov/beego_gorelic
	# go get -v -u github.com/mattbaird/gochimp
	# go get -v -u

release:
	@echo "building release ${OWNER} ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	@godep get
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GOPATH=`godep path`:$GOPATH go build -a -tags netgo -ldflags '-w' -o release/ivy ivy/main.go

old:
	#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o dist/run main.go
	# go build -ldflags "-X main.GitCommit ${GIT_COMMIT}${GIT_DIRTY}" -o run
	# go build -ldflags "-X main.GitCommit ${GIT_COMMIT}${GIT_DIRTY}" -o bin/${BIN_NAME}
	#@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o run main.go

build: clean
	@echo "building ${OWNER} ${BIN_NAME} ${VERSION}"
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o ivy ivy/main.go

clean:
	@test ! -e ./${BIN_NAME} || rm ./${BIN_NAME}

test:
	go test -v ./...

run: swagger docs
	@echo "...............................................................\n"
	@echo $(PROJECT_NAME)
	@echo documentation API open in browser:
	@echo	"	 http://localhost:8080/swagger/\n"
	@echo ...............................................................
	@bee run watchall true -downdoc=true -gendoc=true

docs:
	@bee generate docs

swagger:
	@test -d ./swagger || (wget https://github.com/beego/swagger/archive/v1.tar.gz && tar -xzf v1.tar.gz && mv swagger-1 swagger && rm v1.tar.gz)

# build/$(executable): *.go
# 	mkdir -p build
# 	go build -o build/$(executable)

# build/container: dist/$(executable)
# 	docker build --no-cache -t $(executable) .
# 	mkdir -p build
# 	touch build/container

# dist/$(executable): *.go
# 	mkdir -p dist
# 	GOOS=linux GOARCH=amd64 go build -o dist/$(executable)

# .PHONY: release
# release: build/container
# 	docker tag -f $(tag) $(account)/$(tag)
# 	docker push $(account)/$(tag)
# The entire Docker configuration is solely concerned with point

.PHONY: build dist clean test release run install docs swagger create push destroy save
