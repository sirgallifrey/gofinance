package ui

func CombineClasses(classes []any, rest ...any) []any {
	if classes == nil {
		return rest
	}
	return append(classes, rest)
}
