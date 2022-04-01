package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type userRepository struct {
	db *sql.DB
}

func NewRepositoryUser(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (repository userRepository) Create(user models.User) (uint64, error) {
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

func (repository userRepository) Find(search string) ([]models.User, error) {
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

func (repository userRepository) FindById(id uint64) (models.User, error) {
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

func (repository userRepository) Update(id uint64, user models.User) error {
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

func (repository userRepository) Delete(id uint64) error {
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

func (repository userRepository) FindByEmail(email string) (models.User, error) {
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

func (repository userRepository) Follow(userId, followerId uint64) error {
	statement, err := repository.db.Prepare(
		"INSERT IGNORE INTO FOLLOWERS (USER_ID, FOLLOWER_ID) VALUES (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

func (repository userRepository) UnFollow(userId, followerId uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM FOLLOWERS WHERE USER_ID = ? AND FOLLOWER_ID = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userId, followerId); err != nil {
		return err
	}

	return nil
}

func (repository userRepository) FindFollowers(userId uint64) ([]models.User, error) {
	rows, err := repository.db.Query(
		`SELECT 
			U.ID, 
			U.NAME, 
			U.NICK, 
			U.EMAIL, 
			U.CREATED_AT 
		FROM USERS U 
		INNER JOIN FOLLOWERS F 
		ON U.ID = F.FOLLOWER_ID 
		WHERE F.USER_ID = ?`,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	followers := []models.User{}
	for rows.Next() {
		follower := models.User{}

		if err := rows.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreatedAt,
		); err != nil {
			return nil, err
		}

		followers = append(followers, follower)
	}

	return followers, nil
}

func (repository userRepository) FindFollowings(userId uint64) ([]models.User, error) {
	rows, err := repository.db.Query(
		`SELECT 
			U.ID, 
			U.NAME, 
			U.NICK, 
			U.EMAIL, 
			U.CREATED_AT 
		FROM USERS U 
		INNER JOIN FOLLOWERS F 
		ON U.ID = F.USER_ID 
		WHERE F.FOLLOWER_ID = ?`,
		userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	followings := []models.User{}
	for rows.Next() {
		following := models.User{}

		if err := rows.Scan(
			&following.ID,
			&following.Name,
			&following.Nick,
			&following.Email,
			&following.CreatedAt,
		); err != nil {
			return nil, err
		}

		followings = append(followings, following)
	}

	return followings, nil
}

func (repository userRepository) FindPassword(userId uint64) (string, error) {
	row, err := repository.db.Query("SELECT PASSWORD FROM USERS WHERE ID = ?", userId)
	if err != nil {
		return "", err
	}
	defer row.Close()

	user := models.User{}
	if row.Next() {
		if err = row.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (repository userRepository) ResetPassword(userId uint64, password string) error {
	statement, err := repository.db.Prepare("UPDATE USERS SET PASSWORD = ? WHERE ID = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(password, userId); err != nil {
		return err
	}

	return nil
}
