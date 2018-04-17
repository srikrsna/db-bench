package db

import (
	"errors"
	"github.com/bradfitz/gomemcache/memcache"
	"encoding/gob"
	"bytes"
	"sync"
)

func init() {
	gob.Register(&User{})
}

type memcached struct {
	*memcache.Client

	bytePool sync.Pool
}

func NewMemcached() Store {
	settings := getCredentials()
	mc := memcache.New(settings.Memcached.Server)

	return &memcached{
		mc,
		sync.Pool{New: func() interface{} {
			return new(bytes.Buffer)
		},
		},
	}
}

func (mc *memcached) Add(id string, user User) error {

	data := mc.bytePool.Get().(*bytes.Buffer)
	data.Reset()

	enc := gob.NewEncoder(data)
	if err := enc.Encode(&user); err != nil {
		return err
	}

	mc.Set(&memcache.Item{Key: "foo", Value: data.Bytes()})

	mc.bytePool.Put(data)

	return nil
}

func (mc *memcached) Get(id string) (User, error) {
	_, err := mc.Get(id)
	if err != nil {
		panic(err)
	}
	return User{}, errors.New("not found")
}

func (mc *memcached) Update(user User) error {
	return nil
}

func (mc *memcached) Delete(id string) error {
	return nil
}
