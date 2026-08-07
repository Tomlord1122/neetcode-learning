type Twitter struct {
    time int
	postMap map[int][][]int
	followMap map[int]map[int]bool
}


func Constructor() Twitter {
    return Twitter{
		time: 0, 
		postMap: make(map[int][][]int),
		followMap: make(map[int]map[int]bool),
	}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
    this.postMap[userId] = append(this.postMap[userId], []int{this.time, tweetId})
	this.time++
}


func (this *Twitter) GetNewsFeed(userId int) []int {
	ids := make(map[int]bool)
	ids[userId] = true
	for followeeId := range this.followMap[userId] {
		ids[followeeId] = true
	}

	tweets := [][]int{}
	for id := range ids {
		for _, post := range this.postMap[id] {
			tweets = append(tweets, post)
		}
	}

	sort.Slice(tweets, func(i, j int) bool{
		return tweets[i][0] > tweets[j][0]
	})

	res := make([]int, 0)
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