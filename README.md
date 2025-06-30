# go-nwc

[NWC](https://nwc.dev) Client Go Library

## Features

- Encryption: NIP-44
- Get wallet service info
- Supported methods:
  - Get wallet/node info (`get_info`)
  - Create lightning invoices (`make_invoice`)
  - Pay BOLT11 invoices (`pay_invoice`)
  - Check wallet balance (`get_balance`)
  - Look up invoice/payment info (`lookup_invoice`)
  - List wallet transactions (`list_transactions`)

## Getting Started

`go get https://github.com/rolznz/go-nwc`

### Example

> See [/example/main.go](/example/main.go)

```bash
NWC_URI="nostr+walletconnect://..." go run example/main.go
```

## Development

### Running Tests

Make sure you have a connection that has a balance of 21 sats, and can pay itself (e.g. an Alby Hub sub-wallet)

```bash
NWC_URI="nostr+walletconnect://..." go test ./...
```
