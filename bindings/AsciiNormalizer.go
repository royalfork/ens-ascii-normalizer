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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"asciimap\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"idnamap\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"namehash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0e4cdd59": "idnamap(uint256)",
		"09879962": "namehash(string)",
	},
	Bin: "0x608060405234801561001057600080fd5b506040516106ad3803806106ad83398101604081905261002f91610140565b60005b815181101561012357600082610049836001610225565b815181106100595761005961023d565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016905060005b8383815181106100975761009761023d565b016020015160f81c60ff8216101561010e5760008054600181018255908052602081047f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301805460f885901c601f9093166101000a92830260ff90930219169190911790558061010681610253565b915050610085565b5061011c9050600282610225565b9050610032565b5050610272565b634e487b7160e01b600052604160045260246000fd5b6000602080838503121561015357600080fd5b82516001600160401b038082111561016a57600080fd5b818501915085601f83011261017e57600080fd5b8151818111156101905761019061012a565b604051601f8201601f19908116603f011681019083821181831017156101b8576101b861012a565b8160405282815288868487010111156101d057600080fd5b600093505b828410156101f257848401860151818501870152928501926101d5565b828411156102035760008684830101525b98975050505050505050565b634e487b7160e01b600052601160045260246000fd5b600082198211156102385761023861020f565b500190565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff81036102695761026961020f565b60010192915050565b61042c806102816000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063098799621461003b5780630e4cdd5914610065575b600080fd5b61004e610049366004610275565b610091565b60405161005c929190610326565b60405180910390f35b610078610073366004610383565b61022b565b6040516001600160f81b0319909116815260200161005c565b805160609060009080825b82156101e9576000866100b06001866103b2565b815181106100c0576100c06103c9565b01602001516001600160f81b0319169050601760f91b8190036101285783830360208589010120829060408051602081019390935282015260600160405160208183030381529060405280519060200120915060018461012091906103b2565b9250506101d7565b600160ff1b6001600160f81b031982161061014257600080fd5b6000808260f81c60ff168154811061015c5761015c6103c9565b600091825260209182902091810490910154601f9091166101000a900460f81b90506001600160f81b0319811661019257600080fd5b600160f882901c11156101d45780886101ac6001886103b2565b815181106101bc576101bc6103c9565b60200101906001600160f81b031916908160001a9053505b50505b826101e1816103df565b93505061009c565b82820360208488010120869082906040805160208101939093528201526060016040516020818303038152906040528051906020012094509450505050915091565b6000818154811061023b57600080fd5b9060005260206000209060209182820401919006915054906101000a900460f81b81565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561028757600080fd5b813567ffffffffffffffff8082111561029f57600080fd5b818401915084601f8301126102b357600080fd5b8135818111156102c5576102c561025f565b604051601f8201601f19908116603f011681019083821181831017156102ed576102ed61025f565b8160405282815287602084870101111561030657600080fd5b826020860160208301376000928101602001929092525095945050505050565b604081526000835180604084015260005b818110156103545760208187018101516060868401015201610337565b81811115610366576000606083860101525b50602083019390935250601f91909101601f191601606001919050565b60006020828403121561039557600080fd5b5035919050565b634e487b7160e01b600052601160045260246000fd5b6000828210156103c4576103c461039c565b500390565b634e487b7160e01b600052603260045260246000fd5b6000816103ee576103ee61039c565b50600019019056fea2646970667358221220a2762fffe6893b1e5f6b833dcfbb2a8f4b136bfa2a7e62ccc2fc3040a26a597164736f6c634300080d0033",
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
func DeployAsciiNormalizer(auth *bind.TransactOpts, backend bind.ContractBackend, asciimap []byte) (common.Address, *types.Transaction, *AsciiNormalizer, error) {
	parsed, err := AsciiNormalizerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AsciiNormalizerBin), backend, asciimap)
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
