package torf

import (
	"math/rand"
	"strings"

	tele "gopkg.in/telebot.v3"
)

func randResponse(saidType int) string {
	responsesMap := map[int][]string{
		1: {"有", "没有"},
		2: {"好", "不好"},
		3: {"是", "不是"},
		4: {"尊嘟", "假嘟"},
	}

	if responses, exists := responsesMap[saidType]; exists {
		return responses[rand.Intn(len(responses))]
	}

	return ""
}

func Execute(c tele.Context) error {
	inputText := c.Text()
	var outputText string

	if strings.Contains(inputText, "有没有") {
		outputText = randResponse(1)
	} else if strings.Contains(inputText, "好不好") {
		outputText = randResponse(2)
	} else if strings.Contains(inputText, "是不是") {
		outputText = randResponse(3)
	} else if strings.Contains(inputText, "尊嘟假嘟") {
		outputText = randResponse(4)
	} else {
		return nil
	}

	return c.Reply(outputText)
}
