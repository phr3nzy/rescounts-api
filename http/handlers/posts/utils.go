package posts

import (
	"sort"
)

// sortPosts accepts an array of `Post` and sorts it based on a field
// and direction (asc | desc).
func sortPosts(posts []Post, field, direction string) {
	switch direction {
	case "asc":
		{
			switch field {
			case "reads":
				{
					sort.Slice(posts, func(i, j int) bool {
						return posts[i].Reads < posts[j].Reads
					})
					break
				}
			case "likes":
				{
					sort.Slice(posts, func(i, j int) bool {
						return posts[i].Likes < posts[j].Likes
					})
					break
				}
			case "popularity":
				{
					sort.Slice(posts, func(i, j int) bool {
						return posts[i].Popularity < posts[j].Popularity
					})
					break
				}
			case "id":
			default:
				{
					sort.Slice(posts, func(i, j int) bool {
						return posts[i].ID < posts[j].ID
					})
					break
				}
			}
			break
		}
	case "desc":
	default:
		{
			switch field {
			case "reads":
				{
					sort.Slice(posts, func(i, j int) bool {
						return posts[i].Reads > posts[j].Reads
					})
					break
				}
			case "likes":
				{
					sort.Slice(posts, func(i, j int) bool {
						return posts[i].Likes > posts[j].Likes
					})
					break
				}
			case "popularity":
				{
					sort.Slice(posts, func(i, j int) bool {
						return posts[i].Popularity > posts[j].Popularity
					})
					break
				}
			case "id":
			default:
				{
					sort.Slice(posts, func(i, j int) bool {
						return posts[i].ID > posts[j].ID
					})
					break
				}
			}
			break
		}
	}
}
