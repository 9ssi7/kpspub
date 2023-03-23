package KPSPublic

func (s *service) Verify(config VerifyConfig) (bool, error) {
	res1, err1 := s.VerifyNewId(VerifyNewIdConfig{
		TCIdentityNumber: config.TCIdentityNumber,
		FirstName:        config.FirstName,
		LastName:         config.LastName,
		BirthYear:        config.BirthYear,
		BirthMonth:       config.BirthMonth,
		BirthDay:         config.BirthDay,
		SerialNumber:     config.SerialNumber,
	})
	if err1 != nil {
		return false, err1
	}
	if res1 {
		return true, nil
	}
	res2, err2 := s.VerifyOldId(VerifyOldIdConfig{
		TCIdentityNumber:   config.TCIdentityNumber,
		FirstName:          config.FirstName,
		LastName:           config.LastName,
		BirthYear:          config.BirthYear,
		BirthMonth:         config.BirthMonth,
		BirthDay:           config.BirthDay,
		WalletSerialNumber: config.SerialNumber[:3],
		WalletNumber:       config.SerialNumber[3:6],
	})
	if err2 != nil {
		return false, err2
	}
	return res2, nil
}
