package hydraconfigurator

import (
	"encoding/json"
	"fmt"
	"os"
)

func decodeJsonConfig(v interface{}, filename string) error {
	fmt.Println("Decoding Json")
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	return json.NewDecoder(file).Decode(&v)
}
