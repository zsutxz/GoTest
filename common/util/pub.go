package util

import (
	"encoding/json"
	"log"
)

func Transfer(from, to interface{}) error {
	bytes, err := json.Marshal(&from)
	if err != nil {
		log.Println("Transfer json err = %v", err)
		return err
	}
	err = json.Unmarshal(bytes, &to)
	if err != nil {
		log.Println("Transfer json err = %v", err)
		return err
	}
	return nil
}
