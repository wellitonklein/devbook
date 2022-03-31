package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewRepositoryUser(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO USERS (NAME, NICK, EMAIL, PASSWORD) VALUES(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	response, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastId, err := response.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil

}

func (repository users) Find(search string) ([]models.User, error) {
	search = fmt.Sprintf("%%%s%%", search)

	rows, err := repository.db.Query(
		"SELECT ID, NAME, NICK, EMAIL, CREATED_AT FROM USERS WHERE NAME LIKE ? OR NICK LIKE ?",
		search, search,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		user := models.User{}

		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository users) FindById(id uint64) (models.User, error) {
	rows, err := repository.db.Query(
		"SELECT ID, NAME, NICK, EMAIL, CREATED_AT FROM USERS WHERE ID = ?",
		id,
	)
	if err != nil {
		return models.User{}, err
	}
	defer rows.Close()

	user := models.User{}
	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository users) Update(id uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"UPDATE USERS SET NAME = ?, NICK = ?, EMAIL = ? WHERE ID = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, id); err != nil {
		return err
	}

	return nil
}

func (repository users) Delete(id uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM USERS WHERE ID = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(id); err != nil {
		return err
	}

	return nil
}

func (repository users) FindByEmail(email string) (models.User, error) {
	row, err := repository.db.Query("SELECT ID, PASSWORD FROM USERS WHERE EMAIL = ?", email)
	if err != nil {
		return models.User{}, nil
	}
	defer row.Close()

	user := models.User{}
	if row.Next() {
		if err = row.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}
