package alquilatucancha

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"

	"github.com/rs/zerolog/log"
)

const pageURL = "https://alquilatucancha.com/_next/data/mv2RciMzyj9T0RzhrbCRU/sportclub/523.json"
const baseURL = "https://alquilatucancha.com"

type GetSportClubParams struct {
	Day     string
	BuildID string
}

func ExtractBuildID() (string, error) {
	res, err := http.DefaultClient.Get("https://alquilatucancha.com")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	r := regexp.MustCompile(`"buildId":"([^"]+)"`)

	match := r.FindStringSubmatch(string(body))
	if len(match) > 1 {
		return match[1], nil
	}

	return "", errors.New("build id not found")
}

func GetSportClubAvailability(params GetSportClubParams) (SportClubResponse, error) {
	epURL := baseURL + "/_next/data/" + params.BuildID + "/sportclub/523.json"

	baseURL, err := url.Parse(epURL)
	if err != nil {
		return SportClubResponse{}, fmt.Errorf("failed parsing pageURL %w", err)
	}

	// Define your query parameters
	queryParams := url.Values{}
	queryParams.Add("dia", params.Day)

	// Add the params to the base url
	baseURL.RawQuery = queryParams.Encode()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, baseURL.String(), nil)
	if err != nil {
		return SportClubResponse{}, fmt.Errorf("error building alquilatucancha request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return SportClubResponse{}, fmt.Errorf("error calling alquilatucancha: %w", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK: // success
	default:
		return SportClubResponse{}, fmt.Errorf("unknown status code %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SportClubResponse{}, fmt.Errorf("error reading alquilatucancha body: %w", err)
	}

	var result SportClubResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Error().Msgf("%v", string(body))
		return SportClubResponse{}, fmt.Errorf("error unmarshalling alquilatucancha response: %w", err)
	}

	return result, nil
}
