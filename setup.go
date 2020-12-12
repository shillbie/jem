package linego

import (
	"encoding/json"
	"io/ioutil"
)

type BotSavedData struct {
	LastRevision  int64 `json:"lastRevision"`
	GlobalRev     int64 `json:"globalRevision"`
	IndividualRev int64 `json:"individualRevision"`
}

func NewBotSavedData() *BotSavedData {
	return &BotSavedData{
		LastRevision:  0,
		GlobalRev:     0,
		IndividualRev: 0,
	}
}

func (cl *LineClient) BeforeLogin() {
	cl.SaveData = cl.loadBotData()
}

func (cl *LineClient) loadBotData() *BotSavedData {
	bytes, err := ioutil.ReadFile(cl.Profile.Mid + ".json")
	if err != nil {
		return NewBotSavedData()
	}
	var data *BotSavedData
	if err := json.Unmarshal(bytes, &data); err != nil {
		return NewBotSavedData()
	}
	return data
}

func (cl *LineClient) dumpBotData(data *BotSavedData) {
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile(cl.Profile.Mid+".json", file, 0644)
}
