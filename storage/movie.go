package storage

import (
	"database/sql"

	"github.com/google/uuid"

	"app/models"
)

func InsertMovie(db *sql.DB, movie models.Movie) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
		INSERT INTO movie (
			id,
			title,
			duration,
			description
		) VALUES ($1, $2, $3, $4)
	`

	_, err := db.Exec(query,
		id,
		movie.Title,
		movie.Duration,
		movie.Description,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func GetByIdMovie(db *sql.DB, id string) (models.Movie, error) {

	var (
		movie models.Movie
	)

	query := `
		SELECT
			id,
			title,
			TO_CHAR(duration, 'HH24:MI:SS'),
			description
		FROM movie WHERE id = $1
	`

	err := db.QueryRow(query, id).Scan(
		&movie.Id,
		&movie.Title,
		&movie.Duration,
		&movie.Description,
	)

	if err != nil {
		return models.Movie{}, err
	}

	return movie, nil
}

func GetAllMovie(db *sql.DB) ([]models.Movie, error) {

	var (
		movie []models.Movie
	)

	query := `
		SELECT
			id,
			title,
			TO_CHAR(duration, 'HH24:MI:SS'),
			description
		FROM movie
	`

	rows, err := db.Query(query)
	if err != nil {
		return []models.Movie{}, err
	}

	for rows.Next() {
		var m models.Movie

		err = rows.Scan(
			&m.Id,
			&m.Title,
			&m.Duration,
			&m.Description,
		)
		if err != nil {
			return []models.Movie{}, err
		}

		movie = append(movie, m)
	}

	return movie, nil
}

func Update(db *sql.DB, movie models.Movie) error {

	query := `
		UPDATE movie
			set title=$2,
			    duration=$3,
			    description=$4
		 WHERE id = $1
	`

	_, err := db.Exec(query,
		movie.Id,
		movie.Title,
		movie.Duration,
		movie.Description,
	)

	if err != nil {
		return err
	}

	return nil
}

func Delete(db *sql.DB, id string) error {

	query := `
		DELETE from movie
		 WHERE id = $1
	`

	_, err := db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func GetListMovie(db *sql.DB) ([]models.Movie, error) {

	var (
		movie []models.Movie
	)

	query := `
		SELECT
			id,
			title,
			TO_CHAR(duration, 'HH24:MI:SS'),
			description
		FROM movie
	`

	rows, err := db.Query(query)
	if err != nil {
		return []models.Movie{}, err
	}

	for rows.Next() {
		var m models.Movie

		err = rows.Scan(
			&m.Id,
			&m.Title,
			&m.Duration,
			&m.Description,
		)
		if err != nil {
			return []models.Movie{}, err
		}

		movie = append(movie, m)
	}

	return movie, nil
}
