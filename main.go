package main

import (
	"fmt"
	"log"

	"github.com/gitrunner/gitcollector"
	"github.com/gitrunner/utils"
)

const (
	filePath = "config.json"
)

func main() {
	fmt.Print("git Runner started")
	if err := Runner(); err != nil {
		log.Panic(err.Error())
	}

}

func Runner() (err error) {
	//read data from config
	var config utils.Config
	config, err = utils.GetConfigValues(filePath)
	if err != nil {
		log.Println(err.Error())
		return
	}
	var githandler gitcollector.GitHandler
	// passs config to githandler
	githandler = gitcollector.GetNewGR(config)
	summary, err := githandler.FetchGitPRSummary()
	if err != nil {
		log.Println(err.Error())
		return
	}
	//send mail to server
	return githandler.MailToAdmin(summary)
}
