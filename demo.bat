@echo off

set dir=%~dp0

REM prepare the python dependencies
cd %dir%\py-publisher
pip install -r requirements.txt
if %errorlevel% NEQ 0 exit /b %errorlevel%

REM build the go subscriber
REM note: zmq windows dependencies must be installed beforehand
cd %dir%\go-subscriber
if not exist main.exe (
    call ..\scripts\go_build.bat
    if %errorlevel% NEQ 0 exit /b %errorlevel%
)

REM run the go subscribers in the background
start /b ..\scripts\go_run.bat
start /b ..\scripts\go_run.bat

REM run the python publisher in the foreground
cd %dir%\py-publisher
python main.py
if %errorlevel% NEQ 0 exit /b %errorlevel%

cd %dir%