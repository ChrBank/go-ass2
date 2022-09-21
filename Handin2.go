package main

import (
    "fmt"
    "strings"
)

type packet struct {
    index int
    lifeTime int
    message string
}

func main() {
    split := strings.Split(wholeMessage, "\n")
    
    packets := [45] packet {}

    for i := 0; i < len(split); i++ {
        packets[i] = packet{
                            index: i,
                            lifeTime: 25,
                            message: split[i],
                            }
    }
    
    for i := 0; i < len(split); i++ {
        fmt.Println(packets[i].message)
    }
}

var wholeMessage string =   "Load up on guns, bring your friends\nIt's fun to lose and to pretend\nShe's over bored and self assured\nOh no, I know a dirty word\nHello, hello, hello, how low\nHello, hello, hello, how low\nHello, hello, hello, how low\nHello, hello, hello\nWith the lights out, it's less dangerous\nHere we are now, entertain us\nI feel stupid and contagious\nHere we are now, entertain us\nA mulatto, an albino, a mosquito, my libido\nYeah, hey\nYay\nI'm worse at what I do best\nAnd for this gift I feel blessed\nOur little group has always been\nAnd always will until the end\nHello, hello, hello, how low\nHello, hello, hello, how low\nHello, hello, hello, how low\nHello, hello, hello\nWith the lights out, it's less dangerous\nHere we are now, entertain us\nI feel stupid and contagious\nHere we are now, entertain us\nA mulatto, an albino, a mosquito, my libido\nYeah, hey\nYay\nAnd I forget just why I taste\nOh yeah, I guess it makes me smile\nI found it hard, was hard to find\nOh well, whatever, never mind\nHello, hello, hello, how low\nHello, hello, hello, how low\nHello, hello, hello, how low\nHello, hello, hello\nWith the lights out, it's less dangerous\nHere we are now, entertain us\nI feel stupid and contagious\nHere we are now, entertain us\nA mulatto, an albino, a mosquito, my libido\nA denial, a denial, a denial, a denial, a denial\nA denial, a denial, a denial, a denial"