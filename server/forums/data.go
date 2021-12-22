package forums

import (
	"database/sql"
	"fmt"
	"strings"
)

type Forum struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	TopicKeyword string `json:"topicKeyword"`
	Users        string `json:"users"`
}

type JsonForum struct {
	Name         string   `json:"name"`
	TopicKeyword string   `json:"topicKeyword"`
	Users        []string `json:"users"`
}

type DecodableUser struct {
	Name   string `json:"name"`
	Topics string `json:"topics"`
}

type User struct {
	Name   string   `json:"name"`
	Topics []string `json:"topics"`
}

type Store struct {
	Db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Db: db}
}

func (s *Store) ListForums() ([]*JsonForum, error) {
	rows, err := s.Db.Query("SELECT * FROM \"Forum\"")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*Forum
	for rows.Next() {
		var f Forum
		if err := rows.Scan(&f.Id, &f.Name, &f.TopicKeyword, &f.Users); err != nil {
			return nil, err
		}
		res = append(res, &f)
	}
	if res == nil {
		res = make([]*Forum, 0)
	}

	var jsonRes []*JsonForum
	users, err := s.ListUsers()

	if err != nil {
		return nil, err
	}

	for _, forum := range res {
		var forumRes JsonForum
		forumRes.Name = forum.Name
		forumRes.TopicKeyword = forum.TopicKeyword
		forumRes.Users = []string{}
		for _, user := range users {
			if contains(user.Topics, forum.TopicKeyword) {
				forumRes.Users = append(forumRes.Users, user.Name)
			}
		}
		jsonRes = append(jsonRes, &forumRes)
	}
	if jsonRes == nil {
		jsonRes = make([]*JsonForum, 0)
	}

	return jsonRes, nil
}

func (s *Store) CreateUser(name string, topics []string) error {
	if len(name) < 1 || len(topics) < 1 {
		return fmt.Errorf("user name or topics are not provided")
	}

	var sqlTopics = topics
	for i, topic := range topics {
		sqlTopics[i] = fmt.Sprintf("'%s'", topic)
	}

	_, err := s.Db.Exec(fmt.Sprintf("INSERT INTO \"User\" VALUES ('%s', ARRAY[%s])", name, strings.Join(sqlTopics, ", ")))

	return err
}

func (s *Store) ListUsers() ([]*User, error) {
	rows, err := s.Db.Query("SELECT * FROM \"User\"")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var res []*User

	for rows.Next() {
		var u DecodableUser
		if err := rows.Scan(&u.Name, &u.Topics); err != nil {
			return nil, err
		}

		trimedTopics := u.Topics[1 : len(u.Topics)-1]
		topicsArr := strings.Split(string(trimedTopics), ",")

		var user User
		user.Name = u.Name
		user.Topics = topicsArr

		res = append(res, &user)

	}
	if res == nil {
		res = make([]*User, 0)
	}

	return res, nil
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
