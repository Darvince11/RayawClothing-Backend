package utils

func MergeMap(m1, m2 *map[string]any) map[string]any {
	//This function appends m2 to m1
	result := make(map[string]any)
	for k, v := range *m1 {
		result[k] = v
	}
	for k, v := range *m2 {
		result[k] = v
	}
	return result
}
