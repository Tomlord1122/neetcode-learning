type Twitter struct {
	time int
	tweetMap map[int][][]int // Value struct is (time, tweetId)
	followMap map[int]map[int]bool
}


func Constructor() Twitter {
    return Twitter{
		time: 0,
		tweetMap: make(map[int][][]int),
		followMap: make(map[int]map[int]bool),
	}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
	this.tweetMap[userId] = append(this.tweetMap[userId], []int{this.time, tweetId})
	this.time++
}


func (this *Twitter) GetNewsFeed(userId int) []int {
    // We can split the goal into four part
	// 1. get the user id that userId follow (include himself)
	// 2. get all the tweets with ids from step 1
	// 3. sort in descending order
	// 4. return first ten tweets
	ids := make(map[int]bool)
	ids[userId] = true
	for id := range this.followMap[userId]{
		ids[id] = true
	}

	tweets := [][]int{}
	for id := range ids{
		for _, tweet := range this.tweetMap[id]{
			tweets = append(tweets, tweet)
		}
	}

	sort.Slice(tweets, func(i, j int) bool{
		return tweets[i][0] > tweets[j][0]
	})

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
