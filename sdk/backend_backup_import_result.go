
// backend_backup_import_result automatically generated by SDKgen please do not edit this file manually
// @see https://sdkgen.app


package sdk

// Result of a backup import operation
type BackendBackupImportResult struct {
    Success bool `json:"success"`
    Message string `json:"message"`
    Logs []string `json:"logs"`
}

