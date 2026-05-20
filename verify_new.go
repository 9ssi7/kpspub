package kpspub

import (
	"context"
	"strconv"
)

// VerifyWithNewID verifies identity information using the new Turkish ID card
// serial number format.
//
// The SerialNumber field is sent as TCKKSeriNo. Use Verify instead when the
// caller does not know whether the supplied serial number belongs to a new or
// legacy card.
func VerifyWithNewID(ctx context.Context, config VerifyConfig) (bool, error) {
	return makeRequest(ctx, getBodyForNewID(config))
}

func getBodyForNewID(config VerifyConfig) string {
	return `<?xml version="1.0" encoding="utf-8"?>
	<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	<soap:Body>
	<KisiVeCuzdanDogrula xmlns="http://tckimlik.nvi.gov.tr/WS">
	<TCKimlikNo>` + config.TCIdentityNumber + `</TCKimlikNo>
	<Ad>` + config.FirstName + `</Ad>
	<Soyad>` + config.LastName + `</Soyad>
	<SoyadYok>false</SoyadYok>
	<DogumGun>` + strconv.Itoa(config.BirthDay) + `</DogumGun>
	<DogumGunYok>false</DogumGunYok>
	<DogumAy>` + strconv.Itoa(config.BirthMonth) + `</DogumAy>
	<DogumAyYok>false</DogumAyYok>
	<DogumYil>` + strconv.Itoa(config.BirthYear) + `</DogumYil>
	<TCKKSeriNo>` + config.SerialNumber + `</TCKKSeriNo>
	</KisiVeCuzdanDogrula>
	</soap:Body>
	</soap:Envelope>`
}
