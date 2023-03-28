package filter

func Filter[T interface{}](records []T, f func(T) bool) (filtered []T) {
	for _, record := range records {
		if f(record) {
			filtered = append(filtered, record)
		}
	}
	return
}

func OmitNil[T interface{}](records []*T) []*T {
	return Filter(records, func(v *T) bool { return v != nil })
}
