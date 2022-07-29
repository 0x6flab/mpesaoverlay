package errors

import "errors"

var (
	// ErrAuthentication indicates failure occurred while authenticating.
	ErrAuthentication = errors.New("failed to perform authentication")

	// ErrAuthorization indicates failure occurred while authorizing
	ErrAuthorization = errors.New("failed to perform authorization")

	// ErrInvalidQueryParams indicates invalid query parameters
	ErrInvalidQueryParams = errors.New("invalid query parameters")

	// ErrFailedPublish indicates that publishing message failed.
	ErrFailedPublish = errors.New("failed to publish message")

	// ErrFailedRead indicates that read messages failed.
	ErrFailedRead = errors.New("failed to read messages")

	// ErrFetchHealth indicates that fetching of health check failed.
	ErrFetchHealth = errors.New("failed to fetch health check")

	// ErrCerts indicates error fetching certificates.
	ErrCerts = errors.New("failed to fetch certs data")

	// ErrInvalidCommandID indicates the CommandID is invalid.
	ErrInvalidCommandID = errors.New("invalid command id")
)
