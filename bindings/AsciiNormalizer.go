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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"asciimap\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"idnamap\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"\",\"type\":\"bytes1\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"namehash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"normalize\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"domain\",\"type\":\"string\"}],\"name\":\"valid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0e4cdd59": "idnamap(uint256)",
		"09879962": "namehash(string)",
		"9f10f4b5": "normalize(string)",
		"9791c097": "valid(string)",
	},
	Bin: "0x60806040523480156200001157600080fd5b5060405162000a3538038062000a35833981016040819052620000349162000156565b60005b815181101562000138576000826200005183600162000248565b8151811062000064576200006462000263565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016905060005b838381518110620000a557620000a562000263565b016020015160f81c60ff82161015620001205760008054600181018255908052602081047f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301805460f885901c601f9093166101000a92830260ff909302191691909117905580620001178162000279565b91505062000090565b5062000130905060028262000248565b905062000037565b50506200029b565b634e487b7160e01b600052604160045260246000fd5b600060208083850312156200016a57600080fd5b82516001600160401b03808211156200018257600080fd5b818501915085601f8301126200019757600080fd5b815181811115620001ac57620001ac62000140565b604051601f8201601f19908116603f01168101908382118183101715620001d757620001d762000140565b816040528281528886848701011115620001f057600080fd5b600093505b82841015620002145784840186015181850187015292850192620001f5565b82841115620002265760008684830101525b98975050505050505050565b634e487b7160e01b600052601160045260246000fd5b600082198211156200025e576200025e62000232565b500190565b634e487b7160e01b600052603260045260246000fd5b600060ff821660ff810362000292576200029262000232565b60010192915050565b61078a80620002ab6000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806309879962146100515780630e4cdd591461007b5780639791c097146100a75780639f10f4b5146100ca575b600080fd5b61006461005f366004610504565b6100ea565b604051610072929190610602565b60405180910390f35b61008e610089366004610624565b610294565b6040516001600160f81b03199091168152602001610072565b6100ba6100b536600461063d565b6102c8565b6040519015158152602001610072565b6100dd6100d8366004610504565b61039d565b60405161007291906106af565b80516060906000908082805b83156102475760008761010a6001876106df565b8151811061011a5761011a6106f6565b01602001516001600160f81b0319169050601760f91b819003610186576101428886866104a0565b604080516020810186905290810182905290925060600160405160208183030381529060405280519060200120925060018561017e91906106df565b935050610235565b600160ff1b6001600160f81b03198216106101a057600080fd5b6000808260f81c60ff16815481106101ba576101ba6106f6565b600091825260209182902091810490910154601f9091166101000a900460f81b90506001600160f81b031981166101f057600080fd5b600160f882901c111561023257808961020a6001896106df565b8151811061021a5761021a6106f6565b60200101906001600160f81b031916908160001a9053505b50505b8361023f8161070c565b9450506100f6565b6102528785856104a0565b9050868282604051602001610271929190918252602082015260400190565b604051602081830303815290604052805190602001209550955050505050915091565b600081815481106102a457600080fd5b9060005260206000209060209182820401919006915054906101000a900460f81b81565b6000805b828110156103915760008484838181106102e8576102e86106f6565b909101356001600160f81b031916915050601760f91b81900361030b575061037f565b600160ff1b6001600160f81b03198216111561032c57600092505050610397565b60008160f81c60ff1681548110610345576103456106f6565b90600052602060002090602091828204019190069054906101000a900460f81b60f81c60ff1660011461037d57600092505050610397565b505b8061038981610723565b9150506102cc565b50600190505b92915050565b606060005b82518110156104995760008382815181106103bf576103bf6106f6565b01602001516001600160f81b0319169050601760f91b8190036103e25750610487565b600160ff1b6001600160f81b03198216106103fc57600080fd5b6000808260f81c60ff1681548110610416576104166106f6565b600091825260209182902091810490910154601f9091166101000a900460f81b90506001600160f81b0319811661044c57600080fd5b600160f882901c1115610484578085848151811061046c5761046c6106f6565b60200101906001600160f81b031916908160001a9053505b50505b8061049181610723565b9150506103a2565b5090919050565b600080806104af85602061073c565b905080860191506104e4604051806040016040528087876104d091906106df565b815260200184905280516020909101512090565b9695505050505050565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561051657600080fd5b813567ffffffffffffffff8082111561052e57600080fd5b818401915084601f83011261054257600080fd5b813581811115610554576105546104ee565b604051601f8201601f19908116603f0116810190838211818310171561057c5761057c6104ee565b8160405282815287602084870101111561059557600080fd5b826020860160208301376000928101602001929092525095945050505050565b6000815180845260005b818110156105db576020818501810151868301820152016105bf565b818111156105ed576000602083870101525b50601f01601f19169290920160200192915050565b60408152600061061560408301856105b5565b90508260208301529392505050565b60006020828403121561063657600080fd5b5035919050565b6000806020838503121561065057600080fd5b823567ffffffffffffffff8082111561066857600080fd5b818501915085601f83011261067c57600080fd5b81358181111561068b57600080fd5b86602082850101111561069d57600080fd5b60209290920196919550909350505050565b6020815260006106c260208301846105b5565b9392505050565b634e487b7160e01b600052601160045260246000fd5b6000828210156106f1576106f16106c9565b500390565b634e487b7160e01b600052603260045260246000fd5b60008161071b5761071b6106c9565b506000190190565b600060018201610735576107356106c9565b5060010190565b6000821982111561074f5761074f6106c9565b50019056fea2646970667358221220135e2f0ff4ff8d6d533ab8cabe7ff60e2ff01d2a0ba5a1bf04ca11e833bdc53b64736f6c634300080d0033",
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
