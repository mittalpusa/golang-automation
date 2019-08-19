package stepdefinitions

import (
	"encoding/json"
	"log"

	"github.com/golang-automation/features/helper/api"
	"github.com/logrusorgru/aurora"
	"github.com/yalp/jsonpath"
)

var httpResponse interface{}

/*ResponseFindKey is function to find key of response API*/
func ResponseFindKey(key string) error {
	if err := json.Unmarshal(api.ResponseBody, &httpResponse); err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}

	if _, err := jsonpath.Read(httpResponse, key); err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}

	return nil
}

/*ResponseMatchingValue is function to find and matching key value of response API*/
func ResponseMatchingValue(key string, response string) error {
	http, _ := jsonpath.Prepare(key)

	if err := json.Unmarshal(api.ResponseBody, &httpResponse); err != nil {
		log.Fatalln(aurora.Bold(aurora.Red(err)))
	}

	actualResult, _ := http(httpResponse)
	expectResult := response

	if actualResult != expectResult {
		log.Fatalln("actual status code :", aurora.Bold(aurora.Red(actualResult)))
	}

	return nil
}