package in_memory

import (
	"github.com/bloxapp/KeyVault/core"
	uuid "github.com/google/uuid"
	types "github.com/wealdtech/go-eth2-wallet-types/v2"
)

type InMemStore struct {
	memory         		map[string]interface{}
	attMemory      		map[string]*core.BeaconAttestation
	proposalMemory 		map[string]*core.BeaconBlockHeader
	encryptor	   		types.Encryptor
	encryptionPassword 	[]byte
}

func NewInMemStore(
	) *InMemStore {
	return NewInMemStoreWithEncryptor(nil,nil)
}

func NewInMemStoreWithEncryptor(
	encryptor types.Encryptor,
	password []byte,
	) *InMemStore {
	return &InMemStore{
		memory:         	make(map[string]interface{}),
		attMemory:      	make(map[string]*core.BeaconAttestation),
		proposalMemory: 	make(map[string]*core.BeaconBlockHeader),
		encryptor:			encryptor,
		encryptionPassword:	password,
	}
}

// Name provides the name of the store
func (store *InMemStore) Name() string {
	return "in-memory"
}

func (store *InMemStore) SaveWallet(wallet core.Wallet) error {
	store.memory["wallet"] = wallet
	return nil
}

// will return nil,nil if no wallet was found
func (store *InMemStore) OpenWallet() (core.Wallet,error) {
	if val := store.memory["wallet"]; val != nil {
		ret := val.(core.Wallet)
		ret.SetContext(store.freshContext())
		return ret,nil
	} else {
		return nil,nil
	}
}

// will return an empty array for no accounts
func (store *InMemStore) ListAccounts() ([]core.ValidatorAccount,error) {
	w,err := store.OpenWallet()
	if err != nil {
		return nil,err
	}

	ret := make([]core.ValidatorAccount,0)
	for a := range w.Accounts() {
		ret = append(ret,a)
	}
	return ret,nil
}

func (store *InMemStore) SaveAccount(account core.ValidatorAccount) error {
	store.memory[account.ID().String()] = account
	return nil
}

// will return nil,nil if no account was found
func (store *InMemStore) OpenAccount(accountId uuid.UUID) (core.ValidatorAccount,error) {
	if val := store.memory[accountId.String()]; val != nil {
		return val.(core.ValidatorAccount),nil
	} else {
		return nil,nil
	}
}

func (store *InMemStore) SetEncryptor(encryptor types.Encryptor, password []byte) {
	store.encryptor = encryptor
	store.encryptionPassword = password
}

func (store *InMemStore) freshContext() *core.WalletContext {
	return &core.WalletContext{
		Storage:     store,
	}
}

func (store *InMemStore) canEncrypt() bool {
	if store.encryptor != nil {
		if store.encryptionPassword == nil {
			return false
		}
		return true
	}
	return false
}