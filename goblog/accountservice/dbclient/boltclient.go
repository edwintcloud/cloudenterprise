package dbclient

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/cloudenterprise/goblog/accountservice/model"

	"github.com/boltdb/bolt"
)

// IBoltClient defines the contract we need our BoltClient to fill
type IBoltClient interface {
	OpenBoltDb()
	// QueryAccount(accountID string) (model.Account, error)
	Seed()
}

// BoltClient is the instance of our boltdb client
type BoltClient struct {
	boltDB *bolt.DB
}

// OpenBoltDb opens an instance to the db
func (bc *BoltClient) OpenBoltDb() {
	var err error

	bc.boltDB, err = bolt.Open("accounts.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// Seed starts seeding account
func (bc *BoltClient) Seed() {
	bc.initializeBucket()
	bc.seedAccounts()
}

// Creates an "AccountBucket" in our BoltDB. It will overwrite any existing bucket of the same name.
func (bc *BoltClient) initializeBucket() {
	bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("AccountBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

// Seed(n) make-believe account objects into the AccountBucket bucket.
func (bc *BoltClient) seedAccounts() {

	total := 100
	for i := 0; i < total; i++ {

		// Generate a key 10000 or larger
		key := strconv.Itoa(10000 + i)

		// Create an instance of our Account struct
		acc := model.Account{
			ID:   key,
			Name: "Person_" + strconv.Itoa(i),
		}

		// Serialize the struct into JSON
		jsonBytes, _ := json.Marshal(acc)

		// Write the data to AccountBucket
		bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("AccountBucket"))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
	fmt.Printf("Seeded %v fake accounts... \n", total)
}
