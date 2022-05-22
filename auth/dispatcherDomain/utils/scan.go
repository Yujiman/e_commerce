package utils

import (
	"bufio"
	"io"
	"log"
	"os"
)

func Scan(val *string) {
	scr := bufio.NewScanner(os.Stdin)
	for scr.Scan() {
		line := scr.Text()
		*val = line
		break
	}
	if err := scr.Err(); err != nil {
		if err != io.EOF {
			//log.Println(os.Stderr, err)
			log.Println(err)
		}
	}
}
