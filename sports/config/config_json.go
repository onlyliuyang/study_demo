package config

import (
	"encoding/json"
	"os"
	"strings"
)

func Load(filename string) (config Configuration, err error) {
	var data []byte
	data, err = os.ReadFile(filename)
	if err == nil {
		decoder := json.NewDecoder(strings.NewReader(string(data)))
		m := map[string]interface{}{}
		err = decoder.Decode(&m)
		if err == nil {
			config = &DefaultConfig{configData:m}
		}

	}
	return
}


