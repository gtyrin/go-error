package error

import "github.com/pkg/errors"

// Panic проверяет наличие ошибки и эскалирует ее на уровень выше.
func Panic(e error, context string) {
	if e != nil {
		if len(context) > 0 {
			e = errors.Wrap(e, context)
		}
		panic(e)
	}
}
