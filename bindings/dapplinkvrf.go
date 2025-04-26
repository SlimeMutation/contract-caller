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
	_ = abi.ConvertType
)

// DappLinkVRFMetaData contains all meta data concerning the DappLinkVRF contract.
var DappLinkVRFMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dappLinkAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fulfillRandomWords\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getRequestStatus\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_dappLinkAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lastRequestId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requestIds\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestMapping\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestRandomWords\",\"inputs\":[{\"name\":\"_requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_numWords\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDappLink\",\"inputs\":[{\"name\":\"_dappLinkAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FillRandomWords\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"randomWords\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestSent\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_numWords\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"current\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x6080806040523460aa575f516020610a4f5f395f51905f525460ff8160401c16609b576002600160401b03196001600160401b038216016049575b6040516109a090816100af8239f35b6001600160401b0319166001600160401b039081175f516020610a4f5f395f51905f525581527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80603a565b63f92ee8a960e01b5f5260045ffd5b5f80fdfe60806040526004361015610011575f80fd5b5f3560e01c80631b739ef11461065857806338ba461414610467578063485cc955146102f2578063715018a61461028b57806382e215ab1461025c5780638796ba8c146102255780638da5cb5b146101f1578063996869d0146101ae578063d8a4676f14610111578063f0c28a41146100e9578063f2fde38b146100be5763fc2a88c31461009d575f80fd5b346100ba575f3660031901126100ba576020600154604051908152f35b5f80fd5b346100ba5760203660031901126100ba576100e76100da6107e8565b6100e26108cc565b61085b565b005b346100ba575f3660031901126100ba576002546040516001600160a01b039091168152602090f35b346100ba5760203660031901126100ba57600435805f52600360205260ff60405f205416905f526003602052600160405f200190604051808360208295549384815201905f5260205f20925f5b818110610195575050610173925003836107c6565b61019160405192839215158352604060208401526040830190610828565b0390f35b845483526001948501948794506020909301920161015e565b346100ba5760203660031901126100ba576101c76107e8565b6101cf6108cc565b600280546001600160a01b0319166001600160a01b0392909216919091179055005b346100ba575f3660031901126100ba575f51602061092b5f395f51905f52546040516001600160a01b039091168152602090f35b346100ba5760203660031901126100ba576004355f548110156100ba5761024d6020916107fe565b90549060031b1c604051908152f35b346100ba5760203660031901126100ba576004355f526003602052602060ff60405f2054166040519015158152f35b346100ba575f3660031901126100ba576102a36108cc565b5f51602061092b5f395f51905f5280546001600160a01b031981169091555f906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346100ba5760403660031901126100ba5761030b6107e8565b6024356001600160a01b038116908190036100ba575f51602061094b5f395f51905f525460ff8160401c16159267ffffffffffffffff82168015908161045f575b6001149081610455575b15908161044c575b5061043d5767ffffffffffffffff1982166001175f51602061094b5f395f51905f525561039c9184610411575b506103946108ff565b6100e26108ff565b6bffffffffffffffffffffffff60a01b60025416176002556103ba57005b68ff0000000000000000195f51602061094b5f395f51905f5254165f51602061094b5f395f51905f52557fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602060405160018152a1005b68ffffffffffffffffff191668010000000000000001175f51602061094b5f395f51905f52558461038b565b63f92ee8a960e01b5f5260045ffd5b9050158561035e565b303b159150610356565b85915061034c565b346100ba5760403660031901126100ba5760043560243567ffffffffffffffff81116100ba57366023820112156100ba5780600401359067ffffffffffffffff82116105ef578160051b60208101926104c360405194856107c6565b835260246020840191830101913683116100ba57602401905b8282106106485750506002546001600160a01b0316330390506106035760405191610506836107aa565b6001835260016020840193838552825f52600360205260405f209051151560ff8019835416911617815501925180519067ffffffffffffffff82116105ef57600160401b82116105ef5784548286558083106105c9575b50602001935f5260205f205f5b8281106105b5577ff3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a85856105b06040519283928352604060208401526040830190610828565b0390a1005b60019060208751970196818401550161056a565b855f528260205f2091820191015b8181106105e4575061055d565b5f81556001016105d7565b634e487b7160e01b5f52604160045260245ffd5b60405162461bcd60e51b815260206004820152601860248201527f446170704c696e6b5652462e6f6e6c79446170704c696e6b00000000000000006044820152606490fd5b81358152602091820191016104dc565b346100ba5760403660031901126100ba576004356106746108cc565b60209060405161068483826107c6565b5f81525f3681376001604051610699816107aa565b5f8152848101928352835f526003855260405f209051151560ff801983541691161781550190519081519167ffffffffffffffff83116105ef57600160401b83116105ef5784908254848455808510610782575b5001905f52835f205f5b8381106107705785855f5490600160401b8210156105ef577fe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f01692610742836001606095015f556107fe565b81549060031b9084821b915f19901b19161790558160015560405191825260243590820152306040820152a1005b825182820155918501916001016106f7565b835f5284835f2091820191015b81811061079c57506106ed565b5f815587935060010161078f565b6040810190811067ffffffffffffffff8211176105ef57604052565b90601f8019910116810190811067ffffffffffffffff8211176105ef57604052565b600435906001600160a01b03821682036100ba57565b5f54811015610814575f805260205f2001905f90565b634e487b7160e01b5f52603260045260245ffd5b90602080835192838152019201905f5b8181106108455750505090565b8251845260209384019390920191600101610838565b6001600160a01b031680156108b9575f51602061092b5f395f51905f5280546001600160a01b0319811683179091556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b631e4fbdf760e01b5f525f60045260245ffd5b5f51602061092b5f395f51905f52546001600160a01b031633036108ec57565b63118cdaa760e01b5f523360045260245ffd5b60ff5f51602061094b5f395f51905f525460401c161561091b57565b631afcd79f60e31b5f5260045ffdfe9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a264697066735822122034719ee34b239d40ee2a7115b46be1ad2ea79485e01484b31e25456523a496c364736f6c634300081c0033f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00",
}

// DappLinkVRFABI is the input ABI used to generate the binding from.
// Deprecated: Use DappLinkVRFMetaData.ABI instead.
var DappLinkVRFABI = DappLinkVRFMetaData.ABI

// DappLinkVRFBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DappLinkVRFMetaData.Bin instead.
var DappLinkVRFBin = DappLinkVRFMetaData.Bin

// DeployDappLinkVRF deploys a new Ethereum contract, binding an instance of DappLinkVRF to it.
func DeployDappLinkVRF(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DappLinkVRF, error) {
	parsed, err := DappLinkVRFMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DappLinkVRFBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DappLinkVRF{DappLinkVRFCaller: DappLinkVRFCaller{contract: contract}, DappLinkVRFTransactor: DappLinkVRFTransactor{contract: contract}, DappLinkVRFFilterer: DappLinkVRFFilterer{contract: contract}}, nil
}

// DappLinkVRF is an auto generated Go binding around an Ethereum contract.
type DappLinkVRF struct {
	DappLinkVRFCaller     // Read-only binding to the contract
	DappLinkVRFTransactor // Write-only binding to the contract
	DappLinkVRFFilterer   // Log filterer for contract events
}

// DappLinkVRFCaller is an auto generated read-only Go binding around an Ethereum contract.
type DappLinkVRFCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappLinkVRFTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DappLinkVRFTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappLinkVRFFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DappLinkVRFFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DappLinkVRFSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DappLinkVRFSession struct {
	Contract     *DappLinkVRF      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DappLinkVRFCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DappLinkVRFCallerSession struct {
	Contract *DappLinkVRFCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DappLinkVRFTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DappLinkVRFTransactorSession struct {
	Contract     *DappLinkVRFTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DappLinkVRFRaw is an auto generated low-level Go binding around an Ethereum contract.
type DappLinkVRFRaw struct {
	Contract *DappLinkVRF // Generic contract binding to access the raw methods on
}

// DappLinkVRFCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DappLinkVRFCallerRaw struct {
	Contract *DappLinkVRFCaller // Generic read-only contract binding to access the raw methods on
}

// DappLinkVRFTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DappLinkVRFTransactorRaw struct {
	Contract *DappLinkVRFTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDappLinkVRF creates a new instance of DappLinkVRF, bound to a specific deployed contract.
func NewDappLinkVRF(address common.Address, backend bind.ContractBackend) (*DappLinkVRF, error) {
	contract, err := bindDappLinkVRF(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DappLinkVRF{DappLinkVRFCaller: DappLinkVRFCaller{contract: contract}, DappLinkVRFTransactor: DappLinkVRFTransactor{contract: contract}, DappLinkVRFFilterer: DappLinkVRFFilterer{contract: contract}}, nil
}

// NewDappLinkVRFCaller creates a new read-only instance of DappLinkVRF, bound to a specific deployed contract.
func NewDappLinkVRFCaller(address common.Address, caller bind.ContractCaller) (*DappLinkVRFCaller, error) {
	contract, err := bindDappLinkVRF(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFCaller{contract: contract}, nil
}

// NewDappLinkVRFTransactor creates a new write-only instance of DappLinkVRF, bound to a specific deployed contract.
func NewDappLinkVRFTransactor(address common.Address, transactor bind.ContractTransactor) (*DappLinkVRFTransactor, error) {
	contract, err := bindDappLinkVRF(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFTransactor{contract: contract}, nil
}

// NewDappLinkVRFFilterer creates a new log filterer instance of DappLinkVRF, bound to a specific deployed contract.
func NewDappLinkVRFFilterer(address common.Address, filterer bind.ContractFilterer) (*DappLinkVRFFilterer, error) {
	contract, err := bindDappLinkVRF(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFFilterer{contract: contract}, nil
}

// bindDappLinkVRF binds a generic wrapper to an already deployed contract.
func bindDappLinkVRF(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DappLinkVRFMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DappLinkVRF *DappLinkVRFRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DappLinkVRF.Contract.DappLinkVRFCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DappLinkVRF *DappLinkVRFRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappLinkVRF.Contract.DappLinkVRFTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DappLinkVRF *DappLinkVRFRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DappLinkVRF.Contract.DappLinkVRFTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DappLinkVRF *DappLinkVRFCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DappLinkVRF.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DappLinkVRF *DappLinkVRFTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappLinkVRF.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DappLinkVRF *DappLinkVRFTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DappLinkVRF.Contract.contract.Transact(opts, method, params...)
}

// DappLinkAddress is a free data retrieval call binding the contract method 0xf0c28a41.
//
// Solidity: function dappLinkAddress() view returns(address)
func (_DappLinkVRF *DappLinkVRFCaller) DappLinkAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappLinkVRF.contract.Call(opts, &out, "dappLinkAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DappLinkAddress is a free data retrieval call binding the contract method 0xf0c28a41.
//
// Solidity: function dappLinkAddress() view returns(address)
func (_DappLinkVRF *DappLinkVRFSession) DappLinkAddress() (common.Address, error) {
	return _DappLinkVRF.Contract.DappLinkAddress(&_DappLinkVRF.CallOpts)
}

// DappLinkAddress is a free data retrieval call binding the contract method 0xf0c28a41.
//
// Solidity: function dappLinkAddress() view returns(address)
func (_DappLinkVRF *DappLinkVRFCallerSession) DappLinkAddress() (common.Address, error) {
	return _DappLinkVRF.Contract.DappLinkAddress(&_DappLinkVRF.CallOpts)
}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, uint256[] randomWords)
func (_DappLinkVRF *DappLinkVRFCaller) GetRequestStatus(opts *bind.CallOpts, _requestId *big.Int) (struct {
	Fulfilled   bool
	RandomWords []*big.Int
}, error) {
	var out []interface{}
	err := _DappLinkVRF.contract.Call(opts, &out, "getRequestStatus", _requestId)

	outstruct := new(struct {
		Fulfilled   bool
		RandomWords []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fulfilled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.RandomWords = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, uint256[] randomWords)
func (_DappLinkVRF *DappLinkVRFSession) GetRequestStatus(_requestId *big.Int) (struct {
	Fulfilled   bool
	RandomWords []*big.Int
}, error) {
	return _DappLinkVRF.Contract.GetRequestStatus(&_DappLinkVRF.CallOpts, _requestId)
}

// GetRequestStatus is a free data retrieval call binding the contract method 0xd8a4676f.
//
// Solidity: function getRequestStatus(uint256 _requestId) view returns(bool fulfilled, uint256[] randomWords)
func (_DappLinkVRF *DappLinkVRFCallerSession) GetRequestStatus(_requestId *big.Int) (struct {
	Fulfilled   bool
	RandomWords []*big.Int
}, error) {
	return _DappLinkVRF.Contract.GetRequestStatus(&_DappLinkVRF.CallOpts, _requestId)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_DappLinkVRF *DappLinkVRFCaller) LastRequestId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DappLinkVRF.contract.Call(opts, &out, "lastRequestId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_DappLinkVRF *DappLinkVRFSession) LastRequestId() (*big.Int, error) {
	return _DappLinkVRF.Contract.LastRequestId(&_DappLinkVRF.CallOpts)
}

// LastRequestId is a free data retrieval call binding the contract method 0xfc2a88c3.
//
// Solidity: function lastRequestId() view returns(uint256)
func (_DappLinkVRF *DappLinkVRFCallerSession) LastRequestId() (*big.Int, error) {
	return _DappLinkVRF.Contract.LastRequestId(&_DappLinkVRF.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappLinkVRF *DappLinkVRFCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DappLinkVRF.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappLinkVRF *DappLinkVRFSession) Owner() (common.Address, error) {
	return _DappLinkVRF.Contract.Owner(&_DappLinkVRF.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DappLinkVRF *DappLinkVRFCallerSession) Owner() (common.Address, error) {
	return _DappLinkVRF.Contract.Owner(&_DappLinkVRF.CallOpts)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_DappLinkVRF *DappLinkVRFCaller) RequestIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DappLinkVRF.contract.Call(opts, &out, "requestIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_DappLinkVRF *DappLinkVRFSession) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _DappLinkVRF.Contract.RequestIds(&_DappLinkVRF.CallOpts, arg0)
}

// RequestIds is a free data retrieval call binding the contract method 0x8796ba8c.
//
// Solidity: function requestIds(uint256 ) view returns(uint256)
func (_DappLinkVRF *DappLinkVRFCallerSession) RequestIds(arg0 *big.Int) (*big.Int, error) {
	return _DappLinkVRF.Contract.RequestIds(&_DappLinkVRF.CallOpts, arg0)
}

// RequestMapping is a free data retrieval call binding the contract method 0x82e215ab.
//
// Solidity: function requestMapping(uint256 ) view returns(bool fulfilled)
func (_DappLinkVRF *DappLinkVRFCaller) RequestMapping(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _DappLinkVRF.contract.Call(opts, &out, "requestMapping", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequestMapping is a free data retrieval call binding the contract method 0x82e215ab.
//
// Solidity: function requestMapping(uint256 ) view returns(bool fulfilled)
func (_DappLinkVRF *DappLinkVRFSession) RequestMapping(arg0 *big.Int) (bool, error) {
	return _DappLinkVRF.Contract.RequestMapping(&_DappLinkVRF.CallOpts, arg0)
}

// RequestMapping is a free data retrieval call binding the contract method 0x82e215ab.
//
// Solidity: function requestMapping(uint256 ) view returns(bool fulfilled)
func (_DappLinkVRF *DappLinkVRFCallerSession) RequestMapping(arg0 *big.Int) (bool, error) {
	return _DappLinkVRF.Contract.RequestMapping(&_DappLinkVRF.CallOpts, arg0)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 _requestId, uint256[] _randomWords) returns()
func (_DappLinkVRF *DappLinkVRFTransactor) FulfillRandomWords(opts *bind.TransactOpts, _requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _DappLinkVRF.contract.Transact(opts, "fulfillRandomWords", _requestId, _randomWords)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 _requestId, uint256[] _randomWords) returns()
func (_DappLinkVRF *DappLinkVRFSession) FulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _DappLinkVRF.Contract.FulfillRandomWords(&_DappLinkVRF.TransactOpts, _requestId, _randomWords)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x38ba4614.
//
// Solidity: function fulfillRandomWords(uint256 _requestId, uint256[] _randomWords) returns()
func (_DappLinkVRF *DappLinkVRFTransactorSession) FulfillRandomWords(_requestId *big.Int, _randomWords []*big.Int) (*types.Transaction, error) {
	return _DappLinkVRF.Contract.FulfillRandomWords(&_DappLinkVRF.TransactOpts, _requestId, _randomWords)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address _dappLinkAddress) returns()
func (_DappLinkVRF *DappLinkVRFTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address, _dappLinkAddress common.Address) (*types.Transaction, error) {
	return _DappLinkVRF.contract.Transact(opts, "initialize", initialOwner, _dappLinkAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address _dappLinkAddress) returns()
func (_DappLinkVRF *DappLinkVRFSession) Initialize(initialOwner common.Address, _dappLinkAddress common.Address) (*types.Transaction, error) {
	return _DappLinkVRF.Contract.Initialize(&_DappLinkVRF.TransactOpts, initialOwner, _dappLinkAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address initialOwner, address _dappLinkAddress) returns()
func (_DappLinkVRF *DappLinkVRFTransactorSession) Initialize(initialOwner common.Address, _dappLinkAddress common.Address) (*types.Transaction, error) {
	return _DappLinkVRF.Contract.Initialize(&_DappLinkVRF.TransactOpts, initialOwner, _dappLinkAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DappLinkVRF *DappLinkVRFTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DappLinkVRF.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DappLinkVRF *DappLinkVRFSession) RenounceOwnership() (*types.Transaction, error) {
	return _DappLinkVRF.Contract.RenounceOwnership(&_DappLinkVRF.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DappLinkVRF *DappLinkVRFTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DappLinkVRF.Contract.RenounceOwnership(&_DappLinkVRF.TransactOpts)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x1b739ef1.
//
// Solidity: function requestRandomWords(uint256 _requestId, uint256 _numWords) returns()
func (_DappLinkVRF *DappLinkVRFTransactor) RequestRandomWords(opts *bind.TransactOpts, _requestId *big.Int, _numWords *big.Int) (*types.Transaction, error) {
	return _DappLinkVRF.contract.Transact(opts, "requestRandomWords", _requestId, _numWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x1b739ef1.
//
// Solidity: function requestRandomWords(uint256 _requestId, uint256 _numWords) returns()
func (_DappLinkVRF *DappLinkVRFSession) RequestRandomWords(_requestId *big.Int, _numWords *big.Int) (*types.Transaction, error) {
	return _DappLinkVRF.Contract.RequestRandomWords(&_DappLinkVRF.TransactOpts, _requestId, _numWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x1b739ef1.
//
// Solidity: function requestRandomWords(uint256 _requestId, uint256 _numWords) returns()
func (_DappLinkVRF *DappLinkVRFTransactorSession) RequestRandomWords(_requestId *big.Int, _numWords *big.Int) (*types.Transaction, error) {
	return _DappLinkVRF.Contract.RequestRandomWords(&_DappLinkVRF.TransactOpts, _requestId, _numWords)
}

// SetDappLink is a paid mutator transaction binding the contract method 0x996869d0.
//
// Solidity: function setDappLink(address _dappLinkAddress) returns()
func (_DappLinkVRF *DappLinkVRFTransactor) SetDappLink(opts *bind.TransactOpts, _dappLinkAddress common.Address) (*types.Transaction, error) {
	return _DappLinkVRF.contract.Transact(opts, "setDappLink", _dappLinkAddress)
}

// SetDappLink is a paid mutator transaction binding the contract method 0x996869d0.
//
// Solidity: function setDappLink(address _dappLinkAddress) returns()
func (_DappLinkVRF *DappLinkVRFSession) SetDappLink(_dappLinkAddress common.Address) (*types.Transaction, error) {
	return _DappLinkVRF.Contract.SetDappLink(&_DappLinkVRF.TransactOpts, _dappLinkAddress)
}

// SetDappLink is a paid mutator transaction binding the contract method 0x996869d0.
//
// Solidity: function setDappLink(address _dappLinkAddress) returns()
func (_DappLinkVRF *DappLinkVRFTransactorSession) SetDappLink(_dappLinkAddress common.Address) (*types.Transaction, error) {
	return _DappLinkVRF.Contract.SetDappLink(&_DappLinkVRF.TransactOpts, _dappLinkAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DappLinkVRF *DappLinkVRFTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DappLinkVRF.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DappLinkVRF *DappLinkVRFSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DappLinkVRF.Contract.TransferOwnership(&_DappLinkVRF.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DappLinkVRF *DappLinkVRFTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DappLinkVRF.Contract.TransferOwnership(&_DappLinkVRF.TransactOpts, newOwner)
}

// DappLinkVRFFillRandomWordsIterator is returned from FilterFillRandomWords and is used to iterate over the raw logs and unpacked data for FillRandomWords events raised by the DappLinkVRF contract.
type DappLinkVRFFillRandomWordsIterator struct {
	Event *DappLinkVRFFillRandomWords // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DappLinkVRFFillRandomWordsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappLinkVRFFillRandomWords)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DappLinkVRFFillRandomWords)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DappLinkVRFFillRandomWordsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappLinkVRFFillRandomWordsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappLinkVRFFillRandomWords represents a FillRandomWords event raised by the DappLinkVRF contract.
type DappLinkVRFFillRandomWords struct {
	RequestId   *big.Int
	RandomWords []*big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFillRandomWords is a free log retrieval operation binding the contract event 0xf3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a.
//
// Solidity: event FillRandomWords(uint256 requestId, uint256[] randomWords)
func (_DappLinkVRF *DappLinkVRFFilterer) FilterFillRandomWords(opts *bind.FilterOpts) (*DappLinkVRFFillRandomWordsIterator, error) {

	logs, sub, err := _DappLinkVRF.contract.FilterLogs(opts, "FillRandomWords")
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFFillRandomWordsIterator{contract: _DappLinkVRF.contract, event: "FillRandomWords", logs: logs, sub: sub}, nil
}

// WatchFillRandomWords is a free log subscription operation binding the contract event 0xf3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a.
//
// Solidity: event FillRandomWords(uint256 requestId, uint256[] randomWords)
func (_DappLinkVRF *DappLinkVRFFilterer) WatchFillRandomWords(opts *bind.WatchOpts, sink chan<- *DappLinkVRFFillRandomWords) (event.Subscription, error) {

	logs, sub, err := _DappLinkVRF.contract.WatchLogs(opts, "FillRandomWords")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappLinkVRFFillRandomWords)
				if err := _DappLinkVRF.contract.UnpackLog(event, "FillRandomWords", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFillRandomWords is a log parse operation binding the contract event 0xf3cb4deb0441dd096356debf166f879d78cadc19e4b94053c8bea6d3940de93a.
//
// Solidity: event FillRandomWords(uint256 requestId, uint256[] randomWords)
func (_DappLinkVRF *DappLinkVRFFilterer) ParseFillRandomWords(log types.Log) (*DappLinkVRFFillRandomWords, error) {
	event := new(DappLinkVRFFillRandomWords)
	if err := _DappLinkVRF.contract.UnpackLog(event, "FillRandomWords", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DappLinkVRFInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DappLinkVRF contract.
type DappLinkVRFInitializedIterator struct {
	Event *DappLinkVRFInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DappLinkVRFInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappLinkVRFInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DappLinkVRFInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DappLinkVRFInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappLinkVRFInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappLinkVRFInitialized represents a Initialized event raised by the DappLinkVRF contract.
type DappLinkVRFInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DappLinkVRF *DappLinkVRFFilterer) FilterInitialized(opts *bind.FilterOpts) (*DappLinkVRFInitializedIterator, error) {

	logs, sub, err := _DappLinkVRF.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFInitializedIterator{contract: _DappLinkVRF.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DappLinkVRF *DappLinkVRFFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DappLinkVRFInitialized) (event.Subscription, error) {

	logs, sub, err := _DappLinkVRF.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappLinkVRFInitialized)
				if err := _DappLinkVRF.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_DappLinkVRF *DappLinkVRFFilterer) ParseInitialized(log types.Log) (*DappLinkVRFInitialized, error) {
	event := new(DappLinkVRFInitialized)
	if err := _DappLinkVRF.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DappLinkVRFOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DappLinkVRF contract.
type DappLinkVRFOwnershipTransferredIterator struct {
	Event *DappLinkVRFOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DappLinkVRFOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappLinkVRFOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DappLinkVRFOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DappLinkVRFOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappLinkVRFOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappLinkVRFOwnershipTransferred represents a OwnershipTransferred event raised by the DappLinkVRF contract.
type DappLinkVRFOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DappLinkVRF *DappLinkVRFFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DappLinkVRFOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DappLinkVRF.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFOwnershipTransferredIterator{contract: _DappLinkVRF.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DappLinkVRF *DappLinkVRFFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DappLinkVRFOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DappLinkVRF.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappLinkVRFOwnershipTransferred)
				if err := _DappLinkVRF.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DappLinkVRF *DappLinkVRFFilterer) ParseOwnershipTransferred(log types.Log) (*DappLinkVRFOwnershipTransferred, error) {
	event := new(DappLinkVRFOwnershipTransferred)
	if err := _DappLinkVRF.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DappLinkVRFRequestSentIterator is returned from FilterRequestSent and is used to iterate over the raw logs and unpacked data for RequestSent events raised by the DappLinkVRF contract.
type DappLinkVRFRequestSentIterator struct {
	Event *DappLinkVRFRequestSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DappLinkVRFRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DappLinkVRFRequestSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DappLinkVRFRequestSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DappLinkVRFRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DappLinkVRFRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DappLinkVRFRequestSent represents a RequestSent event raised by the DappLinkVRF contract.
type DappLinkVRFRequestSent struct {
	RequestId *big.Int
	NumWords  *big.Int
	Current   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestSent is a free log retrieval operation binding the contract event 0xe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f016.
//
// Solidity: event RequestSent(uint256 requestId, uint256 _numWords, address current)
func (_DappLinkVRF *DappLinkVRFFilterer) FilterRequestSent(opts *bind.FilterOpts) (*DappLinkVRFRequestSentIterator, error) {

	logs, sub, err := _DappLinkVRF.contract.FilterLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return &DappLinkVRFRequestSentIterator{contract: _DappLinkVRF.contract, event: "RequestSent", logs: logs, sub: sub}, nil
}

// WatchRequestSent is a free log subscription operation binding the contract event 0xe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f016.
//
// Solidity: event RequestSent(uint256 requestId, uint256 _numWords, address current)
func (_DappLinkVRF *DappLinkVRFFilterer) WatchRequestSent(opts *bind.WatchOpts, sink chan<- *DappLinkVRFRequestSent) (event.Subscription, error) {

	logs, sub, err := _DappLinkVRF.contract.WatchLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DappLinkVRFRequestSent)
				if err := _DappLinkVRF.contract.UnpackLog(event, "RequestSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRequestSent is a log parse operation binding the contract event 0xe697eb68c0228bd7d4e553246a2a86e8402d0895e45092ef8ae87b4cfd29f016.
//
// Solidity: event RequestSent(uint256 requestId, uint256 _numWords, address current)
func (_DappLinkVRF *DappLinkVRFFilterer) ParseRequestSent(log types.Log) (*DappLinkVRFRequestSent, error) {
	event := new(DappLinkVRFRequestSent)
	if err := _DappLinkVRF.contract.UnpackLog(event, "RequestSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
