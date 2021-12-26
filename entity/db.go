package entity

//用户
type User struct {
	ID        int64  `db: u_id`
	Username  string `db: username`
	Password  string `db: password`
	Telephone string `db: telephone`
	Email     string `db: email`
	Status    int8   `db: status`
}
