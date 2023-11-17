package misskey

type (
	AccessToken string

	// api/meta
	Meta struct {
		AccessToken AccessToken `json:"i"`
		Detail      bool        `json:"detail"`
	}

	MetaResponse struct {
		MaintainerName string `json:"maintainerName"`
		Version        string `json:"version"`
		Name           string `json:"name"`
		BannerUrl      string `json:"bannerUrl"`
		IconUrl        string `json:"iconUrl"`
	}

	CreateNote struct {
		AccessToken AccessToken `json:"i"`
		Visibility  Visibility  `json:"visibility"`
		Text        string      `json:"text"`
	}

	CreateNoteResponse struct {
		CreatedNote Note `json:"createdNote"`
	}

	// misskey defined types
	Note struct {
		Id         string     `json:"id"`
		CreatedAt  string     `json:"createdAt"`
		UserId     string     `json:"userId"`
		Cw         *string    `json:"cw,omitempty"`
		User       User       `json:"user"`
		Visibility Visibility `json:"visibility"`
	}

	User struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		UserName  string `json:"username"`
		AvatarUrl string `json:"avatarUrl"`
		IsBot     bool   `json:"isBot"`
		IsCat     bool   `json:"isCat"`
	}

	Visibility string
)
