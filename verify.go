package kpspub

import "context"

func Verify(ctx context.Context, cnf VerifyConfig) (bool, error) {
	res, err := VerifyWithNewID(ctx, cnf)
	if err != nil {
		return false, err
	}
	if res {
		return true, nil
	}
	res, err = VerifyWithOldID(ctx, cnf)
	if err != nil {
		return false, err
	}
	return res, nil
}
