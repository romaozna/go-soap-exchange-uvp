package controller

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"main/src/model/json_data"
	"main/src/model/xml_data"
	"main/src/service"
	"net/http"
	"slices"
	"strconv"
	"time"
)

func GetStatus(w http.ResponseWriter, r *http.Request) {
	seconds := time.Now().Unix()
	data, _ := json.Marshal(seconds)
	log.Println("Status:", data)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}

func GetArchive(w http.ResponseWriter, r *http.Request) {
	device := chi.URLParam(r, "deviceId")
	pipeNumber := chi.URLParam(r, "pipeId")
	startTime := r.URL.Query().Get("startTime")
	endTime := r.URL.Query().Get("endTime")

	response := service.GetArchiveByDeviceAndPipe(http.MethodPost, device, pipeNumber, startTime, endTime)
	data := new(xml_data.UvpInfo)
	_ = xml.Unmarshal(response, data)
	dataJson := new(json_data.ArchiveResponse)

	dataJson.StartTime = data.GetArchivesResponse.ArchData.BeginTime
	dataJson.EndTime = data.GetArchivesResponse.ArchData.EndTime
	dataJson.PipeNumber = data.GetArchivesResponse.ArchData.PipeNumber

	s := []string{"Pabs", "Pi", "T", "mwT", "SrawV", "SvSTD", "RoSTD", "RoWRK"}
	jsonValues := make(json_data.Values, len(data.GetArchivesResponse.ArchItems.Value))

	for i, value := range data.GetArchivesResponse.ArchItems.Value {
		if slices.Contains(s, value.Index) {
			jsonValues[i].Comment = value.Comment
			jsonValues[i].Value = value.Text
			jsonValues[i].Measure = value.Measure
			jsonValues[i].Index = value.Index
		}
	}

	var refinedValues = json_data.Values{}
	for _, val := range jsonValues {
		if val.Index != "" {
			refinedValues = append(refinedValues, val)
		}
	}

	dataJson.Val = refinedValues

	log.Println("Response Body:", data)
	log.Println("Response BodyJ:", dataJson)
	marshal, _ := json.Marshal(dataJson)

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(marshal))
	if err != nil {
		log.Println(fmt.Sprintf("Ошибка при записи ответа в функции GetArchive"))
	}
}

func GetDevices(w http.ResponseWriter, r *http.Request) {
	marshal, _ := json.Marshal(service.GetConfig())
	w.Write([]byte(marshal))
}

func GetDeviceById(w http.ResponseWriter, r *http.Request) {
	deviceId, _ := strconv.Atoi(chi.URLParam(r, "deviceId"))
	marshal, _ := json.Marshal(service.GetDeviceById(deviceId))
	w.Write([]byte(marshal))
}
