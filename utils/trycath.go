package utils

import "log"

func TryCatch() {
	if err := recover(); err != nil {
		log.Println(err)
	}
}