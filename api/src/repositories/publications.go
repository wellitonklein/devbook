package repositories

import (
	"api/src/models"
	"database/sql"
)

type publicationRepository struct {
	db *sql.DB
}

func NewRepositoryPublication(db *sql.DB) *publicationRepository {
	return &publicationRepository{db}
}

func (repository publicationRepository) Create(publication models.Publication) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO PUBLICATIONS (TITLE, CONTENT, AUTHOR_ID) VALUES (?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	response, err := statement.Exec(publication.Title, publication.Content, publication.AuthorID)
	if err != nil {
		return 0, err
	}

	lastId, err := response.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}

func (repository publicationRepository) FindById(publicationId uint64) (models.Publication, error) {
	row, err := repository.db.Query(
		`SELECT
			P.*,
			U.NICK
		FROM PUBLICATIONS P
		INNER JOIN USERS U
		ON U.ID = P.AUTHOR_ID
		WHERE P.ID = ?`,
		publicationId,
	)
	if err != nil {
		return models.Publication{}, err
	}
	defer row.Close()

	publication := models.Publication{}

	if row.Next() {
		if err = row.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return models.Publication{}, err
		}
	}

	return publication, nil
}

func (repository publicationRepository) Find(userId uint64) ([]models.Publication, error) {
	rows, err := repository.db.Query(
		`
		SELECT DISTINCT
			P.*,
			U.NICK
		FROM PUBLICATIONS P
		INNER JOIN USERS U
			ON U.ID = P.AUTHOR_ID
		INNER JOIN FOLLOWERS F
			ON P.AUTHOR_ID = F.USER_ID
		WHERE U.ID = ? OR F.FOLLOWER_ID = ?
		ORDER BY 1 DESC
		`,
		userId, userId,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	publications := []models.Publication{}

	if rows.Next() {
		publication := models.Publication{}

		if err = rows.Scan(
			&publication.ID,
			&publication.Title,
			&publication.Content,
			&publication.AuthorID,
			&publication.Likes,
			&publication.CreatedAt,
			&publication.AuthorNick,
		); err != nil {
			return nil, err
		}

		publications = append(publications, publication)
	}

	return publications, nil
}
