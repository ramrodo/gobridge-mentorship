package main

import (
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v2"
)

func main() {
	// Create/Open a temporal database
	db, err := badger.Open(badger.DefaultOptions("/tmp/test1"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Key
	key := "user"

	// Create/Update a key
	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte("ramrodo"))
		return err
	})

	if err != nil {
		log.Fatalf("Error creating/updating the key '%v': %s", key, err)
	}

	// Get the key
	var valCopy []byte
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			valCopy = append([]byte{}, val...)
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error getting the key '%v': %s", key, err)
	}

	fmt.Printf("The key '%v' has value '%s'\n", key, valCopy)

	// Delete the key
	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(key))
		return err
	})

	if err != nil {
		log.Fatalf("Error deleting the key '%v': %s", key, err)
	}
}
