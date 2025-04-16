package utils

// "fmt"
// "time"

// InferTypes is a utility function that takes a row of data and infers the types of the values.
// It returns a slice of strings representing the inferred types of each column in the row.
// The types are represented as strings (e.g., "Bool", "Float64", "String").
//
// Parameters:
// - row: A slice of any type representing a single row of data.
//
// Returns:
// - []string: A slice of strings representing the inferred types of each column.
func InferTypes(row []any) []string {

	types := make([]string, len(row))

	// Loop through each column in the row to infer its type
	for columnIndex, _ := range row {
		// switch v := value.(type) {
		// case bool:
		// 	types[columnIndex] = "Bool"
		// case float64:
		// 	types[columnIndex] = "Float64"
		// case string:
		// 	if _, err := time.Parse("2006-01-02 15:04:05", v); err == nil {
		// 		types[columnIndex] = "DateTime"
		// 	} else if _, err := time.Parse("2006-01-02", v); err == nil {
		// 		types[columnIndex] = "Date"
		// 	} else if _, err := time.Parse(time.RFC3339, v); err == nil {
		// 		types[columnIndex] = "DateTime"
		// 	} else {
		// 		types[columnIndex] = "String"
		// 	}
		// default:
		// 	types[columnIndex] = "String"
		// }

		types[columnIndex] = "String"
	}

	return types
}
