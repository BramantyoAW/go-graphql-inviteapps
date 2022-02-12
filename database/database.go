package database

import(
	"database/sql"
	"log"
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/BramantyoAW/go-graphql-inviteapps/graph/model"
)

var Db *sql.DB

func InitDB(){
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/invitation_apps")
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	Db = db
	log.Println("Connection was sucessfull!!")
  }


// func selectUsersById(id string) (*model.Users, error) {
// 	stmt, err := Db.Prepare("select * from users where id=?")
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}

// 	defer stmt.Close()
// 	rows, err := stmt.Query(id)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}

// 	defer rows.Close()
	
// 	var users model.Users
// 	for rows.Next() {
// 		err = rows.Scan(&users.id, &users.name)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}

// 	defer rows.Close()

// 	if err = rows.Err(); err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}
// 	return &users, nil
// }


func getUsers() {
	stmt, err := Db.Prepare("select * from users")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()
	
	var users model.Users
	for rows.Next() {
		err = rows.Scan(&users.id, &users.name)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &users, nil
}