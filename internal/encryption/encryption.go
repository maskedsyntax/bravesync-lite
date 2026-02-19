package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/maskedsyntax/bravesync-lite/internal/utils"
	"golang.org/x/crypto/argon2"
)

const (
	SaltLen    = 16
	NonceLen   = 12
	KeyLen     = 32
	TagLen     = 16
	Time       = 1
	Memory     = 64 * 1024
	Threads    = 4
)

// Encrypt encrypts data using AES-256-GCM with Argon2id.
func Encrypt(plaintext []byte, password string) ([]byte, error) {
	salt := make([]byte, SaltLen)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err)
	}

	key := deriveKey([]byte(password), salt)
	defer utils.ZeroMem(key)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create gcm: %w", err)
	}

	nonce := make([]byte, NonceLen)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}

	// Output: <salt><nonce><ciphertext><tag>
	// Seal appends the tag to the ciphertext.
	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)

	result := append(salt, nonce...)
	result = append(result, ciphertext...)

	return result, nil
}

// Decrypt decrypts data using AES-256-GCM with Argon2id.
func Decrypt(data []byte, password string) ([]byte, error) {
	if len(data) < SaltLen+NonceLen+TagLen {
		return nil, fmt.Errorf("invalid data format")
	}

	salt := data[:SaltLen]
	nonce := data[SaltLen : SaltLen+NonceLen]
	ciphertext := data[SaltLen+NonceLen:]

	key := deriveKey([]byte(password), salt)
	defer utils.ZeroMem(key)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create gcm: %w", err)
	}

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt: %w", err)
	}

	return plaintext, nil
}

func deriveKey(password []byte, salt []byte) []byte {
	return argon2.IDKey(password, salt, Time, Memory, Threads, KeyLen)
}
