// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AsciiNormalizerMetaData contains all meta data concerning the AsciiNormalizer contract.
var AsciiNormalizerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"r\",\"type\":\"bytes1\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"addRules\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"idnamap\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"namehash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"normalize\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"valid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"47dcf5b6": "addRules(bytes1,uint256)",
		"0e4cdd59": "idnamap(uint256)",
		"09879962": "namehash(string)",
		"9f10f4b5": "normalize(string)",
		"9791c097": "valid(string)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610858806100206000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063098799621461005c5780630e4cdd591461008657806347dcf5b6146100b25780639791c097146100c75780639f10f4b5146100ea575b600080fd5b61006f61006a366004610599565b61010a565b60405161007d929190610697565b60405180910390f35b6100996100943660046106b9565b6102b4565b6040516001600160f81b0319909116815260200161007d565b6100c56100c03660046106d2565b6102e8565b005b6100da6100d536600461070b565b61035d565b604051901515815260200161007d565b6100fd6100f8366004610599565b610432565b60405161007d919061077d565b80516060906000908082805b83156102675760008761012a6001876107ad565b8151811061013a5761013a6107c4565b01602001516001600160f81b0319169050601760f91b8190036101a657610162888686610535565b604080516020810186905290810182905290925060600160405160208183030381529060405280519060200120925060018561019e91906107ad565b935050610255565b600160ff1b6001600160f81b03198216106101c057600080fd5b6000808260f81c60ff16815481106101da576101da6107c4565b600091825260209182902091810490910154601f9091166101000a900460f81b90506001600160f81b0319811661021057600080fd5b600160f882901c111561025257808961022a6001896107ad565b8151811061023a5761023a6107c4565b60200101906001600160f81b031916908160001a9053505b50505b8361025f816107da565b945050610116565b610272878585610535565b9050868282604051602001610291929190918252602082015260400190565b604051602081830303815290604052805190602001209550955050505050915091565b600081815481106102c457600080fd5b9060005260206000209060209182820401919006915054906101000a900460f81b81565b60005b818110156103585760008054600181018255908052602081047f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301805460f886901c601f9093166101000a92830260ff909302191691909117905580610350816107f1565b9150506102eb565b505050565b6000805b8281101561042657600084848381811061037d5761037d6107c4565b909101356001600160f81b031916915050601760f91b8190036103a05750610414565b600160ff1b6001600160f81b0319821611156103c15760009250505061042c565b60008160f81c60ff16815481106103da576103da6107c4565b90600052602060002090602091828204019190069054906101000a900460f81b60f81c60ff166001146104125760009250505061042c565b505b8061041e816107f1565b915050610361565b50600190505b92915050565b606060005b825181101561052e576000838281518110610454576104546107c4565b01602001516001600160f81b0319169050601760f91b819003610477575061051c565b600160ff1b6001600160f81b031982161061049157600080fd5b6000808260f81c60ff16815481106104ab576104ab6107c4565b600091825260209182902091810490910154601f9091166101000a900460f81b90506001600160f81b031981166104e157600080fd5b600160f882901c11156105195780858481518110610501576105016107c4565b60200101906001600160f81b031916908160001a9053505b50505b80610526816107f1565b915050610437565b5090919050565b6000808061054485602061080a565b905080860191506105796040518060400160405280878761056591906107ad565b815260200184905280516020909101512090565b9695505050505050565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156105ab57600080fd5b813567ffffffffffffffff808211156105c357600080fd5b818401915084601f8301126105d757600080fd5b8135818111156105e9576105e9610583565b604051601f8201601f19908116603f0116810190838211818310171561061157610611610583565b8160405282815287602084870101111561062a57600080fd5b826020860160208301376000928101602001929092525095945050505050565b6000815180845260005b8181101561067057602081850181015186830182015201610654565b81811115610682576000602083870101525b50601f01601f19169290920160200192915050565b6040815260006106aa604083018561064a565b90508260208301529392505050565b6000602082840312156106cb57600080fd5b5035919050565b600080604083850312156106e557600080fd5b82356001600160f81b0319811681146106fd57600080fd5b946020939093013593505050565b6000806020838503121561071e57600080fd5b823567ffffffffffffffff8082111561073657600080fd5b818501915085601f83011261074a57600080fd5b81358181111561075957600080fd5b86602082850101111561076b57600080fd5b60209290920196919550909350505050565b602081526000610790602083018461064a565b9392505050565b634e487b7160e01b600052601160045260246000fd5b6000828210156107bf576107bf610797565b500390565b634e487b7160e01b600052603260045260246000fd5b6000816107e9576107e9610797565b506000190190565b60006001820161080357610803610797565b5060010190565b6000821982111561081d5761081d610797565b50019056fea2646970667358221220a92cd1b71144fdabb1241fa9d2ada735ec789ffd8b8872e9c1883998eb5252fe64736f6c634300080d0033",
}

// AsciiNormalizerABI is the input ABI used to generate the binding from.
// Deprecated: Use AsciiNormalizerMetaData.ABI instead.
var AsciiNormalizerABI = AsciiNormalizerMetaData.ABI

// Deprecated: Use AsciiNormalizerMetaData.Sigs instead.
// AsciiNormalizerFuncSigs maps the 4-byte function signature to its string representation.
var AsciiNormalizerFuncSigs = AsciiNormalizerMetaData.Sigs

// AsciiNormalizerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AsciiNormalizerMetaData.Bin instead.
var AsciiNormalizerBin = AsciiNormalizerMetaData.Bin

// DeployAsciiNormalizer deploys a new Ethereum contract, binding an instance of AsciiNormalizer to it.
func DeployAsciiNormalizer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AsciiNormalizer, error) {
	parsed, err := AsciiNormalizerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AsciiNormalizerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AsciiNormalizer{AsciiNormalizerCaller: AsciiNormalizerCaller{contract: contract}, AsciiNormalizerTransactor: AsciiNormalizerTransactor{contract: contract}, AsciiNormalizerFilterer: AsciiNormalizerFilterer{contract: contract}}, nil
}

// AsciiNormalizer is an auto generated Go binding around an Ethereum contract.
type AsciiNormalizer struct {
	AsciiNormalizerCaller     // Read-only binding to the contract
	AsciiNormalizerTransactor // Write-only binding to the contract
	AsciiNormalizerFilterer   // Log filterer for contract events
}

// AsciiNormalizerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AsciiNormalizerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AsciiNormalizerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AsciiNormalizerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AsciiNormalizerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AsciiNormalizerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AsciiNormalizerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AsciiNormalizerSession struct {
	Contract     *AsciiNormalizer  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AsciiNormalizerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AsciiNormalizerCallerSession struct {
	Contract *AsciiNormalizerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AsciiNormalizerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AsciiNormalizerTransactorSession struct {
	Contract     *AsciiNormalizerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AsciiNormalizerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AsciiNormalizerRaw struct {
	Contract *AsciiNormalizer // Generic contract binding to access the raw methods on
}

// AsciiNormalizerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AsciiNormalizerCallerRaw struct {
	Contract *AsciiNormalizerCaller // Generic read-only contract binding to access the raw methods on
}

// AsciiNormalizerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AsciiNormalizerTransactorRaw struct {
	Contract *AsciiNormalizerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAsciiNormalizer creates a new instance of AsciiNormalizer, bound to a specific deployed contract.
func NewAsciiNormalizer(address common.Address, backend bind.ContractBackend) (*AsciiNormalizer, error) {
	contract, err := bindAsciiNormalizer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AsciiNormalizer{AsciiNormalizerCaller: AsciiNormalizerCaller{contract: contract}, AsciiNormalizerTransactor: AsciiNormalizerTransactor{contract: contract}, AsciiNormalizerFilterer: AsciiNormalizerFilterer{contract: contract}}, nil
}

// NewAsciiNormalizerCaller creates a new read-only instance of AsciiNormalizer, bound to a specific deployed contract.
func NewAsciiNormalizerCaller(address common.Address, caller bind.ContractCaller) (*AsciiNormalizerCaller, error) {
	contract, err := bindAsciiNormalizer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AsciiNormalizerCaller{contract: contract}, nil
}

// NewAsciiNormalizerTransactor creates a new write-only instance of AsciiNormalizer, bound to a specific deployed contract.
func NewAsciiNormalizerTransactor(address common.Address, transactor bind.ContractTransactor) (*AsciiNormalizerTransactor, error) {
	contract, err := bindAsciiNormalizer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AsciiNormalizerTransactor{contract: contract}, nil
}

// NewAsciiNormalizerFilterer creates a new log filterer instance of AsciiNormalizer, bound to a specific deployed contract.
func NewAsciiNormalizerFilterer(address common.Address, filterer bind.ContractFilterer) (*AsciiNormalizerFilterer, error) {
	contract, err := bindAsciiNormalizer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AsciiNormalizerFilterer{contract: contract}, nil
}

// bindAsciiNormalizer binds a generic wrapper to an already deployed contract.
func bindAsciiNormalizer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AsciiNormalizerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AsciiNormalizer *AsciiNormalizerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AsciiNormalizer.Contract.AsciiNormalizerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AsciiNormalizer *AsciiNormalizerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AsciiNormalizer.Contract.AsciiNormalizerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AsciiNormalizer *AsciiNormalizerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AsciiNormalizer.Contract.AsciiNormalizerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AsciiNormalizer *AsciiNormalizerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AsciiNormalizer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AsciiNormalizer *AsciiNormalizerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AsciiNormalizer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AsciiNormalizer *AsciiNormalizerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AsciiNormalizer.Contract.contract.Transact(opts, method, params...)
}

// Idnamap is a free data retrieval call binding the contract method 0x0e4cdd59.
//
// Solidity: function idnamap(uint256 ) view returns(bytes1)
func (_AsciiNormalizer *AsciiNormalizerCaller) Idnamap(opts *bind.CallOpts, arg0 *big.Int) ([1]byte, error) {
	var out []interface{}
	err := _AsciiNormalizer.contract.Call(opts, &out, "idnamap", arg0)

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// Idnamap is a free data retrieval call binding the contract method 0x0e4cdd59.
//
// Solidity: function idnamap(uint256 ) view returns(bytes1)
func (_AsciiNormalizer *AsciiNormalizerSession) Idnamap(arg0 *big.Int) ([1]byte, error) {
	return _AsciiNormalizer.Contract.Idnamap(&_AsciiNormalizer.CallOpts, arg0)
}

// Idnamap is a free data retrieval call binding the contract method 0x0e4cdd59.
//
// Solidity: function idnamap(uint256 ) view returns(bytes1)
func (_AsciiNormalizer *AsciiNormalizerCallerSession) Idnamap(arg0 *big.Int) ([1]byte, error) {
	return _AsciiNormalizer.Contract.Idnamap(&_AsciiNormalizer.CallOpts, arg0)
}

// Namehash is a free data retrieval call binding the contract method 0x09879962.
//
// Solidity: function namehash(string domain) view returns(string, bytes32)
func (_AsciiNormalizer *AsciiNormalizerCaller) Namehash(opts *bind.CallOpts, domain string) (string, [32]byte, error) {
	var out []interface{}
	err := _AsciiNormalizer.contract.Call(opts, &out, "namehash", domain)

	if err != nil {
		return *new(string), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// Namehash is a free data retrieval call binding the contract method 0x09879962.
//
// Solidity: function namehash(string domain) view returns(string, bytes32)
func (_AsciiNormalizer *AsciiNormalizerSession) Namehash(domain string) (string, [32]byte, error) {
	return _AsciiNormalizer.Contract.Namehash(&_AsciiNormalizer.CallOpts, domain)
}

// Namehash is a free data retrieval call binding the contract method 0x09879962.
//
// Solidity: function namehash(string domain) view returns(string, bytes32)
func (_AsciiNormalizer *AsciiNormalizerCallerSession) Namehash(domain string) (string, [32]byte, error) {
	return _AsciiNormalizer.Contract.Namehash(&_AsciiNormalizer.CallOpts, domain)
}

// Normalize is a free data retrieval call binding the contract method 0x9f10f4b5.
//
// Solidity: function normalize(string domain) view returns(string)
func (_AsciiNormalizer *AsciiNormalizerCaller) Normalize(opts *bind.CallOpts, domain string) (string, error) {
	var out []interface{}
	err := _AsciiNormalizer.contract.Call(opts, &out, "normalize", domain)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Normalize is a free data retrieval call binding the contract method 0x9f10f4b5.
//
// Solidity: function normalize(string domain) view returns(string)
func (_AsciiNormalizer *AsciiNormalizerSession) Normalize(domain string) (string, error) {
	return _AsciiNormalizer.Contract.Normalize(&_AsciiNormalizer.CallOpts, domain)
}

// Normalize is a free data retrieval call binding the contract method 0x9f10f4b5.
//
// Solidity: function normalize(string domain) view returns(string)
func (_AsciiNormalizer *AsciiNormalizerCallerSession) Normalize(domain string) (string, error) {
	return _AsciiNormalizer.Contract.Normalize(&_AsciiNormalizer.CallOpts, domain)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string domain) view returns(bool)
func (_AsciiNormalizer *AsciiNormalizerCaller) Valid(opts *bind.CallOpts, domain string) (bool, error) {
	var out []interface{}
	err := _AsciiNormalizer.contract.Call(opts, &out, "valid", domain)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string domain) view returns(bool)
func (_AsciiNormalizer *AsciiNormalizerSession) Valid(domain string) (bool, error) {
	return _AsciiNormalizer.Contract.Valid(&_AsciiNormalizer.CallOpts, domain)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string domain) view returns(bool)
func (_AsciiNormalizer *AsciiNormalizerCallerSession) Valid(domain string) (bool, error) {
	return _AsciiNormalizer.Contract.Valid(&_AsciiNormalizer.CallOpts, domain)
}

// AddRules is a paid mutator transaction binding the contract method 0x47dcf5b6.
//
// Solidity: function addRules(bytes1 r, uint256 num) returns()
func (_AsciiNormalizer *AsciiNormalizerTransactor) AddRules(opts *bind.TransactOpts, r [1]byte, num *big.Int) (*types.Transaction, error) {
	return _AsciiNormalizer.contract.Transact(opts, "addRules", r, num)
}

// AddRules is a paid mutator transaction binding the contract method 0x47dcf5b6.
//
// Solidity: function addRules(bytes1 r, uint256 num) returns()
func (_AsciiNormalizer *AsciiNormalizerSession) AddRules(r [1]byte, num *big.Int) (*types.Transaction, error) {
	return _AsciiNormalizer.Contract.AddRules(&_AsciiNormalizer.TransactOpts, r, num)
}

// AddRules is a paid mutator transaction binding the contract method 0x47dcf5b6.
//
// Solidity: function addRules(bytes1 r, uint256 num) returns()
func (_AsciiNormalizer *AsciiNormalizerTransactorSession) AddRules(r [1]byte, num *big.Int) (*types.Transaction, error) {
	return _AsciiNormalizer.Contract.AddRules(&_AsciiNormalizer.TransactOpts, r, num)
}
