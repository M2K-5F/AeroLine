@echo off
swag init -q -g main.go --output ./docs
if %ERRORLEVEL% neq 0 (
    echo.
    echo  [SWAGGER ERROR]
    exit /b %ERRORLEVEL%
)
go build -o tmp\myapp.exe .
if %ERRORLEVEL% neq 0 (
    echo.
    echo  [COMPILE ERROR]
    exit /b %ERRORLEVEL%
)
if exist tmp\myapp.exe (
    tmp\myapp.exe
)