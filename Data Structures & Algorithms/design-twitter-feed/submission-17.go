type Twitter struct {
    time int
	tweetMap map[int][][]int // one user can have multiple tweet
	followMap map[int]map[int]bool
}


func Constructor() Twitter {
    return Twitter{
		time: 0,
		tweetMap: make(map[int][][]int),
		followMap: make(map[int]map[int]bool),
	}
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	following := []int{userId}
    // Get all the user's following accounts, including himself.
	for follow := range this.followMap[userId]{
		if follow != userId{
			following = append(following, follow)
		}
	}
	// Fetch out all the tweets that should present in his timeline
	tweets := [][]int{}
	for _, follow := range following{
		for _, tweet := range this.tweetMap[follow]{
			tweets = append(tweets, tweet)
		}
	}
	// Sort the tweets by time in descending order
	sort.Slice(tweets, func(i, j int) bool{
		return tweets[i][0] > tweets[j][0]
	})
	// Return the first ten tweets
	res := []int{}
	for i := 0; i < 10 && i < len(tweets); i++{
		res = append(res, tweets[i][1])
	}
	return res
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
    this.tweetMap[userId] = append(this.tweetMap[userId], []int{this.time, tweetId})
	this.time++
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
