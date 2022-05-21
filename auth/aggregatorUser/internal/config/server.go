package config

import (
	"errors"
	"os"
	"strconv"
)

func GetServerPort() (int, error) {
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return 0, errors.New("SERVER_PORT environment not valid")
	}

	return port, nil
}
