package main

import (
	"fmt"
	"math/rand"
	"time"
)

type array []string
func main() {
	alphabets :=array{"a","b","c","d","e"}
	fmt.Println(alphabets)	
	fmt.Println(alphabets.shuffle())




}

func (a array)shuffle( )array{
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	r.Shuffle(len(a), func( i, j int){
	a[i], a[j] = a[j], a[i]
	})

	return a

}
