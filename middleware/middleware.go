package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"sync"
	"time"
)

func logger(c *gin.Context) {
	// before request

	t := time.Now()
	c.Set("time", t)

	// handle request
	c.Next()

	// after request

	latency := time.Since(t)
	log.Printf("Latency: %s", latency)

	status := c.Writer.Status()
	log.Printf("Request StatusCode: %d", status)
}

func errorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		for _, err := range c.Errors {
			log.Print(err)
		}
		c.JSON(-1, c.Errors.JSON())
	}
	/*else if len(c.Errors.String()) <= 0 {
		c.
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}*/
}

func RunCronJobs() {
	var lock sync.Mutex
	timer1 := time.NewTicker(time.Second * 10)
	defer timer1.Stop()
	timer2 := time.NewTicker(time.Second * 5)
	defer timer2.Stop()
	for {
		/* run forever */
		select {
		case <-timer1.C:
			go func() {
				lock.Lock()
				defer lock.Unlock()
				println("CronJob :: every 10 seconds")
				/* do things I need done every 10 seconds */
			}()
		case <-timer2.C:
			go func() {
				lock.Lock()
				defer lock.Unlock()
				println("CronJob :: every 5 seconds")
				/* do things I need done every 5 seconds */
			}()
		}
	}
}
