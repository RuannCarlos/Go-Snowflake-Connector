package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"snowflake_connector/go_connector/auth"

	"github.com/joho/godotenv"
	"github.com/snowflakedb/gosnowflake"
)

type Snowflake_User_Info struct {
	property,
	value,
	def,
	description string
}

func main() {
	env := load_envs()
	dsn, err := gosnowflake.DSN(env)
	if err != nil {
		panic(err)
	}

	fmt.Println("Creating Connection...")
	db, err := sql.Open("snowflake", dsn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection Created!")
	test_connection(db)

}

func load_envs() *gosnowflake.Config {
	fmt.Println("Loading environment variables...")
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	config := &gosnowflake.Config{
		Authenticator: gosnowflake.AuthTypeJwt,
		User:          os.Getenv("SNOWFLAKE_USER"),
		Account:       os.Getenv("ACCOUNT"),
		PrivateKey:    auth.Setup_private_key(os.Getenv("PRIVATE_KEY_PATH")),
		Role:          os.Getenv("ROLE"),
		Region:        os.Getenv("REGION"),
	}
	return config
}

func test_connection(db *sql.DB) {
	fmt.Println("Testing connection")
	query, err := db.Query("DESC USER w513180")
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()
	var user_infos []Snowflake_User_Info

	for query.Next() {
		var info Snowflake_User_Info
		if err := query.Scan(&info.property, &info.value, &info.def, &info.description); err != nil {
			log.Fatal(err)
		}
		user_infos = append(user_infos, info)
	}
	fmt.Println(user_infos)
}
