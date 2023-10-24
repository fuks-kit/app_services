# Cloud Services

This repository contains gRPC services for the fuks App

## Command Line Interface

To kickstart your journey with these services, you can use the following commands:

### Start the Server

To initiate the server, set your Google Application Credentials using `credentials.json` and execute:

```bash
go run cmd/server/server.go
```

### Start the Client

Launch the client with one of the following commands:

```bash
go run cmd/cli/cli.go get_events
go run cmd/cli/cli.go get_projects
go run cmd/cli/cli.go get_karlsruher_transfers
```

## Deploy a new release

Prepare a new release by following these steps:

1. Update the changelog in `CHANGELOG.md`
2. Update dependencies `go get -u all`
3. Commit changes `git commit -am "Release vX.X.X"`
4. Push changes `git push`
5. Create a new git tag:
    1. `git tag -a vX.X.X -m "Release vX.X.X"`
    2. `git push origin vX.X.X`
6. Merge `main` branch into `stable` branch

After the release is merged into the `stable` branch, the new release will be automatically deployed by using Google
Cloud Run.

## Generate gRPC Definitions

To generate gRPC definitions, follow these steps:

1. Update APP_DIR Variable
    - In the `proto/Makefile`, update the `APP_DIR` variable to point to the Fuks App directory.

2. Update PROTO_ROOT_DIR Variable
    - If necessary, modify the `PROTO_ROOT_DIR` variable in the `proto/Makefile` to suit your setup.

3. Update gRPC Definitions
    - Make changes to the gRPC definitions in `proto/services.proto` as needed.

4. Generate Code
    - Use the following commands to generate the code:
        - Generate Go code:
          ```bash
          make go
          ```
        - Generate Fuks App code:
          ```bash
          make dart
          ```

These guidelines should help you make the most out of the Fuks Cloud Services repository. Enjoy your journey!
