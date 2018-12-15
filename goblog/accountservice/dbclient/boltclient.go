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
	QueryAccount(accountID string) (model.Account, error)
	Seed()
	Check() bool
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

// QueryAccount queries db for an account
func (bc *BoltClient) QueryAccount(accountID string) (model.Account, error) {

	// Allocate an empty account instance we will use later to populate with JSON using json.Unmarshal
	account := model.Account{}

	// Read an object from the bucket using boltDb.view
	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		// Read the bucket from the db
		b := tx.Bucket([]byte("AccountBucket"))

		// Read the value identified by our accountID supplied as []byte
		accountBytes := b.Get([]byte(accountID))
		if accountBytes == nil {
			return fmt.Errorf("No account found for account %s", accountID)
		}

		// Unmarshal the returned bytes	 into the account struct we created earlier
		json.Unmarshal(accountBytes, &account)

		// return nil to indicate nothing went wrong, e.g no error
		return nil
	})
	// If there was an error, return the error
	if err != nil {
		return model.Account{}, err
	}
	// else return the result and nil error
	return account, nil
}

// Check is a naive healthcheck - just checks to make sure db connection has been initialized
func (bc *BoltClient) Check() bool {
	return bc.boltDB != nil
}
