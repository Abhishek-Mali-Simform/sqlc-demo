// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package models

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const countLookupByDisplayText = `-- name: CountLookupByDisplayText :one
SELECT count(*) FROM lookups WHERE display_text ILIKE $1
`

func (q *Queries) CountLookupByDisplayText(ctx context.Context, displayText string) (int64, error) {
	row := q.db.QueryRow(ctx, countLookupByDisplayText, displayText)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createLookup = `-- name: CreateLookup :one
INSERT INTO lookups(
	lookup_id, table_name, display_order, display_text, is_active, parent_id, internal_key, concurrency_key, create_date, create_user_id, update_date, update_user_id, value_text)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING lookup_id, table_name, display_order, display_text, is_active, parent_id, internal_key, concurrency_key, create_date, create_user_id, update_date, update_user_id, value_text
`

type CreateLookupParams struct {
	LookupID       uuid.UUID
	TableName      string
	DisplayOrder   int32
	DisplayText    string
	IsActive       bool
	ParentID       uuid.NullUUID
	InternalKey    string
	ConcurrencyKey string
	CreateDate     time.Time
	CreateUserID   int32
	UpdateDate     *time.Time
	UpdateUserID   pgtype.Int4
	ValueText      string
}

func (q *Queries) CreateLookup(ctx context.Context, arg CreateLookupParams) (Lookup, error) {
	row := q.db.QueryRow(ctx, createLookup,
		arg.LookupID,
		arg.TableName,
		arg.DisplayOrder,
		arg.DisplayText,
		arg.IsActive,
		arg.ParentID,
		arg.InternalKey,
		arg.ConcurrencyKey,
		arg.CreateDate,
		arg.CreateUserID,
		arg.UpdateDate,
		arg.UpdateUserID,
		arg.ValueText,
	)
	var i Lookup
	err := row.Scan(
		&i.LookupID,
		&i.TableName,
		&i.DisplayOrder,
		&i.DisplayText,
		&i.IsActive,
		&i.ParentID,
		&i.InternalKey,
		&i.ConcurrencyKey,
		&i.CreateDate,
		&i.CreateUserID,
		&i.UpdateDate,
		&i.UpdateUserID,
		&i.ValueText,
	)
	return i, err
}

const getLookup = `-- name: GetLookup :one
SELECT lookup_id, table_name, display_order, display_text, is_active, parent_id, internal_key, concurrency_key, create_date, create_user_id, update_date, update_user_id, value_text FROM lookups WHERE lookup_id = $1
`

func (q *Queries) GetLookup(ctx context.Context, lookupID uuid.UUID) (Lookup, error) {
	row := q.db.QueryRow(ctx, getLookup, lookupID)
	var i Lookup
	err := row.Scan(
		&i.LookupID,
		&i.TableName,
		&i.DisplayOrder,
		&i.DisplayText,
		&i.IsActive,
		&i.ParentID,
		&i.InternalKey,
		&i.ConcurrencyKey,
		&i.CreateDate,
		&i.CreateUserID,
		&i.UpdateDate,
		&i.UpdateUserID,
		&i.ValueText,
	)
	return i, err
}

const getLookupByID = `-- name: GetLookupByID :one
SELECT table_name, display_order, display_text, is_active FROM lookups WHERE lookup_id = $1
`

type GetLookupByIDRow struct {
	TableName    string
	DisplayOrder int32
	DisplayText  string
	IsActive     bool
}

func (q *Queries) GetLookupByID(ctx context.Context, lookupID uuid.UUID) (GetLookupByIDRow, error) {
	row := q.db.QueryRow(ctx, getLookupByID, lookupID)
	var i GetLookupByIDRow
	err := row.Scan(
		&i.TableName,
		&i.DisplayOrder,
		&i.DisplayText,
		&i.IsActive,
	)
	return i, err
}

const listLookups = `-- name: ListLookups :many
SELECT table_name, display_order, display_text, is_active FROM lookups
`

type ListLookupsRow struct {
	TableName    string
	DisplayOrder int32
	DisplayText  string
	IsActive     bool
}

func (q *Queries) ListLookups(ctx context.Context) ([]ListLookupsRow, error) {
	rows, err := q.db.Query(ctx, listLookups)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLookupsRow
	for rows.Next() {
		var i ListLookupsRow
		if err := rows.Scan(
			&i.TableName,
			&i.DisplayOrder,
			&i.DisplayText,
			&i.IsActive,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLookupsByDisplayText = `-- name: ListLookupsByDisplayText :many
SELECT table_name, display_order, display_text, is_active FROM lookups WHERE display_text ILIKE $1
`

type ListLookupsByDisplayTextRow struct {
	TableName    string
	DisplayOrder int32
	DisplayText  string
	IsActive     bool
}

func (q *Queries) ListLookupsByDisplayText(ctx context.Context, displayText string) ([]ListLookupsByDisplayTextRow, error) {
	rows, err := q.db.Query(ctx, listLookupsByDisplayText, displayText)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLookupsByDisplayTextRow
	for rows.Next() {
		var i ListLookupsByDisplayTextRow
		if err := rows.Scan(
			&i.TableName,
			&i.DisplayOrder,
			&i.DisplayText,
			&i.IsActive,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLookupsByTableName = `-- name: ListLookupsByTableName :many
SELECT table_name, display_order, display_text, is_active FROM lookups WHERE table_name ILIKE $1
`

type ListLookupsByTableNameRow struct {
	TableName    string
	DisplayOrder int32
	DisplayText  string
	IsActive     bool
}

func (q *Queries) ListLookupsByTableName(ctx context.Context, tableName string) ([]ListLookupsByTableNameRow, error) {
	rows, err := q.db.Query(ctx, listLookupsByTableName, tableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListLookupsByTableNameRow
	for rows.Next() {
		var i ListLookupsByTableNameRow
		if err := rows.Scan(
			&i.TableName,
			&i.DisplayOrder,
			&i.DisplayText,
			&i.IsActive,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const newLookup = `-- name: NewLookup :one
INSERT INTO lookups(
	table_name, display_order, display_text, is_active, create_date, create_user_id, value_text)
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING lookup_id, table_name, display_order, display_text, is_active, parent_id, internal_key, concurrency_key, create_date, create_user_id, update_date, update_user_id, value_text
`

type NewLookupParams struct {
	TableName    string
	DisplayOrder int32
	DisplayText  string
	IsActive     bool
	CreateDate   time.Time
	CreateUserID int32
	ValueText    string
}

func (q *Queries) NewLookup(ctx context.Context, arg NewLookupParams) (Lookup, error) {
	row := q.db.QueryRow(ctx, newLookup,
		arg.TableName,
		arg.DisplayOrder,
		arg.DisplayText,
		arg.IsActive,
		arg.CreateDate,
		arg.CreateUserID,
		arg.ValueText,
	)
	var i Lookup
	err := row.Scan(
		&i.LookupID,
		&i.TableName,
		&i.DisplayOrder,
		&i.DisplayText,
		&i.IsActive,
		&i.ParentID,
		&i.InternalKey,
		&i.ConcurrencyKey,
		&i.CreateDate,
		&i.CreateUserID,
		&i.UpdateDate,
		&i.UpdateUserID,
		&i.ValueText,
	)
	return i, err
}

const newLookupWithConcurrencyKey = `-- name: NewLookupWithConcurrencyKey :one
INSERT INTO lookups(
	table_name, display_order, display_text, is_active, internal_key, concurrency_key, create_date, create_user_id, value_text)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING lookup_id, table_name, display_order, display_text, is_active, parent_id, internal_key, concurrency_key, create_date, create_user_id, update_date, update_user_id, value_text
`

type NewLookupWithConcurrencyKeyParams struct {
	TableName      string
	DisplayOrder   int32
	DisplayText    string
	IsActive       bool
	InternalKey    string
	ConcurrencyKey string
	CreateDate     time.Time
	CreateUserID   int32
	ValueText      string
}

func (q *Queries) NewLookupWithConcurrencyKey(ctx context.Context, arg NewLookupWithConcurrencyKeyParams) (Lookup, error) {
	row := q.db.QueryRow(ctx, newLookupWithConcurrencyKey,
		arg.TableName,
		arg.DisplayOrder,
		arg.DisplayText,
		arg.IsActive,
		arg.InternalKey,
		arg.ConcurrencyKey,
		arg.CreateDate,
		arg.CreateUserID,
		arg.ValueText,
	)
	var i Lookup
	err := row.Scan(
		&i.LookupID,
		&i.TableName,
		&i.DisplayOrder,
		&i.DisplayText,
		&i.IsActive,
		&i.ParentID,
		&i.InternalKey,
		&i.ConcurrencyKey,
		&i.CreateDate,
		&i.CreateUserID,
		&i.UpdateDate,
		&i.UpdateUserID,
		&i.ValueText,
	)
	return i, err
}

const updateLookup = `-- name: UpdateLookup :one
UPDATE lookups
SET
    table_name = COALESCE(NULLIF($1, ''), table_name),
    display_order = COALESCE($2, display_order),
    display_text = COALESCE(NULLIF($3, ''), display_text),
    is_active = COALESCE($4, is_active),
    parent_id = COALESCE($5, parent_id),
    internal_key = COALESCE(NULLIF($6, ''), internal_key),
    concurrency_key = COALESCE(NULLIF($7, ''), concurrency_key),
    update_date = $8,
    update_user_id = $9,
    value_text = COALESCE(NULLIF($10, ''), value_text)
WHERE
    lookup_id = $11
    RETURNING lookup_id, table_name, display_order, display_text, is_active, parent_id, internal_key, concurrency_key, create_date, create_user_id, update_date, update_user_id, value_text
`

type UpdateLookupParams struct {
	Column1      interface{}
	DisplayOrder int32
	Column3      interface{}
	IsActive     bool
	ParentID     uuid.NullUUID
	Column6      interface{}
	Column7      interface{}
	UpdateDate   *time.Time
	UpdateUserID pgtype.Int4
	Column10     interface{}
	LookupID     uuid.UUID
}

func (q *Queries) UpdateLookup(ctx context.Context, arg UpdateLookupParams) (Lookup, error) {
	row := q.db.QueryRow(ctx, updateLookup,
		arg.Column1,
		arg.DisplayOrder,
		arg.Column3,
		arg.IsActive,
		arg.ParentID,
		arg.Column6,
		arg.Column7,
		arg.UpdateDate,
		arg.UpdateUserID,
		arg.Column10,
		arg.LookupID,
	)
	var i Lookup
	err := row.Scan(
		&i.LookupID,
		&i.TableName,
		&i.DisplayOrder,
		&i.DisplayText,
		&i.IsActive,
		&i.ParentID,
		&i.InternalKey,
		&i.ConcurrencyKey,
		&i.CreateDate,
		&i.CreateUserID,
		&i.UpdateDate,
		&i.UpdateUserID,
		&i.ValueText,
	)
	return i, err
}
