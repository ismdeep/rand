package rand

import "testing"

func TestUsername(t *testing.T) {
	username := Username()
	t.Logf("username = %v", username)
}
