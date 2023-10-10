package sqlx

import (
	"database/sql"
	"time"
)

// ToNullString ...
func ToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

// ToNullInt32 ...
func ToNullInt32(i int32) sql.NullInt32 {
	return sql.NullInt32{
		Int32: i,
		Valid: i != 0,
	}
}

// ToNullInt64 ...
func ToNullInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: i,
		Valid: i != 0,
	}
}

// ToNullFloat64 ...
func ToNullFloat64(f float64) sql.NullFloat64 {
	return sql.NullFloat64{
		Float64: f,
		Valid:   f != 0,
	}
}

// ToNullBool ...
func ToNullBool(b bool) sql.NullBool {
	return sql.NullBool{
		Bool:  b,
		Valid: !b,
	}
}

// ToNullTime ...
func ToNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: !t.IsZero(),
	}
}
