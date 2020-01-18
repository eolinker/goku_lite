package pdao

import (
	"fmt"
	"testing"

	test1 "github.com/eolinker/goku-api-gateway/common/pdao/test/test/test/test"
)

type T1 struct {
}

func (t *T1) Test4() {
	fmt.Println("implement me")
}

type T2 struct {
}

func (t *T2) Test3() {
	fmt.Println("implement me")
}

type T3 struct {
}

func (t *T3) Test2() {

	fmt.Println("implement me")
}

type T4 struct {
}

func (t *T4) Test1() {
	fmt.Println("implement me")
}

func Test(t *testing.T) {

	t.Run("all", func(t *testing.T) {
		var t1 test1.Test
		//var t2 test2.Test = nil
		//var t3 test3.Test  = nil
		//var t4 test4.Test = nil
		Need(&t1)
		//Need(&t2)
		//Need(&t3)
		//Need(&t4)

		var t1v test1.Test = new(T4)
		Set(&t1v)
		//Set(test2.Test(new(T3)))
		//
		//Set(test3.Test(new(T2)))
		//Set(test4.Test(new(T1)))

		t1.Test1()
		//t2.Test2()
		//t3.Test3()
		//t4.Test4()
		Check()

	})

}
