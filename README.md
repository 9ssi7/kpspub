# KPSPublic Go Client

KPSPublic is a SOAP service provided by the Turkish Tax Administration. This is a Go client for KPSPublic.

## Installation

```bash
go get github.com/ssibrahimbas/KPSPublic
```

## Usage

```go
package main

import "github.com/ssibrahimbas/KPSPublic"

func main() {
	kps := KPSPublic.NewService()
	res, err := kps.Verify(KPSPublic.VerifyConfig{
		BirthDay:         <YOUR_BIRTH_DAY>, // 1-31
		BirthMonth:       <YOUR_BIRTH_MONTH>, // 1-12
		BirthYear:        <YOUR_BIRTH_YEAR>, // 1900-2099
		FirstName:        <YOUR_FIRST_NAME>, // "ahmet"
		LastName:         <YOUR_LAST_NAME>, // "yılmaz"
		TCIdentityNumber: <YOUR_TC_IDENTITY_NUMBER>, // "11111111111"
		SerialNumber:     <YOUR_SERIAL_NUMBER>, // "xxxxx"
	})
	if err != nil {
		panic(err)
	}
	println(res) // true or false
}
```

## Contributing

Contributions are always welcome!

## License

[MIT](https://choosealicense.com/licenses/mit/)
