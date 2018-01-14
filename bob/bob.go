package bob

import (
	"regexp"
	"strings"
)

func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	validChar := regexp.MustCompile(`[a-zA-Z0-9]`)
	validLower := regexp.MustCompile(`[a-z]`)
	validWord := regexp.MustCompile(`[a-zA-Z]`)
	validQuestion := regexp.MustCompile(`\?$`)

	hasLowerCase := validLower.MatchString(remark)
	hasWord := validWord.MatchString(remark)
	hasQuestion := validQuestion.MatchString(remark)
	hasChar := validChar.MatchString(remark)

	if hasWord && !hasLowerCase && hasQuestion {
		return "Calm down, I know what I'm doing!"
	}
	if hasWord && !hasLowerCase {
		return "Whoa, chill out!"
	}

	if hasQuestion {
		return "Sure."
	}
	if !hasChar {
		return "Fine. Be that way!"
	}

	return "Whatever."
}
