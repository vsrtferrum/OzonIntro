package readfromconfig

import (
	"encoding/json"
	"os"

	"github.com/vsrtferrum/OzonIntro/internal/transform"
)



func ReadConfig(filename string) (*transform.Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config transform.Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}