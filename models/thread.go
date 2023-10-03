package models

import (
	"time"
)

type Thread struct {
	Id             int
	Uuid           string
	Topic          string
	UserId         int
	CreatedAt      time.Time
	Num            int
	Classification string
}

type UN struct {
	Id       int
	Username string
}

type TP struct {
	Id    int
	Topic string
}

type Result struct {
	UserName []UN
	Topic    []TP
}

// format the CreatedAt date to display nicely on the screen
func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("2006-01-02 15:04:05")
}

// get the number of posts in a thread
func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = ?", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

// get posts to a thread
func (thread *Thread) Posts() (posts []Post, err error) {
	rows, err := Db.Query("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts where thread_id = ?", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// Get all threads in the database and returns it
func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at, num FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt, &conv.Num); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	return
}

func DeleteTopic(id int) (err error) {
	statement := "Delete From threads WHERE id = ?"
	//Db.QueryRow("SELECT id FROM threads WHERE topic = ?", topic).Scan(&num)
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return
}

// 获得用户和帖子的结果集
func Union() (result []Result, err error) {
	//rows, err := Db.Query("SELECT name as name FROM users UNION SELECT topic as topic FROM threads")
	var username []UN
	var topic []TP
	//username := []UN{}
	//topic := []TP{}
	rows, err := Db.Query("SELECT id, name FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := UN{}
		if err = rows.Scan(&conv.Id, &conv.Username); err != nil {
			return
		} else {
			username = append(username, conv)
		}

	}
	rows1, err := Db.Query("SELECT id, topic FROM threads")
	if err != nil {
		return
	}
	for rows1.Next() {
		conv1 := TP{}
		if err = rows1.Scan(&conv1.Id, &conv1.Topic); err != nil {
			return
		} else {
			topic = append(topic, conv1)
		}

	}
	results := Result{
		UserName: username,
		Topic:    topic,
	}
	result = append(result, results)
	return
}

// Get a thread by the UUID
func ThreadByUUID(uuid string) (conv Thread, err error) {
	conv = Thread{}
	err = Db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = ?", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

// Get the user who started this thread
func (thread *Thread) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", thread.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

func LikePlus(id int) (err error) {
	var num int
	err = Db.QueryRow("SELECT num FROM threads WHERE id = ?", id).
		Scan(&num)
	num = num + 1
	statement := "update threads set num = ? where id = ?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(num, id)
	return
}

func Select(classification string) (threads []Thread, err error) {
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at, num FROM threads WHERE classification=? ORDER BY created_at DESC", classification)
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt, &conv.Num); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	return
}
