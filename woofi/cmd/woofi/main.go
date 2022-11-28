package main

import (
	"github.com/nakji-network/woofi-connectors/woofi"
	"github.com/nakji-network/woofi-connectors/woofi/WOOPP"

	"github.com/spf13/pflag"
	_ "go.uber.org/automaxprocs"
)

func main() {
	fromBlock := pflag.Uint64P("from-block", "f", 0, "block number to start backfill from (optional)")
	numBlocks := pflag.Uint64P("num-blocks", "b", 0, "number of blocks to backfill (optional)")

	pflag.Parse()

	conf := &woofi.Config{
		NetworkName: "bsc",
		FromBlock:   *fromBlock,
		NumBlocks:   *numBlocks,
	}

	c := woofi.New(conf)
	c.AddContract(WOOPP.NewContract("0xbf365Ce9cFcb2d5855521985E351bA3bcf77FD3F")) // WooPPV3 (backward compatibility)
	c.Start()
}
