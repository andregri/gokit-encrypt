package utils_test

import (
	"testing"

	"github.com/andregri/gokit-encrypt/utils"
)

func TestEncryptDecrypt(t *testing.T) {
	key := "111023043350789514532147"
	message := "I am a message"

	encrypted := utils.EncryptString(key, message)
	decrypted := utils.DecryptString(key, encrypted)

	if message != decrypted {
		t.Errorf("Expected decrypted: %q, got %q\n", message, decrypted)
	}
}
