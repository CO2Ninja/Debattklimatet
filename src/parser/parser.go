package main

import (
        "database/sql"
        "fmt"
        _ "github.com/lib/pq"
        "log"
		"strings"
		"os"
		"regexp"
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
    rows, err := db.Query("SELECT id, text, user_id, parsed FROM debattklimatet_tweet")
    
    if err != nil {
            log.Fatal(err)
    }
    for rows.Next() {
            //createdat | favoritecount | favorited | id | idstr | retweetcount | retweeted | source | text | user_id
            var id int64
            var text string
            var user_id int64
            var parsed bool
            err := rows.Scan(&id, &text, &user_id, &parsed)
            if  err != nil {
                log.Fatal(err)
            }
            //fmt.Printf("user_id %d tweeted %s\n", user_id, text)

            if err := rows.Err(); err != nil {
            log.Fatal(err)
            }
    
            getPoint := false
            unwantedWords := make([]string, 2)
            unwantedWords[0] = "miljöpartiet"
            // Removes hashtags
            unwantedWords[1] = "(#\\S*)"
            wantedWords := make([]string, 2)
            wantedWords[0] = "hållbar utveckling"
            wantedWords[1] = "miljö([a-z]*)"
            
            tweetText := strings.ToLower(text)

            for i, _ := range unwantedWords {
                tweetText = removeUnwanted(unwantedWords[i], tweetText)
            }
            
            for i, _ := range wantedWords {
                if hasExpression(wantedWords[i], tweetText){
                    fmt.Println("it works!")
                    getPoint = true
                    //status = true
                }
            } 
            if getPoint {
                recountPoints(user_id, db)
            }
    }
    
    
}

	//Removes the specified unwanted expression from the tweet (if the tweet contains the expression)
	func removeUnwanted(expression string, tweet string) string{
	
	reg, err := regexp.Compile ("expression")
    if err != nil {
        fmt.Println("Compile failed: %s", err)
        os.Exit (1)
    }
	return string (reg.ReplaceAllString (tweet, ("")))
					  
	}				  
	
	// Checks if the tweet contains the specified expression 
	func hasExpression (expression string, tweet string) bool{
		r, _ := regexp.Compile("expression")
		return r.Match([]byte("tweet"))										//om något av uttrycken finns, returnera true, annars false
	}
	
	func recountPoints(userId int64, db *sql.DB) {
            //Here we update add one point to the users totalscore column.
            _, err := db.Exec("UPDATE debattklimatet_twitterusers SET totalscore = totalscore + 1 WHERE id=$1", userId)
            if err != nil {     
                        fmt.Println(err)
                    }
                    
		//points int := ...
		//totalPoints int := ...
		//points++
		//totalPoints++
		// Send points and totalPoints back to DB
	}
