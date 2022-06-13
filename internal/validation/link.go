package validation

import "regexp"

type linkValidator struct {
	linkRegExp *regexp.Regexp
}

func (lv *linkValidator) ValidLink(link string) (bool, string) {
	if lv.linkRegExp.MatchString(link) {
		return true, ""
	}

	return false, linkNotValid()
}

func New() *linkValidator {
	return &linkValidator{linkRegExp: regexp.MustCompile(`^(http)|(https)://\w+\.\w{2,}$`)}
}
