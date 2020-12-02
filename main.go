package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/nomad/api"
)

func main() {
	nextGroups := os.Args[1:]
	jobId := os.Getenv("NOMAD_JOB_ID")
	myGroupName := os.Getenv("NOMAD_GROUP_NAME")
	fmt.Printf("next jobs are %#v\n", nextGroups)
	client, err := api.NewClient(&api.Config{
		Namespace: os.Getenv("NOMAD_NAMESPACE"),
	})
	handleError(err)

	jobs := client.Jobs()
	fmt.Printf("getting job: %q\n", jobId)
	job, _, err := jobs.Info(jobId, nil)
	handleError(err)

	for _, g := range nextGroups {
		group := job.LookupTaskGroup(g)
		if group == nil {
			fmt.Fprintf(os.Stderr, "[WARN] could not find next group: %q\n", g)
			continue
		}
		group.Count = i2p(1)
	}
	if g := job.LookupTaskGroup(myGroupName); g != nil {
		g.Count = i2p(0)
	} else {
		fmt.Fprintf(os.Stderr, "[WARN] could not find my group: %q\n", myGroupName)
	}

	_, _, err = jobs.RegisterOpts(job, &api.RegisterOptions{
		EnforceIndex: true,
		ModifyIndex:  *job.JobModifyIndex,
	}, nil)
	handleError(err)
}

func i2p(i int) *int {
	return &i
}

func handleError(err error) {
	if err == nil {
		return
	}
	fmt.Println("error creating client:", err.Error())
	os.Exit(-1)
}