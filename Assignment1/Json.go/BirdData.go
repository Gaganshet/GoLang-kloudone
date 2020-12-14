package main

import (
	"encoding/json"
	"fmt"
) 

func main() {
	birdData := map[string]interface{}{
		"birdSounds": map[string]string{
			"pigeon": "coo",
			"eagle":  "squak",
		},
		"total birds": 2,
	}

	data, _ := json.Marshal(birdData)
	fmt.Println(string(data))
}
