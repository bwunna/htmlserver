package ServerAPI

import (
	"encoding/json"
	"errors"
	"net/http"
)

func checkKey(request *http.Request) (string, error) {
	// checking if key from json is valid
	var keyInJson map[string]interface{}

	err := json.NewDecoder(request.Body).Decode(&keyInJson)
	if err != nil {
		return "", err
	}
	if key, ok := keyInJson["key"].(string); !ok {
		return "", errors.New("key is not valid")
	} else {
		return key, nil
	}

}
