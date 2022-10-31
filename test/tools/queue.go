package tools

type Queue []int

//这里需要使用指针
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

//移除元素
func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

//判断是否为空
func (q *Queue) IsImpty() bool {
	return len((*q)) == 0
}
