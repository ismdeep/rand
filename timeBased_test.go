package rand

import "testing"

func TestTimeBased(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TimeBased()
			t.Logf("got = %v", got)
		})
	}
}
