package ref

func RefValue[T any](v T) *T {
	return &v
}
