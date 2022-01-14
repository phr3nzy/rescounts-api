package posts

import (
	"reflect"
	"testing"
)

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		posts []Post
	}
	tests := []struct {
		name string
		args args
		want []Post
	}{
		{
			name: "remove duplicates",
			args: args{posts: []Post{
				{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
				{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
			}},
			want: []Post{
				{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.posts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortPostsByLikesDescendingly(t *testing.T) {
	type args struct {
		posts     []Post
		field     string
		direction string
	}
	tests := []struct {
		name string
		args args
		want []Post
	}{
		{
			name: "sort posts by likes descendingly",
			args: args{
				posts: []Post{
					{ID: 379, Author: "Ahmed", AuthorID: 96, Likes: 123, Popularity: 459.3, Reads: 12504, Tags: []string{"design", "science"}},
					{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
					{ID: 89, Author: "Mohammed", AuthorID: 77, Likes: 218, Popularity: 68.9, Reads: 504, Tags: []string{"tech", "design"}},
				},
				field:     "likes",
				direction: "desc",
			},
			want: []Post{
				{ID: 89, Author: "Mohammed", AuthorID: 77, Likes: 218, Popularity: 68.9, Reads: 504, Tags: []string{"tech", "design"}},
				{ID: 379, Author: "Ahmed", AuthorID: 96, Likes: 123, Popularity: 459.3, Reads: 12504, Tags: []string{"design", "science"}},
				{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortPosts(tt.args.posts, tt.args.field, tt.args.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortPostsByLikesAscendingly(t *testing.T) {
	type args struct {
		posts     []Post
		field     string
		direction string
	}
	tests := []struct {
		name string
		args args
		want []Post
	}{
		{
			name: "sort posts by likes ascendingly",
			args: args{
				posts: []Post{
					{ID: 379, Author: "Ahmed", AuthorID: 96, Likes: 123, Popularity: 459.3, Reads: 12504, Tags: []string{"design", "science"}},
					{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
					{ID: 89, Author: "Mohammed", AuthorID: 77, Likes: 218, Popularity: 68.9, Reads: 504, Tags: []string{"tech", "design"}},
				},
				field:     "likes",
				direction: "asc",
			},
			want: []Post{
				{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
				{ID: 379, Author: "Ahmed", AuthorID: 96, Likes: 123, Popularity: 459.3, Reads: 12504, Tags: []string{"design", "science"}},
				{ID: 89, Author: "Mohammed", AuthorID: 77, Likes: 218, Popularity: 68.9, Reads: 504, Tags: []string{"tech", "design"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortPosts(tt.args.posts, tt.args.field, tt.args.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortPostsByPopularityDescendingly(t *testing.T) {
	type args struct {
		posts     []Post
		field     string
		direction string
	}
	tests := []struct {
		name string
		args args
		want []Post
	}{
		{
			name: "sort posts by popularity descendingly",
			args: args{
				posts: []Post{
					{ID: 379, Author: "Ahmed", AuthorID: 96, Likes: 123, Popularity: 459.3, Reads: 12504, Tags: []string{"design", "science"}},
					{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
					{ID: 89, Author: "Mohammed", AuthorID: 77, Likes: 218, Popularity: 68.9, Reads: 504, Tags: []string{"tech", "design"}},
				},
				field:     "popularity",
				direction: "desc",
			},
			want: []Post{
				{ID: 379, Author: "Ahmed", AuthorID: 96, Likes: 123, Popularity: 459.3, Reads: 12504, Tags: []string{"design", "science"}},
				{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
				{ID: 89, Author: "Mohammed", AuthorID: 77, Likes: 218, Popularity: 68.9, Reads: 504, Tags: []string{"tech", "design"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortPosts(tt.args.posts, tt.args.field, tt.args.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortPostsByPopularityAscendingly(t *testing.T) {
	type args struct {
		posts     []Post
		field     string
		direction string
	}
	tests := []struct {
		name string
		args args
		want []Post
	}{
		{
			name: "sort posts by popularity ascendingly",
			args: args{
				posts: []Post{
					{ID: 379, Author: "Ahmed", AuthorID: 96, Likes: 123, Popularity: 459.3, Reads: 12504, Tags: []string{"design", "science"}},
					{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
					{ID: 89, Author: "Mohammed", AuthorID: 77, Likes: 218, Popularity: 68.9, Reads: 504, Tags: []string{"tech", "design"}},
				},
				field:     "popularity",
				direction: "asc",
			},
			want: []Post{
				{ID: 89, Author: "Mohammed", AuthorID: 77, Likes: 218, Popularity: 68.9, Reads: 504, Tags: []string{"tech", "design"}},
				{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
				{ID: 379, Author: "Ahmed", AuthorID: 96, Likes: 123, Popularity: 459.3, Reads: 12504, Tags: []string{"design", "science"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortPosts(tt.args.posts, tt.args.field, tt.args.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortPostsByReadsDescendingly(t *testing.T) {
	type args struct {
		posts     []Post
		field     string
		direction string
	}
	tests := []struct {
		name string
		args args
		want []Post
	}{
		{
			name: "sort posts by reads descendingly",
			args: args{
				posts: []Post{
					{ID: 379, Author: "Ahmed", AuthorID: 96, Likes: 123, Popularity: 459.3, Reads: 12504, Tags: []string{"design", "science"}},
					{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
					{ID: 89, Author: "Mohammed", AuthorID: 77, Likes: 218, Popularity: 68.9, Reads: 504, Tags: []string{"tech", "design"}},
				},
				field:     "reads",
				direction: "desc",
			},
			want: []Post{
				{ID: 379, Author: "Ahmed", AuthorID: 96, Likes: 123, Popularity: 459.3, Reads: 12504, Tags: []string{"design", "science"}},
				{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
				{ID: 89, Author: "Mohammed", AuthorID: 77, Likes: 218, Popularity: 68.9, Reads: 504, Tags: []string{"tech", "design"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortPosts(tt.args.posts, tt.args.field, tt.args.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortPostsByReadsAscendingly(t *testing.T) {
	type args struct {
		posts     []Post
		field     string
		direction string
	}
	tests := []struct {
		name string
		args args
		want []Post
	}{
		{
			name: "sort posts by reads ascendingly",
			args: args{
				posts: []Post{
					{ID: 379, Author: "Ahmed", AuthorID: 96, Likes: 123, Popularity: 459.3, Reads: 12504, Tags: []string{"design", "science"}},
					{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
					{ID: 89, Author: "Mohammed", AuthorID: 77, Likes: 218, Popularity: 68.9, Reads: 504, Tags: []string{"tech", "design"}},
				},
				field:     "reads",
				direction: "asc",
			},
			want: []Post{
				{ID: 89, Author: "Mohammed", AuthorID: 77, Likes: 218, Popularity: 68.9, Reads: 504, Tags: []string{"tech", "design"}},
				{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
				{ID: 379, Author: "Ahmed", AuthorID: 96, Likes: 123, Popularity: 459.3, Reads: 12504, Tags: []string{"design", "science"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortPosts(tt.args.posts, tt.args.field, tt.args.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortPostsByIDDescendingly(t *testing.T) {
	type args struct {
		posts     []Post
		field     string
		direction string
	}
	tests := []struct {
		name string
		args args
		want []Post
	}{
		{
			name: "sort posts by reads descendingly",
			args: args{
				posts: []Post{
					{ID: 379, Author: "Ahmed", AuthorID: 96, Likes: 123, Popularity: 459.3, Reads: 12504, Tags: []string{"design", "science"}},
					{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
					{ID: 89, Author: "Mohammed", AuthorID: 77, Likes: 218, Popularity: 68.9, Reads: 504, Tags: []string{"tech", "design"}},
				},
				field:     "id",
				direction: "desc",
			},
			want: []Post{
				{ID: 379, Author: "Ahmed", AuthorID: 96, Likes: 123, Popularity: 459.3, Reads: 12504, Tags: []string{"design", "science"}},
				{ID: 89, Author: "Mohammed", AuthorID: 77, Likes: 218, Popularity: 68.9, Reads: 504, Tags: []string{"tech", "design"}},
				{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortPosts(tt.args.posts, tt.args.field, tt.args.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortPostsByIDAscendingly(t *testing.T) {
	type args struct {
		posts     []Post
		field     string
		direction string
	}
	tests := []struct {
		name string
		args args
		want []Post
	}{
		{
			name: "sort posts by id ascendingly",
			args: args{
				posts: []Post{
					{ID: 379, Author: "Ahmed", AuthorID: 96, Likes: 123, Popularity: 459.3, Reads: 12504, Tags: []string{"design", "science"}},
					{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
					{ID: 89, Author: "Mohammed", AuthorID: 77, Likes: 218, Popularity: 68.9, Reads: 504, Tags: []string{"tech", "design"}},
				},
				field:     "id",
				direction: "asc",
			},
			want: []Post{
				{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
				{ID: 89, Author: "Mohammed", AuthorID: 77, Likes: 218, Popularity: 68.9, Reads: 504, Tags: []string{"tech", "design"}},
				{ID: 379, Author: "Ahmed", AuthorID: 96, Likes: 123, Popularity: 459.3, Reads: 12504, Tags: []string{"design", "science"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortPosts(tt.args.posts, tt.args.field, tt.args.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortPostsDefaults(t *testing.T) {
	type args struct {
		posts     []Post
		field     string
		direction string
	}
	tests := []struct {
		name string
		args args
		want []Post
	}{
		{
			name: "sort posts by reads descendingly",
			args: args{
				posts: []Post{
					{ID: 379, Author: "Ahmed", AuthorID: 96, Likes: 123, Popularity: 459.3, Reads: 12504, Tags: []string{"design", "science"}},
					{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
					{ID: 89, Author: "Mohammed", AuthorID: 77, Likes: 218, Popularity: 68.9, Reads: 504, Tags: []string{"tech", "design"}},
				},
				field:     "",
				direction: "",
			},
			want: []Post{
				{ID: 1, Author: "Osama", AuthorID: 47, Likes: 23, Popularity: 89.9, Reads: 704, Tags: []string{"tech", "science"}},
				{ID: 89, Author: "Mohammed", AuthorID: 77, Likes: 218, Popularity: 68.9, Reads: 504, Tags: []string{"tech", "design"}},
				{ID: 379, Author: "Ahmed", AuthorID: 96, Likes: 123, Popularity: 459.3, Reads: 12504, Tags: []string{"design", "science"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortPosts(tt.args.posts, tt.args.field, tt.args.direction); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}
