package database

import (
	"database/sql"
	sqliteMigrator "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	_ "modernc.org/sqlite"
)

// ModernCDriver implements GORM's dialector interface for modernc.org/sqlite
type ModernCDriver struct {
	Conn      *sql.DB
	DSN       string
	Dialector *sqliteMigrator.Dialector // For migrations
}

// Open creates a new database connection
func OpenModernC(dsn string) gorm.Dialector {
	return &ModernCDriver{
		DSN:       dsn,
		Dialector: sqliteMigrator.Open(dsn).(*sqliteMigrator.Dialector),
	}
}

func (d *ModernCDriver) Name() string {
	return "sqlite"
}

func (d *ModernCDriver) Initialize(db *gorm.DB) (err error) {
	// Initialize the CGO dialector first for migrations
	if err := d.Dialector.Initialize(db); err != nil {
		return err
	}

	// Then initialize our modernc connection
	d.Conn, err = sql.Open("sqlite", d.DSN)
	if err != nil {
		return err
	}

	// Set connection pool settings
	d.Conn.SetMaxOpenConns(1) // SQLite only supports one connection
	d.Conn.SetMaxIdleConns(1)
	db.ConnPool = d.Conn

	return nil
}

func (d *ModernCDriver) Migrator(db *gorm.DB) gorm.Migrator {
	return d.Dialector.Migrator(db)
}

// DataTypeOf returns field's db type
func (d *ModernCDriver) DataTypeOf(field *schema.Field) string {
	switch field.DataType {
	case schema.Bool:
		return "boolean"
	case schema.Int, schema.Uint:
		return "integer"
	case schema.Float:
		return "real"
	case schema.String:
		return "text"
	case schema.Time:
		return "datetime"
	case schema.Bytes:
		return "blob"
	}
	return string(field.DataType)
}

// DefaultValueOf returns field's default value
func (d *ModernCDriver) DefaultValueOf(field *schema.Field) clause.Expression {
	if field.HasDefaultValue && field.DefaultValueInterface != nil {
		return clause.Expr{SQL: "?", Vars: []interface{}{field.DefaultValueInterface}}
	}
	return clause.Expr{SQL: "NULL"}
}

// BindVarTo binds value to stmt
func (d *ModernCDriver) BindVarTo(writer clause.Writer, stmt *gorm.Statement, v interface{}) {
	writer.WriteByte('?')
}

// QuoteTo quotes value to writer
func (d *ModernCDriver) QuoteTo(writer clause.Writer, str string) {
	writer.WriteByte('`')
	writer.WriteString(str)
	writer.WriteByte('`')
}

// Explain returns explain information
func (d *ModernCDriver) Explain(sql string, vars ...interface{}) string {
	return "EXPLAIN QUERY PLAN " + sql
}
