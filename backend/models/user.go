package models

type User struct {
	ID        int
	FirstName string
	LastName  string
	LDAP      string
	Email     string
	Password  string
}
