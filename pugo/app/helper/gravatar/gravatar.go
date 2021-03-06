package gravatar

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// FromEmail convert email to gravatar image path
func FromEmail(email string, size int) string {
	if size <= 0 {
		size = 80
	}
	return fmt.Sprintf("https://www.gravatar.com/avatar/%s?size=%d", email2Hash(email), size)
}

func email2Hash(email string) string {
	hasher := md5.New()
	hasher.Write([]byte(email))
	return hex.EncodeToString(hasher.Sum(nil))
}
