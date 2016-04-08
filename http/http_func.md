从头开始比较pattern是否和path匹配
匹配逻辑
* pattern是否为空,为空 false
* pattern不是以/结尾，直接比对 pattern和path 不等 false
* 判断 path的长度大于等于pattern，并且 pattern和path的前pattern个长度的字符比较相等true

func pathMatch(pattern,path string) bool {
	if len(pattern) == 0 {
		return false
	}

	n := len(pattern)
	if pattern[n-1] != '/' {
		return pattern == path
	} 
	return len(path) >= n && path[0:n] == pattern
}




