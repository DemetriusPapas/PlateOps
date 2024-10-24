// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: food_registry.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createFoodEntry = `-- name: CreateFoodEntry :one
INSERT INTO food_registry (
    name, 
    GTIN, 
    category, 
    description, 
    unit_type, 
    quantity,
    portion_count,
    expiration_after_opening,
    nutrients_per_item,
    calories,
    fats,
    saturated,
    carbs,
    sugars,
    protein,
    fiber,
    sodium,
    created_by
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18
) RETURNING product_id, name, gtin, category, description, unit_type, quantity, portion_count, expiration_after_opening, nutrients_per_item, calories, fats, saturated, carbs, sugars, protein, fiber, sodium, created_at, updated_at, created_by, updated_by
`

type CreateFoodEntryParams struct {
	Name                   string         `json:"name"`
	Gtin                   pgtype.Text    `json:"gtin"`
	Category               pgtype.Text    `json:"category"`
	Description            pgtype.Text    `json:"description"`
	UnitType               string         `json:"unit_type"`
	Quantity               pgtype.Int4    `json:"quantity"`
	PortionCount           pgtype.Int4    `json:"portion_count"`
	ExpirationAfterOpening pgtype.Int4    `json:"expiration_after_opening"`
	NutrientsPerItem       pgtype.Bool    `json:"nutrients_per_item"`
	Calories               pgtype.Numeric `json:"calories"`
	Fats                   pgtype.Numeric `json:"fats"`
	Saturated              pgtype.Numeric `json:"saturated"`
	Carbs                  pgtype.Numeric `json:"carbs"`
	Sugars                 pgtype.Numeric `json:"sugars"`
	Protein                pgtype.Numeric `json:"protein"`
	Fiber                  pgtype.Numeric `json:"fiber"`
	Sodium                 pgtype.Numeric `json:"sodium"`
	CreatedBy              pgtype.UUID    `json:"created_by"`
}

func (q *Queries) CreateFoodEntry(ctx context.Context, arg CreateFoodEntryParams) (FoodRegistry, error) {
	row := q.db.QueryRow(ctx, createFoodEntry,
		arg.Name,
		arg.Gtin,
		arg.Category,
		arg.Description,
		arg.UnitType,
		arg.Quantity,
		arg.PortionCount,
		arg.ExpirationAfterOpening,
		arg.NutrientsPerItem,
		arg.Calories,
		arg.Fats,
		arg.Saturated,
		arg.Carbs,
		arg.Sugars,
		arg.Protein,
		arg.Fiber,
		arg.Sodium,
		arg.CreatedBy,
	)
	var i FoodRegistry
	err := row.Scan(
		&i.ProductID,
		&i.Name,
		&i.Gtin,
		&i.Category,
		&i.Description,
		&i.UnitType,
		&i.Quantity,
		&i.PortionCount,
		&i.ExpirationAfterOpening,
		&i.NutrientsPerItem,
		&i.Calories,
		&i.Fats,
		&i.Saturated,
		&i.Carbs,
		&i.Sugars,
		&i.Protein,
		&i.Fiber,
		&i.Sodium,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const deleteFoodEntry = `-- name: DeleteFoodEntry :exec
DELETE FROM food_registry
WHERE product_id = $1
`

func (q *Queries) DeleteFoodEntry(ctx context.Context, productID int32) error {
	_, err := q.db.Exec(ctx, deleteFoodEntry, productID)
	return err
}

const getFoodEntriesByCategory = `-- name: GetFoodEntriesByCategory :many
SELECT product_id, name, gtin, category, description, unit_type, quantity, portion_count, expiration_after_opening, nutrients_per_item, calories, fats, saturated, carbs, sugars, protein, fiber, sodium, created_at, updated_at, created_by, updated_by
FROM food_registry
WHERE category = $1
`

func (q *Queries) GetFoodEntriesByCategory(ctx context.Context, category pgtype.Text) ([]FoodRegistry, error) {
	rows, err := q.db.Query(ctx, getFoodEntriesByCategory, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FoodRegistry
	for rows.Next() {
		var i FoodRegistry
		if err := rows.Scan(
			&i.ProductID,
			&i.Name,
			&i.Gtin,
			&i.Category,
			&i.Description,
			&i.UnitType,
			&i.Quantity,
			&i.PortionCount,
			&i.ExpirationAfterOpening,
			&i.NutrientsPerItem,
			&i.Calories,
			&i.Fats,
			&i.Saturated,
			&i.Carbs,
			&i.Sugars,
			&i.Protein,
			&i.Fiber,
			&i.Sodium,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.UpdatedBy,
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

const getFoodEntriesByName = `-- name: GetFoodEntriesByName :many
SELECT product_id, name, gtin, category, description, unit_type, quantity, portion_count, expiration_after_opening, nutrients_per_item, calories, fats, saturated, carbs, sugars, protein, fiber, sodium, created_at, updated_at, created_by, updated_by
FROM food_registry
WHERE name = $1
`

func (q *Queries) GetFoodEntriesByName(ctx context.Context, name string) ([]FoodRegistry, error) {
	rows, err := q.db.Query(ctx, getFoodEntriesByName, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FoodRegistry
	for rows.Next() {
		var i FoodRegistry
		if err := rows.Scan(
			&i.ProductID,
			&i.Name,
			&i.Gtin,
			&i.Category,
			&i.Description,
			&i.UnitType,
			&i.Quantity,
			&i.PortionCount,
			&i.ExpirationAfterOpening,
			&i.NutrientsPerItem,
			&i.Calories,
			&i.Fats,
			&i.Saturated,
			&i.Carbs,
			&i.Sugars,
			&i.Protein,
			&i.Fiber,
			&i.Sodium,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.UpdatedBy,
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

const getFoodEntryByGtin = `-- name: GetFoodEntryByGtin :one
SELECT product_id, name, gtin, category, description, unit_type, quantity, portion_count, expiration_after_opening, nutrients_per_item, calories, fats, saturated, carbs, sugars, protein, fiber, sodium, created_at, updated_at, created_by, updated_by
FROM food_registry
WHERE GTIN = $1
`

func (q *Queries) GetFoodEntryByGtin(ctx context.Context, gtin pgtype.Text) (FoodRegistry, error) {
	row := q.db.QueryRow(ctx, getFoodEntryByGtin, gtin)
	var i FoodRegistry
	err := row.Scan(
		&i.ProductID,
		&i.Name,
		&i.Gtin,
		&i.Category,
		&i.Description,
		&i.UnitType,
		&i.Quantity,
		&i.PortionCount,
		&i.ExpirationAfterOpening,
		&i.NutrientsPerItem,
		&i.Calories,
		&i.Fats,
		&i.Saturated,
		&i.Carbs,
		&i.Sugars,
		&i.Protein,
		&i.Fiber,
		&i.Sodium,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const getFoodEntryById = `-- name: GetFoodEntryById :one
SELECT product_id, name, gtin, category, description, unit_type, quantity, portion_count, expiration_after_opening, nutrients_per_item, calories, fats, saturated, carbs, sugars, protein, fiber, sodium, created_at, updated_at, created_by, updated_by
FROM food_registry
WHERE product_id = $1
`

func (q *Queries) GetFoodEntryById(ctx context.Context, productID int32) (FoodRegistry, error) {
	row := q.db.QueryRow(ctx, getFoodEntryById, productID)
	var i FoodRegistry
	err := row.Scan(
		&i.ProductID,
		&i.Name,
		&i.Gtin,
		&i.Category,
		&i.Description,
		&i.UnitType,
		&i.Quantity,
		&i.PortionCount,
		&i.ExpirationAfterOpening,
		&i.NutrientsPerItem,
		&i.Calories,
		&i.Fats,
		&i.Saturated,
		&i.Carbs,
		&i.Sugars,
		&i.Protein,
		&i.Fiber,
		&i.Sodium,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
	)
	return i, err
}

const updateFoodEntry = `-- name: UpdateFoodEntry :exec
UPDATE food_registry
SET
    name = $2,
    GTIN = $3,
    category = $4, 
    description = $5, 
    unit_type = $6, 
    quantity = $7,
    portion_count = $8,
    expiration_after_opening = $9,
    nutrients_per_item = $10,
    calories = $11,
    fats = $12,
    saturated = $13,
    carbs = $14,
    sugars = $15,
    protein = $16,
    fiber = $17,
    sodium = $18,
    updated_at = CURRENT_TIMESTAMP,
    updated_by = $19
WHERE product_id = $1
`

type UpdateFoodEntryParams struct {
	ProductID              int32          `json:"product_id"`
	Name                   string         `json:"name"`
	Gtin                   pgtype.Text    `json:"gtin"`
	Category               pgtype.Text    `json:"category"`
	Description            pgtype.Text    `json:"description"`
	UnitType               string         `json:"unit_type"`
	Quantity               pgtype.Int4    `json:"quantity"`
	PortionCount           pgtype.Int4    `json:"portion_count"`
	ExpirationAfterOpening pgtype.Int4    `json:"expiration_after_opening"`
	NutrientsPerItem       pgtype.Bool    `json:"nutrients_per_item"`
	Calories               pgtype.Numeric `json:"calories"`
	Fats                   pgtype.Numeric `json:"fats"`
	Saturated              pgtype.Numeric `json:"saturated"`
	Carbs                  pgtype.Numeric `json:"carbs"`
	Sugars                 pgtype.Numeric `json:"sugars"`
	Protein                pgtype.Numeric `json:"protein"`
	Fiber                  pgtype.Numeric `json:"fiber"`
	Sodium                 pgtype.Numeric `json:"sodium"`
	UpdatedBy              pgtype.UUID    `json:"updated_by"`
}

func (q *Queries) UpdateFoodEntry(ctx context.Context, arg UpdateFoodEntryParams) error {
	_, err := q.db.Exec(ctx, updateFoodEntry,
		arg.ProductID,
		arg.Name,
		arg.Gtin,
		arg.Category,
		arg.Description,
		arg.UnitType,
		arg.Quantity,
		arg.PortionCount,
		arg.ExpirationAfterOpening,
		arg.NutrientsPerItem,
		arg.Calories,
		arg.Fats,
		arg.Saturated,
		arg.Carbs,
		arg.Sugars,
		arg.Protein,
		arg.Fiber,
		arg.Sodium,
		arg.UpdatedBy,
	)
	return err
}
