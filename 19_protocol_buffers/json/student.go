package student

type Student struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	//Email string "-" // we can neglect the email things
}
