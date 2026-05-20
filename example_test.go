package kpspub_test

import (
	"context"
	"fmt"

	"github.com/salihdev0/kpspub"
)

func ExampleVerify() {
	ok, err := kpspub.Verify(context.Background(), kpspub.VerifyConfig{
		TCIdentityNumber: "10000000146",
		FirstName:        "AHMET",
		LastName:         "YILMAZ",
		BirthYear:        1990,
		BirthMonth:       1,
		BirthDay:         2,
		SerialNumber:     "A12B34567",
	})
	if err != nil {
		// Handle transport, context, or SOAP service errors.
		return
	}

	fmt.Println(ok)
}

func ExampleVerifyWithNewID() {
	ok, err := kpspub.VerifyWithNewID(context.Background(), kpspub.VerifyConfig{
		TCIdentityNumber: "10000000146",
		FirstName:        "AHMET",
		LastName:         "YILMAZ",
		BirthYear:        1990,
		BirthMonth:       1,
		BirthDay:         2,
		SerialNumber:     "A12B34567",
	})
	if err != nil {
		return
	}

	fmt.Println(ok)
}

func ExampleVerifyWithOldID() {
	ok, err := kpspub.VerifyWithOldID(context.Background(), kpspub.VerifyConfig{
		TCIdentityNumber: "10000000146",
		FirstName:        "AHMET",
		LastName:         "YILMAZ",
		BirthYear:        1990,
		BirthMonth:       1,
		BirthDay:         2,
		SerialNumber:     "ABC123",
	})
	if err != nil {
		return
	}

	fmt.Println(ok)
}
