
// backend_database_table_index automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

// This object represents a table index on a relational database
type BackendDatabaseTableIndex struct {
    Name string `json:"name"`
    Unique bool `json:"unique"`
    Columns []string `json:"columns"`
}

