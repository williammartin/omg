package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/williammartin/jsonschema"
	"github.com/williammartin/omg"
)

func main() {
	reflector := &jsonschema.Reflector{AllowAdditionalProperties: false, RequiredFromJSONSchemaTags: true}
	js := reflector.Reflect(&omg.Microservice{})
	jm, err := json.Marshal(js)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println(string(jm))
}
