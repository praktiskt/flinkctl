package tools

func GetKeys(m map[string]interface{}) []string {
	k := []string{}
	for key := range m {
		k = append(k, key)
	}

	return k
}

func GetKey(m map[string]interface{}) string {
	keys := GetKeys(m)
	if len(keys) > 0 {
		return keys[0]
	}
	return ""
}
