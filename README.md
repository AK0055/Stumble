# Stumble
A dating app created using GoLang
<h3>Task Description</h3>
Imagine there is a dating app “Stumble” where there are many
users. There is also a feature for a user to like as many users as he/
she wants. You are given the list of these users in the JSON file.
Each user will have id (int mandatory), name (string mandatory),
location (float mandatory), gender (string optional), email (string
optional).

Also you are given a list of likes. (eg. A likes B- C likes A- B likes A- D
likes C)
- Create a database model in any RDBMS for the dating app
using Golang’s ORM of your choice.
- After creating the database model- import the user’s JSON file
and “likes” data in the DB. 
- Create an endpoint to retrieve all the matches. (Match is when X
likes Y & Y likes X) 
- Given user X and distance k- create an endpoint to retrieve all
the users within distance k from user X. (Assume you are given the
distance of each user in the 1D coordinate system)_
- Given input query q (string)- create an endpoint to retrieve all the
users which have q in their name
(Use localhost and port of your choice for creating endpoints
create api documentation- bonus for deployment)

<h3>Libraries used</h3>
-   "encoding/json" for JSON parsing and encoding
-	"fmt" for basic formatting and printing 
-	"log" for an active logging object that generates lines of output to an io.Writer
-	"math" for using absolute function
-    "net/http" for serving HTTP requests
-	"strconv" for string conversion
-	"strings" for string manipulation
-	"github.com/gorilla/mux" for routing different endpoints 
-	"gorm.io/driver/sqlite" for providing functions of the mysqlite GORM model
-	"gorm.io/gorm" for providing GORM capabilities
)
<h3>API documentation</h3>
- Use the URL: `http://localhost:10000/` to access the homepage
- `http://localhost:10000/matches` to access the matches endpoint
- `http://localhost:10000/distance` to access the distance endpoint
- `http://localhost:10000/substring` to access the substring query endpoint