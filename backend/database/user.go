package database

import "log"

func (db *Database) GetOrCreateSocialUser(name string, email string, avatar string) (string, error) { //, provider string
	// provider
	var user_id string
	err := db.DB.QueryRow(
		`SELECT id FROM users WHERE email = $1;`,
		email).Scan(&user_id)
	if err != nil {
		log.Println(err)
		return "", err
	}
	if user_id != "" {
		return user_id, nil
	}

	err = db.DB.QueryRow(
		`INSERT INTO users (name, email, image) VALUES ($1, $2, $3) RETURNING id;`,
		name, email, avatar).Scan(&user_id)
	if err != nil {
		log.Println(err)
		return "", err
	}

	// TODO: Insert into user_providers table
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

	rows, err := db.DB.Query(
		`SELECT p."name"
		FROM "user_providers" up
		JOIN "providers" p ON up."provider_id" = p."id"
		WHERE up."user_id" = $1`, id)
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
