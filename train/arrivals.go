package train

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"
)

type ArrivalsRequest struct {
	MapID     string `xml:"mapid"`
	StopID    string `xml:"stpid"`
	Max       int    `xml:"max"`
	RouteCode string `xml:"rt"`
}

type APIArrivalsResponse struct {
	XMLName xml.Name `xml:"ctatt"`
	Text    string   `xml:",chardata"`
	Tmst    string   `xml:"tmst"`
	ErrCd   string   `xml:"errCd"`
	ErrNm   string   `xml:"errNm"`
	Eta     []struct {
		Text                 string `xml:",chardata"`
		StationId            string `xml:"staId"`
		StopId               string `xml:"stpId"`
		StationName          string `xml:"staNm"`
		StopDescription      string `xml:"stpDe"`
		RouteNumber          string `xml:"rn"`
		AbbrRouteName        string `xml:"rt"`
		DestinationStationId string `xml:"destSt"`
		DestinationName      string `xml:"destNm"`
		TrainDirectionNum    string `xml:"trDr"`
		PredicationTimestamp string `xml:"prdt"`
		ArrivalTimeRaw       string `xml:"arrT"`
		ArrivalTimeParsed    time.Time
		IsApproaching        string `xml:"isApp"`
		IsScheduled          string `xml:"isSch"`
		IsDelayed            string `xml:"isDly"`
		IsFault              string `xml:"isFlt"`
		Flags                string `xml:"flags"`
		Lat                  string `xml:"lat"`
		Lon                  string `xml:"lon"`
		Heading              string `xml:"heading"`
	} `xml:"eta"`
}

// GetArrivalsAtStation Get arrival predictions for a train
func (c *Client) GetArrivalsAtStation(arrivalRequest ArrivalsRequest) (*APIArrivalsResponse, error) {
	//Create request
	requestURL := "http://lapi.transitchicago.com/api/1.0/ttarrivals.aspx"
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	//Add headers for request
	c.AddAuthKey(req)
	vals := req.URL.Query()
	if arrivalRequest.MapID != "" {
		vals.Add("mapid", arrivalRequest.MapID)
	}
	if arrivalRequest.RouteCode != "" {
		vals.Add("rt", arrivalRequest.RouteCode)
	}
	if arrivalRequest.StopID != "" {
		vals.Add("stpid", arrivalRequest.StopID)
	}
	if arrivalRequest.Max != 0 {
		vals.Add("max", strconv.Itoa(arrivalRequest.Max))
	}
	req.URL.RawQuery = vals.Encode()

	//Read all
	res, err := http.DefaultClient.Do(req)
	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	//Unmarshal
	var responseStruct APIArrivalsResponse
	err = xml.Unmarshal(resBytes, &responseStruct)
	if err != nil {
		return nil, err
	}

	//Process fields
	for i, val := range responseStruct.Eta {
		responseStruct.Eta[i].ArrivalTimeParsed = ConvertCTATime(val.ArrivalTimeRaw)
	}

	//Check for an error
	if responseStruct.ErrCd != "0" {
		return nil, errors.New(responseStruct.ErrNm)
	} else {
		return &responseStruct, nil
	}
}
