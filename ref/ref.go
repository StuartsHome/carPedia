package ref

// helper functions for creating structs and filling pointer variables with constants

func Bool(b bool) *bool {
	return &b
}

func Int(i int) *int {
	return &i
}

func String(s string) *string {
	return &s
}
