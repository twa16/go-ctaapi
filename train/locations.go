package train

import (
	"encoding/xml"
	"io"
	"net/http"
)

type APILocationsResponse struct {
	Resp struct {
		Timestamp   string `json:"tmst"`
		ErrorCode   string `json:"errCd"`
		ErrorNumber any    `json:"errNm"`
		Route       []struct {
			Name  string `json:"@name"`
			Train []struct {
				RunNum              string `json:"rn"`
				DestStation         string `json:"destSt"`
				DestNum             string `json:"destNm"`
				TrainDirection      string `json:"trDr"`
				NextStationID       string `json:"nextStaId"`
				NextStopID          string `json:"nextStpId"`
				NextStationNum      string `json:"nextStaNm"`
				PredictionTimestamp string `json:"prdt"`
				ArrivalTime         string `json:"arrT"`
				IsApproaching       string `json:"isApp"`
				IsDelayed           string `json:"isDly"`
				Flags               any    `json:"flags"`
				Lat                 string `json:"lat"`
				Lon                 string `json:"lon"`
				Heading             string `json:"heading"`
			} `json:"train"`
		} `json:"route"`
	} `json:"ctatt"`
}

func (c *Client) GetLocations(route string) (*APILocationsResponse, error) {
	//Create request
	requestURL := "https://lapi.transitchicago.com/api/1.0/ttpositions.aspx"
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	//Add headers for request
	c.AddAuthKey(req)
	vals := req.URL.Query()
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

	return nil, nil
}
