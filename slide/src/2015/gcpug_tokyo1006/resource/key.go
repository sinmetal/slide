type User struct {
	ID          int // key value
	AccountName string
	NickName    string
	Email       string // uniqueにするのが難しい
}

// Login時に利用するKind
type EmailToUser struct {
	Email    string // key value
	UserID   int
	Password string
}