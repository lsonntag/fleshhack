package config

var conf map[string]interface{}

func init() {
	conf = make(map[string]interface{})
}

func Get(s string) (interface{}, bool) {
	c, ok := conf[s]
	return c, ok
}

func Set(s string, v interface{}) {
	conf[s] = v
}

func GetString(s string) string {
	v, ok := Get(s)
	if !ok {
		return ""
	}

	return v.(string)
}

func GetBool(s string) bool {
	v, ok := Get(s)
	if !ok {
		return false
	}

	return v.(bool)
}
