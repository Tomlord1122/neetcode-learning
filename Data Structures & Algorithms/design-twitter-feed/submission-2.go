type Twitter struct {
	count int
	tweetMap map[int][][]int 
	followMap map[int]map[int]bool
}

// heap item: one tweet + pointer to "next older tweet" of same user
type tweetItem struct{
	count int
	tweetId int
	followeeId int
	index int // index of next older tweet in tweetMap[followeeId]
}

type tweetMaxHeap []tweetItem

func (h tweetMaxHeap) Len() int{
	return len(h)
}

func (h tweetMaxHeap) Less(i, j int) bool{
	return h[i].count > h[j].count
}

func (h tweetMaxHeap) Swap(i, j int){
	h[i], h[j] = h[j], h[i]
}

func (h *tweetMaxHeap) Push(x any){
	*h = append(*h, x.(tweetItem))
}

func (h *tweetMaxHeap) Pop() any{
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}


func Constructor() Twitter {
    return Twitter{
		count: 0,
		tweetMap: make(map[int][][]int),
		followMap: make(map[int]map[int]bool),
	}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
    if this.tweetMap[userId] == nil{
		this.tweetMap[userId] = make([][]int, 0)
	}
	// use increasing count; newer = larger count
	this.tweetMap[userId] = append(this.tweetMap[userId], []int{this.count, tweetId})
	this.count++
}


func (this *Twitter) GetNewsFeed(userId int) []int {
    res := make([]int, 0, 10)

	// ensure maps initialized
	if this.followMap[userId] == nil{
		this.followMap[userId] = make(map[int]bool)
	}
	// user should always see their own tweets
	this.followMap[userId][userId] = true

	h := &tweetMaxHeap{}
	heap.Init(h)

	// Push most recent tweet of each followee into max-heap
	for followeeId := range this.followMap[userId]{
		tweets := this.tweetMap[followeeId]
		if len(tweets) == 0{
			continue
		}
		idx := len(tweets)-1
		count := tweets[idx][0]
		tID := tweets[idx][1]
		heap.Push(h, tweetItem{
			count: count,
			tweetId: tID,
			followeeId: followeeId,
			index: idx - 1, // next older tweet index
		})
	}

	// pop up to 10 most recent tweets, pushing older ones as we go
	for h.Len() > 0 && len(res) < 10{
		item := heap.Pop(h).(tweetItem)
		res = append(res, item.tweetId)

		if item.index >= 0{
			tweets := this.tweetMap[item.followeeId]
			count := tweets[item.index][0]
			tID := tweets[item.index][1]
			heap.Push(h, tweetItem{
				count: count,
				tweetId: tID,
				followeeId: item.followeeId,
				index: item.index - 1,
			})
		}
	}
	return res
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
    if this.followMap[followerId] == nil{
		this.followMap[followerId] = make(map[int]bool)
	}
	this.followMap[followerId][followeeId] = true
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	if this.followMap[followerId] != nil{
		delete(this.followMap[followerId], followeeId)
	}
}
