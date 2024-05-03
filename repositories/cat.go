package repositories

import (
 "database/sql"
 "log"

 "psprint/cat/models"
 "time"
)

type CatRepository struct {
 DB *sql.DB
}

func NewCatRepository(db *sql.DB) CatRepositoryInterface {
 return &CatRepository{DB: db}
}

func (m *CatRepository) InsertCat(post models.InsertCat) bool {
 stmt, err := m.DB.Prepare("INSERT INTO cats (name, gender, user_id) VALUES ($1, $2, $3)")
 if err != nil {
  log.Println(err)
  return false
 }
 defer stmt.Close()
 _, err2 := stmt.Exec(post.Name, post.Gender, post.UserId)
 if err2 != nil {
  log.Println(err2)
  return false
 }
 return true
}

func (m *CatRepository) GetCats() []models.Cat {
    var result []models.Cat

    rows, err := m.DB.Query("SELECT * FROM cats")

    if err != nil {
     log.Println(err)
     return nil
    }
    for rows.Next() {
        var (
         id       uint
         name    string
         gender    string
         user_id    uint
         created_at  time.Time
         updated_at time.Time
        )
        err := rows.Scan(&id, &name, &gender, &user_id, &created_at, &updated_at)
        if err != nil {
         log.Println(err)
        } else {
         cat := models.Cat{ID: id, Name: name, Gender: gender, UserId: user_id, CreatedAt: created_at, UpdatedAt: updated_at}
         result = append(result, cat)
        }
    }

    return result
}