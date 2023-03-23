package KPSPublic

import (
	"strconv"
)

func (s *service) VerifyOldId(config VerifyOldIdConfig) (bool, error) {
	return s.makeRequest(s.getOldIdXML(config))
}

func (s *service) getOldIdXML(config VerifyOldIdConfig) string {
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
	<CuzdanSeri>` + config.WalletSerialNumber + `</CuzdanSeri>
	<CuzdanNo>` + config.WalletNumber + `</CuzdanNo>
	</KisiVeCuzdanDogrula>
	</soap:Body>
	</soap:Envelope>`
}
