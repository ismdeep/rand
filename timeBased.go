package rand

// TimeBased generate with time based
func TimeBased() string {
	return TimeBasedFormat("{datetime}-{hex}")
}
