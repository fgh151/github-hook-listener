package main

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"gopkg.in/go-playground/webhooks.v5/github"
	"net/http"
	"os"
	"os/exec"
)

type Configuration struct {
	Port         string
	Path         string
	GithubSecret string
	ExecCommand  string
}


func deleteHook(payload github.DeletePayload, command string) {
	//@see https://developer.github.com/v3/activity/events/types/#deleteevent
	if payload.RefType == "branch" {
		exec.Command(command, payload.Ref)
	}
}

func errorHandler(err error)  {
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}

func main() {
	configuration := Configuration{}

	errorHandler(gonfig.GetConf(os.Args[1], &configuration))


	hook, _ := github.New(github.Options.Secret(configuration.GithubSecret))
	http.HandleFunc(configuration.Path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent)
		if err != nil {
			os.Stderr.WriteString(err.Error())
		}
		switch payload.(type) {

		case github.DeletePayload:
			deletePayload := payload.(github.DeletePayload)
			deleteHook(deletePayload, configuration.ExecCommand);
			fmt.Printf("%+v", deletePayload)
		}
	})

	errorHandler(http.ListenAndServe(configuration.Port, nil))
}
