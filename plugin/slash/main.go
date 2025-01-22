package slash

import (
	"fmt"
	"strings"
	"unicode"

	tele "gopkg.in/telebot.v3"
)

func isASCII(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func isValid(inputText string) bool {
	if len(inputText) < 2 {
		return false
	}
	if isASCII(inputText[:2]) && !strings.HasPrefix(inputText, "/$") {
		return false
	}
	return true
}

func genLink(c tele.Context) (string, string) {
	genName := func(firstName, lastName string) string {
		if lastName != "" {
			return fmt.Sprintf("%s %s", firstName, lastName)
		}
		return firstName
	}

	senderURI := fmt.Sprintf("tg://user?id=%d", c.Message().Sender.ID)
	senderName := genName(c.Message().Sender.FirstName, c.Message().Sender.LastName)

	// Message is sent on behalf of a Channel or Group
	if c.Message().SenderChat != nil {
		chatID := -1 * (c.Message().SenderChat.ID % 10000000000)
		senderURI = fmt.Sprintf("https://t.me/c/%d", chatID)
		senderName = c.Message().SenderChat.Title
	}

	// Message is NOT a reply to others by default
	replyToURI := ""
	replyToName := "自己"

	// Message is a reply to others
	if c.Message().IsReply() {
		replyToURI = fmt.Sprintf("tg://user?id=%d", c.Message().ReplyTo.Sender.ID)
		replyToName = genName(c.Message().ReplyTo.Sender.FirstName, c.Message().ReplyTo.Sender.LastName)

		if c.Message().ReplyTo.SenderChat != nil {
			chatID := -1 * (c.Message().ReplyTo.SenderChat.ID % 10000000000)
			replyToURI = fmt.Sprintf("https://t.me/c/%d", chatID)
			replyToName = c.Message().ReplyTo.SenderChat.Title
		}
	}

	if len(c.Message().Entities) != 0 {
		if c.Message().Entities[0].Type == "text_mention" {
			replyToURI = fmt.Sprintf("tg://user?id=%d", c.Message().Entities[0].User.ID)
			replyToName = genName(c.Message().Entities[0].User.FirstName, c.Message().Entities[0].User.LastName)
		} else if c.Message().Entities[0].Type == "mention" {
			t := strings.Index(c.Text(), " @")
			if t != -1 {
				pubUserName := c.Text()[t:]
				replyToName = strings.TrimSpace(pubUserName)
			}
			replyToURI = ""
		}
	}

	senderLink := fmt.Sprintf("[%s](%s)", senderName, senderURI)
	replyToLink := fmt.Sprintf("[%s](%s)", replyToName, replyToURI)

	return senderLink, replyToLink
}

func Execute(c tele.Context) error {
	inputText := c.Text()

	if !isValid(inputText) {
		return nil
	}

	actions := strings.SplitN(strings.Replace(inputText, "$", "", 1)[1:], " ", 3)

	if len(actions) != 1 && len(actions) != 2 && len(actions) != 3 {
		return nil
	}

	senderLink, replyToLink := genLink(c)

	outputText := fmt.Sprintf("%s %s了 %s", senderLink, actions[0], replyToLink)
	if len(actions) == 2 || len(actions) == 3 {
		outputText = fmt.Sprintf("%s %s了 %s %s", senderLink, actions[0], replyToLink, actions[1])
	}

	return c.Reply(outputText, &tele.SendOptions{
		ParseMode: "Markdown",
	})
}
