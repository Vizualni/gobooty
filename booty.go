// Package gobooty provides utilities for creating singleton instances for bootstrapping your applications.
package gobooty

import "sync"

func One[T any](b func() T) func() T {
	once := &sync.Once{}
	var instance T
	return func() T {
		once.Do(func() {
			instance = b()
		})
		return instance
	}
}

func Two[T1 any, T2 any](f func() (T1, T2)) func() (T1, T2) {
	once := &sync.Once{}
	var instance1 T1
	var instance2 T2
	return func() (T1, T2) {
		once.Do(func() {
			instance1, instance2 = f()
		})
		return instance1, instance2
	}
}
