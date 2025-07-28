//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"time"
	"sync"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
	mu 		  sync.Mutex
}

const FreeTimeLimitPerUser = 10;

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	start := time.Now();
	process();
	duration := time.Since(start);
	durationInSeconds := int64(duration.Seconds() + 0.5);

	if u.IsPremium {
		return true;
	}

	u.mu.Lock();
	defer u.mu.Unlock();

	if durationInSeconds > FreeTimeLimitPerUser {
		return false;
	}

	if durationInSeconds + u.TimeUsed > FreeTimeLimitPerUser {
		return false;
	}

	u.TimeUsed += durationInSeconds;
	return true;
}

func main() {
	RunMockServer()
}
