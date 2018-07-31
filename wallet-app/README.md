# kWallet

## Getting Started
Requires docker and docker-compose.

Run `make docker-start`.

http://localhost:3000

## Running the Production Build
Run `make docker-build`.

## Component Storybooks
Components are designed and maintained using [Storybook](https://github.com/storybooks/storybook).
To run the storybook locally, run `make docker-storybook`.

http://localhost:6006

## Component Testing
Components are unit tested with [Jest](https://github.com/facebook/jest).

Run `make docker-test` to run all tests.

## Container Testing
Containers will be end to end tested using *TBD*.

## Wallet Backups

### Edge
Private key encryption, storage and backup made possible with [Edge](https://edgesecure.co/) and
[edge-core-js](https://github.com/Airbitz/edge-core-js).

[Documentation](https://developer.airbitz.co/javascript/)

### Account Recovery Warning
If you do not remember your username or password, and you no longer have PIN access, fingerprint access, or recovery questions, Edge will not be able to reset your password, as all data is securely encrypted.

🐨
