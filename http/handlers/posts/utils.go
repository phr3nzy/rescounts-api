package posts

import (
	"sort"
)

// removeDuplicates removes duplicates from posts.
func removeDuplicates(posts []Post) []Post {
	var postsWithoutDuplicates []Post
	allIDs := make(map[int64]bool)

	for _, post := range posts {
		if _, value := allIDs[post.ID]; !value {
			allIDs[post.ID] = true
			postsWithoutDuplicates = append(postsWithoutDuplicates, post)
		}
	}

	return postsWithoutDuplicates
}

// sortPosts accepts an array of `Post` and sorts it based on a field
// and direction (asc | desc). Defaults are `id` and `asc`.
func sortPosts(posts []Post, field, direction string) []Post {
	var postsToSort []Post = posts

	switch field {
	case "likes":
		{
			sort.SliceStable(postsToSort, func(i, j int) bool {
				return postsToSort[i].Likes < postsToSort[j].Likes
			})
		}
	case "popularity":
		{
			sort.SliceStable(postsToSort, func(i, j int) bool {
				return postsToSort[i].Popularity < postsToSort[j].Popularity
			})
		}
	case "reads":
		{
			sort.SliceStable(postsToSort, func(i, j int) bool {
				return postsToSort[i].Reads < postsToSort[j].Reads
			})
		}
	case "id":
		{
			sort.SliceStable(postsToSort, func(i, j int) bool {
				return postsToSort[i].ID < postsToSort[j].ID
			})
		}
	default:
		{
			sort.SliceStable(postsToSort, func(i, j int) bool {
				return postsToSort[i].ID < postsToSort[j].ID
			})
		}
	}

	if direction == "desc" {
		for i, j := 0, len(postsToSort)-1; i < j; i, j = i+1, j-1 {
			postsToSort[i], postsToSort[j] = postsToSort[j], postsToSort[i]
		}
	}

	return postsToSort
}
