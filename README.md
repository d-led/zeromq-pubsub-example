# zeromq-pubsub-example

## Running

in separate consoles:

- in `go-publisher`: `go run .`
- in `go-subscriber`: `go run .` (more than one can be started)
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
