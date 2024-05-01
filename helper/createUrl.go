package helper

import (
	"fmt"
	"os"
)

func CreateUrl(postfix string) string {
	siteDomain := os.Getenv("SITE_URL")
	return fmt.Sprintf("%v/%v",siteDomain,postfix)
}