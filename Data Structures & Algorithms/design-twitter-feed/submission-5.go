type Twitter struct {
	count int
	tweetMap map[int][][]int // userId -> (count, tweetId)
	followMap map[int]map[int]bool // followerId -> followeeId, bool
}


func Constructor() Twitter {
    return Twitter{
		count: 0,
		tweetMap: make(map[int][][]int),
		followMap: make(map[int]map[int]bool),
	}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
    if _, ok := this.tweetMap[userId]; !ok{
		this.tweetMap[userId] = make([][]int, 0)
	}
	this.tweetMap[userId] = append(this.tweetMap[userId], []int{this.count, tweetId})
	// update count
	this.count++
}


func (this *Twitter) GetNewsFeed(userId int) []int {
    allFeeds := [][2]int{}
	followees := []int{userId}
	// 1. Use followMap to get all the followee
	totalFollowee := this.followMap[userId]
	for followee, follow := range totalFollowee{
		if follow && followee != userId {
			followees = append(followees, followee)
		}
	}
	// 2. Append tweetId fron each followee into allFeed
	for _, followee := range followees{
		tweets := this.tweetMap[followee]
		for _, tweet := range tweets{
		allFeeds = append(allFeeds, [2]int{tweet[0], tweet[1]})
		}
	}
	// 3. Sort by count in descending order
	sort.Slice(allFeeds, func(i, j int)bool{
		return allFeeds[i][0] > allFeeds[j][0]
	})
	// 4. Return first 10 tweetId
	res := []int{}

	for i := 0; i < 10 && i < len(allFeeds); i++{
		res = append(res, allFeeds[i][1])
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