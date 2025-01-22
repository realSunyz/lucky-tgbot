package reborn

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	tele "gopkg.in/telebot.v3"
)

type Country struct {
	Name string `json:"country"`
	Pop  int64  `json:"population"`
}

type RebornData struct {
	List     []Country
	TotalPop int64
}

func InitRebornList(filePath string) (*RebornData, error) {
	rawJson, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var countries []Country
	if err = json.Unmarshal(rawJson, &countries); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	totalPop := int64(0)
	for _, country := range countries {
		totalPop += country.Pop
	}

	return &RebornData{
		List:     countries,
		TotalPop: totalPop,
	}, nil
}

func (data *RebornData) randCountry(r *rand.Rand) (Country, error) {
	randNum := r.Int63n(data.TotalPop)

	for _, country := range data.List {
		if randNum < country.Pop {
			return country, nil
		}
		randNum -= country.Pop
	}

	return Country{}, fmt.Errorf("failed to select a country")
}

func randGender(r *rand.Rand) string {
	genders := []string{"男孩子", "女孩子", " MtF", " FtM", "萝莉", "正太", "武装直升机", "沃尔玛购物袋", "狗狗", "猫猫"}
	return genders[r.Intn(len(genders))]
}

func Execute(c tele.Context, data *RebornData) error {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	country, err := data.randCountry(r)
	if err != nil {
		return nil
	}

	gender := randGender(r)

	outputText := fmt.Sprintf("投胎成功！\n你出生在%s，是%s。", country.Name, gender)

	return c.Reply(outputText)
}
