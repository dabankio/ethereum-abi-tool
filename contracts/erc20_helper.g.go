// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ERC20InterfaceABI is the input ABI used to generate the binding from.
const ERC20InterfaceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20InterfaceFuncSigs maps the 4-byte function signature to its string representation.
var ERC20InterfaceFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20InterfaceABIHelper tool for contract abi
type ERC20InterfaceABIHelper struct {
	abi abi.ABI
}

// NewERC20InterfaceABIHelper constructor
func NewERC20InterfaceABIHelper() *ERC20InterfaceABIHelper {
	parsed, _ := abi.JSON(strings.NewReader(ERC20InterfaceABI))
	return &ERC20InterfaceABIHelper{parsed}
}

// PackedAllowance is a free data retrieval call binding the contract method 0xdd62ed3e.
func (_ERC20Interface *ERC20InterfaceABIHelper) PackedAllowance(tokenOwner common.Address, spender common.Address) ([]byte, error) {
	return _ERC20Interface.abi.Pack("allowance", tokenOwner, spender)
}

// UnpackAllowance is a free data retrieval call binding the contract method 0xdd62ed3e.
func (_ERC20Interface *ERC20InterfaceABIHelper) UnpackAllowance(output []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Interface.abi.Unpack(out, "allowance", output)
	return *ret0, err
}

// PackedBalanceOf is a free data retrieval call binding the contract method 0x70a08231.
func (_ERC20Interface *ERC20InterfaceABIHelper) PackedBalanceOf(tokenOwner common.Address) ([]byte, error) {
	return _ERC20Interface.abi.Pack("balanceOf", tokenOwner)
}

// UnpackBalanceOf is a free data retrieval call binding the contract method 0x70a08231.
func (_ERC20Interface *ERC20InterfaceABIHelper) UnpackBalanceOf(output []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Interface.abi.Unpack(out, "balanceOf", output)
	return *ret0, err
}

// PackedTotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
func (_ERC20Interface *ERC20InterfaceABIHelper) PackedTotalSupply() ([]byte, error) {
	return _ERC20Interface.abi.Pack("totalSupply")
}

// UnpackTotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
func (_ERC20Interface *ERC20InterfaceABIHelper) UnpackTotalSupply(output []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Interface.abi.Unpack(out, "totalSupply", output)
	return *ret0, err
}

// PackedApprove is a paid mutator transaction binding the contract method 0x095ea7b3.
func (_ERC20Interface *ERC20InterfaceABIHelper) PackedApprove(spender common.Address, tokens *big.Int) ([]byte, error) {
	return _ERC20Interface.abi.Pack("approve", spender, tokens)
}

// PackedTransfer is a paid mutator transaction binding the contract method 0xa9059cbb.
func (_ERC20Interface *ERC20InterfaceABIHelper) PackedTransfer(to common.Address, tokens *big.Int) ([]byte, error) {
	return _ERC20Interface.abi.Pack("transfer", to, tokens)
}

// PackedTransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
func (_ERC20Interface *ERC20InterfaceABIHelper) PackedTransferFrom(from common.Address, to common.Address, tokens *big.Int) ([]byte, error) {
	return _ERC20Interface.abi.Pack("transferFrom", from, to, tokens)
}
