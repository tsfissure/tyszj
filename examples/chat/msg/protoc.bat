cd %~dp0
@echo off

for %%i in (proto/*.proto) do (
    echo %%i
    protoc -I=proto --gofast_out=pbgo proto/%%i
)

pause
