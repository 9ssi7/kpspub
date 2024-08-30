package kpspublic

const ApiUrl string = "https://tckimlik.nvi.gov.tr/Service/KPSPublicv2.asmx"

// Verify verifies the given information with the Turkish Identity Number
type VerifyConfig struct {

	// TCIdentityNumber is the Turkish Identity Number
	TCIdentityNumber string

	// FirstName is the first name of the person
	FirstName string

	// LastName is the last name of the person
	LastName string

	// BirthYear is the birth year of the person
	BirthYear int

	// BirthMonth is the birth month of the person
	BirthMonth int

	// BirthDay is the birthday of the person
	BirthDay int

	// SerialNumber is the serial number of the ID card
	SerialNumber string
}
