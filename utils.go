package api_utils

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

var TimeNowUTC = func() time.Time {
	return time.Now().UTC()
}

var NewUUIDString = func() string {
	return uuid.New().String()
}

func WriteDocsJSON(title, version, host, basePath, destFileName string) {

	dataByte, err := os.ReadFile("./docs/swagger.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	m := make(map[string]interface{})
	json.Unmarshal(dataByte, &m)

	sInfo := m["info"]
	inf := sInfo.(map[string]interface{})
	inf["title"] = title
	inf["version"] = version
	m["host"] = host
	m["basePath"] = basePath

	of, err := os.Create(fmt.Sprintf("./docs/%v.json", destFileName))
	if err != nil {
		fmt.Println(err)
		return
	}
	mByte, err := json.Marshal(&m)
	of.Write(mByte)
}
