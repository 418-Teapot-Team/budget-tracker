package client

import "time"

type GoverlaOutput struct {
	Data struct {
		Point struct {
			Id    string `json:"id"`
			Rates []struct {
				Id       string `json:"id"`
				Currency struct {
					Alias       string `json:"alias"`
					Name        string `json:"name"`
					Exponent    int    `json:"exponent"`
					CodeAlpha   string `json:"codeAlpha"`
					CodeNumeric int    `json:"codeNumeric"`
					Typename    string `json:"__typename"`
				} `json:"currency"`
				Bid struct {
					Absolute  int       `json:"absolute"`
					Relative  int       `json:"relative"`
					UpdatedAt time.Time `json:"updatedAt"`
					Typename  string    `json:"__typename"`
				} `json:"bid"`
				Ask struct {
					Absolute  int       `json:"absolute"`
					Relative  int       `json:"relative"`
					UpdatedAt time.Time `json:"updatedAt"`
					Typename  string    `json:"__typename"`
				} `json:"ask"`
				Typename string `json:"__typename"`
			} `json:"rates"`
			UpdatedAt time.Time `json:"updatedAt"`
			Typename  string    `json:"__typename"`
		} `json:"point"`
	} `json:"data"`
}

type NbuOfficial struct {
	Txt  string  `json:"txt"`
	Rate float64 `json:"rate"`
	Cc   string  `json:"cc"`
}
