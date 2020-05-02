package main

import (
	"fmt"
	"sync"
)

// define the repository struct (as a singleton) with lowercase letters in order to make it private
// This will disallow instaciating the struct outside the package.
type repository struct {
	items map[string]string
	mu    sync.RWMutex
}

func (r *repository) Get(key string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	item, ok := r.items[key]
	if !ok {
		return "", fmt.Errorf("The '%s' is not presented", key)
	}
	return item, nil
}

func (r *repository) Set(key, data string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[key] = data
}


// CLIENT CALL

// In Golang the static member of the Singleton struct is declared as a global 
// variable in the package that contains this type.
var (
	r *repository
)

// Repository provides a global point of access to repository instance
// Instance(): Singleton
func Repository() *repository {
	if r == nil {
		r = &repository{
			items: make(map[string]string),
		}
	}
	return r
}

func main() {
	r := Repository()

	r.Set("test", "dupa")
	item, err := r.Get("test")
	if err != nil {
		fmt.Println(err)
		return 
	}

	fmt.Println(item)
}