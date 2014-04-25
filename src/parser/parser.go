package main

import (
        "database/sql"
        "fmt"
        _ "github.com/lib/pq"
        "log"
)

var dbURL = "user=co2ninjas dbname=co2ninjas password=co2ninjas12345 host=django-db.cyyapufsikx9.eu-west-1.rds.amazonaws.com port=5432"

func dbConnect(database string, parameters string) *sql.DB {
        db, err := sql.Open(database, parameters)
        if err != nil {
                log.Fatal(err)
        }
        return db
}

func main() {
    db := dbConnect("postgres", dbURL) 
    rows, err := db.Query("SELECT * FROM debattklimatet_tweet")
    
    if err != nil {
            log.Fatal(err)
    }
    for rows.Next() {
            //createdat | favoritecount | favorited | id | idstr | retweetcount | retweeted | source | text | user_id
            var createdat string
            var favoritecount int
            var favorited bool
            var id int64
            var idstr string
            var retweetcount int
            var retweeted bool
            var source string
            var text string
            var user_id int64
            if err := rows.Scan(&createdat, &favoritecount, &favorited, &id, &idstr, &retweetcount, &retweeted, &source, &text, &user_id); err != nil {
                    log.Fatal(err)
            }
            fmt.Printf("user_id %d tweeted %s\n", user_id, text)
    }
    
    if err := rows.Err(); err != nil {
            log.Fatal(err)
    }
}
