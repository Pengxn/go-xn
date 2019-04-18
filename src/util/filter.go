package util

import "reflect"

func fieldSet(fields ...string) map[string]bool {
	set := make(map[string]bool, len(fields))

	for _, s := range fields {
		set[s] = true
	}

	return set
}

// SelectFields will filter your data
// func (s *SearchResult) SelectFields(fields ...string) map[string]interface{} {
func SelectFields(s interface{}, fields ...string) map[string]interface{} {
	fs := fieldSet(fields...)
	rt, rv := reflect.TypeOf(s), reflect.ValueOf(s)
	out := make(map[string]interface{}, rt.NumField())

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		jsonKey := field.Tag.Get("json")

		if fs[jsonKey] {
			out[jsonKey] = rv.Field(i).Interface()
		}
	}

	return out
}
