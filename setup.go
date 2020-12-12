package linego

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

}

func (cl *LineClient) loadBotData() *BotSavedData {
	bytes, err := ioutil.ReadFile(cl.Profile.Mid + ".json")
	if err != nil {
		log.Fatal(err)
	}
	var data *BotSavedData
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatal(err)
	}
	return data
}

func (cl *LineClient) dumpBotData(data *BotSavedData) {
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile(cl.Profile.Mid+".json", file, 0644)
}
