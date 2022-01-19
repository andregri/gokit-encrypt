package helpers

import (
	"context"
	"errors"

	"github.com/andregri/gokit-encrypt/utils"
)

var errEmpty = errors.New("secret key or text should not be empty")

// EncryptServiceInstance is the implementation of the interface for the EncryptService
type EncryptServiceInstance struct{}

// Encrypt text with the given string
func (EncryptServiceInstance) Encrypt(_ context.Context, key, text string) (string, error) {
	return utils.EncryptString(key, text), nil
}

//
func (EncryptServiceInstance) Decrypt(_ context.Context, key, text string) (string, error) {
	if key == "" || text == "" {
		return "", errEmpty
	}
	return utils.DecryptString(key, text), nil
}
