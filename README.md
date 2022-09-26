# go-ass2
qA - Packets is implemented as structs with three fields: index, lifetime and message. Index is used to know the sequence number of the packet. Lifetime is meant to simulate a packets lifetime, even though is has not been implemented. Message is used to hold the packets data.
    We use channels to send packets, requests/answers and metadata. To send packets there has been used a channel which sends packet structs. The requests/answers are sent through a channel which sends communication structs. Communication structs can hold info about the request/answer and metadata.

qB - Our implementation uses thread to simulate the client and server interaction. 
    