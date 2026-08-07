type Twitter struct {
	time int
	followMap map[int]map[int]bool
	tweetMap map[int][][]int // userId -> time, tweetId
}


func Constructor() Twitter {
    return Twitter{
		time: 0,
		followMap: make(map[int]map[int]bool),
		tweetMap: make(map[int][][]int),
	}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
    if _, exist := this.tweetMap[userId]; !exist{
		this.tweetMap[userId] = make([][]int, 0)
	}
	this.tweetMap[userId] = append(this.tweetMap[userId], []int{this.time, tweetId})
	this.time++
}


func (this *Twitter) GetNewsFeed(userId int) []int {
    // get the followee
	followees := this.followMap[userId]
	f := []int{userId}
	for followee, _ := range followees{
		if followee != userId{
			f = append(f, followee)
		}
	}
	// get the tweets
	tweets := [][]int{}
	for _, v := range f{
		tweets = append(tweets, this.tweetMap[v]...)
	}
	// sort the tweets
	sort.Slice(tweets, func(i, j int) bool{
		return tweets[i][0] > tweets[j][0]
	})
	// return first 10 tweets
	res := []int{}
	for i := 0; i < 10 && i < len(tweets); i++{
		res = append(res, tweets[i][1])
	}
	return res
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
    if _, exist := this.followMap[followerId]; !exist{
		this.followMap[followerId] = make(map[int]bool)
	}
	this.followMap[followerId][followeeId] = true
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    if _, exist := this.followMap[followerId]; exist{
		delete(this.followMap[followerId], followeeId)
	}
}
