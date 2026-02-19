package encryption

import (
	"bytes"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	password := "test-password"
	plaintext := []byte("hello world")

	encrypted, err := Encrypt(plaintext, password)
	if err != nil {
		t.Fatalf("failed to encrypt: %v", err)
	}

	decrypted, err := Decrypt(encrypted, password)
	if err != nil {
		t.Fatalf("failed to decrypt: %v", err)
	}

	if !bytes.Equal(plaintext, decrypted) {
		t.Fatalf("decrypted text does not match: expected %v, got %v", string(plaintext), string(decrypted))
	}
}

func TestDecryptFailWithWrongPassword(t *testing.T) {
	password := "correct-password"
	wrongPassword := "wrong-password"
	plaintext := []byte("secret")

	encrypted, _ := Encrypt(plaintext, password)
	_, err := Decrypt(encrypted, wrongPassword)
	if err == nil {
		t.Fatal("expected decryption to fail with wrong password, but it succeeded")
	}
}
