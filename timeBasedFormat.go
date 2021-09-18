package rand

import (
	"fmt"
	"strings"
	"time"
)

type Option struct {
	TimeFormat *string
	HexLen     *int
}

func WithTimeFormat(f string) *Option {
	return &Option{TimeFormat: &f}
}

func WithHexLen(n int) *Option {
	return &Option{HexLen: &n}
}

func TimeBasedFormat(f string, opts ...*Option) string {
	timeFormat := "2006-01-02-150405"
	hexLen := 6

	for _, opt := range opts {
		if opt.TimeFormat != nil {
			timeFormat = *opt.TimeFormat
		}
		if opt.HexLen != nil {
			hexLen = *opt.HexLen
		}
	}

	s := f

	s = strings.ReplaceAll(s, "{datetime}", time.Now().Format(timeFormat))
	s = strings.ReplaceAll(s, "{timestamp}", fmt.Sprintf("%v", time.Now().Unix()))
	s = strings.ReplaceAll(s, "{timestamp_nano}", fmt.Sprintf("%v", time.Now().UnixNano()))
	s = strings.ReplaceAll(s, "{hex}", HexStr(hexLen))

	return s
}
