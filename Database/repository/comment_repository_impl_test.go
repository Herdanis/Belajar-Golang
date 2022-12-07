package repository

import (
	"context"
	"database/sql"
	"fmt"
	"golang_database"
	"golang_database/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "aveor@gmail.com",
		Comment: sql.NullString{String: "erar aerva rba", Valid: true},
	}
	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	result, err := commentRepository.FindById(ctx, 3)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(golang_database.GetConnection())

	ctx := context.Background()
	comments, err := commentRepository.FindAll(ctx)

	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println(comment)
	}
}
