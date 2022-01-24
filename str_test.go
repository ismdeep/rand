package rand

import (
	"testing"
)

func TestStr(t *testing.T) {
	str := Str(10)
	t.Logf("str = %v", str)
}

func TestPIN(t *testing.T) {
	pin := PIN(6)
	t.Logf("pin = %v", pin)
}

func TestPassword(t *testing.T) {
	password := Password(4, 4, 4)
	t.Logf("password = %v", password)
}

func TestPasswordEasyToRemember(t *testing.T) {
	password := PasswordEasyToRemember(4)
	t.Logf("password easy to remember = %v", password)
}
