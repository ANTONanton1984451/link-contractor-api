package validation

func linkNotValid() string {
	return `ссылка должна начинаться с http или https и иметь домен с двумя или больше символами`
}

func pathNotValid() string {
	return `идентификатор должен содержать только буквы верхнего и нижнего регистров, цифры, символы нижнего подчёркивания и тире`
}
