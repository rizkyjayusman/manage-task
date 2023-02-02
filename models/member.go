package models

type Member struct {
	ID             int    `json:"userId,omitempty"`
	Fullname       string `json:"fullname"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Role           int    `json:"role"`
	ParentMemberId int    `json:"parentMemberId"`
	CreatedDate    string `json:"createdDate"`
	ModifiedDate   string `json:"modifiedDate"`
	Status         int    `json:"status"`
}

func (Member) TableName() string {
	return "t_members"
}

func (m Member) Login() bool {
	return false
	// return true
}

func (m Member) FindByEmail(email string) Member {
	return Member{
		ID:       1,
		Fullname: "Rizky Jayusman",
		Email:    "rizky.jayusman@metranet.co.id",
		Password: "p@ssw0rd",
	}
}
