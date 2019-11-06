package goku_template

//Template template
type Template interface {
	Template() interface{}
	Encode(v interface{}) (string, error)
	Decode(org string) (interface{}, error)
}
