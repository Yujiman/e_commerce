package config

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
