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

func GetIntOrDefault(value *int, dflt int) int {
	if value != nil {
		return *value
	} else {
		return dflt
	}
}

func GetStringOrDefault(value *string, dflt string) string {
	if value != nil {
		return *value
	} else {
		return dflt
	}
}

func GetBoolOrDefault(value *bool, dflt bool) bool {
	if value != nil {
		return *value
	} else {
		return dflt
	}
}
