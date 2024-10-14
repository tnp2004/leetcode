package numberOfRecentCalls

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: []int{},
	}
}

func (r *RecentCounter) Ping(t int) int {
	r.queue = append(r.queue, t)
	if len(r.queue) > 1 && t > 3000 {
		for i := len(r.queue) - 2; i >= 0; i-- {
			if t-r.queue[i] > 3000 {
				r.queue = append(r.queue[i+1:])
				break
			}
		}
	}

	return len(r.queue)
}
