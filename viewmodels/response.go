package viewmodels

// Meta holds the metainformation of the microservice
type Meta struct {
	Code    int    `json:"code"`
	Version string `json:"version"`
}

// Error contains a error definition to be rendered
type Error struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// Response holds the response of a microservice,
// it's being reused as response structure in facade-service
type Response struct {
	Meta   Meta        `json:"meta"`
	Data   interface{} `json:"data,omitempty"`
	Errors []Error     `json:"errors,omitempty"`
}

// Message _
type Message struct {
	Message string `json:"message"`
}

// ContentResponse Workaround, It's just for document swagger.
type ContentResponse struct {
	Meta    Meta    `json:"meta"`
	Message Message `json:"data"`
}
