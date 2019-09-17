package queue

type GokuQueue interface {
	Size() int                //大小
	Front() interface{}       //第一个元素
	End() interface{}         //最后一个元素
	IsEmpty() bool            //是否为空
	Enqueue(data interface{}) //入队
	Dequeue() interface{}     //出对
	Clear()                   //清空
}

type Queue struct {
	datastore []interface{}
	theSize   int
}

func (goku_queue *Queue) Clear() {
	goku_queue.datastore = make([]interface{}, 0) //开辟内存
	goku_queue.theSize = 0
}

func NewQueue() *Queue {
	goku_queue := new(Queue)
	goku_queue.Clear()
	return goku_queue

}

func (goku_queue *Queue) Size() int {
	return goku_queue.theSize //大小
}

func (goku_queue *Queue) Front() interface{} {
	if goku_queue.Size() == 0 { //判断是否为空
		return nil
	}
	return goku_queue.datastore[0]
}

func (goku_queue *Queue) End() interface{} {
	if goku_queue.Size() == 0 { //判断是否为空
		return nil
	}
	return goku_queue.datastore[goku_queue.theSize-1]
}

func (goku_queue *Queue) IsEmpty() bool {
	return goku_queue.theSize == 0
}

func (goku_queue *Queue) Enqueue(data interface{}) {
	goku_queue.datastore = append(goku_queue.datastore, data) //入队
	goku_queue.theSize = goku_queue.theSize + 1
}

func (goku_queue *Queue) Dequeue() interface{} {
	size := len(goku_queue.datastore)
	if size < 1 { //判断是否为空
		return nil
	}
	datastore := make([]interface{}, 0)
	data := goku_queue.datastore[0]
	if goku_queue.theSize > 1 {
		datastore = goku_queue.datastore[1:] //截取
	}
	goku_queue.datastore = datastore
	goku_queue.theSize = goku_queue.theSize - 1
	return data
}
