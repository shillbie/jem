package linego

func IsStrInMap(str string, targetMap map[string]int32) bool {
	_, ok := targetMap[str]
	return ok
}
