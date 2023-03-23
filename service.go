package KPSPublic

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

// VerifyNewId verifies the given information with the Turkish Identity Number
type VerifyNewIdConfig struct {

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

	// WalletSerialNumber is the serial number of the new wallet
	SerialNumber string
}

// VerifyOldId verifies the given information with the Turkish Identity Number
type VerifyOldIdConfig struct {

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

	// WalletSerialNumber is the serial number of the old wallet
	WalletSerialNumber string

	// WalletNumber is the number of the old wallet
	WalletNumber string
}

// Service is the interface for the KPSPublic service
type Service interface {

	// Verify verifies the given information with the Turkish Identity Number
	// and returns true if the information is valid
	// try to use VerifyNewId or VerifyOldId instead
	Verify(config VerifyConfig) (bool, error)

	// VerifyNewId verifies the given information with the new Turkish Identity Number
	VerifyNewId(config VerifyNewIdConfig) (bool, error)

	// VerifyOldId verifies the given information with the old Turkish Identity Number
	VerifyOldId(config VerifyOldIdConfig) (bool, error)
}

type service struct {
	url string
}

// New returns a new KPSPublic service
func New() Service {
	return &service{
		url: "https://tckimlik.nvi.gov.tr/Service/KPSPublicv2.asmx",
	}
}
