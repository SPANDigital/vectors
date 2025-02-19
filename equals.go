package vectors

func (a Vector) Equals(b Vector) bool {
	if len(a) != len(b) {
		return false
	}
	for i, component := range a {
		if component != b[i] {
			return false
		}
	}
	return true
}
