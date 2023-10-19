# Cloud Services

This repository contains gRPC services for the fuks App

## Command Line Interface

```bash
# Start the server
export GOOGLE_APPLICATION_CREDENTIALS=credentials.json
go run cmd/server/server.go

# Start the client
go run cmd/cli/cli.go [get_events get_projects get_karlsruher_transfers]
```

## Deploy a new release

Prepare a new release by following these steps:

1. Update the changelog in `CHANGELOG.md`
2. Update dependencies `go get -a all`
3. Commit changes `git commit -am "Release vX.X.X"`
4. Push changes `git push`
5. Create a new git tag:
    1. `git tag -a vX.X.X -m "Release vX.X.X"`
    2. `git push origin vX.X.X`
6. Merge `main` branch into `stable` branch
