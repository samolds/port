// Copyright (C) 2018 Sam Olds

package handler

import (
	"net/http"

	"github.com/samolds/port/template"
)

type link struct {
	Href    string
	Display string
}

func Link(w http.ResponseWriter, r *http.Request) error {
	links := make([]link, 0)
	links = append(links, link{
		Href:    "https://amazon.com/gp/pdp/profile/A2AMJ9HOJK0LCC",
		Display: "Amazon"})
	links = append(links, link{
		Href:    "https://angel.co/samolds",
		Display: "AngelList"})
	links = append(links, link{
		Href:    "http://bitbucket.org/samolds",
		Display: "Bitbucket"})
	links = append(links, link{
		Href:    "http://facebook.com/samolds",
		Display: "Facebook"})
	links = append(links, link{
		Href:    "http://fiverr.com/samolds",
		Display: "Fiverr"})
	links = append(links, link{
		Href:    "http://flickr.com/samolds/sets",
		Display: "Flickr"})
	links = append(links, link{
		Href:    "http://github.com/samolds",
		Display: "GitHub"})
	links = append(links, link{
		Href:    "http://code.google.com/u/samolds",
		Display: "Google Code"})
	links = append(links, link{
		Href:    "http://plus.google.com/+samolds",
		Display: "Google+"})
	links = append(links, link{
		Href:    "http://news.ycombinator.com/user?id=samolds",
		Display: "Hacker News"})
	links = append(links, link{
		Href:    "http://imgur.com/user/samolds",
		Display: "Imgur"})
	links = append(links, link{
		Href:    "http://instagram.com/samuraiolds",
		Display: "Instagram"})
	links = append(links, link{
		Href:    "http://keybase.io/samolds",
		Display: "Keybase"})
	links = append(links, link{
		Href:    "http://last.fm/user/samolds",
		Display: "Last.fm"})
	links = append(links, link{
		Href:    "http://linkedin.com/in/samolds",
		Display: "linkedIn"})
	links = append(links, link{
		Href:    "http://myspace.com/samolds",
		Display: "Myspace"})
	links = append(links, link{
		Href:    "http://reddit.com/user/samolds",
		Display: "Reddit"})
	links = append(links, link{
		Href:    "http://samolds.com",
		Display: "Sam Olds"})
	links = append(links, link{
		Href:    "http://soundcloud.com/samolds",
		Display: "SoundCloud"})
	links = append(links, link{
		Href:    "http://open.spotify.com/user/samolds",
		Display: "Spotify"})
	links = append(links, link{
		Href:    "http://stackoverflow.com/users/1604235/samolds",
		Display: "Stack Overflow"})
	links = append(links, link{
		Href:    "http://twitter.com/samolds",
		Display: "Twitter"})
	links = append(links, link{
		Href:    "http://eng.utah.edu/~samuelo",
		Display: "Utah Student Page"})
	links = append(links, link{
		Href:    "http://youtube.com/samuraiolds",
		Display: "YouTube"})

	return template.Link.Render(w, links)
}
