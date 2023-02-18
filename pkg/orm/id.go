package orm

var userID = 0

func NewUserID() int {
	userID++
	return userID
}
