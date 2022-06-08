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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes1\",\"name\":\"r\",\"type\":\"bytes1\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"addRules\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"namehash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"normalize\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"valid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"47dcf5b6": "addRules(bytes1,uint256)",
		"09879962": "namehash(string)",
		"9f10f4b5": "normalize(string)",
		"9791c097": "valid(string)",
	},
	Bin: "0x608060405234801561001057600080fd5b506107d4806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063098799621461005157806347dcf5b61461007b5780639791c097146100905780639f10f4b5146100b3575b600080fd5b61006461005f36600461052e565b6100d3565b60405161007292919061062c565b60405180910390f35b61008e61008936600461064e565b61027d565b005b6100a361009e366004610687565b6102f2565b6040519015158152602001610072565b6100c66100c136600461052e565b6103c7565b60405161007291906106f9565b80516060906000908082805b8315610230576000876100f3600187610729565b8151811061010357610103610740565b01602001516001600160f81b0319169050601760f91b81900361016f5761012b8886866104ca565b60408051602081018690529081018290529092506060016040516020818303038152906040528051906020012092506001856101679190610729565b93505061021e565b600160ff1b6001600160f81b031982161061018957600080fd5b6000808260f81c60ff16815481106101a3576101a3610740565b600091825260209182902091810490910154601f9091166101000a900460f81b90506001600160f81b031981166101d957600080fd5b600160f882901c111561021b5780896101f3600189610729565b8151811061020357610203610740565b60200101906001600160f81b031916908160001a9053505b50505b8361022881610756565b9450506100df565b61023b8785856104ca565b905086828260405160200161025a929190918252602082015260400190565b604051602081830303815290604052805190602001209550955050505050915091565b60005b818110156102ed5760008054600181018255908052602081047f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301805460f886901c601f9093166101000a92830260ff9093021916919091179055806102e58161076d565b915050610280565b505050565b6000805b828110156103bb57600084848381811061031257610312610740565b909101356001600160f81b031916915050601760f91b81900361033557506103a9565b600160ff1b6001600160f81b031982161115610356576000925050506103c1565b60008160f81c60ff168154811061036f5761036f610740565b90600052602060002090602091828204019190069054906101000a900460f81b60f81c60ff166001146103a7576000925050506103c1565b505b806103b38161076d565b9150506102f6565b50600190505b92915050565b606060005b82518110156104c35760008382815181106103e9576103e9610740565b01602001516001600160f81b0319169050601760f91b81900361040c57506104b1565b600160ff1b6001600160f81b031982161061042657600080fd5b6000808260f81c60ff168154811061044057610440610740565b600091825260209182902091810490910154601f9091166101000a900460f81b90506001600160f81b0319811661047657600080fd5b600160f882901c11156104ae578085848151811061049657610496610740565b60200101906001600160f81b031916908160001a9053505b50505b806104bb8161076d565b9150506103cc565b5090919050565b600080806104d9856020610786565b9050808601915061050e604051806040016040528087876104fa9190610729565b815260200184905280516020909101512090565b9695505050505050565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561054057600080fd5b813567ffffffffffffffff8082111561055857600080fd5b818401915084601f83011261056c57600080fd5b81358181111561057e5761057e610518565b604051601f8201601f19908116603f011681019083821181831017156105a6576105a6610518565b816040528281528760208487010111156105bf57600080fd5b826020860160208301376000928101602001929092525095945050505050565b6000815180845260005b81811015610605576020818501810151868301820152016105e9565b81811115610617576000602083870101525b50601f01601f19169290920160200192915050565b60408152600061063f60408301856105df565b90508260208301529392505050565b6000806040838503121561066157600080fd5b82356001600160f81b03198116811461067957600080fd5b946020939093013593505050565b6000806020838503121561069a57600080fd5b823567ffffffffffffffff808211156106b257600080fd5b818501915085601f8301126106c657600080fd5b8135818111156106d557600080fd5b8660208285010111156106e757600080fd5b60209290920196919550909350505050565b60208152600061070c60208301846105df565b9392505050565b634e487b7160e01b600052601160045260246000fd5b60008282101561073b5761073b610713565b500390565b634e487b7160e01b600052603260045260246000fd5b60008161076557610765610713565b506000190190565b60006001820161077f5761077f610713565b5060010190565b6000821982111561079957610799610713565b50019056fea2646970667358221220ca545ed46b37c0b2f374cb74903c4206cd95f2d56b144b0bbf07ba31cde8b35f64736f6c634300080d0033",
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

// StringsMetaData contains all meta data concerning the Strings contract.
var StringsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220bd8d574c32f5ace7281316275cd96361baca90fafa6722ded60aace9c17a025764736f6c634300080d0033",
}

// StringsABI is the input ABI used to generate the binding from.
// Deprecated: Use StringsMetaData.ABI instead.
var StringsABI = StringsMetaData.ABI

// StringsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StringsMetaData.Bin instead.
var StringsBin = StringsMetaData.Bin

// DeployStrings deploys a new Ethereum contract, binding an instance of Strings to it.
func DeployStrings(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Strings, error) {
	parsed, err := StringsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StringsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// Strings is an auto generated Go binding around an Ethereum contract.
type Strings struct {
	StringsCaller     // Read-only binding to the contract
	StringsTransactor // Write-only binding to the contract
	StringsFilterer   // Log filterer for contract events
}

// StringsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StringsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StringsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StringsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StringsSession struct {
	Contract     *Strings          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StringsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StringsCallerSession struct {
	Contract *StringsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StringsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StringsTransactorSession struct {
	Contract     *StringsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StringsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StringsRaw struct {
	Contract *Strings // Generic contract binding to access the raw methods on
}

// StringsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StringsCallerRaw struct {
	Contract *StringsCaller // Generic read-only contract binding to access the raw methods on
}

// StringsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StringsTransactorRaw struct {
	Contract *StringsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrings creates a new instance of Strings, bound to a specific deployed contract.
func NewStrings(address common.Address, backend bind.ContractBackend) (*Strings, error) {
	contract, err := bindStrings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// NewStringsCaller creates a new read-only instance of Strings, bound to a specific deployed contract.
func NewStringsCaller(address common.Address, caller bind.ContractCaller) (*StringsCaller, error) {
	contract, err := bindStrings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StringsCaller{contract: contract}, nil
}

// NewStringsTransactor creates a new write-only instance of Strings, bound to a specific deployed contract.
func NewStringsTransactor(address common.Address, transactor bind.ContractTransactor) (*StringsTransactor, error) {
	contract, err := bindStrings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StringsTransactor{contract: contract}, nil
}

// NewStringsFilterer creates a new log filterer instance of Strings, bound to a specific deployed contract.
func NewStringsFilterer(address common.Address, filterer bind.ContractFilterer) (*StringsFilterer, error) {
	contract, err := bindStrings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StringsFilterer{contract: contract}, nil
}

// bindStrings binds a generic wrapper to an already deployed contract.
func bindStrings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StringsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.StringsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transact(opts, method, params...)
}
