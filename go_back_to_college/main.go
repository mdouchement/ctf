package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	irc "github.com/fluffle/goirc/client"
	"github.com/mdouchement/ctf"
)

// http://www.root-me.org/en/Challenges/Programming/Go-back-to-college-147
func main() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	c := irc.SimpleClient("GeorgeAbitbol")
	c.EnableStateTracking()

	quit := make(chan bool)
	c.HandleFunc(irc.DISCONNECTED, func(conn *irc.Conn, line *irc.Line) { quit <- true })

	c.HandleFunc(irc.CONNECTED, func(conn *irc.Conn, line *irc.Line) { conn.Join("#root-me_challenge") })
	c.HandleFunc(irc.JOIN, func(conn *irc.Conn, line *irc.Line) {
		time.Sleep(1 * time.Second)
		log.Info("Candy -> !ep1")
		c.Privmsg("Candy", "!ep1")
	})

	c.HandleFunc(irc.PRIVMSG, func(conn *irc.Conn, line *irc.Line) {
		if line.Nick == "Candy" {
			if strings.Contains(line.Text(), " / ") {
				log.Info(line.Text())
				number := processResult(line.Text())
				log.Info(fmt.Sprintf("Candy -> !ep1 -rep %.1f", number))
				c.Privmsg("Candy", fmt.Sprintf("!ep1 -rep %.1f", number))
			} else {
				log.Info(line.Text())
				quit <- true
			}
		}
	})

	err := c.ConnectTo("irc.root-me.org:6667")
	check(err)

	log.Info("Wait for disconnect")
	<-quit
}

func processResult(response string) float64 {
	numbers := []float64{}
	for _, number := range strings.Split(response, " / ") {
		i, err := strconv.Atoi(number)
		check(err)
		numbers = append(numbers, float64(i))
	}
	sqrt := math.Sqrt(numbers[0])
	return extendedmath.RoundPlus(sqrt*numbers[1], 1)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
