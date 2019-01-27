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

	Describe("Info", func() {
		Describe("the microservice version", func() {
			It("is required", func() {
				microservice := &Microservice{Info: &Info{}}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement("info: version is required"))
			})

			It("must be a semantic version", func() {
				microservice := &Microservice{Info: &Info{Version: "not-valid"}}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement("info.version: Does not match pattern '[0-9]*\\.[0-9]*\\.[0-9]*'"))
			})
		})

		Describe("the title", func() {
			It("is required", func() {
				microservice := &Microservice{Info: &Info{}}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement("info: title is required"))
			})
		})

		Describe("the description", func() {
			It("is required", func() {
				microservice := &Microservice{Info: &Info{}}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement("info: description is required"))
			})
		})

		Describe("the license", func() {
			It("is required", func() {
				microservice := &Microservice{Info: &Info{}}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement("info: license is required"))
			})

			It("requires a name", func() {
				microservice := &Microservice{Info: &Info{License: &License{}}}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement("info.license: name is required"))
			})

			It("requires a url", func() {
				microservice := &Microservice{Info: &Info{License: &License{}}}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement("info.license: url is required"))
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
