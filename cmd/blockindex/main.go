package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/rickyseezy/block/internal/application"
	"github.com/rickyseezy/block/internal/config"
	"github.com/rickyseezy/block/internal/datasources/remote"
	"github.com/rickyseezy/block/internal/interfaces/rest"
	"github.com/rickyseezy/block/pkg/abi"
	"log"
)

// @title           Blockchain Indexes API
// @version         1.0
// @description     Simple server to interact with Ethereum Blockchain.
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /
func main() {
	c := config.New()
	var r *gin.Engine
	if c.AppEnv == config.Production {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		r.Use(gin.Recovery())
	} else {
		r = gin.Default()
	}

	ethClient, err := ethclient.Dial(fmt.Sprintf("https://%s:%s@%s/%s", c.InfuraProjectID, c.InfuraProjectSecret,
		c.InfuraURI, c.InfuraProjectID))
	if err != nil {
		log.Fatalln("Oops! There was a problem", err)
	} else {
		fmt.Println("Success! you are connected to Ropsten Network")
	}

	contract, err := abi.NewContract(common.HexToAddress(c.ContractAddress), ethClient)
	if err != nil {
		log.Fatalf("Failed to instantiate a Contract: %v", err)
	}

	blockRepo := remote.NewBlockRepository(ethClient)
	groupRepo := remote.NewGroupRepository(contract)
	indexRepo := remote.NewIndexRepository(contract)

	blockIndexApp := application.NewBlockIndex(blockRepo, groupRepo, indexRepo)
	s := rest.NewServer(r, blockIndexApp, c.ServerPort)
	log.Fatal(s.Start())
}
