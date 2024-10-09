// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type FoodRegistry struct {
	ProductID              int32
	Name                   string
	Gtin                   pgtype.Text
	Category               pgtype.Text
	Description            pgtype.Text
	UnitType               string
	Quantity               pgtype.Int4
	PortionCount           pgtype.Int4
	ExpirationAfterOpening pgtype.Int4
	NutrientsPerItem       pgtype.Bool
	Calories               pgtype.Numeric
	Fats                   pgtype.Numeric
	Saturated              pgtype.Numeric
	Carbs                  pgtype.Numeric
	Sugars                 pgtype.Numeric
	Protein                pgtype.Numeric
	Fiber                  pgtype.Numeric
	Sodium                 pgtype.Numeric
	CreatedAt              pgtype.Timestamp
	UpdatedAt              pgtype.Timestamp
	CreatedBy              pgtype.UUID
	UodatedBy              pgtype.UUID
}

type Inventory struct {
	ID                     int32
	UserID                 pgtype.UUID
	ProductID              pgtype.Int4
	ExpirationDate         pgtype.Date
	ExpirationAfterOpening pgtype.Int4
	IsOpened               pgtype.Bool
	MaxQuantity            float64
	CurrQuantity           float64
	PurchaseDate           pgtype.Timestamp
	OpenedDate             pgtype.Timestamp
	CreatedAt              pgtype.Timestamp
	UpdatedAt              pgtype.Timestamp
}

type Recipe struct {
	ID                  int32
	UserID              pgtype.UUID
	Name                string
	Private             pgtype.Bool
	Description         pgtype.Text
	Instructions        pgtype.Text
	UnitType            string
	Quantity            pgtype.Int4
	NutrientsPerPortion pgtype.Bool
	Calories            pgtype.Numeric
	Fats                pgtype.Numeric
	Saturated           pgtype.Numeric
	Carbs               pgtype.Numeric
	Sugars              pgtype.Numeric
	Protein             pgtype.Numeric
	Fiber               pgtype.Numeric
	Sodium              pgtype.Numeric
	CreatedAt           pgtype.Timestamp
	UpdatedAt           pgtype.Timestamp
}

type RecipeIngredient struct {
	ID        int32
	RecipeID  pgtype.Int4
	ProductID pgtype.Int4
	Quantity  float64
}

type User struct {
	ID            pgtype.UUID
	Username      string
	Email         string
	PasswordHash  string
	CreatedAt     pgtype.Timestamp
	UpdatedAt     pgtype.Timestamp
	LastLogin     pgtype.Timestamp
	ActiveAccount pgtype.Bool
}