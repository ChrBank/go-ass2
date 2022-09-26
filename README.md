# go-ass2
To run this program, use the command $go run Handin2.go

qA - Packets is implemented as structs with three fields: index, lifetime and message. Index is used to know the sequence number of the packet. Lifetime is meant to simulate a packets lifetime, even though is has not been implemented. Message is used to hold the packets data.
We use channels to send packets, requests/answers and metadata. To send packets there has been used a channel which sends packet structs. The requests/answers are sent through a channel which sends communication structs. Communication structs can hold info about the request/answer and metadata.

qB - Our implementation uses thread to simulate the client and server interaction.
It is not realistic to use threads since they are running on the same local PC.

qC - The server looks at the sequence number of the packet and inserts it correctly into an array.

qD - We do not handle message loss. To be able to handle message loss their would need to be an array/set of the packets that have been recieved from the server. This would need to be done from another goroutine. The reasoning is that if the user tries to listen to which packet has been recieved there can occur a deadlock since the server will not send a response.

qE - The 3-way handshake has two main functions. To make sure that both sides know that they are ready to transfer data and to make sure that both sides agree on the initial sequence numbers. The sequence number field is used to keep track of packets successfully delivered. In order for the connection to work, both sides needs to verify that they can send to the other side. If, for example, it was only a 2-way handshake, then only one party would be able to send data reliably. The 3-way handshake allows both parties to send data reliably.