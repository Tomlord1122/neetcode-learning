type Twitter struct {
    time int
	tweetMap map[int][][]int // value is (tweet, time) pair
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
    this.tweetMap[userId] = append(this.tweetMap[userId], []int{tweetId, this.time})
	this.time++
}


func (this *Twitter) GetNewsFeed(userId int) []int {
    // 1. get the users that userId follow (including himself)
	// 2. get all the tweetId by these users
	// 3. sort these tweets by time in descedning order
	// 4. return first 10 tweets
	users := []int{userId}
	for follow := range this.followMap[userId]{
		if follow != userId{
			users = append(users, follow)
		}
	}

	tweets := [][]int{}
	for _, user := range users{
		for _, tweet := range this.tweetMap[user]{
			tweets = append(tweets, tweet)
		}
	}

	sort.Slice(tweets, func(i, j int) bool{
		return tweets[i][1] > tweets[j][1]
	})

	res := []int{}
	for i := 0; i < 10 && i < len(tweets);i++{
		res = append(res, tweets[i][0])
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
    if account, exist := this.followMap[followerId]; exist{
		delete(account, followeeId)
	}
}
