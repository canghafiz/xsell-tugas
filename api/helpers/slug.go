package helpers

import (
	"regexp"
	"strings"
)

func GenerateSlug(s string) string {
	if s == "" {
		return ""
	}
	slug := strings.ToLower(s)
	reg, _ := regexp.Compile("[^a-z0-9\\s]+")
	slug = reg.ReplaceAllString(slug, " ")
	slug = strings.Join(strings.Fields(slug), "_")
	return slug
}
