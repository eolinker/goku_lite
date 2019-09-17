package monitor_key

type MonitorValues []int64


func (a MonitorValues) Add(key MonitorKeyType)  {
	index := int(key)
	if index < len(a){
		a[index]++
	}
}
func (a MonitorValues)Get(key MonitorKeyType) int64  {
	if a==nil{
		return 0
	}
	index := int(key)
	if index < len(a){
		return a[index]
	}
	return 0
}
func (a MonitorValues)Append(args ...MonitorValues)  {
	if len(args) == 0 {
		return
	}

	for _,arg:=range args{

		for i:=range arg{
			a[i]+= arg[i]
		}
	}
	return
}

func MakeValue() MonitorValues {
	return make(MonitorValues,MonitorKeyTypeSize)
}