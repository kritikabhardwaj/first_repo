package api

//Data emp model
type Data struct {
	tableName struct{} `sql:"emp"`
	ID        int      `sql:"user_id" json:"user_id"`
	Name      string   `sql:"user_name" json:"user_name" validate:"required"`
	Password  string   `sql:"password" json:"password" validate:"min=8"`
	Email     string   `sql:"email" json:"email" validate:"required,email"`
	//HouseNo   int      `sql:"house_no" json:"house_no"`
}
