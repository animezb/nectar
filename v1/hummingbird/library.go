package hummingbird

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type LibraryEntry struct {
	Id              int                `json:"id"`
	EpisodesWatched int                `json:"episodes_watched"`
	LastWatched     time.Time          `json:"last_watched"`
	RewatchedTimes  int                `json:"rewated_times"`
	Notes           string             `json:"notes"`
	NotesPresent    bool               `json:"notes_present"`
	Status          string             `json:"status"`
	Private         bool               `json:"private"`
	Rewatching      bool               `json:"rewatching"`
	Anime           Anime              `json:"anime"`
	Rating          LibraryEntryRating `json:"rating"`
}

type LibraryEntryRating struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func GetLibrary(username string) ([]LibraryEntry, error) {
	var library []LibraryEntry
	var e apiError

	if req, err := http.NewRequest("GET", "http://hummingbird.me/api/v1/users/"+username+"/library", nil); err == nil {
		req.Header.Add("User-Agent", userAgent)
		if resp, err := defaultClient.Do(req); err == nil {
			decoder := json.NewDecoder(resp.Body)
			if resp.StatusCode == 200 {
				if err := decoder.Decode(&library); err == nil {
					return library, nil
				} else {
					return library, err
				}
			} else {
				if err := decoder.Decode(&e); err == nil {
					return library, e
				} else {
					return library, fmt.Errorf("Failed to retrieve library.")
				}
			}
		} else {
			return library, err
		}
	} else {
		return library, err
	}
}

func RemoveLibaryEntryById(id int, token AuthenicationToken) (bool, error) {
	return RemoveLibaryEntryBySlug(fmt.Sprintf("%d", id), token)
}

func RemoveLibaryEntryBySlug(slug string, token AuthenicationToken) (bool, error) {
	var result bool
	var e apiError
	result = false

	if req, err := http.NewRequest("POST", "http://hummingbird.me/api/v1/libraries/"+slug+"/remove?auth_token="+token.Token(), nil); err == nil {
		req.Header.Add("User-Agent", userAgent)
		if resp, err := defaultClient.Do(req); err == nil {
			decoder := json.NewDecoder(resp.Body)
			if resp.StatusCode == 200 {
				if err := decoder.Decode(&result); err == nil {
					return result, nil
				} else {
					return result, err
				}
			} else {
				if err := decoder.Decode(&e); err == nil {
					return result, e
				} else {
					return result, fmt.Errorf("Failed to retrieve result.")
				}
			}
		} else {
			return result, err
		}
	} else {
		return result, err
	}
}
