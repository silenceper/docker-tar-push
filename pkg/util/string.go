package util

import "strings"

//ParseImageAndTag parse image repo
func ParseImageAndTag(repo string) (string, string) {
	arr := strings.Split(repo, ":")
	return arr[0], arr[1]
}
