package config

import (
	"errors"
	"os"
	"strconv"
)

type Argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

func GetDefaultArgon2idParams() *Argon2Params {
	return &Argon2Params{
		Memory:      64 * 1024,
		Iterations:  1,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}
}

func GetArgon2idParams() (*Argon2Params, error) {
	memory, err := strconv.Atoi(os.Getenv("ARGON2_MEMORY_MB"))
	if err != nil {
		return nil, errors.New("ARGON2_MEMORY_MB environment not valid")
	}
	memory32 := uint32(memory) * 1024

	iterations, err := strconv.Atoi(os.Getenv("ARGON2_ITERATIONS"))
	if err != nil {
		return nil, errors.New("ARGON2_ITERATIONS environment not valid")
	}

	parallelism, err := strconv.Atoi(os.Getenv("ARGON2_PARALLELISM"))
	if err != nil {
		return nil, errors.New("ARGON2_PARALLELISM environment not valid")
	}

	saltL, err := strconv.Atoi(os.Getenv("ARGON2_SALT_LENGTH"))
	if err != nil {
		return nil, errors.New("ARGON2_SALT_LENGTH environment not valid")
	}

	keyL, err := strconv.Atoi(os.Getenv("ARGON2_KEY_LENGTH"))
	if err != nil {
		return nil, errors.New("ARGON2_KEY_LENGTH environment not valid")
	}

	return &Argon2Params{
		Memory:      memory32,
		Iterations:  uint32(iterations),
		Parallelism: uint8(parallelism),
		SaltLength:  uint32(saltL),
		KeyLength:   uint32(keyL),
	}, nil
}
