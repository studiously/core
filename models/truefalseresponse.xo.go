// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// TrueFalseResponse represents a row from 'public.true_false_responses'.
type TrueFalseResponse struct {
	ID         int64         `json:"id"`          // id
	QuestionID int64         `json:"question_id"` // question_id
	AuthorID   sql.NullInt64 `json:"author_id"`   // author_id
	Response   bool          `json:"response"`    // response

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the TrueFalseResponse exists in the database.
func (tfr *TrueFalseResponse) Exists() bool {
	return tfr._exists
}

// Deleted provides information if the TrueFalseResponse has been deleted from the database.
func (tfr *TrueFalseResponse) Deleted() bool {
	return tfr._deleted
}

// Insert inserts the TrueFalseResponse to the database.
func (tfr *TrueFalseResponse) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if tfr._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO public.true_false_responses (` +
		`id, question_id, author_id, response` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)`

	// run query
	XOLog(sqlstr, tfr.ID, tfr.QuestionID, tfr.AuthorID, tfr.Response)
	err = db.QueryRow(sqlstr, tfr.ID, tfr.QuestionID, tfr.AuthorID, tfr.Response).Scan(&tfr.ID)
	if err != nil {
		return err
	}

	// set existence
	tfr._exists = true

	return nil
}

// Update updates the TrueFalseResponse in the database.
func (tfr *TrueFalseResponse) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !tfr._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if tfr._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.true_false_responses SET (` +
		`question_id, author_id, response` +
		`) = ( ` +
		`$1, $2, $3` +
		`) WHERE id = $4`

	// run query
	XOLog(sqlstr, tfr.QuestionID, tfr.AuthorID, tfr.Response, tfr.ID)
	_, err = db.Exec(sqlstr, tfr.QuestionID, tfr.AuthorID, tfr.Response, tfr.ID)
	return err
}

// Save saves the TrueFalseResponse to the database.
func (tfr *TrueFalseResponse) Save(db XODB) error {
	if tfr.Exists() {
		return tfr.Update(db)
	}

	return tfr.Insert(db)
}

// Upsert performs an upsert for TrueFalseResponse.
//
// NOTE: PostgreSQL 9.5+ only
func (tfr *TrueFalseResponse) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if tfr._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.true_false_responses (` +
		`id, question_id, author_id, response` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, question_id, author_id, response` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.question_id, EXCLUDED.author_id, EXCLUDED.response` +
		`)`

	// run query
	XOLog(sqlstr, tfr.ID, tfr.QuestionID, tfr.AuthorID, tfr.Response)
	_, err = db.Exec(sqlstr, tfr.ID, tfr.QuestionID, tfr.AuthorID, tfr.Response)
	if err != nil {
		return err
	}

	// set existence
	tfr._exists = true

	return nil
}

// Delete deletes the TrueFalseResponse from the database.
func (tfr *TrueFalseResponse) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !tfr._exists {
		return nil
	}

	// if deleted, bail
	if tfr._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.true_false_responses WHERE id = $1`

	// run query
	XOLog(sqlstr, tfr.ID)
	_, err = db.Exec(sqlstr, tfr.ID)
	if err != nil {
		return err
	}

	// set deleted
	tfr._deleted = true

	return nil
}

// TrueFalseQuestionByAuthorID returns the TrueFalseQuestion associated with the TrueFalseResponse's AuthorID (author_id).
//
// Generated from foreign key 'true_false_author_fkey'.
func (tfr *TrueFalseResponse) TrueFalseQuestionByAuthorID(db XODB) (*TrueFalseQuestion, error) {
	return TrueFalseQuestionByID(db, tfr.AuthorID.Int64)
}

// TrueFalseQuestionByQuestionID returns the TrueFalseQuestion associated with the TrueFalseResponse's QuestionID (question_id).
//
// Generated from foreign key 'true_false_question_fkey'.
func (tfr *TrueFalseResponse) TrueFalseQuestionByQuestionID(db XODB) (*TrueFalseQuestion, error) {
	return TrueFalseQuestionByID(db, tfr.QuestionID)
}

// TrueFalseResponseByID retrieves a row from 'public.true_false_responses' as a TrueFalseResponse.
//
// Generated from index 'true_false_responses_pkey'.
func TrueFalseResponseByID(db XODB, id int64) (*TrueFalseResponse, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, question_id, author_id, response ` +
		`FROM public.true_false_responses ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	tfr := TrueFalseResponse{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&tfr.ID, &tfr.QuestionID, &tfr.AuthorID, &tfr.Response)
	if err != nil {
		return nil, err
	}

	return &tfr, nil
}
