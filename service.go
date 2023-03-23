package KPSPublic

type VerifyConfig struct {
	TCIdentityNumber string
	FirstName        string
	LastName         string
	BirthYear        int
	BirthMonth       int
	BirthDay         int
	SerialNumber     string
}
type VerifyNewIdConfig struct {
	TCIdentityNumber string
	FirstName        string
	LastName         string
	BirthYear        int
	BirthMonth       int
	BirthDay         int
	SerialNumber     string
}
type VerifyOldIdConfig struct {
	TCIdentityNumber   string
	FirstName          string
	LastName           string
	BirthYear          int
	BirthMonth         int
	BirthDay           int
	WalletSerialNumber string
	WalletNumber       string
}

type Service interface {
	Verify(config VerifyConfig) (bool, error)
	VerifyNewId(config VerifyNewIdConfig) (bool, error)
	VerifyOldId(config VerifyOldIdConfig) (bool, error)
}

type service struct {
	url string
}

func NewService() Service {
	return &service{
		url: "https://tckimlik.nvi.gov.tr/Service/KPSPublicv2.asmx",
	}
}
