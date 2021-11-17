package studentModel

type Student struct {
	Id     int64  `json:"id,omitempty"`
	Fname  string `json:"fName"`
	Lname  string `json:"lName"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Age    int    `json:"age"`
}
