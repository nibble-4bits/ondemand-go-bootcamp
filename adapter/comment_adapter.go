package adapter

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/nibble-4bits/ondemand-go-bootcamp/entity"
)

var (
	ErrCommentNotFoundByID = errors.New("no comment found by ID")
)

type commentAdapter struct {
	csvDataSource  CSVDataSource
	httpDataSource HTTPDataSource
	dataStore      DataStore
	comments       []entity.Comment
}

// NewCommentAdapter receives a data source and will try to fetch the
// list of comments from a data source.
//
// If successful, an instance of *commentAdapter will be returned.
// Otherwise and error will be returned.
func NewCommentAdapter(csvSource CSVDataSource, httpSource HTTPDataSource, store DataStore) (*commentAdapter, error) {
	adapter := &commentAdapter{csvDataSource: csvSource, httpDataSource: httpSource, dataStore: store}

	if err := adapter.loadComments(); err != nil {
		return nil, err
	}

	return adapter, nil
}

func (a *commentAdapter) loadComments() error {
	csvRecords, err := a.csvDataSource.ReadCollection()
	if err != nil {
		return err
	}

	for _, v := range csvRecords {
		c := entity.Comment{}

		c.ID.ParseInt(v[0], -1)
		c.PostID.ParseInt(v[1], -1)
		c.Name = v[2]
		c.Email = v[3]
		c.Body = v[4]

		a.comments = append(a.comments, c)
	}

	return nil
}

func (a *commentAdapter) saveRecord(comment *entity.Comment) {
	record := []string{
		fmt.Sprint(comment.ID),
		fmt.Sprint(comment.PostID),
		comment.Name,
		comment.Email,
		comment.Body,
	}

	err := a.dataStore.SaveRecord(record)
	if err != nil {
		log.Printf("Error trying to save record to data store: %v", err)
	}
}

func (a *commentAdapter) getCommentByIDFromAPI(id int) (*entity.Comment, error) {
	comment := &entity.Comment{}
	endpoint := fmt.Sprintf("https://jsonplaceholder.typicode.com/comments/%v", id)
	resp, err := a.httpDataSource.ReadItem(endpoint)
	if err != nil {
		return nil, err
	}

	// All bad requests to the https://jsonplaceholder.typicode.com/comments/{id} endpoint
	// return a 404 status code. Doesn't matter if the id parameter is invalid.
	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("%w %v", ErrCommentNotFoundByID, id)
	}

	err = json.Unmarshal(resp.Body, comment)
	if err != nil {
		return nil, err
	}

	// Save the comment in a CSV file, so next time we don't have
	// to fetch it from the API.
	a.saveRecord(comment)

	return comment, nil
}

// GetByID searches for a comment with the given id parameter.
//
// If the search is successful, a pointer to the found Comment is returned.
// Otherwise an error is returned.
func (a *commentAdapter) GetByID(id int) (*entity.Comment, error) {
	// Search for comments in our local cache
	for _, comment := range a.comments {
		if id == int(comment.ID) {
			return &comment, nil
		}
	}

	// If a comment is not found in the local cache, try fetching it from the JSON Placeholder API
	comment, err := a.getCommentByIDFromAPI(id)
	if err != nil {
		return nil, err
	}

	// Append the fetched comment to our local cache
	a.comments = append(a.comments, *comment)

	return comment, nil
}
