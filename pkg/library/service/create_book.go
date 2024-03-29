package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

	"library-api/pkg/library"
)

func (s *service) CreateBook(book *library.Book, bookFile multipart.File) (*library.Book, error) {
	if len(book.Title) < 3 || len(book.Title) > 80 {
		return nil, fmt.Errorf(
			`%w: the field "name" must be between 3 and 80 characters long`,
			ErrInvalidName,
		)
	}

	if len(book.Description) < 8 {
		return nil, fmt.Errorf(
			`%w: field "description" must be longer than 8 characters`,
			ErrInvalidDescription,
		)
	}

	if _, err := time.Parse(onlyDateLayout, book.ReleaseDate); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidDate, err)
	}

	bookModel := &library.Book{
		Title:       book.Title,
		Description: book.Description,
		Writer:      book.Writer,
		Gender:      book.Gender,
		ReleaseDate: book.ReleaseDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*50)
	defer cancel()

	if err := s.libraryRepo.CreateBook(bookModel, bookFile, s.uploadToBucket(ctx, bookModel, bookFile)); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrCreateBook, err)
	}

	return bookModel, nil
}

func (s *service) uploadToBucket(
	ctx context.Context,
	bookInfo *library.Book,
	bookFile multipart.File,
) func() error {
	return func() error {
		return s.storage.Upload(ctx, bookInfo, bookFile)
	}
}
