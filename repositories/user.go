package repositories

import (
 "database/sql"
 "log"

 "psprint/cat/models"
 "time"
)

type UserRepository struct {
 DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepositoryInterface {
 return &UserRepository{DB: db}
}

func (m *UserRepository) InsertUser(post models.InsertUser) bool {
 stmt, err := m.DB.Prepare("INSERT INTO users (username, password) VALUES ($1, $2)")
 if err != nil {
  log.Println(err)
  return false
 }
 defer stmt.Close()
 _, err2 := stmt.Exec(post.Username, post.Password)
 if err2 != nil {
  log.Println(err2)
  return false
 }
 return true
}

func (m *UserRepository) FindUser(username string) *models.User {
    rows, err := m.DB.Query("SELECT * FROM users WHERE username = $1 LIMIT 1", username)

    if err != nil {
     log.Println(err)
     return nil
    }

    for rows.Next() {
        var (
         id       uint
         username    string
         password    string
         created_at  time.Time
         updated_at time.Time
        )
        err := rows.Scan(&id, &username, &password, &created_at, &updated_at)
        if err != nil {
         log.Println(err)
         return nil
        } else {
         user := models.User{ID: id, Username: username, Password: password, CreatedAt: created_at, UpdatedAt: updated_at}
         return &user
        }
    }

    return nil
}