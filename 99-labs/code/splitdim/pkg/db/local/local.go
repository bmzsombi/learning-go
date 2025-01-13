package local

import (
	"errors"
	"math"
	"sort"
	"splitdim/pkg/api"
	"sync"
)

// localDB is a simple implementation of the DataLayer API.
type localDB struct {
	// accounts maintains the balance for each user name
	accounts map[string]int
	// The read-write mutex makes sure concurrent access is safe.
	mu sync.RWMutex
}

// NewDataLayer creates a new database of accounts.
func NewDataLayer() api.DataLayer {
	return &localDB{accounts: make(map[string]int)}
}

func (db *localDB) Transfer(t api.Transfer) error {
	if t.Sender == t.Receiver {
		err := errors.New("sender and receiver are the same, cancelling transaction")
		return err
	}
	db.mu.Lock()
	defer db.mu.Unlock()
	_, exist := db.accounts[t.Sender]
	if !exist {
		db.accounts[t.Sender] = 0
		db.accounts[t.Sender] += t.Amount
	} else {
		db.accounts[t.Sender] += t.Amount
	}
	_, exist = db.accounts[t.Receiver]
	if !exist {
		db.accounts[t.Receiver] = 0
		db.accounts[t.Receiver] -= t.Amount
	} else {
		db.accounts[t.Receiver] -= t.Amount
	}
	return nil
}
func (db *localDB) AccountList() ([]api.Account, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	ret := []api.Account{}
	for key, val := range db.accounts {
		account := api.Account{
			Holder:  key,
			Balance: val,
		}
		ret = append(ret, account)
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Holder < ret[j].Holder
	})
	return ret, nil
}
func (db *localDB) Clear() ([]api.Transfer, error) {
	db.mu.Lock()
	//defer db.mu.Unlock()
	var accountDebts int
	for _, val := range db.accounts {
		accountDebts += val
	}
	if accountDebts != 0 {
		accountBalanceError := errors.New("account balance is inconsistent")
		return nil, accountBalanceError
	}
	//tempAcc := db.accounts
	tempAcc := make(map[string]int)
	for key, value := range db.accounts {
		tempAcc[key] = value
	}
	db.mu.Unlock()
	transfers := make([]api.Transfer, 0)
	for {
		cleared := true
		for sender, balance := range tempAcc {
			if balance < 0 {
				for receiver, receiverBalance := range tempAcc {
					if receiverBalance > 0 {
						transferAmount := int(math.Min(float64(-balance), float64(receiverBalance)))
						transfers = append(transfers, api.Transfer{
							Sender:   sender,
							Receiver: receiver,
							Amount:   transferAmount,
						})

						tempAcc[sender] += transferAmount
						tempAcc[receiver] -= transferAmount

						cleared = false

						if tempAcc[sender] == 0 {
							break
						}
					}
				}
			}
		}
		if cleared {
			break
		}
	}

	return transfers, nil
}
func (db *localDB) Reset() error {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.accounts = make(map[string]int)
	return nil
}
