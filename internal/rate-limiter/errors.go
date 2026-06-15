package ratelimiter

import "fmt"

type ErrUnparsableRateLimitHeader struct {
	Header string
}

func NewErrUnparsableRateLimitHeader(header string) *ErrUnparsableRateLimitHeader {
	return &ErrUnparsableRateLimitHeader{
		Header: header,
	}
}

func (e *ErrUnparsableRateLimitHeader) Error() string {
	return fmt.Sprintf("unparsable rate limit header: %s", e.Header)
}
