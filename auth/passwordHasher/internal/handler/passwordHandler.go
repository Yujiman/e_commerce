package handler

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/Yujiman/e_commerce/auth/passwordHasher/internal/config"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"golang.org/x/crypto/argon2"
)

var (
	// ErrInvalidHash in returned by ComparePasswordAndHash if the provided
	// hash isn't in the expected format.
	ErrInvalidHash = errors.New("argon2id: hash is not in the correct format")

	// ErrIncompatibleVersion in returned by ComparePasswordAndHash if the
	// provided hash was created using a different version of Argon2.
	ErrIncompatibleVersion = errors.New("argon2id: incompatible version of argon2")
)

var mu = sync.Mutex{}

func CreateHash(password string, params *config.Argon2Params) (hash string, err error) {
	salt, err := generateRandomBytes(params.SaltLength)
	if err != nil {
		return "", status.Error(codes.Code(500), "PasswordHasher:CreateHash err="+err.Error())
	}

	mu.Lock()
	key := argon2.IDKey([]byte(password), salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)
	mu.Unlock()

	base64Salt := base64.RawStdEncoding.EncodeToString(salt)
	base64Key := base64.RawStdEncoding.EncodeToString(key)

	hash = fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, params.Memory, params.Iterations, params.Parallelism, base64Salt, base64Key,
	)

	return hash, nil
}

func Validate(password, hash string) (match bool, err error) {
	match, _, err = checkHash(password, hash)
	return match, err
}

// CheckHash is like Validate, except it also returns the params that the hash was
// created with. This can be useful if you want to update your hash params over time (which you
// should).
func checkHash(password, hash string) (match bool, params *config.Argon2Params, err error) {
	params, salt, key, err := decodeHash(hash)
	if err != nil {
		return false, nil, err
	}

	mu.Lock()
	otherKey := argon2.IDKey([]byte(password), salt, params.Iterations, params.Memory, params.Parallelism, params.KeyLength)
	mu.Unlock()

	keyLen := int32(len(key))
	otherKeyLen := int32(len(otherKey))

	if subtle.ConstantTimeEq(keyLen, otherKeyLen) == 0 {
		return false, params, nil
	}
	if subtle.ConstantTimeCompare(key, otherKey) == 1 {
		return true, params, nil
	}
	return false, params, nil
}

// decodeHash expects a hash created from this package, and parses it to return the params used to
// create it, as well as the salt and key (password hash).
func decodeHash(hash string) (params *config.Argon2Params, salt, key []byte, err error) {
	vals := strings.Split(hash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	params = &config.Argon2Params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &params.Memory, &params.Iterations, &params.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	params.SaltLength = uint32(len(salt))

	key, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	params.KeyLength = uint32(len(key))

	return params, salt, key, nil
}

func generateRandomBytes(len uint32) ([]byte, error) {
	bytes := make([]byte, len)
	_, err := rand.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
