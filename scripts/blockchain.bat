@echo off
REM Navicate to the desired directory
REM cd ../

for /L %%i in (1, 1, 5) do (
    start cmd /k "go run ./cmd/main/main.go"
)

REM Keep the scriptwindow open to view output
pause
