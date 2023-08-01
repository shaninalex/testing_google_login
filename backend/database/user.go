package database

import (
	"errors"
	"log"
)

func (db *Database) GetOrCreateSocialUser(name, email, avatar, provider string) (string, error) {
	// provider
	var user_id string
	err := db.DB.QueryRow(
		`SELECT id FROM users WHERE email = $1;`,
		email).Scan(&user_id)
	if err != nil {
		err = db.DB.QueryRow(
			`INSERT INTO users (name, email, image) VALUES ($1, $2, $3) RETURNING id;`,
			name, email, avatar).Scan(&user_id)
		if err != nil {
			return "", err
		}
		_ = db.DB.QueryRow(
			`INSERT INTO user_providers (user_id, provider) VALUES ($1, $2);`,
			user_id, provider).Scan(&user_id)
		return user_id, nil
	}

	return user_id, nil
}

func (db *Database) GetUser(id string) (*User, error) {
	var user User
	err := db.DB.QueryRow(
		`SELECT id, name, email, image FROM users WHERE id = $1;`, id).Scan(
		&user.Id, &user.Name, &user.Email, &user.Image)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var providers []string

	rows, err := db.DB.Query(`SELECT p."name" FROM "user_providers" up
		JOIN "providers" p ON up."provider" = p."name" WHERE up."user_id" = $1`, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var p string
		rows.Scan(&p)
		providers = append(providers, p)
	}

	user.Providers = providers
	return &user, nil
}

func (db *Database) GetUserByEmail(email string) (*User, error) {
	var user User
	err := db.DB.QueryRow(
		`SELECT id, name, email, image FROM users WHERE id = $1;`, email).Scan(
		&user.Id, &user.Name, &user.Email, &user.Image)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var providers []string

	rows, err := db.DB.Query(`SELECT p."name" FROM "user_providers" up
		JOIN "providers" p ON up."provider_id" = p."id" WHERE up."user_id" = $1`, user.Id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var p string
		rows.Scan(&p)
		providers = append(providers, p)
	}

	user.Providers = providers
	return &user, nil
}

func (db *Database) RegularUserCreate(payload *RegisterPayload) (*User, error) {
	var user User
	err := db.DB.QueryRow(
		`SELECT id, name, email, image FROM users WHERE email = $1;`, payload.Email).Scan(
		&user.Id, &user.Name, &user.Email, &user.Image)
	if err != nil {
		log.Println("user not found, creating one")
		password_hash, err := GenerateFromPassword(payload.Password)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		err = db.DB.QueryRow(
			`INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3) RETURNING id;`,
			payload.Name, payload.Email, password_hash).Scan(&user.Id)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		log.Println("user created")

		// TODO: Insert into user_providers table (local)
		return &user, nil
	}

	return nil, errors.New("User is already exists")

}
