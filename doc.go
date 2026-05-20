// Package kpspub provides a small Go client for the Turkish KPSPublic SOAP
// identity card verification service.
//
// The package verifies a Turkish identity number together with name, birth date,
// and identity card serial information by calling the public NVI KPS endpoint.
// Use Verify for the common flow: it first tries the newer Turkish ID card
// serial format and then falls back to the legacy identity card format.
//
// Callers should pass a context with an appropriate timeout or cancellation
// policy. The package sends the provided personal data to the remote government
// service and does not perform local identity validation.
package kpspub
