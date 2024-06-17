package service

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"main/src/model/json_data"
	xmlData "main/src/model/xml_data"
	"net/http"
	"os"
	"time"
)

const DateFormat string = "20060102"

func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func FillRequest(r *xmlData.ArchiveRequest, pipe string, startTime string, endTime string) {
	r.GetArchives.Options.PipeNumber = pipe
	r.GetArchives.Options.ReturnMeasure = "1"
	r.GetArchives.Options.ReturnHeader = "0"
	r.GetArchives.Options.ReturnComment = "1"
	r.GetArchives.Options.UserMeasure = "0"
	r.GetArchives.Options.StartTime = startTime
	r.GetArchives.Options.StopTime = endTime
}

func GetArchiveByDeviceAndPipe(method string, device string, pipe string, startTime string, endTime string) []byte {
	client := httpClient()
	endpoint := "http://192.168.3.222:80/xmldata.asmx"
	var requestBody = new(xmlData.ArchiveRequest)
	FillRequest(requestBody, "1", "2024-05-20T10:00:00.000", "2024-05-20T12:00:00.000")
	data, err := xml.Marshal(requestBody)

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request to API endpoint. %+v", err)
	}

	// Close the connection to reuse it
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}

	return body
}

func getConfig() *[]json_data.Device {
	f, err := os.ReadFile("config.json")
	if err != nil {
		log.Println(err)
	}
	data := new([]json_data.Device)
	json.Unmarshal([]byte(f), &data)
	return data
}