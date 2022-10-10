package models

type Token struct {
	AccessToken string `json:"accessToken"`
}

type Analysis struct {
	Category string `json:"category"`
	Tag      string `json:"tag"`
}

type FilesName struct {
	ThubmnailId string `json:"thubmnailId"`
	VideoId     string `json:"videoId"`
	Size        int    `json:"size"`
}

type File struct {
	Title    string   `json:"title"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
	Files    *FilesName
}

type Files []*File

type User struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
}

type PostAnalysis struct {
	AccessToken   string   `json:"accessToken"`
	Category      string   `json:"category"`
	Tags          []string `json:"tags"`
	Date          string   `json:"date"`
	Duration      int      `json:"duration"`
	TotalDuration int      `json:"total_duration"`
	UserId        int      `json:"userId"`
}

type Data struct {
	Category interface{} `json:"category"`
	Tag      interface{} `json:"tag"`
}

type ResponseMessage struct {
	Data   *Data
	Status int `json:"status"`
}
