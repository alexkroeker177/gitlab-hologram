package main

import (
	asapp "blank-sparkable/app"
	"context"
	"github.com/Bitspark/go-bitnode/api/wsApi"
	"github.com/Bitspark/go-bitnode/bitnode"
	"github.com/Bitspark/go-bitnode/library"
	"github.com/Bitspark/go-bitnode/store"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	localAddress := os.Getenv("BITNODE_LOCAL_ADDRESS")
	remoteNodeAddress := os.Getenv("BITNODE_REMOTE_ADDRESS")

	node := bitnode.NewNode()
	node.AddMiddlewares(library.GetMiddlewares())

	dom := bitnode.NewDomain()
	dom, _ = dom.AddDomain("hub")

	// Prepare node connections.
	nodeConns := wsApi.NewNodeConns(node, remoteNodeAddress)

	// Prepare node.
	localDom, err := dom.AddDomain("asapp")
	if err != nil {
		log.Fatal(err)
	}
	if err := localDom.LoadFromFile("./asapp/domain.yml"); err != nil {
		log.Fatal(err)
	}
	if err := localDom.Compile(); err != nil {
		log.Fatal(err)
	}

	asappDom := &asapp.Domain{
		Domain: dom,
		Node:   node,
	}

	// Read store.
	st1 := store.NewStore("store")
	if err := st1.Read("."); err != nil {
		log.Println(err)
	} else {
		// Load node.
		if err := nodeConns.Load(st1, dom); err != nil {
			log.Fatalf("Error loading node: %v", err)
		} else {
			log.Printf("Loaded node from %s", ".")
		}
	}

	creds := bitnode.Credentials{}

	var acc *asapp.BlankSparkable

	if len(node.Systems(creds)) == 0 {
		var err error
		acc, err = asappDom.NewPersonioAccount()
		if err != nil {
			log.Fatal(err)
		}

		// Make computer system the root system.
		node.SetSystem(acc.Native())
	} else {
		log.Printf("Found %d startup systems", len(node.Systems(creds)))

		// Get the system from the node.
		accSys := node.System(creds)

		acc = &asapp.BlankSparkable{
			System: accSys,
		}
	}

	// Add the custom BlankSparkable implementation.
	if err := acc.Init(); err != nil {
		log.Fatal(err)
	}

	// Create server.
	server := wsApi.NewServer(nodeConns, localAddress)

	stored := make(chan error)

	go func() {
		log.Println(server.Listen())

		// Create store.
		st := store.NewStore("store")

		// Store node.
		if err := nodeConns.Store(st); err != nil {
			stored <- err
			return
		}

		// Write node store.
		if err := st.Write("."); err != nil {
			log.Println(err)
			stored <- err
			return
		}

		stored <- nil
	}()

	log.Printf("Listening on %s...", server.Address())

	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)
	<-cancelChan

	log.Println("Stopping...")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}

	if err := <-stored; err != nil {
		log.Printf("Error storing node: %v", err)
	}

	time.Sleep(1 * time.Second)
}