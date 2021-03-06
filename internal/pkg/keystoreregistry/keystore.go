package keystoreregistry

import (
	"context"
	"encoding/json"
	"log"
)

type Keystore interface {
	// Used to clean up any resources being used
	// ... such as a connection to etcd.
	Close() error

	// Removes any key starting with the given prefix.
	// An error is returned if nothing was deleted,
	// which some users may choose to safely ignore.
	CleanPrefix(prefix string) error

	// Atomically add all the key value pairs
	//
	// If an error occurs no keyvalues are written.
	// Error is returned if any key already exists.
	Add(keyValues []KeyValue) error

	// Update the specifed key values, atomically
	//
	// If ModRevision is 0, it is ignored.
	// Otherwise if the revisions of any key doesn't
	// match the current revision of that key, the update fails.
	// When update fails an error is returned and no keyValues are updated
	Update(keyValues []KeyValueVersion) error

	// Delete the specifed key values, atomically
	//
	// Similar to update, checks ModRevision matches current key,
	// ignores ModRevision if not zero.
	// If any keys are not currently present, the request fails.
	// Deletes no keys if an error is returned
	DeleteAll(keyValues []KeyValueVersion) error

	// Get all key values for a given prefix.
	GetAll(prefix string) ([]KeyValueVersion, error)

	// Get all keys for a given prefix.
	Get(key string) (KeyValueVersion, error)

	// Get callback on all changes related to the given prefix.
	//
	// When a key is created for the first time, old is an empty value,
	// and new.CreateRevision == new.ModRevision
	// This starts watching from the current version, rather than replaying old events
	// Returns the revision that the watch is starting on
	WatchPrefix(prefix string, onUpdate func(old *KeyValueVersion, new *KeyValueVersion))

	// Watch given key
	//
	// When callback returns true, stop watch the key
	WatchKey(ctxt context.Context, key string, onUpdate func(old *KeyValueVersion, new *KeyValueVersion))

	// Add a key, and remove it when calling process dies
	// Error is returned if the key already exists
	KeepAliveKey(key string) error
}

type KeyValue struct {
	Key   string
	Value string // TODO: should this be []byte? Or have a json parsed version?
}

type KeyValueVersion struct {
	Key            string
	Value          string
	CreateRevision int64
	ModRevision    int64
}

func (kvv KeyValueVersion) String() string {
	return toJson(kvv)
}

func toJson(message interface{}) string {
	b, error := json.Marshal(message)
	if error != nil {
		log.Fatal(error)
	}
	return string(b)
}
