package response

import "fmt"

func GetError(reason string) []byte {
	return []byte(fmt.Sprintf("{ \"status\": \"error\", \"reason\": \"%s\" }", reason))
}
