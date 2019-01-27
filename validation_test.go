package omg_test

import (
	"encoding/json"

	"github.com/alecthomas/jsonschema"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xeipuuv/gojsonschema"

	. "github.com/williammartin/omg"
)

var _ = Describe("Schema Validation", func() {
	Describe("the microservice", func() {
		Describe("the omg version", func() {
			It("is required", func() {
				microservice := &Microservice{}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement("(root): omg is required"))
			})
		})

		Describe("the microservice info", func() {
			It("is required", func() {
				microservice := &Microservice{}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement("(root): info is required"))
			})
		})
	})
})

func validate(microservice *Microservice) (bool, []string) {
	reflector := &jsonschema.Reflector{AllowAdditionalProperties: false, RequiredFromJSONSchemaTags: true}
	js := reflector.Reflect(&Microservice{})
	jm, err := json.Marshal(js)
	Expect(err).NotTo(HaveOccurred())

	schemaLoader := gojsonschema.NewStringLoader(string(jm))
	documentLoader := gojsonschema.NewGoLoader(microservice)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	Expect(err).NotTo(HaveOccurred())

	validationErrors := []string{}
	for _, e := range result.Errors() {
		validationErrors = append(validationErrors, e.String())
	}

	return result.Valid(), validationErrors
}
