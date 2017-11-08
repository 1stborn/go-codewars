pushd "%~dp0"
set GOPATH=%CD%
set COMPILER_PATH="
if "%GO_HOME%" neq "" (
    if exist "%GO_HOME%\bin\go.exe" (
        set COMPILER_PATH="%GO_HOME%\bin\"
    )
)
cd src
call "%COMPILER_PATH:"=%go" build -o ../MyStrategy.exe  2>../compilation.log
popd
