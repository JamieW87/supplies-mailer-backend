package postgres

import "github.com/google/uuid"

func (pg PostgresStore) InsertUser(name, email, phone string) (*uuid.UUID, error) {

	//SELECT user_id FROM users WHERE email = 'user_email@example.com';

	//If not exists

	//Insert into users table, get id

	return nil, nil
}

func (pg PostgresStore) InsertUserCategory(user uuid.UUID, category string) {

	// SELECT id from categories where name = category

	//Insert into user_categories (userId categoryId)

}
