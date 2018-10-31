package common

import "time"

func CeilTimeByModule(blockStarts time.Time, mod time.Duration) time.Time {
	interval := mod.Nanoseconds()
	intervalSecs := blockStarts.Unix()
	intervalNano := blockStarts.UnixNano() - intervalSecs*int64(time.Second)
	intervalMilli := intervalNano / interval
	intervalNano = intervalNano % interval
	if intervalNano > 0 {
		intervalMilli++
	}
	intervalMilli *= interval

	return time.Unix(blockStarts.Unix(), intervalMilli*int64(time.Nanosecond))
}