package util

var ParisMap = make(map[string]interface{})

func ParisMap_Get(key string) interface{} {
	v, exist := ParisMap[key]
	if exist {
		return v
	} else {
		return nil
	}
}

func ParisMap_Put(key string, value interface{}) {
	ParisMap[key] = value
}

func ParisMap_Del(key string) {
	ParisMap[key] = nil
}

func ParisMap_DelAll() {
	ParisMap = make(map[string]interface{})
}
