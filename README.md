# Cloud Services

This repository contains gRPC services for the fuks App

## Deploy a new release

Prepare a new release by following these steps:

1. Update the changelog in `CHANGELOG.md`
2. Create a new git tag:
    1. `git tag -a vX.X.X -m "Release vX.X.X"`
    2. `git push origin vX.X.X`
3. Merge `main` branch into `stable` branch
