# KPSPublic Go Client

KPSPublic Go Client is a small Go package for verifying Turkish identity card
information against the public NVI KPS SOAP service.

The package sends the Turkish identity number, name, birth date, and identity
card serial information to the KPSPublic endpoint and returns whether the remote
service accepted the supplied data.

## Features

- Simple `context.Context`-aware API.
- Support for both new Turkish ID card serial numbers and legacy identity card
  serial values.
- No third-party dependencies.

## Requirements

- Go 1.23 or newer.
- Network access to the KPSPublic SOAP endpoint.

## Installation

```bash
go get github.com/salihdev0/kpspub
```

## Usage

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/salihdev0/kpspub"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	isVerified, err := kpspub.Verify(ctx, kpspub.VerifyConfig{
		TCIdentityNumber: "10000000146",
		FirstName:        "AHMET",
		LastName:         "YILMAZ",
		BirthYear:        1990,
		BirthMonth:       1,
		BirthDay:         2,
		SerialNumber:     "A12B34567",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(isVerified)
}
```

`Verify` tries the new ID card format first. If the remote service returns
`false` without an error, it falls back to the legacy card format.

## Configuration

| Field | Description |
| --- | --- |
| `TCIdentityNumber` | 11-digit Turkish identity number. |
| `FirstName` | First name as registered with NVI. |
| `LastName` | Last name as registered with NVI. |
| `BirthYear` | Birth year. |
| `BirthMonth` | Birth month, from `1` to `12`. |
| `BirthDay` | Birth day, from `1` to `31`. |
| `SerialNumber` | Identity card serial number. For new cards, pass the full TCKK serial number. For legacy cards, pass at least six characters: the first three are used as `CuzdanSeri`, and the next three are used as `CuzdanNo`. |

## API

### `Verify`

```go
ok, err := kpspub.Verify(ctx, config)
```

Use this as the default verification function. It tries both supported serial
number formats.

### `VerifyWithNewID`

```go
ok, err := kpspub.VerifyWithNewID(ctx, config)
```

Use this when you know the serial number belongs to a new Turkish ID card. The
`SerialNumber` value is sent as `TCKKSeriNo`.

### `VerifyWithOldID`

```go
ok, err := kpspub.VerifyWithOldID(ctx, config)
```

Use this when you know the serial number belongs to a legacy identity card.
`SerialNumber` must contain at least six characters. Shorter values will panic
because the legacy request body splits the value into `CuzdanSeri` and
`CuzdanNo`.

## Errors and limitations

- `error` reports request creation, network, context, or response body read
  failures.
- A `false` result means the KPSPublic service did not verify the supplied data.
- The package does not validate identity data locally before sending the request.
- The package uses `http.DefaultClient`; pass a context with a timeout or
  cancellation policy.
- The data handled by this package is sensitive personal data. Avoid logging
  raw request values or responses.

## Contributing

Contributions are always welcome!

## License

[MIT](https://choosealicense.com/licenses/mit/)
