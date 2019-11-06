package interpreter

import "fmt"

//GrammarError grammarError
type GrammarError string

func (e GrammarError) Error() string {
	return fmt.Sprint("Invalid:", string(e))
}
func (e GrammarError) String() string {
	return e.Error()
}
