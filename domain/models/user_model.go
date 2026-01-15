package models

type User struct {
	ID         string
	FirebaseID string
	Role       string
}

type UserDetails struct {
	ID        string
	User_ID   string
	User_Name string
	Email     string
}
