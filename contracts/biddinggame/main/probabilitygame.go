package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"math/rand"
	"strings"
	"time"

	ethereum "github.com/energicryptocurrency/energi"
	"github.com/energicryptocurrency/energi/accounts/abi"
	"github.com/energicryptocurrency/energi/accounts/abi/bind"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/contracts/biddinggame/contract"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/crypto"
	"github.com/energicryptocurrency/energi/ethclient"
)

const (
	N_Bids       = 10
	MAX_DURATION = 15
)

type Game struct {
	Duration time.Duration
}

type LogBid struct {
	Bidder            common.Address
	Bid               *big.Int
	HighestBidder     common.Address
	HighestBid        *big.Int
	HighestBindingBid *big.Int
}

type LogWithdrawal struct {
	Withdrawer        common.Address
	WithdrawalAccount common.Address
}

//Game run code
func (g *Game) Run(instance *contract.Bid, client *ethclient.Client, owneraddress common.Address, fromaddress1 common.Address, fromaddress2 common.Address, fromaddress3 common.Address, done <-chan interface{}, in <-chan int64) {
	go func() {
		counter := 0
		for {
			select {
			case <-done:

				privateKey1, err := crypto.HexToECDSA("7920ca01d3d1ac463dfd55b5ddfdcbb64ae31830f31be045ce2d51a305516a37")
				if err != nil {
					log.Fatal(err)
				}

				publicKey1 := privateKey1.Public()
				publicKeyECDSA1, ok := publicKey1.(*ecdsa.PublicKey)
				if !ok {
					log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
				}

				fromAddress1 := crypto.PubkeyToAddress(*publicKeyECDSA1)
				nonce1, err := client.PendingNonceAt(context.Background(), fromAddress1)
				if err != nil {
					log.Fatal(err)
				}

				gasPrice1, err := client.SuggestGasPrice(context.Background())
				if err != nil {
					log.Fatal(err)
				}

				auth1 := bind.NewKeyedTransactor(privateKey1)
				auth1.Nonce = big.NewInt(int64(nonce1))
				auth1.Value = big.NewInt(10)       // in wei
				auth1.GasLimit = uint64(300000000) // in units
				auth1.GasPrice = gasPrice1
				log.Println("Bidding is over")
				instance.CancelBid(auth1)
				value1, err := instance.FundsByBidder(&bind.CallOpts{}, fromaddress1)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(value1)
				value2, err := instance.FundsByBidder(&bind.CallOpts{}, fromaddress2)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(value2)
				instance.Withdraw(auth1)
				return
				//Game timer has expired
			case <-time.After(15 * time.Second):
				log.Println("bidding round has completed")

				privateKey1, err := crypto.HexToECDSA("7920ca01d3d1ac463dfd55b5ddfdcbb64ae31830f31be045ce2d51a305516a37")
				if err != nil {
					log.Fatal(err)
				}

				publicKey1 := privateKey1.Public()
				publicKeyECDSA1, ok := publicKey1.(*ecdsa.PublicKey)
				if !ok {
					log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
				}

				fromAddress1 := crypto.PubkeyToAddress(*publicKeyECDSA1)
				nonce1, err := client.PendingNonceAt(context.Background(), fromAddress1)
				if err != nil {
					log.Fatal(err)
				}

				gasPrice1, err := client.SuggestGasPrice(context.Background())
				if err != nil {
					log.Fatal(err)
				}

				auth1 := bind.NewKeyedTransactor(privateKey1)
				auth1.Nonce = big.NewInt(int64(nonce1))
				auth1.Value = big.NewInt(10)       // Correct nonce for transaction sequence
				auth1.GasLimit = uint64(300000000) // Gas price computation
				auth1.GasPrice = gasPrice1

				log.Println("Bidding is over")
				//Bidding game is canceled
				instance.CancelBid(auth1)
				//Display funds bidded by each bidder
				value1, err := instance.FundsByBidder(&bind.CallOpts{}, fromaddress1)

				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(value1)
				value2, err := instance.FundsByBidder(&bind.CallOpts{}, fromaddress2)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(value2)
				//Withdraw game commission by game owner
				instance.Withdraw(auth1)
				//Reward distribution at the end of game
				//each player withdraws, only highest bidder gets cumulative bids amount at the end of game timeout
				privateKey2, err := crypto.HexToECDSA("bb63b692f9d8f21f0b978b596dc2b8611899f053d68aec6c1c20d1df4f5b6ee2")
				if err != nil {
					log.Fatal(err)
				}

				publicKey2 := privateKey2.Public()
				publicKeyECDSA2, ok := publicKey2.(*ecdsa.PublicKey)
				if !ok {
					log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
				}

				fromAddress2 := crypto.PubkeyToAddress(*publicKeyECDSA2)
				nonce2, err := client.PendingNonceAt(context.Background(), fromAddress2)
				if err != nil {
					log.Fatal(err)
				}

				gasPrice2, err := client.SuggestGasPrice(context.Background())
				if err != nil {
					log.Fatal(err)
				}

				auth2 := bind.NewKeyedTransactor(privateKey2)
				auth2.Nonce = big.NewInt(int64(nonce2))
				auth2.Value = big.NewInt(10)
				auth2.GasLimit = uint64(300000000)
				auth2.GasPrice = gasPrice2
				log.Println("Player 1 tries to withdraw")
				instance.Withdraw(auth2)

				privateKey3, err := crypto.HexToECDSA("2f615ea53711e0d91390e97cdd5ce97357e345e441aa95d255094164f44c8652")
				if err != nil {
					log.Fatal(err)
				}

				publicKey3 := privateKey3.Public()
				publicKeyECDSA3, ok := publicKey3.(*ecdsa.PublicKey)
				if !ok {
					log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
				}

				fromAddress3 := crypto.PubkeyToAddress(*publicKeyECDSA3)
				nonce3, err := client.PendingNonceAt(context.Background(), fromAddress3)
				if err != nil {
					log.Fatal(err)
				}

				gasPrice3, err := client.SuggestGasPrice(context.Background())
				if err != nil {
					log.Fatal(err)
				}

				auth3 := bind.NewKeyedTransactor(privateKey3)
				auth3.Nonce = big.NewInt(int64(nonce3))
				auth3.Value = big.NewInt(10)

				auth3.GasLimit = uint64(300000000) // in units
				auth3.GasPrice = gasPrice3

				log.Println("Player 2 tries to withdraw")
				instance.Withdraw(auth3)
				return
			case v, ok := <-in:
				//bid generator channel is closed
				if !ok {
					log.Println("bidding has closed")
					return
				}

				value2, err := instance.GetHighestBid(&bind.CallOpts{})

				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(value2)

				if counter%2 == 0 {

					privateKey1, err := crypto.HexToECDSA("bb63b692f9d8f21f0b978b596dc2b8611899f053d68aec6c1c20d1df4f5b6ee2")
					if err != nil {
						log.Fatal(err)
					}

					publicKey1 := privateKey1.Public()
					publicKeyECDSA1, ok := publicKey1.(*ecdsa.PublicKey)
					if !ok {
						log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
					}

					fromAddress1 := crypto.PubkeyToAddress(*publicKeyECDSA1)
					nonce1, err := client.PendingNonceAt(context.Background(), fromAddress1)
					if err != nil {
						log.Fatal(err)
					}

					gasPrice1, err := client.SuggestGasPrice(context.Background())
					if err != nil {
						log.Fatal(err)
					}

					auth1 := bind.NewKeyedTransactor(privateKey1)
					auth1.Nonce = big.NewInt(int64(nonce1))

					auth1.Value = big.NewInt(v) //Derive bid from bid generator

					auth1.GasLimit = uint64(300000000) // in units
					auth1.GasPrice = gasPrice1

					log.Println("Placed bid 1")
					instance.PlaceBid(auth1)
					value1, err := instance.FundsByBidder(&bind.CallOpts{}, fromaddress1)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(value1)

				} else {

					privateKey1, err := crypto.HexToECDSA("2f615ea53711e0d91390e97cdd5ce97357e345e441aa95d255094164f44c8652")
					if err != nil {
						log.Fatal(err)
					}

					publicKey1 := privateKey1.Public()
					publicKeyECDSA1, ok := publicKey1.(*ecdsa.PublicKey)
					if !ok {
						log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
					}

					fromAddress1 := crypto.PubkeyToAddress(*publicKeyECDSA1)
					nonce1, err := client.PendingNonceAt(context.Background(), fromAddress1)
					if err != nil {
						log.Fatal(err)
					}

					gasPrice1, err := client.SuggestGasPrice(context.Background())
					if err != nil {
						log.Fatal(err)
					}

					auth1 := bind.NewKeyedTransactor(privateKey1)
					auth1.Nonce = big.NewInt(int64(nonce1))

					auth1.Value = big.NewInt(v) //Derive bid form bid generator

					auth1.GasLimit = uint64(300000000) // in units
					auth1.GasPrice = gasPrice1

					log.Println("Placed bid 2")
					instance.PlaceBid(auth1)
					value1, err := instance.FundsByBidder(&bind.CallOpts{}, fromaddress2)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(value1)
				}
				counter++
			}
		}
	}()
}

func eventEmission(contractAddress common.Address) {

	client, err := ethclient.Dial("ws://localhost:8545/ws")
	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	//logBidSig := []byte("LogBid(address,uint,address,uint,uint)")
//	logWithdrawSig := []byte("LogWithdrawal(address,address,uint)")
//	logBidSigHash := crypto.Keccak256Hash(logBidSig)
//	logWithdrawSigHash := crypto.Keccak256Hash(logWithdrawSig)
//	fmt.Printf(logBidSigHash.Hex())
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(contract.BidABI)))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
			fmt.Printf("Log Index: %d\n", vLog.Index)
			fmt.Printf(vLog.Topics[0].Hex())

			if(vLog.Topics[0].Hex() == "0xf152f4ff5e488c55370a2d53925a55055228ebd8ec95bd0251bbb299e48786b0"){

				var bidEvent LogBid

				err := contractAbi.Unpack(&bidEvent, "LogBid", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
		        fmt.Printf("\n\n\n******LogBid Event Emitted******\n\n")
				fmt.Printf("Bidder: %s\n", bidEvent.Bidder.Hex())
				fmt.Printf("Bid: %s\n", bidEvent.Bid.String())
				fmt.Printf("HighestBidder: %s\n", bidEvent.HighestBidder.Hex())
				fmt.Printf("Highestbid: %s\n", bidEvent.HighestBid.String())
 				fmt.Printf("HighestBindingBid: %s\n", bidEvent.HighestBindingBid.String())
				fmt.Printf("Bidder: %s\n", bidEvent.Bidder.Hex())
				fmt.Printf("\n**************\n")

			} else if (vLog.Topics[0].Hex() ==  "0x0ec497a8ae5b1ba29c60470ef651def995fac3deebbdcc56c47a4e5f51a4c2bd") {

				fmt.Printf("Log Name: Withdrawal\n")

				var withdrawEvent LogWithdrawal

				err := contractAbi.Unpack(&withdrawEvent, "LogWithdrawal", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Printf("Withdrawer: %s\n", withdrawEvent.Withdrawer.Hex())
				fmt.Printf("Withdrawal Account: %s\n", withdrawEvent.WithdrawalAccount.Hex())

			} else {

			}

		}
	}
}

func NewGame(duration int) *Game {
	return &Game{
		Duration: time.Duration(duration) * time.Second,
	}
}

//Bidding game simulation

func main() {

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	//Bid generator pattern
	//Bid is placed from array of bids
	//Player 1 and player 2 are picking bids from generator channel and placing the bids
	bidgenerator := func(done chan interface{}, limit int) <-chan int64 {
		outStream := make(chan int64)
		bidarray := []int{50, 60, 70, 80, 90, 100, 110, 120, 130, 140}
		go func() {
			for i := 0; i < limit; i++ {
				randomIndex := rand.Intn(len(bidarray))
				pick := bidarray[randomIndex]
				time.Sleep(1 * time.Second)
				select {
				case <-done:
					return
				case outStream <- int64(pick):
					fmt.Println("Bid picked", pick)
				}

			}

		}()

		return outStream

	}

	privateKey, err := crypto.HexToECDSA("7920ca01d3d1ac463dfd55b5ddfdcbb64ae31830f31be045ce2d51a305516a37")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce)) //Correct Nonce for each transaction
	auth.Value = big.NewInt(10)
	auth.GasLimit = uint64(300000000)
	auth.GasPrice = gasPrice
	fmt.Println(gasPrice)
	//Bidding engine initialisation
	address, _, instance, err := contract.DeployBid(auth, client, fromAddress, big.NewInt(10), big.NewInt(5), big.NewInt(15), "f80981fc8a46c8b499ffe8ca6d1b2bb69bad563846f498da73c848891887d586")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address.Hex())

	go eventEmission(address)

	privateKey1, err := crypto.HexToECDSA("bb63b692f9d8f21f0b978b596dc2b8611899f053d68aec6c1c20d1df4f5b6ee2")
	if err != nil {
		log.Fatal(err)
	}

	publicKey1 := privateKey1.Public()
	publicKeyECDSA1, ok := publicKey1.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress1 := crypto.PubkeyToAddress(*publicKeyECDSA1)
	nonce1, err := client.PendingNonceAt(context.Background(), fromAddress1)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice1, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth1 := bind.NewKeyedTransactor(privateKey1)
	auth1.Nonce = big.NewInt(int64(nonce1)) //Correct nonce for each transaction
	auth1.Value = big.NewInt(20)            // bid value
	auth1.GasLimit = uint64(300000000)
	auth1.GasPrice = gasPrice1
	fmt.Println(gasPrice1)

	privateKey2, err := crypto.HexToECDSA("2f615ea53711e0d91390e97cdd5ce97357e345e441aa95d255094164f44c8652")
	if err != nil {
		log.Fatal(err)
	}

	publicKey2 := privateKey2.Public()
	publicKeyECDSA2, ok := publicKey2.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress2 := crypto.PubkeyToAddress(*publicKeyECDSA2)
	nonce2, err := client.PendingNonceAt(context.Background(), fromAddress2)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice2, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth2 := bind.NewKeyedTransactor(privateKey2)
	auth2.Nonce = big.NewInt(int64(nonce2))
	auth2.Value = big.NewInt(30)       // bid value
	auth2.GasLimit = uint64(300000000) // Gas configuration for transaction in transaction options
	auth2.GasPrice = gasPrice2         // Incorrect configuration can lead to block gas limit exceeded, insufficient funds of owner to do transaction errors
	fmt.Println(gasPrice2)

	privateKey3, err := crypto.HexToECDSA("7d52c3f6477e1507d54a826833169ad169a56e02ffc49a1801218a7d87ca50bd")
	if err != nil {
		log.Fatal(err)
	}

	publicKey3 := privateKey2.Public()
	publicKeyECDSA3, ok := publicKey3.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress3 := crypto.PubkeyToAddress(*publicKeyECDSA3)
	nonce3, err := client.PendingNonceAt(context.Background(), fromAddress3)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice3, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth3 := bind.NewKeyedTransactor(privateKey3)
	auth3.Nonce = big.NewInt(int64(nonce3))
	auth3.Value = big.NewInt(40) // bid value
	auth3.GasLimit = uint64(300000000)
	auth3.GasPrice = gasPrice3
	fmt.Println(gasPrice3)

	privateKey4, err := crypto.HexToECDSA("6aecd44fcb79d4b68f1ee2b2c706f8e9a0cd06b0de4729fe98cfed8886315256")
	if err != nil {
		log.Fatal(err)
	}

	publicKey4 := privateKey4.Public()
	publicKeyECDSA4, ok := publicKey4.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress4 := crypto.PubkeyToAddress(*publicKeyECDSA4)
	nonce4, err := client.PendingNonceAt(context.Background(), fromAddress4)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice4, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth4 := bind.NewKeyedTransactor(privateKey4)
	auth4.Nonce = big.NewInt(int64(nonce4))
	auth4.Value = big.NewInt(40) // bid value
	auth4.GasLimit = uint64(300000000)
	auth4.GasPrice = gasPrice
	fmt.Println(gasPrice4)
	//Player 1 placing bids
	value, err := instance.PlaceBid(auth1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
	//Player 2 placing bids
	value, err = instance.PlaceBid(auth2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value)
	//Display bids placed by bidder
	value1, err := instance.FundsByBidder(&bind.CallOpts{}, fromAddress1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value1)
	//Display higher bid placed so far
	value2, err := instance.GetHighestBid(&bind.CallOpts{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(value2)

	done := make(chan interface{})
	//N bids for maximum bids to place in game
	in := bidgenerator(done, N_Bids)

	//Max duration of game
	game := NewGame(MAX_DURATION)

	game.Run(instance, client, fromAddress, fromAddress1, fromAddress2, fromAddress3, done, in)

	time.Sleep(200 * time.Second)

}
