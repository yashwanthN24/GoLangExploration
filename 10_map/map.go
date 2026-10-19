package main

import (
	"fmt"
	"maps"
)

// maps are hashtable/ map in c++ , object in js , dict in python
func main(){
	// creating a map 

	// map[keytype]valuetype 

	// so here creating a map of key as string and value as string
	m := make(map[string]string)

	// Setting an element 
	m["name"] = "golang"

	m["area"] = "backend"
	// get an element 
	fmt.Println(m["name"] , m["area"])

	// IMP : if key does not exists in the map then it returns zero value (i.e for string "" , for int 0 and for bool false )
	fmt.Println(m["phone"])


	m2 := make(map[string]int) // creating a map with key as string and value as int 
	m2["age"] = 30
	fmt.Println(m2["yash"]) /// gets zero value for int its 0 

	fmt.Println(len(m2)) // gives the count number of entries in map 

	// deleting key value pair from map 
	fmt.Println(m2)
	delete(m2 , "age")

	fmt.Println(m)
	fmt.Println(m2)

	// to clear the entire map 
	fmt.Println(m)
	clear(m)

	fmt.Println(m)

	// other way of creating a map 
	newMap := map[string]int{"price": 40 , "phones": 3}
	fmt.Println(newMap)

	// use this when you know n advance all key value pairs

	// exists check
    v , ok := newMap["phones"]

	fmt.Println(v)
	if ok {
		fmt.Println("all ok")
	}else{
		fmt.Println("not ok")
	}


	map1 := map[string]int{"prices": 40 , "phones" : 3}
	// map2 := map[string]int{"prices": 40 , "phones" : 3}
	map2 := map[string]int{"prices": 40 , "phones" : 4}

	fmt.Println(maps.Equal(map1 , map2))

}
