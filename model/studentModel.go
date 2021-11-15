package studentModel

type Student struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
