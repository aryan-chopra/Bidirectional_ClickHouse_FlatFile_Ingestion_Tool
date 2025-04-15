package utils

import "time"

func InferTypes(row []any) []string {

	types := make([]string, len(row))

	for columnIndex, value := range row {
		switch v := value.(type) {
		case bool:
			types[columnIndex] = "Bool"
		case float64:
			types[columnIndex] = "Float64"
		case string:
			if _, err := time.Parse("2006-01-02 15:04:05", v); err == nil {
				types[columnIndex] = "DateTime"
			} else if _, err := time.Parse("2006-01-02", v); err == nil {
				types[columnIndex] = "Date"
			} else if _, err := time.Parse(time.RFC3339, v); err == nil {
				types[columnIndex] = "DateTime"
			} else {
				types[columnIndex] = "String"
			}
		default:
			types[columnIndex] = "String"
		}
	}

	return types
}
