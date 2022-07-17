package validation

import "regexp"

type linkValidator struct {
	linkRegExp *regexp.Regexp
	pathRegExp *regexp.Regexp
}

// ValidLink импелементация интерфейса юзкейса создания ссылки
func (lv *linkValidator) ValidLink(link string) (bool, string) {
	if lv.linkRegExp.MatchString(link) {
		return true, ""
	}

	return false, linkNotValid()
}

// ValidPath импелементация интерфейса юзкейса создания пути для укороченной ссылки
func (lv *linkValidator) ValidPath(path string) (bool, string) {
	if lv.pathRegExp.MatchString(path) {
		return true, ""
	}

	return false, pathNotValid()
}

func New() *linkValidator {
	return &linkValidator{
		linkRegExp: regexp.MustCompile(`^(http)|(https)://\w+\.\w{2,}$`),
		pathRegExp: regexp.MustCompile(`^[A-Za-z0-9_-]+$`),
	}
}
