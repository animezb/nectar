package hummingbird

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Anime struct {
	Id             int    `json:"id"`
	Slug           string `json:"slug"`
	Status         string `json:"status"`
	Url            string `json:"url"`
	Title          string `json:"title"`
	AlternateTitle string `json:"alternate_title"`
	EpisodeCount   int    `json:"episode_count"`
	CoverImage     string `json:"cover_image"`
	Synopsis       string `json:"synopsis"`
	ShowType       string `json:"show_type"`
	Genres         []struct {
		Name string `json:"name"`
	} `json:"genres"`

	FavRank int `json:"fav_rank"`
	FavId   int `json:"fav_id"`
}

func GetAnimeById(id int) (Anime, error) {
	return GetAnimeBySlug(fmt.Sprintf("%d", id))
}

func GetAnimeBySlug(slug string) (Anime, error) {
	var anime Anime
	var e apiError

	if req, err := http.NewRequest("GET", "http://hummingbird.me/api/v1/anime/"+slug, nil); err == nil {
		req.Header.Add("User-Agent", userAgent)
		if resp, err := defaultClient.Do(req); err == nil {
			decoder := json.NewDecoder(resp.Body)
			if resp.StatusCode == 200 {
				if err := decoder.Decode(&anime); err == nil {
					return anime, nil
				} else {
					return anime, err
				}
			} else {
				if err := decoder.Decode(&e); err == nil {
					return anime, e
				} else {
					return anime, fmt.Errorf("Failed to retrieve anime.")
				}
			}
		} else {
			return anime, err
		}
	} else {
		return anime, err
	}
}

func FindAnime(query string) ([]Anime, error) {
	var anime []Anime
	var e apiError
	v := url.Values{}
	v.Set("query", query)

	if req, err := http.NewRequest("GET", "http://hummingbird.me/api/v1/search/anime?"+v.Encode(), nil); err == nil {
		req.Header.Add("User-Agent", userAgent)
		if resp, err := defaultClient.Do(req); err == nil {
			decoder := json.NewDecoder(resp.Body)
			if resp.StatusCode == 200 {
				if err := decoder.Decode(&anime); err == nil {
					return anime, nil
				} else {
					return anime, err
				}
			} else {
				if err := decoder.Decode(&e); err == nil {
					return anime, e
				} else {
					return anime, fmt.Errorf("Failed to retrieve anime.")
				}
			}
		} else {
			return anime, err
		}
	} else {
		return anime, err
	}
}
