@echo on

:start

ping /n 3 114.114.114.114 | findstr "TTL=" && goto next || goto start

:next
cmd /k "cd /d C:\Users\jenkins\dev\project\zagent && git pull && go get all && go run cmd\agent-vm\main.go -t vm"

pause
