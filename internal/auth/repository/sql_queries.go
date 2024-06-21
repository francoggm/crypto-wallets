package repository

const (
	getUserByEmail = "SELECT id, username, email, password, role, created_at, updated_at, last_login FROM users WHERE email=$1"

	getUserByID = "SELECT id, username, email, password, role, created_at, updated_at, last_login FROM users WHERE id=$1"

	insertUser = "INSERT INTO users (username, email, password, role) VALUES (:username, :email, :password, :role) RETURNING id"

	updateLastLogin = "UPDATE users SET last_login = $2 WHERE id = $1"
)
