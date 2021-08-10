package domain

type User struct {
	Id          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
	IPAddress   string `json:"id_address"`
}

func (user User) NewUser(u User) User {
	return User{
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		Username:    u.Username,
		PhoneNumber: u.PhoneNumber,
		Email:       u.Email,
		Gender:      u.Gender,
		IPAddress:   u.IPAddress,
	}
}
