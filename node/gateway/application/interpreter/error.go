package interpreter

import "fmt"

type GrammarError string

func (e GrammarError)Error()string  {
	return fmt.Sprint("Invalid:",string(e))
}
func (e GrammarError)String()string  {
	return e.Error()
}
