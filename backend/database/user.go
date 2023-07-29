package database

import "log"

func (db *Database) CreateSocialUser(name string, email string, avatar string) (string, error) { //, provider string
	// provider
	var user_id string
	err := db.DB.QueryRow(
		`INSERT INTO users (name, email, image) VALUES ($1, $2, $3) RETURNING id;`,
		name, email, avatar).Scan(&user_id)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// TODO: Insert into user_providers table
	return user_id, nil
}
