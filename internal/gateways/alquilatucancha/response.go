package alquilatucancha

type SportClubResponse struct {
	PageProps struct {
		Sportclub struct {
			ID            int    `json:"id"`
			Permalink     string `json:"permalink"`
			Name          string `json:"name"`
			Logo          string `json:"logo"`
			LogoURL       string `json:"logo_url"`
			Background    string `json:"background"`
			BackgroundURL string `json:"background_url"`
			Location      struct {
				Name     string `json:"name"`
				City     string `json:"city"`
				Lat      string `json:"lat"`
				Lng      string `json:"lng"`
				Geohash9 string `json:"geohash_9"`
			} `json:"location"`
			Zone struct {
				ID       int    `json:"id"`
				Name     string `json:"name"`
				FullName string `json:"full_name"`
				Placeid  string `json:"placeid"`
				Country  struct {
					ID      int    `json:"id"`
					Name    string `json:"name"`
					IsoCode string `json:"iso_code"`
				} `json:"country"`
			} `json:"zone"`
			Props struct {
				Sponsor  bool   `json:"sponsor"`
				Favorite bool   `json:"favorite"`
				Stars    string `json:"stars"`
				Payment  bool   `json:"payment"`
			} `json:"props"`
			Attributes []string `json:"attributes"`
			Openhours  []struct {
				DayOfWeek int  `json:"day_of_week"`
				OpenTime  int  `json:"open_time"`
				CloseTime int  `json:"close_time"`
				Open      bool `json:"open"`
			} `json:"openhours"`
			Courts []struct {
				ID         int    `json:"id"`
				Name       string `json:"name"`
				Attributes struct {
					Floor  string `json:"floor"`
					Light  bool   `json:"light"`
					Roofed bool   `json:"roofed"`
				} `json:"attributes"`
				Sports []struct {
					ID                int    `json:"id"`
					ParentID          int    `json:"parent_id"`
					Name              string `json:"name"`
					PlayersMax        int    `json:"players_max"`
					Order             int    `json:"order"`
					DefaultDuration   int    `json:"default_duration"`
					DivisibleDuration int    `json:"divisible_duration"`
					Icon              string `json:"icon"`
					Pivot             struct {
						CourtID int `json:"court_id"`
						SportID int `json:"sport_id"`
						Enabled int `json:"enabled"`
					} `json:"pivot"`
				} `json:"sports"`
				Disclaimer string `json:"disclaimer"`
			} `json:"courts"`
			CourtSports []struct {
				ID                int         `json:"id"`
				Name              string      `json:"name"`
				Order             int         `json:"order"`
				Icon              string      `json:"icon"`
				IconResource      string      `json:"iconResource"`
				DefaultDuration   int         `json:"default_duration"`
				DivisibleDuration int         `json:"divisible_duration"`
				Durations         []int       `json:"durations"`
				Children          interface{} `json:"children"`
				ParentID          int         `json:"parent_id"`
				PlayersMax        int         `json:"players_max"`
				PlayersMin        int         `json:"players_min"`
			} `json:"court_sports"`
		} `json:"sportclub"`
		BookingsBySport struct {
			Num2 []struct {
				ID            int    `json:"id"`
				Permalink     string `json:"permalink"`
				Name          string `json:"name"`
				Logo          string `json:"logo"`
				LogoURL       string `json:"logo_url"`
				Background    string `json:"background"`
				BackgroundURL string `json:"background_url"`
				Location      struct {
					Name     string `json:"name"`
					City     string `json:"city"`
					Lat      string `json:"lat"`
					Lng      string `json:"lng"`
					Geohash9 string `json:"geohash_9"`
				} `json:"location"`
				Zone struct {
					ID       int    `json:"id"`
					Name     string `json:"name"`
					FullName string `json:"full_name"`
					Placeid  string `json:"placeid"`
					Country  struct {
						ID      int    `json:"id"`
						Name    string `json:"name"`
						IsoCode string `json:"iso_code"`
					} `json:"country"`
				} `json:"zone"`
				Props struct {
					Sponsor  bool   `json:"sponsor"`
					Favorite bool   `json:"favorite"`
					Stars    string `json:"stars"`
					Payment  bool   `json:"payment"`
				} `json:"props"`
				Attributes []string `json:"attributes"`
				Openhours  []struct {
					DayOfWeek int  `json:"day_of_week"`
					OpenTime  int  `json:"open_time"`
					CloseTime int  `json:"close_time"`
					Open      bool `json:"open"`
				} `json:"openhours"`
				Courts []struct {
					ID         int    `json:"id"`
					Name       string `json:"name"`
					Attributes struct {
						Floor  string `json:"floor"`
						Light  bool   `json:"light"`
						Roofed bool   `json:"roofed"`
						Beelup bool   `json:"beelup"`
					} `json:"attributes"`
					Sports []struct {
						ID                int    `json:"id"`
						ParentID          int    `json:"parent_id"`
						Name              string `json:"name"`
						PlayersMax        int    `json:"players_max"`
						Order             int    `json:"order"`
						DefaultDuration   int    `json:"default_duration"`
						DivisibleDuration int    `json:"divisible_duration"`
						Icon              string `json:"icon"`
						Pivot             struct {
							CourtID int `json:"court_id"`
							SportID int `json:"sport_id"`
							Enabled int `json:"enabled"`
						} `json:"pivot"`
					} `json:"sports"`
					Disclaimer string `json:"disclaimer"`
					Available  []struct {
						Price    int    `json:"price"`
						Duration int    `json:"duration"`
						Datetime string `json:"datetime"`
						Start    string `json:"start"`
						End      string `json:"end"`
						Priority int    `json:"_priority"`
					} `json:"available"`
				} `json:"courts"`
				Priority int `json:"_priority"`
			} `json:"2"`
		} `json:"bookingsBySport"`
		Day      string `json:"day"`
		Hour     string `json:"hour"`
		SportIds []int  `json:"sportIds"`
		Location struct {
			Value   string `json:"value"`
			Display string `json:"display"`
		} `json:"location"`
		SportOptions []struct {
			Children        []interface{} `json:"children"`
			Value           int           `json:"value"`
			Display         string        `json:"display"`
			SvgResource     string        `json:"svgResource"`
			Svg             string        `json:"Svg"`
			Order           int           `json:"order"`
			PlayersMax      int           `json:"playersMax"`
			PlayersMin      int           `json:"playersMin"`
			DefaultDuration int           `json:"defaultDuration"`
		} `json:"sportOptions"`
		SportclubID       int         `json:"sportclubId"`
		AvailableSportIds []int       `json:"availableSportIds"`
		PlaceID           string      `json:"placeId"`
		LocationName      string      `json:"locationName"`
		PlaceSearched     string      `json:"placeSearched"`
		Error             interface{} `json:"error"`
		QueryDate         interface{} `json:"queryDate"`
	} `json:"pageProps"`
	NSSP bool `json:"__N_SSP"`
}
