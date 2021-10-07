package main

import "time"

type Sample struct {
	ApiVersion string `json:"apiVersion"`
	Data       struct {
		Updated      time.Time `json:"updated"`
		TotalItems   int       `json:"totalItems"`
		StartIndex   int       `json:"startIndex"`
		ItemsPerPage int       `json:"itemsPerPage"`
		Items        []struct {
			Id          string    `json:"id"`
			Uploaded    time.Time `json:"uploaded"`
			Updated     time.Time `json:"updated"`
			Uploader    string    `json:"uploader"`
			Category    string    `json:"category"`
			Title       string    `json:"title"`
			Description string    `json:"description"`
			Tags        []string  `json:"tags"`
			Thumbnail   struct {
				Default   string `json:"default"`
				HqDefault string `json:"hqDefault"`
			} `json:"thumbnail"`
			Player struct {
				Default string `json:"default"`
			} `json:"player"`
			Content struct {
				Field1 string `json:"1"`
				Field2 string `json:"5"`
				Field3 string `json:"6"`
			} `json:"content"`
			Duration      int     `json:"duration"`
			AspectRatio   string  `json:"aspectRatio"`
			Rating        float64 `json:"rating"`
			RatingCount   int     `json:"ratingCount"`
			ViewCount     int     `json:"viewCount"`
			FavoriteCount int     `json:"favoriteCount"`
			CommentCount  int     `json:"commentCount"`
			Status        struct {
				Value  string `json:"value"`
				Reason string `json:"reason"`
			} `json:"status"`
			AccessControl struct {
				Syndicate    string `json:"syndicate"`
				CommentVote  string `json:"commentVote"`
				Rate         string `json:"rate"`
				List         string `json:"list"`
				Comment      string `json:"comment"`
				Embed        string `json:"embed"`
				VideoRespond string `json:"videoRespond"`
			} `json:"accessControl"`
		} `json:"items"`
	} `json:"data"`
}
