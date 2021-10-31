package internal

// CompareUsers compares allowed and active users and returns the list of users to be added and deleted
// Since go doesn't have sets by default, using the technique mentioned in https://www.davidkaya.com/sets-in-golang/
// There could be better ways of doing this, but for around 200 users, this should be good enough
func CompareUsers(allowedUsers []string, activeUsers []string) ([]string, []string) {
	var usersToAdd, usersToRemove []string
	var allowedSet = make(map[string]struct{})
	var activeSet = make(map[string]struct{})
	var exists = struct{}{}
	for _, u := range allowedUsers {
		allowedSet[u] = exists
	}
	for _, u := range activeUsers {
		activeSet[u] = exists
	}
	for _, u := range allowedUsers {
		_, isPresent := activeSet[u]
		if !isPresent {
			usersToAdd = append(usersToAdd, u)
		}
	}
	for _, u := range activeUsers {
		_, isPresent := allowedSet[u]
		if !isPresent {
			usersToRemove = append(usersToRemove, u)
		}
	}
	return usersToAdd, usersToRemove
}
