package main

//pass send only channel to the function so that it acts as just a producer
//unidirectional channel with only send capability
func sendData(channelInstance chan <- int) {
	for i:=0; i<100 ; i++ {
	channelInstance<-i + 10
	}

	//close the channel (sender one)
	close(channelInstance)
}

//pass receive only channel to the function so that it acts as just a receiver
//unidirectional channel with only receive capability
//consume the data unless channel is empty
func receiveData(channelInstance <- chan int) {
	valueToCheck1,ok1:= <-channelInstance
	if ok1 {
		println("The data on the channel is : ",valueToCheck1)
		println("Whether data consumed? ",ok1)
	}

	for value:= range channelInstance {
		println("The value consumed from the channel is : ", value)
	}

	valueToCheck2,ok2:= <-channelInstance
	println("The data on the channel is : ",valueToCheck2)
	println("Whether data consumed? ",ok2)
}

func main() {
	channelInstance:=make(chan int)
	go sendData(channelInstance)
	receiveData(channelInstance)
}
