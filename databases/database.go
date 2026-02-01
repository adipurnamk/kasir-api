package databases

import (
    "context"
    "database/sql"
    "log"
    "strings"
    "time"

    // use pgx stdlib driver
    _ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB(connectionString string) (*sql.DB, error) {
    // Mask password for logs
    masked := connectionString
    if at := strings.Index(masked, "@"); at != -1 {
        if proto := strings.Index(masked, "://"); proto != -1 {
            masked = masked[:proto+3] + "*****" + masked[at:]
        }
    }
    log.Println("Connecting to DB:", masked)

    // open with pgx driver
    db, err := sql.Open("pgx", connectionString)
    if err != nil {
        return nil, err
    }

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    if err := db.PingContext(ctx); err != nil {
        return nil, err
    }

    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(5)

    log.Println("Database connected successfully")
    return db, nil
}