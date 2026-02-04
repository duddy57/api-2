package location

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type NominatimResponse struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func (n *NominatimResponse) UnmarshalJSON(data []byte) error {
	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Parse latitude
	if latVal, ok := raw["lat"]; ok {
		switch v := latVal.(type) {
		case float64:
			n.Lat = v
		case string:
			lat, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return fmt.Errorf("invalid latitude value: %v", v)
			}
			n.Lat = lat
		default:
			return fmt.Errorf("unexpected type for latitude: %T", v)
		}
	}

	// Parse longitude
	if lonVal, ok := raw["lon"]; ok {
		switch v := lonVal.(type) {
		case float64:
			n.Lon = v
		case string:
			lon, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return fmt.Errorf("invalid longitude value: %v", v)
			}
			n.Lon = lon
		default:
			return fmt.Errorf("unexpected type for longitude: %T", v)
		}
	}

	return nil
}

func GeocodeAddress(ctx context.Context, street, number, neighborhood, city, state, postalCode, country string) (float64, float64, error) {
	destination := fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s",
		street, number, neighborhood, city, state, postalCode, country)

	url := fmt.Sprintf("%s?q=%s&format=json&limit=1", os.Getenv("NOMINATIM_URL"), url.QueryEscape(destination))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "OlideskAPI/1.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to make geocoding request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, 0, fmt.Errorf("nominatim API returned status %d", resp.StatusCode)
	}

	var results []NominatimResponse
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return 0, 0, fmt.Errorf("failed to decode nominatim response: %w", err)
	}

	if len(results) == 0 {
		return 0, 0, fmt.Errorf("endereço não encontrado: %s", destination)
	}

	lat := results[0].Lat
	lng := results[0].Lon

	if lat < -90 || lat > 90 {
		return 0, 0, fmt.Errorf("invalid latitude: %f", lat)
	}
	if lng < -180 || lng > 180 {
		return 0, 0, fmt.Errorf("invalid longitude: %f", lng)
	}

	return lat, lng, nil
}
