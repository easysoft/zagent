@echo on

:start

ping /n 3 114.114.114.114 | findstr "TTL=" && goto next || goto start

:next
cmd /k "cd /d C:\Users\jenkins\dev\project\zagent && git fetch --all && git reset --hard origin/main && git pull && go env -w GOPROXY=https://goproxy.io && go get all && go run cmd\agent-vm\main.go -t vm"

pause
