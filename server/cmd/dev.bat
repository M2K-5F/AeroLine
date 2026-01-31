@echo off
swag init -g main.go --output ./docs
go build -o tmp\myapp.exe .
if exist tmp\myapp.exe (
    tmp\myapp.exe
)