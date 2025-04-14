package utils

func InferTypes (row []any) []string {
	types := make([]string, len(row))
	
	for columnIndex, value := range row {
		switch value.(type) {
			case bool:
				types[columnIndex] = "Bool"
			case float64:
				types[columnIndex] = "Float64"
			case string:
				types[columnIndex] = "String"
			default:
				types[columnIndex] = "String"
		}
	}
	
	return types
}
