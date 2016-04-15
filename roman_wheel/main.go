package main

import (
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	irc "github.com/fluffle/goirc/client"
)

func main() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	timer := time.Now()

	c := irc.SimpleClient("GeorgeAbitbol")
	c.DisableStateTracking()
	c.Config().Timeout = 100 * time.Millisecond

	quit := make(chan bool)
	c.HandleFunc(irc.DISCONNECTED, func(conn *irc.Conn, line *irc.Line) { quit <- true })

	c.HandleFunc(irc.CONNECTED, func(conn *irc.Conn, line *irc.Line) { conn.Join("#root-me_challenge") })
	c.HandleFunc(irc.JOIN, func(conn *irc.Conn, line *irc.Line) {
		log.Info("Candy -> !ep3")
		timer = time.Now()
		c.Privmsg("Candy", "!ep3")
	})

	firstMessage := true
	c.HandleFunc(irc.PRIVMSG, func(conn *irc.Conn, line *irc.Line) {
		if line.Nick == "Candy" {
			if firstMessage {
				str := processResult(line.Text())
				c.Privmsg("Candy", fmt.Sprintf("!ep3 -rep %s", str))
				log.Info("Elapsed time to get response: ", time.Since(timer))
				log.Info(line.Text())
				log.Info(fmt.Sprintf("Candy -> !ep3 -rep %s", str))
				firstMessage = false
			} else {
				log.Info(line.Text())
				log.Info("Total elapsed time: ", time.Since(timer))
				quit <- true
			}
		}
	})

	err := c.ConnectTo("irc.root-me.org:6667")
	check(err)

	log.Info("Wait for disconnect")
	<-quit
}

func processResult(response string) string {
	raw := make([]byte, len(response), len(response))
	copy(raw, response)
	for i, b := range raw {
		raw[i] = rot13(b)
	}
	return string(raw)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func rot13(x byte) byte {
	capital := x >= 'A' && x <= 'Z'
	if !capital && (x < 'a' || x > 'z') {
		return x // Not a letter
	}

	x += 13
	if capital && x > 'Z' || !capital && x > 'z' {
		x -= 26
	}
	return x
}
