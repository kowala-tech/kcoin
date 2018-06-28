# Notifications service

[![Build Status](http://ci.kowala.io/api/badges/kowala-tech/kcoin/status.svg)](http://ci.kowala.io/kowala-tech/kcoin)

This service monitors the blockchain and send notifications to users

### Send a test transaction using web3

```
web3.personal.unlockAccount("0x007ccffb7916f37f7aeef05e8096ecfbe55afc2f", "")
web3.eth.sendTransaction({
  from: "0x007ccffb7916f37f7aeef05e8096ecfbe55afc2f",
  to: "0x99429f64cf4d5837620dcc293c1a537d58729b68",
  value: 10000000
})
```

# API cli

### Usage

Example:
```
go run cmd/api-cli/main.go -addr localhost:3000 -o register -w 0x99429f64cf4d5837620dcc293c1a537d58729b68 -e your-email@email.com
```
