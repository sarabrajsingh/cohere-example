package main

import (
	"fmt"
	"log"

	"github.com/cohere-ai/cohere-go"
)

func main() {
	log.Println("creating client")
	co, err := cohere.CreateClient("lg4hLLJXV9nPWw9lDxNuyCbSKFThiVfAGyIoacLq")
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Println("successfully created client")

	examples := []cohere.Example{
		{Text: "How do I find my insurance policy?", Label: "Finding policy details"},
		{Text: "How do I download a copy of my insurance policy?", Label: "Finding policy details"},
		{Text: "How do I find my policy effective date?", Label: "Finding policy details"},
		{Text: "When does my insurance policy end?", Label: "Finding policy details"},
		{Text: "Could you please tell me the date my policy becomes effective?", Label: "Finding policy details"},
		{Text: "How do I sign up for electronic filing?", Label: "Change account settings"},
		{Text: "How do I change my policy?", Label: "Change account settings"},
		{Text: "How do I sign up for direct deposit?", Label: "Change account settings"},
		{Text: "I want direct deposit. Can you help with that?", Label: "Change account settings"},
		{Text: "Could you deposit money into my account rather than mailing me a physical cheque?", Label: "Change account settings"},
		{Text: "How do I file an insurance claim?", Label: "Filing a claim and viewing status"},
		{Text: "How do I file a reimbursement claim?", Label: "Filing a claim and viewing status"},
		{Text: "How do I check my claim status?", Label: "Filing a claim and viewing status"},
		{Text: "When will my claim be reimbursed?", Label: "Filing a claim and viewing status"},
		{Text: "I filed my claim 2 weeks ago but I still haven't received a deposit for it.", Label: "Filing a claim and viewing status"},
		{Text: "I want to cancel my policy immediately! This is nonsense.", Label: "Cancelling coverage"},
		{Text: "Could you please help my end my insurance coverage? Thank you.", Label: "Cancelling coverage"},
		{Text: "Your service sucks. I'm switching providers. Cancel my coverage.", Label: "Cancelling coverage"},
		{Text: "Hello there! How do I cancel my coverage?", Label: "Cancelling coverage"},
		{Text: "How do I delete my account?", Label: "Cancelling coverage"},
	}
	inputs := []string{
		"Hi My Name is Sarabraj, can we be friends please?",
	}

	log.Println("starting classification")
	response, err := co.Classify(cohere.ClassifyOptions{
		Inputs:   inputs,
		Examples: examples,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("classification done")

	fmt.Println(response)
}
