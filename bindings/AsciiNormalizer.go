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

// ENSMetaData contains all meta data concerning the ENS contract.
var ENSMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"02571be3": "owner(bytes32)",
	},
}

// ENSABI is the input ABI used to generate the binding from.
// Deprecated: Use ENSMetaData.ABI instead.
var ENSABI = ENSMetaData.ABI

// Deprecated: Use ENSMetaData.Sigs instead.
// ENSFuncSigs maps the 4-byte function signature to its string representation.
var ENSFuncSigs = ENSMetaData.Sigs

// ENS is an auto generated Go binding around an Ethereum contract.
type ENS struct {
	ENSCaller     // Read-only binding to the contract
	ENSTransactor // Write-only binding to the contract
	ENSFilterer   // Log filterer for contract events
}

// ENSCaller is an auto generated read-only Go binding around an Ethereum contract.
type ENSCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ENSTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ENSFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ENSSession struct {
	Contract     *ENS              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ENSCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ENSCallerSession struct {
	Contract *ENSCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ENSTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ENSTransactorSession struct {
	Contract     *ENSTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ENSRaw is an auto generated low-level Go binding around an Ethereum contract.
type ENSRaw struct {
	Contract *ENS // Generic contract binding to access the raw methods on
}

// ENSCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ENSCallerRaw struct {
	Contract *ENSCaller // Generic read-only contract binding to access the raw methods on
}

// ENSTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ENSTransactorRaw struct {
	Contract *ENSTransactor // Generic write-only contract binding to access the raw methods on
}

// NewENS creates a new instance of ENS, bound to a specific deployed contract.
func NewENS(address common.Address, backend bind.ContractBackend) (*ENS, error) {
	contract, err := bindENS(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ENS{ENSCaller: ENSCaller{contract: contract}, ENSTransactor: ENSTransactor{contract: contract}, ENSFilterer: ENSFilterer{contract: contract}}, nil
}

// NewENSCaller creates a new read-only instance of ENS, bound to a specific deployed contract.
func NewENSCaller(address common.Address, caller bind.ContractCaller) (*ENSCaller, error) {
	contract, err := bindENS(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ENSCaller{contract: contract}, nil
}

// NewENSTransactor creates a new write-only instance of ENS, bound to a specific deployed contract.
func NewENSTransactor(address common.Address, transactor bind.ContractTransactor) (*ENSTransactor, error) {
	contract, err := bindENS(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ENSTransactor{contract: contract}, nil
}

// NewENSFilterer creates a new log filterer instance of ENS, bound to a specific deployed contract.
func NewENSFilterer(address common.Address, filterer bind.ContractFilterer) (*ENSFilterer, error) {
	contract, err := bindENS(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ENSFilterer{contract: contract}, nil
}

// bindENS binds a generic wrapper to an already deployed contract.
func bindENS(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ENSABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENS *ENSRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ENS.Contract.ENSCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENS *ENSRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENS.Contract.ENSTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENS *ENSRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENS.Contract.ENSTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENS *ENSCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ENS.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENS *ENSTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENS.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENS *ENSTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENS.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_ENS *ENSCaller) Owner(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ENS.contract.Call(opts, &out, "owner", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_ENS *ENSSession) Owner(node [32]byte) (common.Address, error) {
	return _ENS.Contract.Owner(&_ENS.CallOpts, node)
}

// Owner is a free data retrieval call binding the contract method 0x02571be3.
//
// Solidity: function owner(bytes32 node) view returns(address)
func (_ENS *ENSCallerSession) Owner(node [32]byte) (common.Address, error) {
	return _ENS.Contract.Owner(&_ENS.CallOpts, node)
}

// ENSAsciiNormalizerMetaData contains all meta data concerning the ENSAsciiNormalizer contract.
var ENSAsciiNormalizerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractENS\",\"name\":\"_ens\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"asciimap\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ens\",\"outputs\":[{\"internalType\":\"contractENS\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"idnamap\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"lookup\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"namehash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3f15457f": "ens()",
		"0e4cdd59": "idnamap(uint256)",
		"f67187ac": "lookup(string)",
		"09879962": "namehash(string)",
	},
	Bin: "0x608060405234801561001057600080fd5b5060405161080d38038061080d83398101604081905261002f9161015c565b600080546001600160a01b0319166001600160a01b0384161781555b815181101561013e57600082610062836001610264565b815181106100725761007261027c565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016905060005b8383815181106100b0576100b061027c565b016020015160f81c60ff8216101561012957600180548082018255600091909152602081047fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601805460f885901c601f9093166101000a92830260ff90930219169190911790558061012181610292565b91505061009e565b506101379050600282610264565b905061004b565b5050506102b1565b634e487b7160e01b600052604160045260246000fd5b6000806040838503121561016f57600080fd5b82516001600160a01b038116811461018657600080fd5b602084810151919350906001600160401b03808211156101a557600080fd5b818601915086601f8301126101b957600080fd5b8151818111156101cb576101cb610146565b604051601f8201601f19908116603f011681019083821181831017156101f3576101f3610146565b81604052828152898684870101111561020b57600080fd5b600093505b8284101561022d5784840186015181850187015292850192610210565b8284111561023e5760008684830101525b8096505050505050509250929050565b634e487b7160e01b600052601160045260246000fd5b600082198211156102775761027761024e565b500190565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff81036102a8576102a861024e565b60010192915050565b61054d806102c06000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806309879962146100515780630e4cdd591461007b5780633f15457f146100a7578063f67187ac146100d2575b600080fd5b61006461005f366004610366565b610104565b604051610072929190610417565b60405180910390f35b61008e610089366004610474565b61029a565b6040516001600160f81b03199091168152602001610072565b6000546100ba906001600160a01b031681565b6040516001600160a01b039091168152602001610072565b6100e56100e0366004610366565b6102ce565b604080516001600160a01b039093168352602083019190915201610072565b805160609060009080825b8215610258576000866101236001866104a3565b81518110610133576101336104ba565b01602001516001600160f81b0319169050601760f91b81900361019b5783830360208589010120829060408051602081019390935282015260600160405160208183030381529060405280519060200120915060018461019391906104a3565b925050610246565b600160ff1b6001600160f81b03198216106101b557600080fd5b600060018260f81c60ff16815481106101d0576101d06104ba565b600091825260208083209082040154601f9091166101000a900460f81b915060ff1660f882901c0361020157600080fd5b600160f882901c111561024357808861021b6001886104a3565b8151811061022b5761022b6104ba565b60200101906001600160f81b031916908160001a9053505b50505b82610250816104d0565b93505061010f565b82820360208488010120869082906040805160208101939093528201526060016040516020818303038152906040528051906020012094509450505050915091565b600181815481106102aa57600080fd5b9060005260206000209060209182820401919006915054906101000a900460f81b81565b6000806102da83610104565b6000546040516302571be360e01b8152600481018390529193506001600160a01b031691506302571be390602401602060405180830381865afa158015610325573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061034991906104e7565b9150915091565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561037857600080fd5b813567ffffffffffffffff8082111561039057600080fd5b818401915084601f8301126103a457600080fd5b8135818111156103b6576103b6610350565b604051601f8201601f19908116603f011681019083821181831017156103de576103de610350565b816040528281528760208487010111156103f757600080fd5b826020860160208301376000928101602001929092525095945050505050565b604081526000835180604084015260005b818110156104455760208187018101516060868401015201610428565b81811115610457576000606083860101525b50602083019390935250601f91909101601f191601606001919050565b60006020828403121561048657600080fd5b5035919050565b634e487b7160e01b600052601160045260246000fd5b6000828210156104b5576104b561048d565b500390565b634e487b7160e01b600052603260045260246000fd5b6000816104df576104df61048d565b506000190190565b6000602082840312156104f957600080fd5b81516001600160a01b038116811461051057600080fd5b939250505056fea26469706673582212207692456a955bb98377c922288e72dfb5ffec6f116aca382ab7361765f7ea456064736f6c634300080d0033",
}

// ENSAsciiNormalizerABI is the input ABI used to generate the binding from.
// Deprecated: Use ENSAsciiNormalizerMetaData.ABI instead.
var ENSAsciiNormalizerABI = ENSAsciiNormalizerMetaData.ABI

// Deprecated: Use ENSAsciiNormalizerMetaData.Sigs instead.
// ENSAsciiNormalizerFuncSigs maps the 4-byte function signature to its string representation.
var ENSAsciiNormalizerFuncSigs = ENSAsciiNormalizerMetaData.Sigs

// ENSAsciiNormalizerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ENSAsciiNormalizerMetaData.Bin instead.
var ENSAsciiNormalizerBin = ENSAsciiNormalizerMetaData.Bin

// DeployENSAsciiNormalizer deploys a new Ethereum contract, binding an instance of ENSAsciiNormalizer to it.
func DeployENSAsciiNormalizer(auth *bind.TransactOpts, backend bind.ContractBackend, _ens common.Address, asciimap []byte) (common.Address, *types.Transaction, *ENSAsciiNormalizer, error) {
	parsed, err := ENSAsciiNormalizerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ENSAsciiNormalizerBin), backend, _ens, asciimap)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ENSAsciiNormalizer{ENSAsciiNormalizerCaller: ENSAsciiNormalizerCaller{contract: contract}, ENSAsciiNormalizerTransactor: ENSAsciiNormalizerTransactor{contract: contract}, ENSAsciiNormalizerFilterer: ENSAsciiNormalizerFilterer{contract: contract}}, nil
}

// ENSAsciiNormalizer is an auto generated Go binding around an Ethereum contract.
type ENSAsciiNormalizer struct {
	ENSAsciiNormalizerCaller     // Read-only binding to the contract
	ENSAsciiNormalizerTransactor // Write-only binding to the contract
	ENSAsciiNormalizerFilterer   // Log filterer for contract events
}

// ENSAsciiNormalizerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ENSAsciiNormalizerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSAsciiNormalizerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ENSAsciiNormalizerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSAsciiNormalizerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ENSAsciiNormalizerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ENSAsciiNormalizerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ENSAsciiNormalizerSession struct {
	Contract     *ENSAsciiNormalizer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ENSAsciiNormalizerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ENSAsciiNormalizerCallerSession struct {
	Contract *ENSAsciiNormalizerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ENSAsciiNormalizerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ENSAsciiNormalizerTransactorSession struct {
	Contract     *ENSAsciiNormalizerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ENSAsciiNormalizerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ENSAsciiNormalizerRaw struct {
	Contract *ENSAsciiNormalizer // Generic contract binding to access the raw methods on
}

// ENSAsciiNormalizerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ENSAsciiNormalizerCallerRaw struct {
	Contract *ENSAsciiNormalizerCaller // Generic read-only contract binding to access the raw methods on
}

// ENSAsciiNormalizerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ENSAsciiNormalizerTransactorRaw struct {
	Contract *ENSAsciiNormalizerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewENSAsciiNormalizer creates a new instance of ENSAsciiNormalizer, bound to a specific deployed contract.
func NewENSAsciiNormalizer(address common.Address, backend bind.ContractBackend) (*ENSAsciiNormalizer, error) {
	contract, err := bindENSAsciiNormalizer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ENSAsciiNormalizer{ENSAsciiNormalizerCaller: ENSAsciiNormalizerCaller{contract: contract}, ENSAsciiNormalizerTransactor: ENSAsciiNormalizerTransactor{contract: contract}, ENSAsciiNormalizerFilterer: ENSAsciiNormalizerFilterer{contract: contract}}, nil
}

// NewENSAsciiNormalizerCaller creates a new read-only instance of ENSAsciiNormalizer, bound to a specific deployed contract.
func NewENSAsciiNormalizerCaller(address common.Address, caller bind.ContractCaller) (*ENSAsciiNormalizerCaller, error) {
	contract, err := bindENSAsciiNormalizer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ENSAsciiNormalizerCaller{contract: contract}, nil
}

// NewENSAsciiNormalizerTransactor creates a new write-only instance of ENSAsciiNormalizer, bound to a specific deployed contract.
func NewENSAsciiNormalizerTransactor(address common.Address, transactor bind.ContractTransactor) (*ENSAsciiNormalizerTransactor, error) {
	contract, err := bindENSAsciiNormalizer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ENSAsciiNormalizerTransactor{contract: contract}, nil
}

// NewENSAsciiNormalizerFilterer creates a new log filterer instance of ENSAsciiNormalizer, bound to a specific deployed contract.
func NewENSAsciiNormalizerFilterer(address common.Address, filterer bind.ContractFilterer) (*ENSAsciiNormalizerFilterer, error) {
	contract, err := bindENSAsciiNormalizer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ENSAsciiNormalizerFilterer{contract: contract}, nil
}

// bindENSAsciiNormalizer binds a generic wrapper to an already deployed contract.
func bindENSAsciiNormalizer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ENSAsciiNormalizerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENSAsciiNormalizer *ENSAsciiNormalizerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ENSAsciiNormalizer.Contract.ENSAsciiNormalizerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENSAsciiNormalizer *ENSAsciiNormalizerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENSAsciiNormalizer.Contract.ENSAsciiNormalizerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENSAsciiNormalizer *ENSAsciiNormalizerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENSAsciiNormalizer.Contract.ENSAsciiNormalizerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ENSAsciiNormalizer *ENSAsciiNormalizerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ENSAsciiNormalizer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ENSAsciiNormalizer *ENSAsciiNormalizerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ENSAsciiNormalizer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ENSAsciiNormalizer *ENSAsciiNormalizerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ENSAsciiNormalizer.Contract.contract.Transact(opts, method, params...)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() view returns(address)
func (_ENSAsciiNormalizer *ENSAsciiNormalizerCaller) Ens(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ENSAsciiNormalizer.contract.Call(opts, &out, "ens")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() view returns(address)
func (_ENSAsciiNormalizer *ENSAsciiNormalizerSession) Ens() (common.Address, error) {
	return _ENSAsciiNormalizer.Contract.Ens(&_ENSAsciiNormalizer.CallOpts)
}

// Ens is a free data retrieval call binding the contract method 0x3f15457f.
//
// Solidity: function ens() view returns(address)
func (_ENSAsciiNormalizer *ENSAsciiNormalizerCallerSession) Ens() (common.Address, error) {
	return _ENSAsciiNormalizer.Contract.Ens(&_ENSAsciiNormalizer.CallOpts)
}

// Idnamap is a free data retrieval call binding the contract method 0x0e4cdd59.
//
// Solidity: function idnamap(uint256 ) view returns(bytes1)
func (_ENSAsciiNormalizer *ENSAsciiNormalizerCaller) Idnamap(opts *bind.CallOpts, arg0 *big.Int) ([1]byte, error) {
	var out []interface{}
	err := _ENSAsciiNormalizer.contract.Call(opts, &out, "idnamap", arg0)

	if err != nil {
		return *new([1]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)

	return out0, err

}

// Idnamap is a free data retrieval call binding the contract method 0x0e4cdd59.
//
// Solidity: function idnamap(uint256 ) view returns(bytes1)
func (_ENSAsciiNormalizer *ENSAsciiNormalizerSession) Idnamap(arg0 *big.Int) ([1]byte, error) {
	return _ENSAsciiNormalizer.Contract.Idnamap(&_ENSAsciiNormalizer.CallOpts, arg0)
}

// Idnamap is a free data retrieval call binding the contract method 0x0e4cdd59.
//
// Solidity: function idnamap(uint256 ) view returns(bytes1)
func (_ENSAsciiNormalizer *ENSAsciiNormalizerCallerSession) Idnamap(arg0 *big.Int) ([1]byte, error) {
	return _ENSAsciiNormalizer.Contract.Idnamap(&_ENSAsciiNormalizer.CallOpts, arg0)
}

// Lookup is a free data retrieval call binding the contract method 0xf67187ac.
//
// Solidity: function lookup(string domain) view returns(address owner, bytes32 node)
func (_ENSAsciiNormalizer *ENSAsciiNormalizerCaller) Lookup(opts *bind.CallOpts, domain string) (struct {
	Owner common.Address
	Node  [32]byte
}, error) {
	var out []interface{}
	err := _ENSAsciiNormalizer.contract.Call(opts, &out, "lookup", domain)

	outstruct := new(struct {
		Owner common.Address
		Node  [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Node = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Lookup is a free data retrieval call binding the contract method 0xf67187ac.
//
// Solidity: function lookup(string domain) view returns(address owner, bytes32 node)
func (_ENSAsciiNormalizer *ENSAsciiNormalizerSession) Lookup(domain string) (struct {
	Owner common.Address
	Node  [32]byte
}, error) {
	return _ENSAsciiNormalizer.Contract.Lookup(&_ENSAsciiNormalizer.CallOpts, domain)
}

// Lookup is a free data retrieval call binding the contract method 0xf67187ac.
//
// Solidity: function lookup(string domain) view returns(address owner, bytes32 node)
func (_ENSAsciiNormalizer *ENSAsciiNormalizerCallerSession) Lookup(domain string) (struct {
	Owner common.Address
	Node  [32]byte
}, error) {
	return _ENSAsciiNormalizer.Contract.Lookup(&_ENSAsciiNormalizer.CallOpts, domain)
}

// Namehash is a free data retrieval call binding the contract method 0x09879962.
//
// Solidity: function namehash(string domain) view returns(string, bytes32)
func (_ENSAsciiNormalizer *ENSAsciiNormalizerCaller) Namehash(opts *bind.CallOpts, domain string) (string, [32]byte, error) {
	var out []interface{}
	err := _ENSAsciiNormalizer.contract.Call(opts, &out, "namehash", domain)

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
func (_ENSAsciiNormalizer *ENSAsciiNormalizerSession) Namehash(domain string) (string, [32]byte, error) {
	return _ENSAsciiNormalizer.Contract.Namehash(&_ENSAsciiNormalizer.CallOpts, domain)
}

// Namehash is a free data retrieval call binding the contract method 0x09879962.
//
// Solidity: function namehash(string domain) view returns(string, bytes32)
func (_ENSAsciiNormalizer *ENSAsciiNormalizerCallerSession) Namehash(domain string) (string, [32]byte, error) {
	return _ENSAsciiNormalizer.Contract.Namehash(&_ENSAsciiNormalizer.CallOpts, domain)
}
