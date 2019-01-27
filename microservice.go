package omg

// Microservice is the the top level configuration
type Microservice struct {
	// OMG specifies the version of the OMG specification with which this microservice complies.
	OMG int `json:"omg,omitempty"`

	// Info specifies general information about the microservice.
	Info *Info `json:"info,omitempty"`
}
