# Cosmos Remote Signer (PoC)

This is a minimal remote signer prototype for CometBFT.

## Features
- TCP JSON API
- Vote signing
- Vote extension support (stub)
- Basic slashing protection logic

## Architecture
Client → TCP → Remote Signer → Signature Response

The signer:
- validates request
- enforces slashing protection rules
- signs vote bytes using ed25519
- returns signature + public key

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
