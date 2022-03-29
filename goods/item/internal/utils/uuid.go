package utils

import uuid "github.com/satori/go.uuid"

func CheckUuid(ids ...string) error {
	for _, id := range ids {
		_, err := uuid.FromString(id)
		if err != nil {
			return err
		}
	}

	return nil
}

func GenerateUuid() uuid.UUID {
	newUuid, _ := uuid.NewV4()
	return uuid.Must(newUuid, nil)
}
