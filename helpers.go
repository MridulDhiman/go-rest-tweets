package main

import "slices"

func GetTweetById (id string) (*Tweet, bool) {
actualId := slices.IndexFunc(tweets, func (t Tweet)  bool {
	return t.Id == id
})

if actualId == -1 {
	return nil, false
}

return &tweets[actualId], true
}

func RemoveTweet(slice []Tweet, s int) []Tweet {
    return append(slice[:s], slice[s+1:]...)
}