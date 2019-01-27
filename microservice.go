package omg

// Microservice is the the top level configuration
type Microservice struct {
	// OMG specifies the version of the OMG specification with which this microservice complies.
	OMG int `json:"omg,omitempty"`

	// Info specifies general information about the microservice.
	Info *Info `json:"info,omitempty"`

	// Lifecycle specifies the lifecycle of the microservice.
	Lifecycle *Lifecycle `json:"lifecycle,omitempty"`

	// Actions specifies how to interact with the microservice.
	Actions Actions `json:"actions,omitempty"`

	// Environment specifies environment variables that should be exposed to the microservice.
	Environment Environment `json:"environment,omitempty"`
}
