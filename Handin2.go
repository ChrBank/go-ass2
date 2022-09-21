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

type communication struct {
    message bool
    packetChan chan packet
    packetAmount int
}

func main() {
    split := strings.Split(wholeMessage, "\n")
    
    packets := make([]packet, len(split))

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
    
    comChan := make (chan communication)
    
    go server(comChan)
    go client(comChan, packets)
}

func server(comChan chan communication) {
    inUse := false
    var response = communication {}
    for true {
        m := <- comChan
        if (m.request && !inUse) {                              // Check request to use me
            inUse = true
            response.request = true
            comChan <- response                                 // Give myself to client
            m := <- communication                               
            if (m.request) {                                    // If received necessary info then proceed
                packetChannel = m.packetChan
                receivedPackets := make ([]packet, m.packetAmount)
                for i := 0; i < m.packetAmount; i++ {
                    p := <- packetChannel
                    comChan <- communication {}
                }
            }
        } else {
            response.request = false
        }
    }
}

func client(comChan chan communication, packets packets[]) {
    packetChannel := make (chan packet)
    var response = communication {}
    for true {
        communication <- communication {request: true}          // Request to use server
        m := <- communication
        if (m.request) {                                        // Server free to use
            response.packetChan = packetChannel
            response.packetAmount = len(packets)
            communication <- response                           // Sends needed packet channel and amount of packets
        }
    }
}

// Change request to message and to int to send which packets received

var wholeMessage string =   "Load up on guns, bring your friends\nIt's fun to lose and to pretend\nShe's over bored and self assured\nOh no, I know a dirty word\nHello, hello, hello, how low\nHello, hello, hello, how low\nHello, hello, hello, how low\nHello, hello, hello\nWith the lights out, it's less dangerous\nHere we are now, entertain us\nI feel stupid and contagious\nHere we are now, entertain us\nA mulatto, an albino, a mosquito, my libido\nYeah, hey\nYay\nI'm worse at what I do best\nAnd for this gift I feel blessed\nOur little group has always been\nAnd always will until the end\nHello, hello, hello, how low\nHello, hello, hello, how low\nHello, hello, hello, how low\nHello, hello, hello\nWith the lights out, it's less dangerous\nHere we are now, entertain us\nI feel stupid and contagious\nHere we are now, entertain us\nA mulatto, an albino, a mosquito, my libido\nYeah, hey\nYay\nAnd I forget just why I taste\nOh yeah, I guess it makes me smile\nI found it hard, was hard to find\nOh well, whatever, never mind\nHello, hello, hello, how low\nHello, hello, hello, how low\nHello, hello, hello, how low\nHello, hello, hello\nWith the lights out, it's less dangerous\nHere we are now, entertain us\nI feel stupid and contagious\nHere we are now, entertain us\nA mulatto, an albino, a mosquito, my libido\nA denial, a denial, a denial, a denial, a denial\nA denial, a denial, a denial, a denial"