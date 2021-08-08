package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)


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
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    // if err != nil {
    //         log.Fatal(err)
    // }
    // fmt.Println(dir)
	filePath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(filePath)  // for example /home/user

	ex, err := os.Executable()
	fmt.Println(ex)
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(filePath)
	fmt.Println(exPath)
	fmt.Println(filePath + path)
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
		fmt.Println(k)
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		fmt.Println(v)
		result[k] = v
	}
	return result, nil
}

// ever singleton will use sync Once
var instance Database

func getSingletonDatabase() Database {
	once.Do(func() {
		db := singletonDatabase{}
		heros, err := readData("\\heros.txt")
		if err == nil {
			db.heros = heros
		}
		instance = &db
	})
	return instance
}

func main() {
	db := getSingletonDatabase()
	hero1Level := db.getLevel("Sheena")
	fmt.Println("Sheena level is", hero1Level)
}