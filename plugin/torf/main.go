package torf

import (
	"math/rand"
	"strings"

	tele "gopkg.in/telebot.v3"
)

func randResponse(saidType int, r *rand.Rand) string {
	responsesMap := map[int][]string{
		1: {"有", "没有"},
		2: {"好", "不好"},
		3: {"是", "不是"},
		4: {"尊嘟", "假嘟"},
	}

	if responses, exists := responsesMap[saidType]; exists {
		return responses[r.Intn(len(responses))]
	}

	return ""
}

func Execute(c tele.Context, r *rand.Rand) error {
	inputText := c.Text()
	var outputText string

	if strings.Contains(inputText, "有没有") {
		outputText = randResponse(1, r)
	} else if strings.Contains(inputText, "好不好") {
		outputText = randResponse(2, r)
	} else if strings.Contains(inputText, "是不是") {
		outputText = randResponse(3, r)
	} else if strings.Contains(inputText, "尊嘟假嘟") {
		outputText = randResponse(4, r)
	} else {
		return nil
	}

	return c.Reply(outputText)
}
