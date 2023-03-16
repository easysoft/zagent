@echo off
title zagent-install

:: 开始获取管理员权限
setlocal
set uac=~uac_permission_tmp_%random%
md "%SystemRoot%\system32\%uac%" 2>nul
if %errorlevel%==0 ( rd "%SystemRoot%\system32\%uac%" >nul 2>nul ) else (
    echo set uac = CreateObject^("Shell.Application"^)>"%temp%\%uac%.vbs"
    echo uac.ShellExecute "%~s0","","","runas",1 >>"%temp%\%uac%.vbs"
    echo WScript.Quit >>"%temp%\%uac%.vbs"
    "%temp%\%uac%.vbs" /f
    del /f /q "%temp%\%uac%.vbs" & exit )
endlocal

set Url=https://pkg.qucheng.com/zenagent/app/zagent-vm.exe

set Save=%TEMP%
set AgentDir=C:\Users\%username%\zagent
if not exist %Save% (mkdir %Save%)

for %%a in ("%Url%") do set "FileName=%%~nxa"

if exist %Save%\%FileName% (del %Save%\%FileName%)
if not defined Save set "Save=%cd%"
(
echo Download Wscript.Arguments^(0^),Wscript.Arguments^(1^)
echo Sub Download^(url,target^)
echo   Const adTypeBinary = 1
echo   Const adSaveCreateOverWrite = 2
echo   Dim http,ado
echo   Set http = CreateObject^("Msxml2.ServerXMLHTTP"^)
echo   http.open "GET",url,False
echo   http.send
echo   Set ado = createobject^("Adodb.Stream"^)
echo   ado.Type = adTypeBinary
echo   ado.Open
echo   ado.Write http.responseBody
echo   ado.SaveToFile target
echo   ado.Close
echo End Sub)>DownloadFile.vbs

echo downloading：%FileName%
DownloadFile.vbs "%Url%" "%Save%\%FileName%"
::del DownloadFile.vbs

IF NOT EXIST %AgentDir% (
		MD %AgentDir%
	)
echo Move %FileName% to %AgentDir%\
move %Save%\%FileName% %AgentDir%\%FileName%

:beginInstall
del DownloadFile.vbs

echo %AgentDir%\%FileName% -p 55201 -secret %1 -s %2
if exist "%AgentDir%\%FileName%" (echo start zagent：%AgentDir%\%FileName% & start cmd /k ""%AgentDir%\%FileName%" -p 55201 -secret %1 -s %2") else (echo error：%AgentDir%\%FileName% not exist！& exit)

echo Add to Start Menu

(
echo @echo on
echo cmd /k "cd /d %AgentDir% && %FileName% -p 55201 -secret %1 -s %2"
echo pause)>zagent-vm-start.bat

echo move zagent-vm-start.bat to "C:\ProgramData\Microsoft\Windows\Start Menu\Programs\Startup\"
move .\zagent-vm-start.bat "C:\ProgramData\Microsoft\Windows\Start Menu\Programs\Startup\zagent-vm-start.bat"


echo install success
