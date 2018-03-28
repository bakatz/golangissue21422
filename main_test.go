package golangissue21422

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	migrationContents, err := ioutil.ReadFile("./some_migration.sql")
	if err != nil {
		fmt.Println("Something weird happened, exiting with an error code :(")
		os.Exit(-1)
	}

	fmt.Println(fmt.Sprintf("Hey, I ran! The current time is %s and the contents of the sql file is: %s", time.Now(), migrationContents))
	m.Run()
	os.Exit(0)
}
