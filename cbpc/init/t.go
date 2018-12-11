package main

import(
	"fmt"
	"sort"
	"math/rand"
)

type Hero struct {
	Name string
	Age int
}

type HeroSlice []Hero

func (hs HeroSlice) Len() int{
	return len(hs)
}

func (hs HeroSlice) Less(i,j int) bool {
	return hs[i].Age>hs[j].Age
}

func (hs HeroSlice) Swap(i,j int) {
	temp:=hs[i]
	hs[i]=hs[j]
	hs[j]=temp
}

func main(){
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero:=Hero{
			Name:fmt.Sprintf("英雄～%d",rand.Intn(100)),
			Age:rand.Intn(100),
		}
		heroes=append(heroes,hero)
	}

	for _, v := range heroes {
		fmt.Println(v)
	}

	sort.Sort(heroes)
fmt.Println("--------------------------")
	for _, v := range heroes {
		fmt.Println(v)
	}
}