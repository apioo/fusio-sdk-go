
// backend_database_table_foreign_key_constraint automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

// This object represents a foreign key constraint on a relational database
type BackendDatabaseTableForeignKeyConstraint struct {
    Name string `json:"name"`
    ForeignTable string `json:"foreignTable"`
    LocalColumnNames []string `json:"localColumnNames"`
    ForeignColumnNames []string `json:"foreignColumnNames"`
}

