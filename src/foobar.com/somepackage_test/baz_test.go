package somepackage

import "errors"

func Foo() error {
	return errors.New("hello")
}

func bar() {
	foo()
}
