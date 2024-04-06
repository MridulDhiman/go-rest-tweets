package main

import (
	"github.com/gofiber/fiber/v2"
    "fmt"
	"github.com/google/uuid"
	"slices"
)



func HelloController(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}


func GetAllNotes (c *fiber.Ctx) error {
  return c.JSON(tweets)
}

func CreateNewNote  (c* fiber.Ctx) error {

	tweet := new(Tweet)
	id := uuid.New()
	tweet.Id = id.String()


	if  err:= c.BodyParser(tweet); err!= nil {
		return err;
	}


tweets = append(tweets, *tweet)
fmt.Printf("Tweet ID: %v\n", tweet.Id);
fmt.Printf("Tweet Content: %v\n", tweet.Content);
fmt.Printf("Tweet Author: %v\n", tweet.Author);


return nil
}



func GetSingleTweet (c* fiber.Ctx) error {
	id:= c.Params("id")
    
	tweet, flag := GetTweetById(id)

	if !flag {
		fmt.Println("ID does not matches")
		return fiber.ErrBadRequest
	}

	return c.JSON(*tweet)
}

func UpdateTweet (c* fiber.Ctx) error {
	id := c.Params("id")

	tweet, flag := GetTweetById(id)

	if !flag {
		fmt.Println("ID does not matches")
		return fiber.ErrBadRequest
	}

	reqBody := new(UpdateRequestBody)
	

	if  err:= c.BodyParser(reqBody); err!= nil {
		return err;
	}

	fmt.Println(reqBody.Content)
	tweet.Content = reqBody.Content

	return c.JSON(*tweet)
}

func DeleteTweet (c *fiber.Ctx) error {
	id:= c.Params("id")

	
	_, flag := GetTweetById(id)

	if !flag {
		fmt.Println("ID does not matches")
		return fiber.ErrBadRequest
	}

	actualId := slices.IndexFunc(tweets, func (t Tweet)  bool {
		return t.Id == id
	})

	remainingTweets:= RemoveTweet(tweets, actualId)

	tweets = remainingTweets

return c.JSON(remainingTweets)

}