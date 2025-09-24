package accounting

type Warnings []Warning

type Warning struct {
	Message string `json:"Message,omitempty" xml:"Message,omitempty"`
}
