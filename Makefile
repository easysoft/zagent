VERSION=1.0.0
PROJECT=zv-server
PACKAGE=${PROJECT}-${VERSION}
BINARY=zv-server
BIN_DIR=bin
BIN_ZIP_DIR=${BIN_DIR}/zip/${PROJECT}/${VERSION}/
BIN_ZIP_RELAT=../../../zip/${PROJECT}/${VERSION}/
BIN_OUT=${BIN_DIR}/${PROJECT}/${VERSION}/
BIN_WIN64=${BIN_OUT}win64/zv-server/
BIN_WIN32=${BIN_OUT}win32/zv-server/
BIN_LINUX=${BIN_OUT}linux/zv-server/
BIN_MAC=${BIN_OUT}mac/zv-server/
QINIU_DIR=/Users/aaron/work/zentao/qiniu/
QINIU_DIST_DIR=${QINIU_DIR}${PROJECT}/${VERSION}/

default: prepare_res compile_all copy_files package

win64: prepare_res compile_win64 copy_files package
win32: prepare_res compile_win32 copy_files package
linux: prepare_res compile_linux copy_files package
mac: prepare_res compile_mac copy_files package
upload: upload_to

prepare_res:
	@echo 'start prepare res'
	@go-bindata -o=res/server/res.go -pkg=serverRes res/server/...
	@rm -rf ${BIN_DIR}

compile_all: compile_win64 compile_win32 compile_linux compile_mac

compile_win64:
	@echo 'start compile win64'
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ${BIN_WIN64}zv-server.exe cmd/server/main.go

compile_win32:
	@echo 'start compile win32'
	@CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o ${BIN_WIN32}zv-server.exe cmd/server/main.go

compile_linux:
	@echo 'start compile linux'
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BIN_LINUX}zv-server cmd/server/main.go

compile_mac:
	@echo 'start compile mac'
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BIN_MAC}zv-server cmd/server/main.go

copy_files:
	@echo 'start copy files'
	@cp -r {cmd/server/server.yml,cmd/server/perms.yml,cmd/server/rbac_model.conf} bin
	@for subdir in `ls ${BIN_OUT}`; \
	    do cp -r {bin/server.yml,bin/perms.yml,bin/rbac_model.conf} "${BIN_OUT}$${subdir}/zv-server"; done

package:
	@echo 'start package'
	@find . -name .DS_Store -print0 | xargs -0 rm -f
	@for subdir in `ls ${BIN_OUT}`; do mkdir -p ${BIN_DIR}/zip/${PROJECT}/${VERSION}/$${subdir}; done

	@cd ${BIN_OUT} && \
		for subdir in `ls ./`; do cd $${subdir} && zip -r ${BIN_ZIP_RELAT}$${subdir}/${BINARY}.zip "${BINARY}" && cd ..; done

upload_to:
	@echo 'upload...'
	@find ${QINIU_DIR} -name ".DS_Store" -type f -delete
	@qshell qupload2 --src-dir=${QINIU_DIR} --bucket=download --thread-count=10 --log-file=qshell.log \
					 --skip-path-prefixes=zd,ztf,zmanager --rescan-local --overwrite --check-hash
