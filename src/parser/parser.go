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
            var parsed bool
            var relevant bool
            if err := rows.Scan(&createdat, &favoritecount, &favorited, &id, &idstr, &retweetcount, &retweeted, &source, &text, &user_id, &parsed, &relevant); err != nil {
                    log.Fatal(err)
            }
            fmt.Printf("user_id %d tweeted %s\n", user_id, text)
    }
    
    if err := rows.Err(); err != nil {
            log.Fatal(err)
    }
	
		getPoint bool := false
		unwantedWords := make([]string, 2)
		unwantedWords[0] = "miljöpartiet"
		// Removes hashtags
		unwanteedWords[1] = "(#\S*)"
		wantedWords := make([]string, 2)
		wantedWords[0} = "hållbar utveckling"
		wantedWords[1] = "miljö([a-z]*)"
		
		text = toLower(text)

		for i int = 0, i < len(unwantedWords), i++ {
			text := removeUnwanted(unwantedWords[i], text)
		}
		
		for i int = 0, i < len(wantedWords), i++ {
			if hasExpression(wantedWord[i], text){
				getPoint = true
				//status = true
			}
		}
		
		if getPoint {
			recount(user_id, db)
		}
		
	}
}

	//Removes the specified unwanted expression from the tweet (if the tweet contains the expression)
	func removeUnwanted(string expression, string tweet) string{
	
	reg, error := regexp.Compile ("expression")
    if error != nil {
        fmt.Printf ("Compile failed: %s", error.String ())
        os.Exit (1)
    }
	output := ""
		output = string (reg.ReplaceAll (strings.Bytes (tweet),
    			      strings.Bytes ("")))
					  
	}				  
	
	// Checks if the tweet contains the specified expression 
	func hasExpression (string expression, string tweet) bool{
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
