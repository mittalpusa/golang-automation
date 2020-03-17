package step

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/golang-automation/features/helper"
	api "github.com/golang-automation/features/helper/api"
	"github.com/golang-automation/features/helper/message"
	"github.com/yalp/jsonpath"
)

/*JSONValue struct for set json value*/
type JSONValue struct {
	variable interface{}
}

var jsonResponse, actualResult interface{}

func decryptJSONResponse() error {
	if err := json.Unmarshal(api.ResponseBody, &jsonResponse); err != nil {
		log.Panicln(fmt.Errorf("REASON: %s", err))
	}

	return nil
}

/*ResponseFindPath is function to find path of response API*/
func ResponseFindPath(path string) error {
	decryptJSONResponse()

	countpath, _ := jsonpath.Read(jsonResponse, path)

	if err := len(countpath.([]interface{})); err == 0 {
		log.Panicln(fmt.Errorf("REASON: %s", strconv.Itoa(err)))
	}

	return nil
}

func getJSONValue(path string) {
	decryptJSONResponse()

	HTTPJson, _ := jsonpath.Prepare(path)
	actualResult, _ = HTTPJson(jsonResponse)
}

/*ResponseMatchingValue is function to find and matching path value of response API*/
func ResponseMatchingValue(path string, expectResult string) error {
	getJSONValue(path)

	helper.AssertEqual(expectResult, actualResult, message.NotMatchValue(actualResult))

	return nil
}

/*ResponseDataType is function to find and matching path value with data type*/
func ResponseDataType(path string, expectType string) error {
	var actualType string

	decryptJSONResponse()

	HTTPJson, _ := jsonpath.Prepare(path)
	actualResult, _ := HTTPJson(jsonResponse)

	helper.AssertEqual(expectType, actualType, message.NotMatchDataType(actualResult.(string)))

	return nil
}

/*CollectsJSON function to keep jsonpath value*/
func CollectsJSON(path string, value string) error {
	getJSONValue(path)

	// still on research
	realValue := actualResult.(string)
	realVariable := value

	fmt.Println(realValue)
	fmt.Println(realVariable)

	return nil
}