package hummingbird

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type AuthenicationToken string

func (t AuthenicationToken) Token() string {
	return string(t)
}

type User struct {
	Name                    string     `json:"name"`
	Waifu                   string     `json:"waifu"`
	WaifuOrHusbando         string     `json:"waifu_or_husbando"`
	WaiguSlug               string     `json:"waigu_slug"`
	WaifuCharId             string     `json:"waifu_char_id"`
	Location                string     `json:"location"`
	URL                     string     `json:"url"`
	Website                 string     `json:"website"`
	Avatar                  string     `json:"avatar"`
	AvatarSmall             string     `json:"avatar_small"`
	CoverImage              string     `json:"cover_image"`
	About                   string     `json:"about"`
	Bio                     string     `json:"bio"`
	Karma                   float64    `json:"karma"`
	LifeSpentOnAnime        int        `json:"life_spent_on_anime"`
	ShowAdultContent        bool       `json:"show_adult_content"`
	TitleLanguagePreference string     `json:"title_language_preference"`
	LastLibraryUpdate       time.Time  `json:"last_library_update"`
	Online                  bool       `json:"online"`
	Following               bool       `json:"following"`
	Favorites               []Favorite `json:"favorites"`
}

type Favorite struct {
	Id        int       `json:"id"`
	UserId    int       `json:"user_id"`
	ItemId    int       `json:"item_id"`
	ItemType  string    `json:"item_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	FavRank   int       `json:"fav_rank"`
}

type Story struct {
	Id              int        `json:"id"`
	StoryType       string     `json:"story_type"`
	User            User       `json:"user"`
	UpdatedAt       time.Time  `json:"updated_at"`
	SelfPost        bool       `json:"self_post"`
	Poster          User       `json:"poster"`
	Media           Anime      `json:"anime"`
	SubStoriesCount int        `json:"substories_count"`
	SubStories      []SubStory `json:"substories"`
}

type SubStory struct {
	Id            int         `json:"id"`
	SubStoryType  string      `json:"substory_type"`
	CreatedAt     time.Time   `json:"created_at"`
	Comment       string      `json:"comment"`
	EpisodeNumber int         `json:"episode_number,string"`
	FollowedUser  User        `json:"followed_user"`
	NewStatus     string      `json:"new_status"`
	Service       string      `json:"service"`
	Permissions   interface{} `json:"permissions"`
}

func AuthenicateUser(email string, username string, password string) (AuthenicationToken, error) {
	var e apiError
	var token AuthenicationToken
	if email == "" && username == "" {
		return "", fmt.Errorf("Empty email and username.")
	}
	if password == "" {
		return "", fmt.Errorf("Empty password.")
	}
	v := url.Values{}
	if email != "" {
		v.Set("email", email)
	}
	if username != "" {
		v.Set("username", username)
	}
	v.Set("password", password)

	if req, err := http.NewRequest("POST", "http://hummingbird.me/api/v1/users/authenticate?"+v.Encode(), nil); err == nil {
		req.Header.Add("User-Agent", userAgent)
		if resp, err := defaultClient.Do(req); err == nil {
			decoder := json.NewDecoder(resp.Body)
			if resp.StatusCode == 200 || resp.StatusCode == 201 {
				if err := decoder.Decode(&token); err == nil {
					return token, nil
				} else {
					return token, err
				}
			} else {
				if err := decoder.Decode(&e); err == nil {
					return token, e
				} else {
					return token, err
				}
			}
		} else {
			return token, err
		}
	} else {
		return token, err
	}
}

func GetUser(username string) (User, error) {
	var user User
	var e apiError

	if req, err := http.NewRequest("GET", "http://hummingbird.me/api/v1/users/"+username, nil); err == nil {
		req.Header.Add("User-Agent", userAgent)
		if resp, err := defaultClient.Do(req); err == nil {
			decoder := json.NewDecoder(resp.Body)
			if resp.StatusCode == 200 {
				if err := decoder.Decode(&user); err == nil {
					return user, nil
				} else {
					return user, err
				}
			} else {
				if err := decoder.Decode(&e); err == nil {
					return user, e
				} else {
					return user, fmt.Errorf("Failed to retrieve user.")
				}
			}
		} else {
			return user, err
		}
	} else {
		return user, err
	}
}

func GetUserFeed(username string) ([]Story, error) {
	var stories []Story
	var e apiError

	if req, err := http.NewRequest("GET", "http://hummingbird.me/api/v1/users/"+username+"/feed", nil); err == nil {
		req.Header.Add("User-Agent", userAgent)
		if resp, err := defaultClient.Do(req); err == nil {
			decoder := json.NewDecoder(resp.Body)
			if resp.StatusCode == 200 {
				if err := decoder.Decode(&stories); err == nil {
					return stories, nil
				} else {
					return stories, err
				}
			} else {
				if err := decoder.Decode(&e); err == nil {
					return stories, e
				} else {
					return stories, fmt.Errorf("Failed to retrieve stories.")
				}
			}
		} else {
			return stories, err
		}
	} else {
		return stories, err
	}
}

func GetUserFavoriteAnime(username string) ([]Anime, error) {
	var favorites []Anime
	var e apiError

	if req, err := http.NewRequest("GET", "http://hummingbird.me/api/v1/users/"+username+"/favorite_anime", nil); err == nil {
		req.Header.Add("User-Agent", userAgent)
		if resp, err := defaultClient.Do(req); err == nil {
			decoder := json.NewDecoder(resp.Body)
			if resp.StatusCode == 200 {
				if err := decoder.Decode(&favorites); err == nil {
					return favorites, nil
				} else {
					return favorites, err
				}
			} else {
				if err := decoder.Decode(&e); err == nil {
					return favorites, e
				} else {
					return favorites, fmt.Errorf("Failed to retrieve favorites.")
				}
			}
		} else {
			return favorites, err
		}
	} else {
		return favorites, err
	}
}
