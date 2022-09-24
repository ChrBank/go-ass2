package main

import (
    "fmt"
    "strings"
    "math/rand"
    "time"
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
    packetReceived int
}

func main() {
    split := strings.Split(wholeMessage, "\n")
    
    packets := make([]packet, len(split))
    
    fmt.Println("-----------------------------------------------PreShuffle")

    for i := 0; i < len(split); i++ {
        packets[i] = packet{
                            index: i,
                            lifeTime: 25,
                            message: split[i],
                            }
    }
    
    rand.Seed(time.Now().UnixNano())                                    // To make shuffle non consistent
    rand.Shuffle(len(packets), func(i,j int) { 
            packets[i], packets[j] = packets[j], packets[i]})          // Shuffle made to imitate uncertainty when sending packets
    
    for i := 0; i < len(split); i++ {
        fmt.Println(packets[i].message)
    }
    
    comChan := make (chan communication)
    
    go server(comChan)
    go client(comChan, packets)
    
    time.Sleep(5 * time.Second)
}

func server(comChan chan communication) {
    packetChannel := make (chan packet)
    inUse := false
    var response = communication {}
    
    for true {
        m := <- comChan
        if (m.message && !inUse) {                                      // Check request to use me
            inUse = true
            response.message = true
            response.packetChan = packetChannel                         // Put packet channel into response
            comChan <- response                                         // Send response
            m := <- comChan                               
            if (m.message) {                                            // If received necessary info then proceed
                receivedPackets := make ([]packet, m.packetAmount)
                for i := 0; i < m.packetAmount; i++ {
                    p := <- packetChannel
                    receivedPackets[p.index] = p                        // Puts the packet into the correct place
                    fmt.Println("-------------Packet", p.index, "Received")
                    comChan <- communication {packetReceived: p.index}  // Respond to received packet
                }
                
                for i := 0; i < len(receivedPackets); i++ {
                        fmt.Println(receivedPackets[i].message)
                }
                
                inUse = false
            }
        } else {
            response.message = false
        }
        break;
    }
}

func client(comChan chan communication, packets []packet) {
    for true {
    comChan <- communication {message: true}                            // Request to use server
        m := <- comChan
        if (m.message) {                                                // Server free to use
            comChan <- communication {
                    message: true, packetAmount: len(packets)}          // Sends amount of packets
            
            var packetChannel = m.packetChan
            
            for i := 0; i < len(packets); i++ {
                packetChannel <- packets[i]
                fmt.Println("----------Packet", packets[i].index, "Sent")
                
                r := <- comChan
                if (r.packetReceived == 1) {
                    continue
                }
            }
        }
    }
}

// Change request to message and to int to send which packets received

var wholeMessage string = "Load up on guns, bring your friends\nIt's fun to lose and to pretend\nShe's over bored and self assured\nOh no, I know a dirty word\nHello, hello, hello, how low\nHello, hello, hello, how low\nHello, hello, hello, how low\nHello, hello, hello\nWith the lights out, it's less dangerous\nHere we are now, entertain us\nI feel stupid and contagious\nHere we are now, entertain us\nA mulatto, an albino, a mosquito, my libido\nYeah, hey\nYay\nI'm worse at what I do best\nAnd for this gift I feel blessed\nOur little group has always been\nAnd always will until the end\nHello, hello, hello, how low\nHello, hello, hello, how low\nHello, hello, hello, how low\nHello, hello, hello\nWith the lights out, it's less dangerous\nHere we are now, entertain us\nI feel stupid and contagious\nHere we are now, entertain us\nA mulatto, an albino, a mosquito, my libido\nYeah, hey\nYay\nAnd I forget just why I taste\nOh yeah, I guess it makes me smile\nI found it hard, was hard to find\nOh well, whatever, never mind\nHello, hello, hello, how low\nHello, hello, hello, how low\nHello, hello, hello, how low\nHello, hello, hello\nWith the lights out, it's less dangerous\nHere we are now, entertain us\nI feel stupid and contagious\nHere we are now, entertain us\nA mulatto, an albino, a mosquito, my libido\nA denial, a denial, a denial, a denial, a denial\nA denial, a denial, a denial, a denial"