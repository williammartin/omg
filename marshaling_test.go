package omg_test

import (
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	yaml "gopkg.in/yaml.v2"

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
		It("can unmarshal an OMG microservice with lifecycle hooks", func() {
			microservice := loadMicroservice("assets/lifecycle.yml")

			lifecycle := microservice.Lifecycle
			Expect(lifecycle.Startup.Command).To(Equal("startup"))
			Expect(lifecycle.Shutdown.Command).To(Equal("shutdown"))
			Expect(lifecycle.Shutdown.Timeout).To(Equal(60))
		})

		It("can unmarshal an OMG microservice with string array commands", func() {
			microservice := loadMicroservice("assets/array_lifecycle.yml")

			lifecycle := microservice.Lifecycle
			Expect(lifecycle.Startup.Command).To(ConsistOf("startup", "now"))
			Expect(lifecycle.Shutdown.Command).To(ConsistOf("shutdown", "sometime", "please"))
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
