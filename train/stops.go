package train

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type CityOfChicagoStop struct {
	StopID                 string `json:"stop_id"`
	DirectionID            string `json:"direction_id"`
	StopName               string `json:"stop_name"`
	StationName            string `json:"station_name"`
	StationDescriptiveName string `json:"station_descriptive_name"`
	MapID                  string `json:"map_id"`
	Ada                    bool   `json:"ada"`
	Red                    bool   `json:"red"`
	Blue                   bool   `json:"blue"`
	Green                  bool   `json:"g"`
	Brown                  bool   `json:"brn"`
	Purple                 bool   `json:"p"`
	Pexp                   bool   `json:"pexp"`
	Yellow                 bool   `json:"y"`
	Pink                   bool   `json:"pnk"`
	Orange                 bool   `json:"o"`
	Location               struct {
		Latitude     string `json:"latitude"`
		Longitude    string `json:"longitude"`
		HumanAddress string `json:"human_address"`
	} `json:"location"`
	ComputedRegionAwafS7Ux string `json:":@computed_region_awaf_s7ux"`
	ComputedRegion6MkvF3Dw string `json:":@computed_region_6mkv_f3dw"`
	ComputedRegionVrxfVc4K string `json:":@computed_region_vrxf_vc4k"`
	ComputedRegionBdys3D7I string `json:":@computed_region_bdys_3d7i"`
	ComputedRegion43Wa7Qmu string `json:":@computed_region_43wa_7qmu"`
}

// GetStopsFromChicagoData Get stops from the City of Chicago online dataset
func GetStopsFromChicagoData() ([]CityOfChicagoStop, error) {
	requestURL := "https://data.cityofchicago.org/resource/8pix-ypme.json"

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var stops []CityOfChicagoStop
	err = json.Unmarshal(resBody, &stops)
	if err != nil {
		return nil, err
	}
	return stops, nil
}

func GetStopsByName(stopSet []CityOfChicagoStop, name string) []CityOfChicagoStop {
	var results []CityOfChicagoStop
	for _, stop := range stopSet {
		if strings.Contains(stop.StationDescriptiveName, name) {
			results = append(results, stop)
		}
	}
	return results
}

func GetStopById(stopSet []CityOfChicagoStop, id string) *CityOfChicagoStop {
	for _, stop := range stopSet {
		if stop.MapID == id {
			return &stop
		}
	}
	return nil
}
