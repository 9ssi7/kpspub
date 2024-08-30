package kpspub

import (
	"context"
	"strconv"
)

func VerifyWithOldID(ctx context.Context, config VerifyConfig) (bool, error) {
	return makeRequest(ctx, getBodyForOldID(config))
}

func getBodyForOldID(config VerifyConfig) string {
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
	<CuzdanSeri>` + config.SerialNumber[:3] + `</CuzdanSeri>
	<CuzdanNo>` + config.SerialNumber[3:6] + `</CuzdanNo>
	</KisiVeCuzdanDogrula>
	</soap:Body>
	</soap:Envelope>`
}
