# zeromq-pubsub-example

## Running

in separate consoles:

- in `go-publisher`: `go run .`
  - on Windows: install the dependencies as described below, and then `..\scripts\go_build.bat`
    to build. Then, run `..\scripts\go_run.bat` to launch the compiled program.
- in `go-subscriber`: `go run .` (more than one can be started)
  - on Windows: same as for `go-publisher`
- alternatively, start the python publisher in `py-publisher`:
  - install the dependencies: `pip3 install -r requirements.txt`
  - run: `python3 main.go`

### In Docker

```bash
docker-compose up --build --force-recreate
```

## Dependencies

### OSX

- `brew install zeromq libsodium czmq`

### Windows

- install [vcpkg](https://github.com/microsoft/vcpkg#quick-start-windows)
- set the environment variable `VCPKG_ROOT` to the bootstrapped vcpkg location 
- add the `VCPKG_ROOT` location to `PATH`
- install MinGW: `choco install /y mingw`
- [install](https://github.com/zeromq/goczmq/issues/229#issuecomment-1019070347) czmq: `vcpkg install --triplet=x64-windows czmq`
  - rename `%VCPKG_ROOT%\installed\x64-windows\lib\libzmq-mt-4_3_4.lib` to `...\libzmq.lib` for CGO to find it
- install libsodium: `vcpkg install --triplet=x64-windows libsodium`