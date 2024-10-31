package main

import "time"

type Cacher interface {
	Set([]byte, []byte, time.Duration) error
	Get([]byte) ([]byte, error)
	Exist([]byte) bool
	Remove([]byte) error
}
