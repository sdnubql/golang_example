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




