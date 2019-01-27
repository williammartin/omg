package omg_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xeipuuv/gojsonschema"

	"github.com/williammartin/jsonschema"
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

	Describe("Actions", func() {
		It("requires one of format or http", func() {
			actions := Actions{
				"action": &Action{},
			}

			microservice := &Microservice{Actions: actions}
			valid, errors := validate(microservice)
			Expect(valid).To(BeFalse())
			Expect(errors).To(ContainElement("actions.action: Must validate one and only one schema (oneOf)"))
		})

		Describe("http", func() {
			It("requires a method", func() {
				actions := Actions{
					"action": &Action{
						HTTP: &HTTP{},
					},
				}

				microservice := &Microservice{Actions: actions}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement("actions.action.http: method is required"))
			})

			It("requires a port", func() {
				actions := Actions{
					"action": &Action{
						HTTP: &HTTP{},
					},
				}

				microservice := &Microservice{Actions: actions}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement("actions.action.http: port is required"))
			})

			It("requires a path", func() {
				actions := Actions{
					"action": &Action{
						HTTP: &HTTP{},
					},
				}

				microservice := &Microservice{Actions: actions}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement("actions.action.http: path is required"))
			})
		})

		Describe("format", func() {
			It("requires a command", func() {
				actions := Actions{
					"action": &Action{
						Format: &Format{},
					},
				}

				microservice := &Microservice{Actions: actions}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement("actions.action.format: command is required"))
			})
		})

		Describe("the arguments", func() {
			It("requires a type", func() {
				actions := Actions{
					"action": &Action{
						Arguments: Arguments{
							"arg": &Argument{},
						},
					},
				}

				microservice := &Microservice{Actions: actions}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement("actions.action.arguments.arg: type is required"))
			})

			It("requires arguments must be of a particular type", func() {
				actions := Actions{
					"action": &Action{
						Arguments: Arguments{
							"arg": &Argument{
								Type: "not-valid",
							},
						},
					},
				}

				microservice := &Microservice{Actions: actions}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement(
					"actions.action.arguments.arg.type: actions.action.arguments.arg.type must be one of the following: \"int\", \"float\", \"string\", \"list\", \"map\", \"boolean\", \"enum\"",
				))
			})

			It("requires a location", func() {
				actions := Actions{
					"action": &Action{
						Arguments: Arguments{
							"arg": &Argument{},
						},
					},
				}

				microservice := &Microservice{Actions: actions}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement("actions.action.arguments.arg: in is required"))
			})

			It("requires arguments must have a particular location", func() {
				actions := Actions{
					"action": &Action{
						Arguments: Arguments{
							"arg": &Argument{
								In: "not-valid",
							},
						},
					},
				}

				microservice := &Microservice{Actions: actions}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement(
					"actions.action.arguments.arg.in: actions.action.arguments.arg.in must be one of the following: \"requestBody\", \"query\", \"path\"",
				))
			})
		})

		Describe("the output", func() {
			It("requires a type", func() {
				actions := Actions{
					"action": &Action{
						Output: &Output{},
					},
				}

				microservice := &Microservice{Actions: actions}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement("actions.action.output: type is required"))
			})

			It("requires that output must be of a particular type", func() {
				actions := Actions{
					"action": &Action{
						Output: &Output{
							Type: "not-valid",
						},
					},
				}

				microservice := &Microservice{Actions: actions}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement(
					"actions.action.output.type: actions.action.output.type must be one of the following: \"int\", \"float\", \"string\", \"list\", \"map\", \"boolean\", \"object\"",
				))
			})

			It("requires that properties have a type", func() {
				actions := Actions{
					"action": &Action{
						Output: &Output{
							Properties: map[string]*Property{
								"prop": &Property{},
							},
						},
					},
				}

				microservice := &Microservice{Actions: actions}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement(
					"actions.action.output.properties.prop: type is required",
				))
			})

			It("requires that properties must be of a particular type", func() {
				actions := Actions{
					"action": &Action{
						Output: &Output{
							Properties: map[string]*Property{
								"prop": &Property{
									Type: "not-valid",
								},
							},
						},
					},
				}

				microservice := &Microservice{Actions: actions}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement(
					"actions.action.output.properties.prop.type: actions.action.output.properties.prop.type must be one of the following: \"int\", \"float\", \"string\", \"list\", \"map\", \"boolean\", \"object\"",
				))
			})
		})
	})

	Describe("Environment", func() {
		Describe("the environment variables", func() {
			It("requires a type", func() {
				environment := Environment{
					"var": &Variable{},
				}

				microservice := &Microservice{Environment: environment}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement("environment.var: type is required"))
			})
		})

		It("requires variables must be of a particular type", func() {
			environment := Environment{
				"var": &Variable{
					Type: "not-valid",
				},
			}

			microservice := &Microservice{Environment: environment}
			valid, errors := validate(microservice)
			Expect(valid).To(BeFalse())
			Expect(errors).To(ContainElement(
				"environment.var.type: environment.var.type must be one of the following: \"int\", \"float\", \"string\", \"boolean\"",
			))
		})
	})

	Describe("Volumes", func() {
		Describe("the volume", func() {
			It("requires a type", func() {
				volumes := Volumes{
					"disk": &Volume{},
				}

				microservice := &Microservice{Volumes: volumes}
				valid, errors := validate(microservice)
				Expect(valid).To(BeFalse())
				Expect(errors).To(ContainElement("volumes.disk: target is required"))
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
