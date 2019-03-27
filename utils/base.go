package utils

//定义key value结构体，实现对map使用sort方法排序
type KeyValue struct {
	Key   int
	Value string
}

type KeyValues []KeyValue

func (u KeyValues) Len() int {
	return len(u)
}

func (u KeyValues) Less(i, j int) bool {
	return u[i].Key < u[j].Key
}

func (u KeyValues) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
