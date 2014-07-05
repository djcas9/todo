package todo

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"os/user"
	"path"
	"strconv"
)

func hasArgs(args []string, expected int) bool {
	switch expected {
	case -1:
		if len(args) > 0 {
			return true
		} else {
			return false
		}
	default:
		if len(args) == expected {
			return true
		} else {
			return false
		}
	}
}

func OpenTodoDB() *leveldb.DB {
	usr, err := user.Current()

	if err != nil {
		log.Fatal(err)
	}

	db, err := leveldb.OpenFile(path.Join(usr.HomeDir, DATABASE_PATH), nil)
	return db
}

func ShowAction(c *cli.Context) {
	if !hasArgs(c.Args(), 1) {
		log.Fatal("Incorrect number of arguments for the 'show' command")
	}

	id, _ := strconv.ParseInt(c.Args()[0], 0, 64)

	key, _, err := GetById(id)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", key)
}

func RemoveAction(c *cli.Context) {
	if !hasArgs(c.Args(), 1) {
		log.Fatal("Incorrect number of arguments for the 'remove' command")
	}

	id, _ := strconv.ParseInt(c.Args()[0], 0, 64)

	key, db, err := GetById(id)

	if err != nil {
		log.Fatal(err)
	}

	db.Delete(key, nil)
}

func AddAction(c *cli.Context) {
	if !hasArgs(c.Args(), 2) {
		log.Fatal("Incorrect number of arguments for the 'add' command")
	}

	db := OpenTodoDB()

	err := db.Put([]byte(c.Args()[0]), []byte(c.Args()[1]), nil)

	if err != nil {
		log.Fatal(err)
	}
}

func GetById(id int64) ([]byte, *leveldb.DB, error) {
	db := OpenTodoDB()
	var index int64 = 0

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()

		index++

		if id == index {
			return key, db, nil
		}
	}

	iter.Release()
	err := iter.Error()

	if err != nil {
		log.Fatal(err)
	}

	return []byte(""), db, fmt.Errorf("No todo found for this id")
}

func ListAction(c *cli.Context) {
	db := OpenTodoDB()
	var index int = 0

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := iter.Key()
		value := iter.Value()

		index++
		fmt.Printf("ID: %d %s %s\n", index, key, value)
	}

	iter.Release()
	err := iter.Error()

	if err != nil {
		log.Fatal(err)
	}
}
