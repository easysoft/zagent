VERSION=1.0.0
PROJECT=zagent
PACKAGE=${PROJECT}-${VERSION}
BINARY_SERVER=server
BINARY_HOST=host
BINARY_VM=vm
BIN_DIR=bin
BIN_ZIP_DIR=${BIN_DIR}/zip/${PROJECT}/${VERSION}/
BIN_ZIP_RELAT=../../../zip/${PROJECT}/${VERSION}/
BIN_OUT=${BIN_DIR}/${PROJECT}/${VERSION}/
BIN_SERVER_WIN64=${BIN_OUT}win64/
BIN_SERVER_WIN32=${BIN_OUT}win32/
BIN_SERVER_LINUX=${BIN_OUT}linux/
BIN_SERVER_MAC=${BIN_OUT}mac/
QINIU_DIR=/Users/aaron/work/zentao/qiniu/
QINIU_DIST_DIR=${QINIU_DIR}${PROJECT}/${VERSION}/

help: ## this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

host: prepare_res_host compile_host copy_host_files package_host
vm: prepare_res_vm compile_vm copy_vm_files package_vm

prepare_res_host:
	@echo 'start prepare host res'
	@go-bindata -o=res/host/res.go -pkg=hostRes res/host/...
	@rm -rf ${BIN_DIR}

prepare_res_vm:
	@echo 'start prepare vm res'
	@go-bindata -o=res/vm/res.go -pkg=vmRes res/vm/...
	@rm -rf ${BIN_DIR}

compile_host: compile_host_linux
compile_vm: compile_vm_win64 compile_vm_win32 compile_vm_linux compile_vm_mac

compile_host_linux:
	@echo 'start compile host linux'
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ${BIN_SERVER_LINUX}${BINARY_HOST}/host cmd/host/main.go

compile_vm_win64:
	@echo 'start compile vm win64'
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BIN_SERVER_WIN64}${BINARY_VM}/vm.exe cmd/vm/main.go
compile_vm_win32:
	@echo 'start compile vm win32'
	@CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ${BIN_SERVER_WIN32}${BINARY_VM}/vm.exe cmd/vm/main.go
compile_vm_linux:
	@echo 'start compile vm linux'
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BIN_SERVER_LINUX}${BINARY_VM}/vm cmd/vm/main.go
compile_vm_mac:
	@echo 'start compil vme mac'
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BIN_SERVER_MAC}${BINARY_VM}/vm cmd/vm/main.go

copy_host_files:
	@echo 'start copy host files'

copy_vm_files:
	@echo 'start copy vm files'

package_host:
	@echo 'start package host'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@for subdir in `ls ${BIN_OUT_HOST}`; do mkdir -p ${QINIU_DIST_DIR}$${platform}; done

	@cd ${BIN_OUT} && \
		for platform in `ls ./`; \
			do  cd $${platform} && \
				zip -r ${QINIU_DIST_DIR}$${platform}/${BINARY_HOST}.zip "${BINARY_HOST}" && \
				md5sum ${QINIU_DIST_DIR}$${platform}/${BINARY_HOST}.zip | awk '{print $$1}' | \
					xargs echo > ${QINIU_DIST_DIR}$${platform}/${BINARY_HOST}.zip.md5 && \
				cd ..; \
			done

package_vm:
	@echo 'start package vm'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@for subdir in `ls ${BIN_OUT}`; do mkdir -p ${QINIU_DIST_DIR}$${platform}; done

	@cd ${BIN_OUT} && \
		for platform in `ls ./`; \
			do  cd $${platform} && \
				zip -r ${QINIU_DIST_DIR}$${platform}/${BINARY_VM}.zip "${BINARY_VM}" && \
				md5sum ${QINIU_DIST_DIR}$${platform}/${BINARY_VM}.zip | awk '{print $$1}' | \
					xargs echo > ${QINIU_DIST_DIR}$${platform}/${BINARY_VM}.zip.md5 && \
				cd ..; \
			done

upload_to:
	@echo 'upload...'
	@find ${QINIU_DIR} -name ".DS_Store" -type f -delete
	@qshell qupload2 --src-dir=${QINIU_DIR} --bucket=download --thread-count=10 --log-file=qshell.log \
					 --skip-path-prefixes=zd,ztf,zmanager,deeptest --rescan-local --overwrite --check-hash

fmt: ## fmt
	gofmt -s -w .
	goimports -w .
	@echo gofmt -l
	@OUTPUT=`gofmt -l . 2>&1`; \
	if [ "$$OUTPUT" ]; then \
		echo "gofmt must be run on the following files:"; \
        echo "$$OUTPUT"; \
        exit 1; \
    fi

golint: ## lint

	@echo golangci-lint run --skip-files \".*test.go\" -v ./...
	@OUTPUT=`command -v golangci-lint >/dev/null 2>&1 && golangci-lint run --skip-files ".*test.go"  -v ./... 2>&1`; \
	if [ "$$OUTPUT" ]; then \
		echo "golint errors:"; \
		echo "$$OUTPUT"; \
	fi

default: fmt golint ## fmt code
