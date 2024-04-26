package reqmodels

type TestModel struct {
	FirstName string `json:"firstName" binding:"required,eq=pranav"`
	LastName  string `json:"lastName" binding:"required,eq=shukla"`
}
