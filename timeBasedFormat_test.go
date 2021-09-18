package rand

import (
	"testing"
	"time"
)

func TestTimeBasedFormat(t *testing.T) {
	got1 := TimeBasedFormat("test-{datetime}-{hex}", WithTimeFormat(time.RFC3339), WithHexLen(8))
	t.Logf("got = %v", got1)

	got2 := TimeBasedFormat("test-{timestamp}-{hex}", WithTimeFormat(time.RFC3339), WithHexLen(8))
	t.Logf("got = %v", got2)

	got3 := TimeBasedFormat("test-{timestamp_nano}-{hex}", WithTimeFormat(time.RFC3339), WithHexLen(8))
	t.Logf("got = %v", got3)
}
