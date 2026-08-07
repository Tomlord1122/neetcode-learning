type Twitter struct {
    time int
	tweetMap map[int][][]int // userId -> (count, tweetId)
	followMap map[int]map[int]bool // followerId -> followeeId -> bool
}


func Constructor() Twitter {
    return Twitter{
		time: 0,
		tweetMap: make(map[int][][]int),
		followMap: make(map[int]map[int]bool),
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
    // Get all the followee from followMap
	followees := []int{userId}
	f := this.followMap[userId]
	for followee, follow := range f{
		if follow && followee != userId{
			followees = append(followees, followee)
		}
	}
	// Get all the tweets from tweetMap
	tweets := [][]int{}
	for _, user := range followees{
		for _, tweet := range this.tweetMap[user]{
			tweets = append(tweets, []int{tweet[0], tweet[1]}) 
		}
	}
	// Sort by time in descending order
	sort.Slice(tweets, func(i, j int) bool{
		return tweets[i][0] > tweets[j][0]
	})
	// Return the first 10 tweetId
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
