package misskey

import "time"

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
	CreatedNote struct {
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

	// TODO: nullだった項目はanyになっているので適切に修正する
	Note struct {
		Type string `json:"type"`
		Body struct {
			ID   string `json:"id"`
			Type string `json:"type"`
			Body struct {
				ID        string    `json:"id"`
				CreatedAt time.Time `json:"createdAt"`
				UserID    string    `json:"userId"`
				User      struct {
					ID                string `json:"id"`
					Name              string `json:"name"`
					Username          string `json:"username"`
					Host              any    `json:"host"`
					AvatarURL         string `json:"avatarUrl"`
					AvatarBlurhash    string `json:"avatarBlurhash"`
					AvatarDecorations []struct {
						ID  string `json:"id"`
						URL string `json:"url"`
					} `json:"avatarDecorations"`
					IsBot  bool `json:"isBot"`
					IsCat  bool `json:"isCat"`
					Emojis struct {
					} `json:"emojis"`
					OnlineStatus string `json:"onlineStatus"`
					BadgeRoles   []struct {
						Name         string `json:"name"`
						IconURL      string `json:"iconUrl"`
						DisplayOrder int    `json:"displayOrder"`
					} `json:"badgeRoles"`
				} `json:"user"`
				Text               string `json:"text"`
				Cw                 any    `json:"cw"`
				Visibility         string `json:"visibility"`
				LocalOnly          bool   `json:"localOnly"`
				ReactionAcceptance any    `json:"reactionAcceptance"`
				RenoteCount        int    `json:"renoteCount"`
				RepliesCount       int    `json:"repliesCount"`
				Reactions          struct {
				} `json:"reactions"`
				ReactionEmojis struct {
				} `json:"reactionEmojis"`
				ReactionAndUserPairCache []any  `json:"reactionAndUserPairCache"`
				FileIds                  []any  `json:"fileIds"`
				Files                    []any  `json:"files"`
				ReplyID                  any    `json:"replyId"`
				RenoteID                 string `json:"renoteId"`
				ClippedCount             int    `json:"clippedCount"`
				Renote                   struct {
					ID        string    `json:"id"`
					CreatedAt time.Time `json:"createdAt"`
					UserID    string    `json:"userId"`
					User      struct {
						ID                string `json:"id"`
						Name              string `json:"name"`
						Username          string `json:"username"`
						Host              any    `json:"host"`
						AvatarURL         string `json:"avatarUrl"`
						AvatarBlurhash    string `json:"avatarBlurhash"`
						AvatarDecorations []any  `json:"avatarDecorations"`
						IsBot             bool   `json:"isBot"`
						IsCat             bool   `json:"isCat"`
						Emojis            struct {
						} `json:"emojis"`
						OnlineStatus string `json:"onlineStatus"`
						BadgeRoles   []struct {
							Name         string `json:"name"`
							IconURL      string `json:"iconUrl"`
							DisplayOrder int    `json:"displayOrder"`
						} `json:"badgeRoles"`
					} `json:"user"`
					Text               string `json:"text"`
					Cw                 any    `json:"cw"`
					Visibility         string `json:"visibility"`
					LocalOnly          bool   `json:"localOnly"`
					ReactionAcceptance any    `json:"reactionAcceptance"`
					RenoteCount        int    `json:"renoteCount"`
					RepliesCount       int    `json:"repliesCount"`
					Reactions          struct {
						Utukusii                 int `json:":utukusii@.:"`
						Subarasii2               int `json:":subarasii2@.:"`
						Utsukushii               int `json:":utsukushii@.:"`
						KotobaNiDekinaiUtukusisa int `json:":kotoba_ni_dekinai_utukusisa@.:"`
					} `json:"reactions"`
					ReactionEmojis struct {
					} `json:"reactionEmojis"`
					ReactionAndUserPairCache []string `json:"reactionAndUserPairCache"`
					Tags                     []string `json:"tags"`
					FileIds                  []string `json:"fileIds"`
					Files                    []struct {
						ID          string    `json:"id"`
						CreatedAt   time.Time `json:"createdAt"`
						Name        string    `json:"name"`
						Type        string    `json:"type"`
						Md5         string    `json:"md5"`
						Size        int       `json:"size"`
						IsSensitive bool      `json:"isSensitive"`
						Blurhash    string    `json:"blurhash"`
						Properties  struct {
							Width  int `json:"width"`
							Height int `json:"height"`
						} `json:"properties"`
						URL          string `json:"url"`
						ThumbnailURL string `json:"thumbnailUrl"`
						Comment      any    `json:"comment"`
						FolderID     any    `json:"folderId"`
						Folder       any    `json:"folder"`
						UserID       any    `json:"userId"`
						User         any    `json:"user"`
					} `json:"files"`
					ReplyID      any `json:"replyId"`
					RenoteID     any `json:"renoteId"`
					ClippedCount int `json:"clippedCount"`
				} `json:"renote"`
			} `json:"body"`
		} `json:"body"`
	}

	Visibility string
)
