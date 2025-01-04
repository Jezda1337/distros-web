package distro

import (
	"encoding/json"
	"os"
)

func LoadDistrosFromFile(filePath string) ([]Distro, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var distros []Distro
	if err := json.Unmarshal(data, &distros); err != nil {
		return nil, err
	}

	return distros, nil
}
