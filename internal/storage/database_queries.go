package storage

var (
	getPostsList = `SELECT id, comments
		FROM posts;`
	getPostById = `SELECT id, text, comments
		FROM posts
		WHERE id = $1;`
	getCommentsByPost = `SELECT comment_id, post_id, text
		FROM comments
		WHERE post_id = $1;`
	writeComment = `INSERT INTO comments
		(post_id, text)
		VALUES ($1, $2)
		RETURNING comment_id;`
	writeReference = `INSERT INTO comment_ref
		(original_comment, reference_comment)
		VALUES ($1, $2);`
	writePost = `INERT INTO posts
		(text, comments)
		VALUES ($1, $2)
		RETURNING id;`
	getCommentsStatus =`SELECT comments
		FROM posts
		WHERE id = $1`
)
