package main

type Iterator interface {
	HasNext()
	Next()
}