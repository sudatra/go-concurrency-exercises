//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, tweets chan<- *Tweet) {
    defer func() {
        close(tweets);
    }()

    for {
        tweet, err := stream.Next();
        if err == ErrEOF {
            return;
        }

        tweets <- tweet;
    }
}

func consumer(tweets <-chan *Tweet, done chan<- struct{}) {
    for t := range tweets {
        if t.IsTalkingAboutGo() {
            fmt.Println(t.Username, "\ttweets about golang")
        } else {
            fmt.Println(t.Username, "\tdoes not tweet about golang")
        }
    }

    done <- struct{}{};
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	// channels
	tweetsChannel := make(chan *Tweet, 100);
	doneChannel := make(chan struct{});

	// Producer
	go producer(stream, tweetsChannel);

	// Consumer
	go consumer(tweetsChannel, doneChannel);

	<-doneChannel;

	fmt.Printf("Process took %s\n", time.Since(start))
}
