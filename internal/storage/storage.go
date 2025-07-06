package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	pool *pgxpool.Pool
}

func NewDB(dbURL string) (*Storage, error) {
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		return nil, err
	}

	_, err = pool.Exec(context.Background(), `
		create table if not exists url (
		    id serial primary key,
		    alias text not null unique,
		    url text not null
		);`)
	if err != nil {
		return nil, fmt.Errorf("create table error: %w", err)
	}

	return &Storage{pool: pool}, nil
}

func (s *Storage) SaveUrl(originalURL string, alias string) error {
	_, err := s.pool.Exec(context.Background(), `
		insert into url (alias, url) values ($1, $2);`, alias, originalURL)
	if err != nil {
		return fmt.Errorf("insert failed: %w", err)
	}
	return nil
}

func (s *Storage) GetUrl(alias string) (string, error) {
	var url string
	err := s.pool.QueryRow(context.Background(), `
		select url from url where alias = $1;`, alias).Scan(&url)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", fmt.Errorf("url %q not found", alias)
		}
		return "", fmt.Errorf("failed to get url: %w", err)
	}
	return url, nil
}

func (s *Storage) DeleteUrl(alias string) error {
	cmtg, err := s.pool.Exec(context.Background(), `
		delete from url where alias = $1`, alias)
	if err != nil {
		return fmt.Errorf("failed to delete url: %w", err)
	}

	if cmtg.RowsAffected() == 0 {
		return fmt.Errorf("alias %s not found", alias)
	}

	return nil
}
