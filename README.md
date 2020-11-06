# ethereum-abi-tool

## Usage

```bash
git clone --depth=1 https://github.com/dabankio/go-ethereum.git
cd go-ethereum/cmd/abigen
go build -o xabigen
mv xabigen $GOPATH/bin

#gen abi helper file
xabigen --sol contracts/erc20.sol -pkg contracts --tplgo ./abi_helper.tpl --out contracts/erc20_helper.g.go
```

```go
//generated code usage

helper := contracts.NewERC20InterfaceABIHelper()
data, err := helper.PackedTransfer(from, to, tokens)

tx := types.NewTransaction(nonce, contractAddress, value, gasLimit, gasPrice, data)

rlpTx, err := ethtx.EncodeRLP()
```