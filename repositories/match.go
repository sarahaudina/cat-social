package repositories

import (
 "database/sql"
 "log"

 "psprint/cat/models"
 "time"
)

type MatchRepository struct {
 DB *sql.DB
}

func NewMatchRepository(db *sql.DB) MatchRepositoryInterface {
 return &MatchRepository{DB: db}
}

func (m *MatchRepository) InsertMatch(post models.InsertMatch) bool {
 stmt, err := m.DB.Prepare("INSERT INTO matchs (status, issuer, receiver) VALUES ($1, $2, $3)")
 if err != nil {
  log.Println(err)
  return false
 }
 defer stmt.Close()
 _, err2 := stmt.Exec(post.Status, post.Issuer, post.Receiver)
 if err2 != nil {
  log.Println(err2)
  return false
 }
 return true
}

func (m *MatchRepository) GetMatch() []models.Match {
    var result []models.Match

    rows, err := m.DB.Query("SELECT * FROM matchs") // todo: fix this

    if err != nil {
     log.Println(err)
     return nil
    }
    for rows.Next() {
        var (
         id       uint
         status    string
         issuer    uint
         receiver    uint
         created_at  time.Time
         updated_at time.Time
        )
        err := rows.Scan(&id, &status, &issuer, &receiver, &created_at, &updated_at)
        if err != nil {
         log.Println(err)
        } else {
         match := models.Match{ID: id, Status: status, Issuer: issuer, Receiver: receiver, CreatedAt: created_at, UpdatedAt: updated_at}
         result = append(result, match)
        }
    }

    return result
}

func (m *MatchRepository) UpdateMatch(post models.UpdateMatch) bool {
    stmt, err := m.DB.Prepare("UPDATE matchs SET status = $1 WHERE id = $2;")
    if err != nil {
     log.Println(err)
     return false
    }
    defer stmt.Close()
    _, err2 := stmt.Exec(post.Status, post.ID)
    if err2 != nil {
     log.Println(err2)
     return false
    }
    return true
}

func (m *MatchRepository) DeleteMatch(matchId uint) bool {
    stmt, err := m.DB.Prepare("DELETE FROM matchs WHERE id = $1;")
    if err != nil {
     log.Println(err)
     return false
    }
    defer stmt.Close()
    _, err2 := stmt.Exec(matchId)
    if err2 != nil {
     log.Println(err2)
     return false
    }
    return true
}