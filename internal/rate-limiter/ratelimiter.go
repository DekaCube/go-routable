package ratelimiter

import (
	"math/rand/v2"
	"regexp"
	"strconv"
	"sync"
	"time"
)

const bucketRefilDuration = 60

// example rate limit limit limit header
// "fetch";r=263;t=0
var limitHeaderExpr = regexp.MustCompile(`(?P<requestype>fetch|create);r=(?P<budget>\d+)`)

type Limiter struct {
	fetchCurrentBudget int64
	fetchReserveLimit  int64
	fetchMaxLimit      int64
	fetchResetAt       time.Time

	createCurrentBudget int64
	createReserveLimit  int64
	createMaxLimit      int64
	createResetAt       time.Time

	lock sync.RWMutex
}

func NewRateLimiter(fetchReserveLimit, fetchMaxLimit, createReserveLimit, createMaxLimit int64) *Limiter {
	return &Limiter{
		fetchCurrentBudget: fetchMaxLimit,
		fetchReserveLimit:  fetchReserveLimit,
		fetchMaxLimit:      fetchMaxLimit,
		fetchResetAt:       time.Now().Add(bucketRefilDuration * time.Second),

		createCurrentBudget: createMaxLimit,
		createReserveLimit:  createReserveLimit,
		createMaxLimit:      createMaxLimit,
		createResetAt:       time.Now().Add(bucketRefilDuration * time.Second),
	}
}

func (l *Limiter) Update(s string) error {
	l.lock.Lock()
	defer l.lock.Unlock()

	matches := limitHeaderExpr.FindStringSubmatch(s)
	if len(matches) < 3 {
		return NewErrUnparsableRateLimitHeader(s)
	}

	requestType := matches[limitHeaderExpr.SubexpIndex("requestype")]
	budgetStr := matches[limitHeaderExpr.SubexpIndex("budget")]

	budget, err := strconv.ParseInt(budgetStr, 10, 64)
	if err != nil {
		return NewErrUnparsableRateLimitHeader(s)
	}
	expiry := time.Now().Add(bucketRefilDuration * time.Second)

	switch requestType {
	case "fetch":
		l.fetchResetAt = expiry
		l.fetchCurrentBudget = budget
	case "create":
		l.createResetAt = expiry
		l.createCurrentBudget = budget
	default:
		return NewErrUnparsableRateLimitHeader(s)
	}

	return nil
}

func jitter(max time.Duration) time.Duration {
	return time.Duration(rand.N(int64(max)))
}
