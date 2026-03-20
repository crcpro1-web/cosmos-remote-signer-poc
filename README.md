# Cosmos Remote Signer (PoC)

This is a minimal remote signer prototype for CometBFT.

## Features

* TCP JSON API
* Vote signing
* Vote extension support (stub)
* Basic slashing protection logic

## Architecture

Client → TCP → Remote Signer → Signature Response

The signer:

* validates incoming requests
* enforces slashing protection rules (prevents double-signing)
* signs vote bytes using ed25519
* returns signature + public key

## API

### Request

```json
{
  "type": "sign_vote",
  "sign_bytes": "base64",
  "height": 1,
  "round": 0,
  "step": "precommit"
}
```

### Response

```json
{
  "status": "ok",
  "signature": "<base64_signature>",
  "pub_key": "<base64_public_key>"
}
```

## Important

This is **NOT production-ready**.

### Missing:

* full CometBFT priv_validator protocol
* protobuf parsing
* secure key storage
* full slashing protection
* network security (TLS / authentication)

> This PoC demonstrates how a custom signer can safely replace TMKMS by supporting vote extensions and enforcing strict slashing protection.

## Run

```bash
export SIGNER_PRIVATE_KEY="BASE64_KEY"
go run main.go signer.go types.go
```

## Test

```bash
echo '{"type":"sign_vote","sign_bytes":"dGVzdA==","height":10,"round":0,"step":"precommit"}' | nc localhost 26659
```
## ⚠️ Security Warning / Disclaimer

This repository contains a Proof-of-Concept (PoC) implementation of a Cosmos/CometBFT remote signer.

**Important notes:**

- **Do not use real validator keys** in this repository.
- `.pem`, `.der`, and `state.json` files are excluded via `.gitignore` and should **never be committed**.
- This code is **not production-ready**:
  - No TLS or authentication
  - No secure key storage
  - No full slashing protection
- Use this PoC only for **testing, learning, or demonstration purposes**.

No full slashing protection

Use this PoC only for testing, learning, or demonstration purposes.
