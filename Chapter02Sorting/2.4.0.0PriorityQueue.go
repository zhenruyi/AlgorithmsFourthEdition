package Chapter02Sorting

type MaxPriorityQueue struct {
	priorityQueue []int
	length int
}

func (pq *MaxPriorityQueue) IsEmpty() bool {
	return pq.length == 0
}

func (pq *MaxPriorityQueue) Size() int {
	return pq.length
}

// 上浮，在插入操作之后
func (pq *MaxPriorityQueue) swim(k int) {
	for k > 0 && less(pq.priorityQueue, k-1/2, k) {
		exch(pq.priorityQueue, k-1/2, k)
		k = k-1 / 2
	}
}

//下沉，恢复堆
func (pq *MaxPriorityQueue) sink(k int) {
	for 2*k+1 < pq.length {
		j := 2*k + 1
		if j+1 < pq.length && less(pq.priorityQueue, j, j+1) {
			j++
		}
		if !less(pq.priorityQueue, k, j) {
			break
		}
		exch(pq.priorityQueue, k, j)
		k = j
	}
}

func (pq *MaxPriorityQueue) Insert(item int) {
	pq.priorityQueue[pq.length] = item
	pq.swim(pq.length)
	pq.length++
}

func (pq *MaxPriorityQueue) DelMax() int {
	max := pq.priorityQueue[0]
	exch(pq.priorityQueue, 0, pq.length-1)
	pq.length--
	pq.priorityQueue = pq.priorityQueue[:pq.length]
	pq.sink(0)
	return max
}
