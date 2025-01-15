package kvstore

import (
	"fmt"
	"math"
	"resilient"
	"strconv"
	"time"

	clientapi "kvstore/pkg/api"
	"kvstore/pkg/client"

	"splitdim/pkg/api"
)

type kvstore struct {
	client.Client
}

// NewDataLayer creates a new database of scores.
func NewDataLayer(kvStoreAddr string) api.DataLayer {
	return &kvstore{Client: client.NewClient(kvStoreAddr)}
}

func (db *kvstore) setBalanceForUser(user string, amount int) resilient.Closure {
	return func() error { return db.setBalance(user, amount) }
}

func (db *kvstore) setBalance(user string, amount int) error {
	vv, err := db.Get(user)
	if err != nil {
		return fmt.Errorf("error happened during api call: %w", err)
	}
	balance, _ := strconv.Atoi(vv.Value)         // Convert to integer: should be safe
	vv.Value = fmt.Sprintf("%d", balance+amount) // Set the new balance in the value
	vkv := clientapi.VersionedKeyValue{user, vv} // Create a VersionedKeyValue
	err = db.Put(vkv)
	if err != nil {
		return fmt.Errorf("couldn't put element to db: %w", err)
	}
	return nil
}

func (db *kvstore) Transfer(t api.Transfer) error {
	if t.Sender == "" || t.Receiver == "" || t.Sender == t.Receiver {
		return fmt.Errorf("invalid transfer")
	}

	var defaultBackoff = resilient.Backoff{
		Base:      150 * time.Millisecond,
		Cap:       2 * time.Second,
		Jitter:    3,
		NumTrials: 6,
	}

	senderClosure := db.setBalanceForUser(t.Sender, t.Amount)
	decoreatedIncreaseClosure := resilient.WithRetry(senderClosure, defaultBackoff)
	err := decoreatedIncreaseClosure()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	receiverClosure := db.setBalanceForUser(t.Receiver, -t.Amount)
	decoratedReceiverClosure := resilient.WithRetry(receiverClosure, defaultBackoff)
	err = decoratedReceiverClosure()
	if err != nil {
		senderClosure2 := db.setBalanceForUser(t.Sender, -t.Amount)
		decoratedSenderClosure := resilient.WithRetry(senderClosure2, defaultBackoff)
		err = decoratedSenderClosure()
		return fmt.Errorf("%w", err)
	} /*
		for {
			err := db.setBalance(t.Sender, -t.Amount)
			if err == nil {
				break
			}
		}
		for {
			err := db.setBalance(t.Receiver, -t.Amount)
			if err == nil {
				break
			}
		}*/
	return nil
}

func (db *kvstore) AccountList() ([]api.Account, error) {
	accounts, err := db.List()
	if err != nil {
		return []api.Account{}, fmt.Errorf("cannot list accounts %w", err)
	}
	ret := []api.Account{}
	for _, account := range accounts {
		balance, err := strconv.Atoi(account.Value)
		if err != nil {
			return nil, fmt.Errorf("invalid balance for account %s: %w", account.Key, err)
		}
		ret = append(ret, api.Account{
			Holder:  account.Key,
			Balance: balance,
		})
	}
	return ret, nil
}

func (db *kvstore) Clear() ([]api.Transfer, error) {
	/*accounts, err := db.List()
	if err != nil {
		return []api.Transfer{}, err
	}

	sum := 0

	for _, account := range accounts {
		balance, _ := strconv.Atoi(account.Value)
		sum += balance
	}

	if sum != 0 {
		return []api.Transfer{}, fmt.Errorf("db is incosistent")
	}

	tempAcc := make(map[string]int)

	transfers := []api.Transfer{}

	for sender, balance := range tempAcc {
		if balance >= 0 {
			continue
		}

		for receiver, receiverBalance := range tempAcc {
			if receiverBalance <= 0 {
				continue
			}

			transferAmount := min(-balance, receiverBalance)
			transfers = append(transfers, api.Transfer{Sender: sender, Receiver: receiver, Amount: transferAmount})

			tempAcc[sender] += transferAmount
			tempAcc[receiver] -= transferAmount

			if tempAcc[sender] == 0 {
				break
			}
		}
	}

	return transfers, nil
	*/

	//kvs := make(map[int]api.Account)
	accounts, err := db.List()
	if err != nil {
		return []api.Transfer{}, fmt.Errorf("failed to list accounts: %w", err)
	}

	tempAcc := make(map[string]int)
	//sum := 0
	for _, account := range accounts {
		balance, _ := strconv.Atoi(account.Value)
		tempAcc[account.Key] += balance
	}
	transfers := make([]api.Transfer, 0)
	for sender, balance := range tempAcc {
		if balance >= 0 {
			continue
		}

		for receiver, receiverBalance := range tempAcc {
			if receiverBalance <= 0 {
				continue
			}

			transferAmount := int(math.Min(float64(-balance), float64(receiverBalance)))
			transfers = append(transfers, api.Transfer{Sender: sender, Receiver: receiver, Amount: transferAmount})

			tempAcc[sender] += transferAmount
			tempAcc[receiver] -= transferAmount

			if tempAcc[sender] == 0 {
				break
			}
		}
	}
	/*transfers, err := db.Clear()
	if err != nil {
		return []api.Transfer{}, fmt.Errorf("failed to clear accounts: %w", err)
	}*/
	return transfers, nil

}

// Reset sets all balances to zero.
func (db *kvstore) Reset() error {
	return db.Client.Reset()
}
