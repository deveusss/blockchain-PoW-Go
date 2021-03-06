package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"

	"Deveuspow/Deveuscore/block"
	"Deveuspow/Deveuscore/blockchain"
	"Deveuspow/Deveuscore/consensus"
	"Deveuspow/Deveuscore/explorer"
	"Deveuspow/Deveuscore/mempool"
	"Deveuspow/Deveuscore/miner"
	"Deveuspow/Deveuscore/network"
	"Deveuspow/Deveuscore/network/cli"
	"Deveuspow/Deveuscore/network/message"
	"Deveuspow/Deveuscore/transaction"
	"Deveuspow/Deveuscore/wallet"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	cli_app "github.com/urfave/cli"
	"google.golang.org/grpc"
)

const targetDifficulty = "000000FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"

var (
	addNode      string
	listen       string
	httpListen   string
	minerAddress string
	runMining    bool
)

func main() {
	app := cli_app.NewApp()
	app.Name = "\nDeveusPoW"
	app.Version = "0.1.1"
	app.Usage = "Run DeveusPoW as a daemon"
	app.Flags = []cli_app.Flag{
		cli_app.StringFlag{
			Name:        "add-node",
			Usage:       "this flag used to connect to an another node",
			Destination: &addNode,
		},
		cli_app.StringFlag{
			Name:        "listen",
			Usage:       "this flag starts the node on given listen address",
			Value:       "0.0.0.0:10000",
			Destination: &listen,
		},
		cli_app.StringFlag{
			Name:        "http-listen",
			Usage:       "this flag starts a block explorer on given listen address",
			Destination: &httpListen,
		},
		cli_app.StringFlag{
			Name:        "miner-address",
			Usage:       "an address for mining reward ( for mining mode only )",
			Value:       "12PLGoQb9usohESQAeB6rrYXPcua9M36Pc",
			Destination: &minerAddress,
		},
		cli_app.BoolFlag{
			Name:        "run-mining",
			Usage:       "this flag enables mining mode",
			Destination: &runMining,
		},
	}
	app.Action = func(c *cli_app.Context) error {
		run()
		return nil
	}
	app.Run(os.Args)
}

func run() {
	l, err := net.Listen("tcp", listen)
	if err != nil {
		log.Fatalf("[network]: failed to listen: %v", err)
	}

	ctx := context.Background()
	storage, _ := blockchain.NewMemoryStorage()
	memPool := mempool.New()

	targetBytes, _ := hex.DecodeString(targetDifficulty)
	targetInt := new(big.Int).SetBytes(targetBytes)

	solver := block.NewDSHA256Solver(targetInt)
	signer := transaction.NewSimpleSigner(storage)
	verifier := transaction.NewSimpleVerifier(storage)

	validator := consensus.NewBaseValidator(storage, memPool, solver, verifier)
	walletStorage := wallet.NewFileStorage("")

	w, _ := walletStorage.Init(ctx)
	walletStorage.Write(ctx, w)

	service := &network.Node{
		Validator:   validator,
		MemPool:     memPool,
		Storage:     storage,
		PeerManager: network.NewPeerManager(),
		Wallet:      w,
		Signer:      signer,
	}

	if addNode != "" {
		go connectToNode(ctx, service)
	}

	if httpListen != "" {
		go runBlockExplorer(storage, memPool)
	}

	if runMining {
		toAddr, err := btcutil.DecodeAddress(minerAddress, &chaincfg.MainNetParams)
		if err != nil {
			log.Fatal("[miner]: failed to parse miner address")
			//log.Fatal(err.Error)
		}

		minerAddrPKH, ok := toAddr.(*btcutil.AddressPubKeyHash)
		if !ok {
			log.Fatal("[miner]: miner address is not PKH address")
		}

		m := miner.New(storage, memPool, solver, *minerAddrPKH)
		foundBlock := make(chan block.Block)
		defer close(foundBlock)

		go m.Run(ctx, foundBlock)
		go func() {
			for {
				b := <-foundBlock
				blockData, err := json.Marshal(b)
				if err != nil {
					log.Println("[miner]: failed to serialize block")
				} else {
					service.PeerManager.Send(ctx, network.NewAddBlock(blockData))
				}
			}
		}()
	}

	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	message.RegisterMessageServiceServer(grpcServer, service)
	cli.RegisterCliServiceServer(grpcServer, service)
	log.Printf("[network]: node started at %q\n", listen)
	grpcServer.Serve(l)
}

func connectToNode(ctx context.Context, s *network.Node) {
	clientConn, err := grpc.Dial(addNode, grpc.WithInsecure())
	if err != nil {
		log.Printf("[network]: failed to create a client for %q\n", addNode)
		os.Exit(1)
	}
	client := message.NewMessageServiceClient(clientConn)
	stream, err := client.Message(ctx)
	if err != nil {
		log.Printf("[network]: failed to connect to %q\n", addNode)
		os.Exit(1)
	}
	log.Printf("[network]: connected to a node %q\n", addNode)
	s.ServeClient(stream)
}

func runBlockExplorer(s blockchain.Storage, m *mempool.MemPool) {
	be, err := explorer.New(s, m)
	if err != nil {
		log.Printf("[network]: failed to start blockexplorer: %v\n", err)
		os.Exit(1)
	}

	http.HandleFunc("/", be.ViewIndex)
	http.HandleFunc("/tx/", be.ViewTX)
	http.HandleFunc("/block/", be.ViewBlock)
	log.Printf("[network]: blockexplorer started at %q\n", httpListen)
	http.ListenAndServe(httpListen, nil)
}
