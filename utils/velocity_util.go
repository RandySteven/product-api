package utils

func CheckOldAndNew(old, new string) string {
	if new == "" {
		return old
	}
	return new
}
