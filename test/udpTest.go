package main

import (
	"crypto/sha1"
	"io"
	"log"
	"time"

	"github.com/xtaci/kcp-go/v5"
	"golang.org/x/crypto/pbkdf2"
)

func main() {
	client("127.0.0.1:9000")
}

func client(addrees string) {
	key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	block, _ := kcp.NewAESBlockCrypt(key)

	// wait for server to become ready
	time.Sleep(time.Second)

	// dial to the echo server
	if sess, err := kcp.DialWithOptions(addrees, block, 10, 3); err == nil {
		for {
			data := time.Now().String()
			buf := make([]byte, len(data))
			log.Println("client sent:", data)
			if _, err := sess.Write([]byte(data)); err == nil {
				// read back the data
				if _, err := io.ReadFull(sess, buf); err == nil {
					log.Println("client recv:", string(buf))
				} else {
					log.Fatal(err)
				}
			} else {
				log.Fatal(err)
			}
			time.Sleep(time.Second)
		}
	} else {
		log.Fatal(err)
	}
}
