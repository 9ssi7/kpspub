package kpspub

// ApiUrl is the KPSPublic SOAP endpoint used for identity card verification.
const ApiUrl string = "https://tckimlik.nvi.gov.tr/Service/KPSPublicv2.asmx"

// VerifyConfig contains the identity data sent to the KPSPublic verification
// service.
//
// All fields are required. TCIdentityNumber should contain the 11-digit Turkish
// identity number. BirthDay, BirthMonth, and BirthYear should describe the
// person's date of birth. SerialNumber is interpreted as the full TCKK serial
// number for new ID cards, or as a six-character legacy serial value where the
// first three characters are CuzdanSeri and the next three characters are
// CuzdanNo.
type VerifyConfig struct {

	// TCIdentityNumber is the 11-digit Turkish identity number.
	TCIdentityNumber string

	// FirstName is the person's first name as registered with NVI.
	FirstName string

	// LastName is the person's last name as registered with NVI.
	LastName string

	// BirthYear is the person's birth year.
	BirthYear int

	// BirthMonth is the person's birth month, from 1 to 12.
	BirthMonth int

	// BirthDay is the person's birth day, from 1 to 31.
	BirthDay int

	// SerialNumber is the identity card serial number.
	SerialNumber string
}
