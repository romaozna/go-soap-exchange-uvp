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
	"strconv"
	"time"
)

const DateFormat string = "20060102"

func GetArchiveByDeviceAndPipe(method string, device string, pipe string, startTime string, endTime string) []byte {
	client := httpClient()
	deviceId, _ := strconv.Atoi(device)
	endpoint := getEndpointFromDevice(deviceId)
	var requestBody = new(xmlData.ArchiveRequest)
	fillArchiveRequest(requestBody, pipe, startTime, endTime)
	data, err := xml.Marshal(requestBody)

	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending request to API endpoint. %+v\n", err)
	}

	// Close the connection to reuse it
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Couldn't parse response body. %+v", err)
	}

	return body
}

func GetConfig() *[]json_data.Device {
	f, err := os.ReadFile("config.json")
	if err != nil {
		log.Println(err)
	}
	data := new([]json_data.Device)
	json.Unmarshal([]byte(f), &data)
	return data
}

func GetDeviceById(deviceId int) json_data.Device {
	devices := GetConfig()
	for _, device := range *devices {
		if device.Device == deviceId {
			return device
		}
	}
	return json_data.Device{}
}

func getEndpointFromDevice(deviceId int) (endpoint string) {
	device := GetDeviceById(deviceId)
	endpoint = device.Endpoint
	return
}

func fillArchiveRequest(r *xmlData.ArchiveRequest, pipe string, startTime string, endTime string) {
	r.GetArchives.Options.PipeNumber = pipe
	r.GetArchives.Options.ReturnMeasure = "1"
	r.GetArchives.Options.ReturnHeader = "0"
	r.GetArchives.Options.ReturnComment = "1"
	r.GetArchives.Options.UserMeasure = "0"
	r.GetArchives.Options.StartTime = startTime
	r.GetArchives.Options.StopTime = endTime
}

func httpClient() *http.Client {
	client := &http.Client{Timeout: 20 * time.Second}
	return client
}
