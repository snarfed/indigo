package countstore

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

const (
	PeriodTotal = "total"
	PeriodDay   = "day"
	PeriodHour  = "hour"
)

// CountStore is an interface for storing incrementing event counts, bucketed into periods.
// It is implemented by MemCountStore and by RedisCountStore.
//
// Period bucketing works on the basis of the current date (as determined mid-call).
// See the `Period*` consts for the available period types.
//
// The "GetCount" and "Increment" methods perform actual counting.
// The "*Distinct" methods have a different behavior:
// "IncrementDistinct" marks a value as seen at least once,
// and "GetCountDistinct" asks how _many_ values have been seen at least once.
//
// The exact implementation and precision of the "*Distinct" methods may vary:
// in the MemCountStore implementation, it is precise (it's based on large maps);
// in the RedisCountStore implementation, it uses the Redis "pfcount" feature,
// which is based on a HyperLogLog datastructure which has probablistic properties
// (see https://redis.io/commands/pfcount/ ).
type CountStore interface {
	GetCount(ctx context.Context, name, val, period string) (int, error)
	Increment(ctx context.Context, name, val string) error
	// TODO: batch increment method
	GetCountDistinct(ctx context.Context, name, bucket, period string) (int, error)
	IncrementDistinct(ctx context.Context, name, bucket, val string) error
}

func periodBucket(name, val, period string) string {
	switch period {
	case PeriodTotal:
		return fmt.Sprintf("%s/%s", name, val)
	case PeriodDay:
		t := time.Now().UTC().Format(time.DateOnly)
		return fmt.Sprintf("%s/%s/%s", name, val, t)
	case PeriodHour:
		t := time.Now().UTC().Format(time.RFC3339)[0:13]
		return fmt.Sprintf("%s/%s/%s", name, val, t)
	default:
		slog.Warn("unhandled counter period", "period", period)
		return fmt.Sprintf("%s/%s", name, val)
	}
}
