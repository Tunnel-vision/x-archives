package model


type Comment struct {
	Model
	ArticleID uint64
	AuthorId uint64
	Content string
	ParentCommentID uint64
	IP string
	UserAgent string
	PushedAt string

	AuthorName string
	AuthorAvatarURL string
	AuthorURL string
	BlogID uint64
}

// SyncCommentAuthorID is the id of sync comment bot.
const SyncCommentAuthorID = math.MaxInt32