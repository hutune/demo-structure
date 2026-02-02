package utils

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"

	libDecimal "github.com/shopspring/decimal"
)

// Decimal type
type Decimal libDecimal.Decimal

var DecimalZero = Decimal(libDecimal.Zero)

// Helper to convert to decimal.Decimal
func (d *Decimal) unwrap() libDecimal.Decimal {
	return libDecimal.Decimal(*d)
}

// region: ======= JSON support =======

func (d Decimal) MarshalJSON() ([]byte, error) {
	f, _ := d.unwrap().Round(2).Float64()
	return json.Marshal(f)
}

func (d *Decimal) UnmarshalJSON(data []byte) error {
	f, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		*d = Decimal(libDecimal.Zero)
		return nil
	}
	*d = Decimal(libDecimal.NewFromFloat(f))
	return nil
}

// endregion: ======= JSON support =======

// region: ======= GORM support =======

// Implement Valuer for GORM (database storage)
func (d Decimal) Value() (driver.Value, error) {
	return d.unwrap().String(), nil // Store as string in DB
}

// Implement Scanner for GORM (database retrieval)
func (d *Decimal) Scan(value interface{}) error {
	var dec libDecimal.Decimal
	var err error

	switch v := value.(type) {
	case string:
		dec, err = libDecimal.NewFromString(v)
	case float64:
		dec = libDecimal.NewFromFloat(v)
	case []byte:
		dec, err = libDecimal.NewFromString(string(v))
	default:
		return fmt.Errorf("unsupported scan type for Decimal: %T", value)
	}

	if err != nil {
		return err
	}

	*d = Decimal(dec)
	return nil
}

// endregion: ======= GORM support =======

// region: ======= Math decimal support =======

func (d Decimal) Add(other Decimal) Decimal {
	return Decimal(d.unwrap().Add(other.unwrap()))
}

func (d Decimal) Sub(other Decimal) Decimal {
	return Decimal(d.unwrap().Sub(other.unwrap()))
}

// Mul returns d * d2
func (d Decimal) Mul(other Decimal) Decimal {
	return Decimal(d.unwrap().Mul(other.unwrap()))
}

// Div returns d / d2
func (d Decimal) Div(other Decimal) Decimal {
	return Decimal(d.unwrap().Div(other.unwrap()))
}

// endregion: ======= Math decimal support =======

// region: ======= Parse decimal support =======

// Float64 returns the nearest float64 value for d and a bool indicating whether f represents d exactly. For more details, see the documentation for big.Rat.Float64
func (d Decimal) Float64() (f float64, isExact bool) {
	return d.unwrap().Float64()
}

// InexactFloat64 returns the nearest float64 value for d. It doesn't indicate if the returned value represents d exactly.
func (d Decimal) InexactFloat64() (f float64) {
	return d.unwrap().InexactFloat64()
}

// NewFromString returns a new Decimal from a string representation. Trailing zeroes are not trimmed.
func DecimalNewFromString(s string) (Decimal, error) {
	dec, err := libDecimal.NewFromString(s)
	return Decimal(dec), err
}

// NewFromFloat returns a new Decimal from a float64.
func DecimalNewFromFloat(f float64) Decimal {
	return Decimal(libDecimal.NewFromFloat(f))
}

// NewFromInt returns a new Decimal from an int64.
func DecimalNewFromInt(i int64) Decimal {
	return Decimal(libDecimal.New(i, 0))
}

func (d Decimal) String() string {
	return d.unwrap().String()
}

// endregion: ======= Parse decimal support =======

// region: ======= Other decimal support =======

func (d Decimal) IsZero() bool {
	return d == DecimalZero
}

func (d Decimal) IsPositive() bool {
	return d.unwrap().IsPositive()
}

func (d Decimal) IsNegative() bool {
	return d.unwrap().IsNegative()
}

// Floor returns the nearest integer value less than or equal to d.
func (d Decimal) Floor() Decimal {
	return Decimal(d.unwrap().Floor())
}

// Ceil returns the nearest integer value greater than or equal to d.
func (d Decimal) Ceil() Decimal {
	return Decimal(d.unwrap().Ceil())
}

// Cmp returns -1, 0, or 1 if d <, ==, or > d2.
func (d Decimal) Cmp(d2 Decimal) int {
	return d.unwrap().Cmp(d2.unwrap())
}

// endregion: ======= Other decimal support =======
