package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/Hundemeier/go-sacn/sacn"
)

//KeyEvent is an abstract event for storing the events information
//If Value is 1 then this is a DOWN event, 0 is UP and 2 is REPEATED
type KeyEvent struct {
	Code  uint16
	Value int32
}

func main() {
	multicast := flag.Bool("multicast", true,
		"set wether multicast should be used for sending out the sACN packets")
	universe := flag.Uint("universe", 1, "the sACn universe to use")
	verbose := flag.Bool("verbose", false, "enables output of more information while a key was pressed")
	destination := flag.String("destination", "", "Set the unicast destination")

	flag.Parse()

	if *universe > 65535 {
		log.Fatalf("The given universe of %v is too high!", *universe)
	}

	log.Println("Starting Keylogger...")
	in := getKeylogger()

	log.Println("Starting sACN...")
	cid := [16]byte{}
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := range cid {
		//make the CID random:
		cid[i] = byte(rand.Int())
	}
	trans, err := sacn.NewTransmitter("", cid, "key2sACN")
	logErr(err)
	trans.SetMulticast(uint16(*universe), *multicast)
	errs := trans.SetDestinations(uint16(*universe), []string{*destination})
	for _, v := range errs {
		logErr(v)
	}
	sacn, err := trans.Activate(uint16(*universe))
	logErr(err)

	log.Println("Quit with Ctrl+C. Listening for keys...")
	if *verbose {
		log.Println("<keyCode> <state> -> <DMX channel> <DMX value>")
	}
	data := [512]byte{}
	for i := range in {
		//check if we have a key down or up
		if i.Value == 1 {
			//Key down
			data[i.Code] = 255
			sacn <- data
			if *verbose {
				log.Printf("%v DOWN -> %v 100%%", i.Code+1, i.Code+1)
			}
		} else if i.Value == 0 {
			//Key UP
			data[i.Code] = 0
			sacn <- data
			if *verbose {
				log.Printf("%v UP   -> %v 0%%", i.Code+1, i.Code+1)
			}
		}

	}
}

func logErr(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}
