package omg_test

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/williammartin/omg"
)

var _ = Describe("OMG", func() {
	It("can unmarshal a minimal OMG microservice", func() {
		microservice := loadMicroservice("assets/minimal.yml")
		Expect(microservice.OMG).To(Equal(1))

		info := microservice.Info
		Expect(info.Version).To(Equal("0.0.1"))
		Expect(info.Title).To(Equal("MinimalMicroservice"))
		Expect(info.Description).To(Equal("A minimal microservice"))

		contact := info.Contact
		Expect(contact.Name).To(Equal("John Doe"))
		Expect(contact.URL).To(Equal("example.com/minimal"))
		Expect(contact.Email).To(Equal("minimal@example.com"))

		license := info.License
		Expect(license.Name).To(Equal("MIT"))
		Expect(license.URL).To(Equal("example.com/MIT"))
	})

	Describe("lifecycle hooks", func() {
		It("can unmarshal a microservice with lifecycle hooks", func() {
			microservice := loadMicroservice("assets/lifecycle.yml")

			lifecycle := microservice.Lifecycle
			Expect(lifecycle.Startup.Command).To(Equal("startup"))
			Expect(lifecycle.Shutdown.Command).To(Equal("shutdown"))
			Expect(lifecycle.Shutdown.Timeout).To(Equal(60))
		})

		It("can unmarshal a microservice with string array commands", func() {
			microservice := loadMicroservice("assets/array_lifecycle.yml")

			lifecycle := microservice.Lifecycle
			Expect(lifecycle.Startup.Command).To(ConsistOf("startup", "now"))
			Expect(lifecycle.Shutdown.Command).To(ConsistOf("shutdown", "sometime", "please"))
		})
	})

	Describe("actions", func() {
		It("can unmarshal various types of actions", func() {
			microservice := loadMicroservice("assets/actions.yml")
			actions := microservice.Actions

			By("unmarshaling format actions")
			fooAction := actions["foo"]
			Expect(fooAction.Format.Command).To(ConsistOf("foo", "command"))

			By("unmarshaling http actions")
			barAction := actions["bar"]
			Expect(barAction.HTTP.Port).To(Equal(8080))
			Expect(barAction.HTTP.Method).To(Equal("POST"))
			Expect(barAction.HTTP.Path).To(Equal("/bar"))
			Expect(barAction.HTTP.ContentType).To(Equal("application/json"))

			By("unmarshaling common fields")
			commonAction := actions["common"]
			Expect(commonAction.Help).To(Equal("common action help"))

			ditArg := commonAction.Arguments["dit"]
			Expect(ditArg.Help).To(Equal("arguments help"))
			Expect(ditArg.Type).To(Equal("int"))
			Expect(ditArg.In).To(Equal("path"))
			Expect(ditArg.Required).To(Equal(true))
			Expect(ditArg.Pattern).To(Equal(".*"))
			Expect(ditArg.Enum).To(ConsistOf("first", "second", "third"))
			Expect(ditArg.Range.Min).To(Equal(1))
			Expect(ditArg.Range.Max).To(Equal(2))

			output := commonAction.Output
			Expect(output.Type).To(Equal("string"))
			Expect(output.ContentType).To(Equal("text/plain"))
			Expect(output.Properties["thing"].Type).To(Equal("boolean"))
		})
	})

	Describe("environment", func() {
		It("can unmarshal a microservice with environment variable requirements", func() {
			microservice := loadMicroservice("assets/environment.yml")
			environment := microservice.Environment

			Expect(environment["ENV_VAR"].Type).To(Equal("string"))
			Expect(environment["ENV_VAR"].Pattern).To(Equal(".*"))
			Expect(environment["ENV_VAR"].Required).To(BeTrue())
			Expect(environment["ENV_VAR"].Help).To(Equal("env var help"))
		})

	})
})

func loadMicroservice(filepath string) *Microservice {
	Expect(filepath).To(BeAnExistingFile())

	bytes, err := ioutil.ReadFile(filepath)
	Expect(err).NotTo(HaveOccurred())

	var microservice Microservice
	Expect(yaml.Unmarshal(bytes, &microservice)).To(Succeed())

	return &microservice
}
