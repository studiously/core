// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Unit represents a row from 'public.units'.
type Unit struct {
	ID      int64          `json:"id"`       // id
	ClassID int64          `json:"class_id"` // class_id
	Name    sql.NullString `json:"name"`     // name
	Weight  int            `json:"weight"`   // weight

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Unit exists in the database.
func (u *Unit) Exists() bool {
	return u._exists
}

// Deleted provides information if the Unit has been deleted from the database.
func (u *Unit) Deleted() bool {
	return u._deleted
}

// Insert inserts the Unit to the database.
func (u *Unit) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO public.units (` +
		`id, class_id, name, weight` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)`

	// run query
	XOLog(sqlstr, u.ID, u.ClassID, u.Name, u.Weight)
	err = db.QueryRow(sqlstr, u.ID, u.ClassID, u.Name, u.Weight).Scan(&u.ID)
	if err != nil {
		return err
	}

	// set existence
	u._exists = true

	return nil
}

// Update updates the Unit in the database.
func (u *Unit) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if u._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.units SET (` +
		`class_id, name, weight` +
		`) = ( ` +
		`$1, $2, $3` +
		`) WHERE id = $4`

	// run query
	XOLog(sqlstr, u.ClassID, u.Name, u.Weight, u.ID)
	_, err = db.Exec(sqlstr, u.ClassID, u.Name, u.Weight, u.ID)
	return err
}

// Save saves the Unit to the database.
func (u *Unit) Save(db XODB) error {
	if u.Exists() {
		return u.Update(db)
	}

	return u.Insert(db)
}

// Upsert performs an upsert for Unit.
//
// NOTE: PostgreSQL 9.5+ only
func (u *Unit) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.units (` +
		`id, class_id, name, weight` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, class_id, name, weight` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.class_id, EXCLUDED.name, EXCLUDED.weight` +
		`)`

	// run query
	XOLog(sqlstr, u.ID, u.ClassID, u.Name, u.Weight)
	_, err = db.Exec(sqlstr, u.ID, u.ClassID, u.Name, u.Weight)
	if err != nil {
		return err
	}

	// set existence
	u._exists = true

	return nil
}

// Delete deletes the Unit from the database.
func (u *Unit) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return nil
	}

	// if deleted, bail
	if u._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.units WHERE id = $1`

	// run query
	XOLog(sqlstr, u.ID)
	_, err = db.Exec(sqlstr, u.ID)
	if err != nil {
		return err
	}

	// set deleted
	u._deleted = true

	return nil
}

// Class returns the Class associated with the Unit's ClassID (class_id).
//
// Generated from foreign key 'units_class_id_fkey'.
func (u *Unit) Class(db XODB) (*Class, error) {
	return ClassByID(db, u.ClassID)
}

// UnitByID retrieves a row from 'public.units' as a Unit.
//
// Generated from index 'units_pkey'.
func UnitByID(db XODB, id int64) (*Unit, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, class_id, name, weight ` +
		`FROM public.units ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	u := Unit{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&u.ID, &u.ClassID, &u.Name, &u.Weight)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
