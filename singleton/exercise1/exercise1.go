package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

var instance Database

type Database interface {
	getLevel(name string) int
}

var once sync.Once

type singletonDatabase struct {
	heros map[string]int
}

func (s *singletonDatabase) getLevel(name string) int {
	return s.heros[name]
}

func readData(path string) (map[string]int, error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)

	file, err := os.Open(exPath + path)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

func GetSingletonDatabase() Database {
	once.Do(func() {
		db := singletonDatabase{}
		heros, err := readData(".\\heros.txt")
		if err == nil {
			db.heros = heros
		}
		instance = &db
	})
	return instance
}