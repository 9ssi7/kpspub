package kpspub

import "context"

// Verify verifies identity card information with KPSPublic.
//
// It first calls VerifyWithNewID using the full SerialNumber as TCKKSeriNo. If
// that returns false without an error, it calls VerifyWithOldID using the
// legacy serial split. The returned bool reports whether either service check
// accepted the supplied data.
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
