# Cosmos Remote Signer (PoC)

This is a minimal remote signer prototype for CometBFT.

## Features
- TCP JSON API
- Vote signing
- Vote extension support (stub)
- Basic slashing protection logic

## Important
This is NOT production-ready.

Missing:
- full CometBFT priv_validator protocol
- protobuf parsing
- secure key storage
- full slashing protection

## Run

```bash
export SIGNER_PRIVATE_KEY="BASE64_KEY"
go run main.go signer.go types.go
