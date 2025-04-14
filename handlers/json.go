package handlers

import (
	"encoding/json"
	"os"
)

func SaveJson(file string, data any) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = f.Write(b)

	return err
}

func OpenJson(file string, data any) (err error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		return
	}

	return
}
