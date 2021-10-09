package model

// Marvel Api Response Struct
type MarvelResponse struct {
	Code            int    `json:"code"`
	Status          string `json:"status"`
	Copyright       string `json:"copyright"`
	AttributionText string `json:"attributionText"`
	AttributionHTML string `json:"attributionHTML"`
	Etag            string `json:"etag"`
	Data            struct {
		Offset  int `json:"offset"`
		Limit   int `json:"limit"`
		Total   int `json:"total"`
		Count   int `json:"count"`
		Results []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Modified    string `json:"modified"`
			Thumbnail   struct {
				Path      string `json:"path"`
				Extension string `json:"extension"`
			} `json:"thumbnail"`
			ResourceURI string `json:"resourceURI"`
			Comics      struct {
				Available     int    `json:"available"`
				CollectionURI string `json:"collectionURI"`
				Items         []struct {
					ResourceURI string `json:"resourceURI"`
					Name        string `json:"name"`
				} `json:"items"`
				Returned int `json:"returned"`
			} `json:"comics"`
			Series struct {
				Available     int    `json:"available"`
				CollectionURI string `json:"collectionURI"`
				Items         []struct {
					ResourceURI string `json:"resourceURI"`
					Name        string `json:"name"`
				} `json:"items"`
				Returned int `json:"returned"`
			} `json:"series"`
			Stories struct {
				Available     int    `json:"available"`
				CollectionURI string `json:"collectionURI"`
				Items         []struct {
					ResourceURI string `json:"resourceURI"`
					Name        string `json:"name"`
					Type        string `json:"type"`
				} `json:"items"`
				Returned int `json:"returned"`
			} `json:"stories"`
			Events struct {
				Available     int    `json:"available"`
				CollectionURI string `json:"collectionURI"`
				Items         []struct {
					ResourceURI string `json:"resourceURI"`
					Name        string `json:"name"`
				} `json:"items"`
				Returned int `json:"returned"`
			} `json:"events"`
			Urls []struct {
				Type string `json:"type"`
				URL  string `json:"url"`
			} `json:"urls"`
		} `json:"results"`
	} `json:"data"`
}
