// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// MyStakeClientInfo is an auto generated low-level Go binding around an user-defined struct.
type MyStakeClientInfo struct {
	Ts     *big.Int
	Result string
}

// MyStakeStake is an auto generated low-level Go binding around an user-defined struct.
type MyStakeStake struct {
	DepositAmount           *big.Int
	StakeCreationTime       *big.Int
	LastUploadTime          *big.Int
	Returned                bool
	AlreadyWithdrawedAmount *big.Int
	AllUploadCount          *big.Int
	AllRewardAmount         *big.Int
	RewardAmount            *big.Int
	RewardUpload            *big.Int
	AllReturnAmount         *big.Int
	WalletAddr              common.Address
	LastPw                  *big.Int
	AlldayPw                *big.Int
	IsUsed                  bool
}

// PoolABI is the input ABI used to generate the binding from.
const PoolABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"alpha\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"bindTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"burnFee\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"cashout\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swarmContract\",\"type\":\"address\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"depositStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getAllInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"deposit_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake_creation_time\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"last_upload_time\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"returned\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"alreadyWithdrawedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"AllUploadCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"AllRewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardUpload\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"AllReturnAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"walletAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lastPw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"alldayPw\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isUsed\",\"type\":\"bool\"}],\"internalType\":\"structMyStake.Stake\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getAllUploadCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllUser\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getClientInfos\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ts\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"internalType\":\"structMyStake.ClientInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"swarmContract\",\"type\":\"address\"}],\"name\":\"getNCTRContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalDepositAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUploadCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUploadTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getallpw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rp\",\"type\":\"uint256\"}],\"name\":\"setAlpha\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ga\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_uploadCount\",\"type\":\"uint256\"}],\"name\":\"setUploadCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_uploadTime\",\"type\":\"uint256\"}],\"name\":\"setUploadTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"totalCashout\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"totalReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ts\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tsStr\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"sign\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_hw\",\"type\":\"uint256\"}],\"name\":\"uploadInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// Alpha is a free data retrieval call binding the contract method 0xdb1d0fd5.
//
// Solidity: function alpha() view returns(uint256)
func (_Pool *PoolCaller) Alpha(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "alpha")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Alpha is a free data retrieval call binding the contract method 0xdb1d0fd5.
//
// Solidity: function alpha() view returns(uint256)
func (_Pool *PoolSession) Alpha() (*big.Int, error) {
	return _Pool.Contract.Alpha(&_Pool.CallOpts)
}

// Alpha is a free data retrieval call binding the contract method 0xdb1d0fd5.
//
// Solidity: function alpha() view returns(uint256)
func (_Pool *PoolCallerSession) Alpha() (*big.Int, error) {
	return _Pool.Contract.Alpha(&_Pool.CallOpts)
}

// DepositStatus is a free data retrieval call binding the contract method 0xf3205dfa.
//
// Solidity: function depositStatus(address _addr) view returns(bool)
func (_Pool *PoolCaller) DepositStatus(opts *bind.CallOpts, _addr common.Address) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "depositStatus", _addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositStatus is a free data retrieval call binding the contract method 0xf3205dfa.
//
// Solidity: function depositStatus(address _addr) view returns(bool)
func (_Pool *PoolSession) DepositStatus(_addr common.Address) (bool, error) {
	return _Pool.Contract.DepositStatus(&_Pool.CallOpts, _addr)
}

// DepositStatus is a free data retrieval call binding the contract method 0xf3205dfa.
//
// Solidity: function depositStatus(address _addr) view returns(bool)
func (_Pool *PoolCallerSession) DepositStatus(_addr common.Address) (bool, error) {
	return _Pool.Contract.DepositStatus(&_Pool.CallOpts, _addr)
}

// FeeAmount is a free data retrieval call binding the contract method 0x69e15404.
//
// Solidity: function feeAmount() view returns(uint256)
func (_Pool *PoolCaller) FeeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "feeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeAmount is a free data retrieval call binding the contract method 0x69e15404.
//
// Solidity: function feeAmount() view returns(uint256)
func (_Pool *PoolSession) FeeAmount() (*big.Int, error) {
	return _Pool.Contract.FeeAmount(&_Pool.CallOpts)
}

// FeeAmount is a free data retrieval call binding the contract method 0x69e15404.
//
// Solidity: function feeAmount() view returns(uint256)
func (_Pool *PoolCallerSession) FeeAmount() (*big.Int, error) {
	return _Pool.Contract.FeeAmount(&_Pool.CallOpts)
}

// GetAllInfo is a free data retrieval call binding the contract method 0x7e16c30c.
//
// Solidity: function getAllInfo(address _addr) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,bool))
func (_Pool *PoolCaller) GetAllInfo(opts *bind.CallOpts, _addr common.Address) (MyStakeStake, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getAllInfo", _addr)

	if err != nil {
		return *new(MyStakeStake), err
	}

	out0 := *abi.ConvertType(out[0], new(MyStakeStake)).(*MyStakeStake)

	return out0, err

}

// GetAllInfo is a free data retrieval call binding the contract method 0x7e16c30c.
//
// Solidity: function getAllInfo(address _addr) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,bool))
func (_Pool *PoolSession) GetAllInfo(_addr common.Address) (MyStakeStake, error) {
	return _Pool.Contract.GetAllInfo(&_Pool.CallOpts, _addr)
}

// GetAllInfo is a free data retrieval call binding the contract method 0x7e16c30c.
//
// Solidity: function getAllInfo(address _addr) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,bool))
func (_Pool *PoolCallerSession) GetAllInfo(_addr common.Address) (MyStakeStake, error) {
	return _Pool.Contract.GetAllInfo(&_Pool.CallOpts, _addr)
}

// GetAllUploadCount is a free data retrieval call binding the contract method 0xb3dce268.
//
// Solidity: function getAllUploadCount(address _addr) view returns(uint256)
func (_Pool *PoolCaller) GetAllUploadCount(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getAllUploadCount", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAllUploadCount is a free data retrieval call binding the contract method 0xb3dce268.
//
// Solidity: function getAllUploadCount(address _addr) view returns(uint256)
func (_Pool *PoolSession) GetAllUploadCount(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.GetAllUploadCount(&_Pool.CallOpts, _addr)
}

// GetAllUploadCount is a free data retrieval call binding the contract method 0xb3dce268.
//
// Solidity: function getAllUploadCount(address _addr) view returns(uint256)
func (_Pool *PoolCallerSession) GetAllUploadCount(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.GetAllUploadCount(&_Pool.CallOpts, _addr)
}

// GetAllUser is a free data retrieval call binding the contract method 0x47531df8.
//
// Solidity: function getAllUser() view returns(address[])
func (_Pool *PoolCaller) GetAllUser(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getAllUser")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllUser is a free data retrieval call binding the contract method 0x47531df8.
//
// Solidity: function getAllUser() view returns(address[])
func (_Pool *PoolSession) GetAllUser() ([]common.Address, error) {
	return _Pool.Contract.GetAllUser(&_Pool.CallOpts)
}

// GetAllUser is a free data retrieval call binding the contract method 0x47531df8.
//
// Solidity: function getAllUser() view returns(address[])
func (_Pool *PoolCallerSession) GetAllUser() ([]common.Address, error) {
	return _Pool.Contract.GetAllUser(&_Pool.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Pool *PoolCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Pool *PoolSession) GetBalance() (*big.Int, error) {
	return _Pool.Contract.GetBalance(&_Pool.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Pool *PoolCallerSession) GetBalance() (*big.Int, error) {
	return _Pool.Contract.GetBalance(&_Pool.CallOpts)
}

// GetClientInfos is a free data retrieval call binding the contract method 0xac4ec70e.
//
// Solidity: function getClientInfos(address _addr) view returns((uint256,string)[])
func (_Pool *PoolCaller) GetClientInfos(opts *bind.CallOpts, _addr common.Address) ([]MyStakeClientInfo, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getClientInfos", _addr)

	if err != nil {
		return *new([]MyStakeClientInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]MyStakeClientInfo)).(*[]MyStakeClientInfo)

	return out0, err

}

// GetClientInfos is a free data retrieval call binding the contract method 0xac4ec70e.
//
// Solidity: function getClientInfos(address _addr) view returns((uint256,string)[])
func (_Pool *PoolSession) GetClientInfos(_addr common.Address) ([]MyStakeClientInfo, error) {
	return _Pool.Contract.GetClientInfos(&_Pool.CallOpts, _addr)
}

// GetClientInfos is a free data retrieval call binding the contract method 0xac4ec70e.
//
// Solidity: function getClientInfos(address _addr) view returns((uint256,string)[])
func (_Pool *PoolCallerSession) GetClientInfos(_addr common.Address) ([]MyStakeClientInfo, error) {
	return _Pool.Contract.GetClientInfos(&_Pool.CallOpts, _addr)
}

// GetDepositAmount is a free data retrieval call binding the contract method 0xb8ba16fd.
//
// Solidity: function getDepositAmount(address _addr) view returns(uint256)
func (_Pool *PoolCaller) GetDepositAmount(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getDepositAmount", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositAmount is a free data retrieval call binding the contract method 0xb8ba16fd.
//
// Solidity: function getDepositAmount(address _addr) view returns(uint256)
func (_Pool *PoolSession) GetDepositAmount(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.GetDepositAmount(&_Pool.CallOpts, _addr)
}

// GetDepositAmount is a free data retrieval call binding the contract method 0xb8ba16fd.
//
// Solidity: function getDepositAmount(address _addr) view returns(uint256)
func (_Pool *PoolCallerSession) GetDepositAmount(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.GetDepositAmount(&_Pool.CallOpts, _addr)
}

// GetNCTRContract is a free data retrieval call binding the contract method 0x36666f7a.
//
// Solidity: function getNCTRContract(address swarmContract) view returns(address)
func (_Pool *PoolCaller) GetNCTRContract(opts *bind.CallOpts, swarmContract common.Address) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getNCTRContract", swarmContract)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNCTRContract is a free data retrieval call binding the contract method 0x36666f7a.
//
// Solidity: function getNCTRContract(address swarmContract) view returns(address)
func (_Pool *PoolSession) GetNCTRContract(swarmContract common.Address) (common.Address, error) {
	return _Pool.Contract.GetNCTRContract(&_Pool.CallOpts, swarmContract)
}

// GetNCTRContract is a free data retrieval call binding the contract method 0x36666f7a.
//
// Solidity: function getNCTRContract(address swarmContract) view returns(address)
func (_Pool *PoolCallerSession) GetNCTRContract(swarmContract common.Address) (common.Address, error) {
	return _Pool.Contract.GetNCTRContract(&_Pool.CallOpts, swarmContract)
}

// GetRewardAmount is a free data retrieval call binding the contract method 0x44a040f5.
//
// Solidity: function getRewardAmount(address _addr) view returns(uint256)
func (_Pool *PoolCaller) GetRewardAmount(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getRewardAmount", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardAmount is a free data retrieval call binding the contract method 0x44a040f5.
//
// Solidity: function getRewardAmount(address _addr) view returns(uint256)
func (_Pool *PoolSession) GetRewardAmount(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.GetRewardAmount(&_Pool.CallOpts, _addr)
}

// GetRewardAmount is a free data retrieval call binding the contract method 0x44a040f5.
//
// Solidity: function getRewardAmount(address _addr) view returns(uint256)
func (_Pool *PoolCallerSession) GetRewardAmount(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.GetRewardAmount(&_Pool.CallOpts, _addr)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Pool *PoolCaller) GetTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getTokenAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Pool *PoolSession) GetTokenAddress() (common.Address, error) {
	return _Pool.Contract.GetTokenAddress(&_Pool.CallOpts)
}

// GetTokenAddress is a free data retrieval call binding the contract method 0x10fe9ae8.
//
// Solidity: function getTokenAddress() view returns(address)
func (_Pool *PoolCallerSession) GetTokenAddress() (common.Address, error) {
	return _Pool.Contract.GetTokenAddress(&_Pool.CallOpts)
}

// GetTotalDepositAmount is a free data retrieval call binding the contract method 0x3c9b97fc.
//
// Solidity: function getTotalDepositAmount() view returns(uint256)
func (_Pool *PoolCaller) GetTotalDepositAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getTotalDepositAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDepositAmount is a free data retrieval call binding the contract method 0x3c9b97fc.
//
// Solidity: function getTotalDepositAmount() view returns(uint256)
func (_Pool *PoolSession) GetTotalDepositAmount() (*big.Int, error) {
	return _Pool.Contract.GetTotalDepositAmount(&_Pool.CallOpts)
}

// GetTotalDepositAmount is a free data retrieval call binding the contract method 0x3c9b97fc.
//
// Solidity: function getTotalDepositAmount() view returns(uint256)
func (_Pool *PoolCallerSession) GetTotalDepositAmount() (*big.Int, error) {
	return _Pool.Contract.GetTotalDepositAmount(&_Pool.CallOpts)
}

// GetUploadCount is a free data retrieval call binding the contract method 0xca27cf04.
//
// Solidity: function getUploadCount() view returns(uint256)
func (_Pool *PoolCaller) GetUploadCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getUploadCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUploadCount is a free data retrieval call binding the contract method 0xca27cf04.
//
// Solidity: function getUploadCount() view returns(uint256)
func (_Pool *PoolSession) GetUploadCount() (*big.Int, error) {
	return _Pool.Contract.GetUploadCount(&_Pool.CallOpts)
}

// GetUploadCount is a free data retrieval call binding the contract method 0xca27cf04.
//
// Solidity: function getUploadCount() view returns(uint256)
func (_Pool *PoolCallerSession) GetUploadCount() (*big.Int, error) {
	return _Pool.Contract.GetUploadCount(&_Pool.CallOpts)
}

// GetUploadTime is a free data retrieval call binding the contract method 0xa0a58af7.
//
// Solidity: function getUploadTime() view returns(uint256)
func (_Pool *PoolCaller) GetUploadTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getUploadTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUploadTime is a free data retrieval call binding the contract method 0xa0a58af7.
//
// Solidity: function getUploadTime() view returns(uint256)
func (_Pool *PoolSession) GetUploadTime() (*big.Int, error) {
	return _Pool.Contract.GetUploadTime(&_Pool.CallOpts)
}

// GetUploadTime is a free data retrieval call binding the contract method 0xa0a58af7.
//
// Solidity: function getUploadTime() view returns(uint256)
func (_Pool *PoolCallerSession) GetUploadTime() (*big.Int, error) {
	return _Pool.Contract.GetUploadTime(&_Pool.CallOpts)
}

// Getallpw is a free data retrieval call binding the contract method 0xcd53d9be.
//
// Solidity: function getallpw() view returns(uint256)
func (_Pool *PoolCaller) Getallpw(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getallpw")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Getallpw is a free data retrieval call binding the contract method 0xcd53d9be.
//
// Solidity: function getallpw() view returns(uint256)
func (_Pool *PoolSession) Getallpw() (*big.Int, error) {
	return _Pool.Contract.Getallpw(&_Pool.CallOpts)
}

// Getallpw is a free data retrieval call binding the contract method 0xcd53d9be.
//
// Solidity: function getallpw() view returns(uint256)
func (_Pool *PoolCallerSession) Getallpw() (*big.Int, error) {
	return _Pool.Contract.Getallpw(&_Pool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pool *PoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pool *PoolSession) Owner() (common.Address, error) {
	return _Pool.Contract.Owner(&_Pool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pool *PoolCallerSession) Owner() (common.Address, error) {
	return _Pool.Contract.Owner(&_Pool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pool *PoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pool *PoolSession) Paused() (bool, error) {
	return _Pool.Contract.Paused(&_Pool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pool *PoolCallerSession) Paused() (bool, error) {
	return _Pool.Contract.Paused(&_Pool.CallOpts)
}

// TotalCashout is a free data retrieval call binding the contract method 0x0443d346.
//
// Solidity: function totalCashout(address _addr) view returns(uint256)
func (_Pool *PoolCaller) TotalCashout(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "totalCashout", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalCashout is a free data retrieval call binding the contract method 0x0443d346.
//
// Solidity: function totalCashout(address _addr) view returns(uint256)
func (_Pool *PoolSession) TotalCashout(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.TotalCashout(&_Pool.CallOpts, _addr)
}

// TotalCashout is a free data retrieval call binding the contract method 0x0443d346.
//
// Solidity: function totalCashout(address _addr) view returns(uint256)
func (_Pool *PoolCallerSession) TotalCashout(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.TotalCashout(&_Pool.CallOpts, _addr)
}

// TotalReward is a free data retrieval call binding the contract method 0x99908f74.
//
// Solidity: function totalReward(address _addr) view returns(uint256)
func (_Pool *PoolCaller) TotalReward(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "totalReward", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReward is a free data retrieval call binding the contract method 0x99908f74.
//
// Solidity: function totalReward(address _addr) view returns(uint256)
func (_Pool *PoolSession) TotalReward(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.TotalReward(&_Pool.CallOpts, _addr)
}

// TotalReward is a free data retrieval call binding the contract method 0x99908f74.
//
// Solidity: function totalReward(address _addr) view returns(uint256)
func (_Pool *PoolCallerSession) TotalReward(_addr common.Address) (*big.Int, error) {
	return _Pool.Contract.TotalReward(&_Pool.CallOpts, _addr)
}

// BindTokenAddress is a paid mutator transaction binding the contract method 0xba1075f5.
//
// Solidity: function bindTokenAddress(address _tokenAddress) returns()
func (_Pool *PoolTransactor) BindTokenAddress(opts *bind.TransactOpts, _tokenAddress common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "bindTokenAddress", _tokenAddress)
}

// BindTokenAddress is a paid mutator transaction binding the contract method 0xba1075f5.
//
// Solidity: function bindTokenAddress(address _tokenAddress) returns()
func (_Pool *PoolSession) BindTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Pool.Contract.BindTokenAddress(&_Pool.TransactOpts, _tokenAddress)
}

// BindTokenAddress is a paid mutator transaction binding the contract method 0xba1075f5.
//
// Solidity: function bindTokenAddress(address _tokenAddress) returns()
func (_Pool *PoolTransactorSession) BindTokenAddress(_tokenAddress common.Address) (*types.Transaction, error) {
	return _Pool.Contract.BindTokenAddress(&_Pool.TransactOpts, _tokenAddress)
}

// BurnFee is a paid mutator transaction binding the contract method 0xfce589d8.
//
// Solidity: function burnFee() payable returns()
func (_Pool *PoolTransactor) BurnFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "burnFee")
}

// BurnFee is a paid mutator transaction binding the contract method 0xfce589d8.
//
// Solidity: function burnFee() payable returns()
func (_Pool *PoolSession) BurnFee() (*types.Transaction, error) {
	return _Pool.Contract.BurnFee(&_Pool.TransactOpts)
}

// BurnFee is a paid mutator transaction binding the contract method 0xfce589d8.
//
// Solidity: function burnFee() payable returns()
func (_Pool *PoolTransactorSession) BurnFee() (*types.Transaction, error) {
	return _Pool.Contract.BurnFee(&_Pool.TransactOpts)
}

// Cashout is a paid mutator transaction binding the contract method 0xa9e732bb.
//
// Solidity: function cashout(uint256 _amount) payable returns()
func (_Pool *PoolTransactor) Cashout(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "cashout", _amount)
}

// Cashout is a paid mutator transaction binding the contract method 0xa9e732bb.
//
// Solidity: function cashout(uint256 _amount) payable returns()
func (_Pool *PoolSession) Cashout(_amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Cashout(&_Pool.TransactOpts, _amount)
}

// Cashout is a paid mutator transaction binding the contract method 0xa9e732bb.
//
// Solidity: function cashout(uint256 _amount) payable returns()
func (_Pool *PoolTransactorSession) Cashout(_amount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.Cashout(&_Pool.TransactOpts, _amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x2e2d2984.
//
// Solidity: function deposit(uint256 _amount, address walletAddr, address swarmContract) returns()
func (_Pool *PoolTransactor) Deposit(opts *bind.TransactOpts, _amount *big.Int, walletAddr common.Address, swarmContract common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "deposit", _amount, walletAddr, swarmContract)
}

// Deposit is a paid mutator transaction binding the contract method 0x2e2d2984.
//
// Solidity: function deposit(uint256 _amount, address walletAddr, address swarmContract) returns()
func (_Pool *PoolSession) Deposit(_amount *big.Int, walletAddr common.Address, swarmContract common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Deposit(&_Pool.TransactOpts, _amount, walletAddr, swarmContract)
}

// Deposit is a paid mutator transaction binding the contract method 0x2e2d2984.
//
// Solidity: function deposit(uint256 _amount, address walletAddr, address swarmContract) returns()
func (_Pool *PoolTransactorSession) Deposit(_amount *big.Int, walletAddr common.Address, swarmContract common.Address) (*types.Transaction, error) {
	return _Pool.Contract.Deposit(&_Pool.TransactOpts, _amount, walletAddr, swarmContract)
}

// Redeem is a paid mutator transaction binding the contract method 0xbe040fb0.
//
// Solidity: function redeem() returns(bool)
func (_Pool *PoolTransactor) Redeem(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "redeem")
}

// Redeem is a paid mutator transaction binding the contract method 0xbe040fb0.
//
// Solidity: function redeem() returns(bool)
func (_Pool *PoolSession) Redeem() (*types.Transaction, error) {
	return _Pool.Contract.Redeem(&_Pool.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0xbe040fb0.
//
// Solidity: function redeem() returns(bool)
func (_Pool *PoolTransactorSession) Redeem() (*types.Transaction, error) {
	return _Pool.Contract.Redeem(&_Pool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pool *PoolTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pool *PoolSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pool.Contract.RenounceOwnership(&_Pool.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pool *PoolTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pool.Contract.RenounceOwnership(&_Pool.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_Pool *PoolTransactor) Reward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "reward")
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_Pool *PoolSession) Reward() (*types.Transaction, error) {
	return _Pool.Contract.Reward(&_Pool.TransactOpts)
}

// Reward is a paid mutator transaction binding the contract method 0x228cb733.
//
// Solidity: function reward() returns()
func (_Pool *PoolTransactorSession) Reward() (*types.Transaction, error) {
	return _Pool.Contract.Reward(&_Pool.TransactOpts)
}

// SetAlpha is a paid mutator transaction binding the contract method 0x0c17d42c.
//
// Solidity: function setAlpha(uint256 rp) returns()
func (_Pool *PoolTransactor) SetAlpha(opts *bind.TransactOpts, rp *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setAlpha", rp)
}

// SetAlpha is a paid mutator transaction binding the contract method 0x0c17d42c.
//
// Solidity: function setAlpha(uint256 rp) returns()
func (_Pool *PoolSession) SetAlpha(rp *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetAlpha(&_Pool.TransactOpts, rp)
}

// SetAlpha is a paid mutator transaction binding the contract method 0x0c17d42c.
//
// Solidity: function setAlpha(uint256 rp) returns()
func (_Pool *PoolTransactorSession) SetAlpha(rp *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetAlpha(&_Pool.TransactOpts, rp)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 ga) returns()
func (_Pool *PoolTransactor) SetFee(opts *bind.TransactOpts, ga *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setFee", ga)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 ga) returns()
func (_Pool *PoolSession) SetFee(ga *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetFee(&_Pool.TransactOpts, ga)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 ga) returns()
func (_Pool *PoolTransactorSession) SetFee(ga *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetFee(&_Pool.TransactOpts, ga)
}

// SetUploadCount is a paid mutator transaction binding the contract method 0x27cf4264.
//
// Solidity: function setUploadCount(uint256 _uploadCount) returns()
func (_Pool *PoolTransactor) SetUploadCount(opts *bind.TransactOpts, _uploadCount *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setUploadCount", _uploadCount)
}

// SetUploadCount is a paid mutator transaction binding the contract method 0x27cf4264.
//
// Solidity: function setUploadCount(uint256 _uploadCount) returns()
func (_Pool *PoolSession) SetUploadCount(_uploadCount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetUploadCount(&_Pool.TransactOpts, _uploadCount)
}

// SetUploadCount is a paid mutator transaction binding the contract method 0x27cf4264.
//
// Solidity: function setUploadCount(uint256 _uploadCount) returns()
func (_Pool *PoolTransactorSession) SetUploadCount(_uploadCount *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetUploadCount(&_Pool.TransactOpts, _uploadCount)
}

// SetUploadTime is a paid mutator transaction binding the contract method 0x84e96291.
//
// Solidity: function setUploadTime(uint256 _uploadTime) returns()
func (_Pool *PoolTransactor) SetUploadTime(opts *bind.TransactOpts, _uploadTime *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "setUploadTime", _uploadTime)
}

// SetUploadTime is a paid mutator transaction binding the contract method 0x84e96291.
//
// Solidity: function setUploadTime(uint256 _uploadTime) returns()
func (_Pool *PoolSession) SetUploadTime(_uploadTime *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetUploadTime(&_Pool.TransactOpts, _uploadTime)
}

// SetUploadTime is a paid mutator transaction binding the contract method 0x84e96291.
//
// Solidity: function setUploadTime(uint256 _uploadTime) returns()
func (_Pool *PoolTransactorSession) SetUploadTime(_uploadTime *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.SetUploadTime(&_Pool.TransactOpts, _uploadTime)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pool *PoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pool *PoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pool.Contract.TransferOwnership(&_Pool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pool *PoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pool.Contract.TransferOwnership(&_Pool.TransactOpts, newOwner)
}

// UploadInfo is a paid mutator transaction binding the contract method 0xc3309d2a.
//
// Solidity: function uploadInfo(uint256 ts, string tsStr, bytes32 sign, string result, uint256 _hw) returns()
func (_Pool *PoolTransactor) UploadInfo(opts *bind.TransactOpts, ts *big.Int, tsStr string, sign [32]byte, result string, _hw *big.Int) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "uploadInfo", ts, tsStr, sign, result, _hw)
}

// UploadInfo is a paid mutator transaction binding the contract method 0xc3309d2a.
//
// Solidity: function uploadInfo(uint256 ts, string tsStr, bytes32 sign, string result, uint256 _hw) returns()
func (_Pool *PoolSession) UploadInfo(ts *big.Int, tsStr string, sign [32]byte, result string, _hw *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.UploadInfo(&_Pool.TransactOpts, ts, tsStr, sign, result, _hw)
}

// UploadInfo is a paid mutator transaction binding the contract method 0xc3309d2a.
//
// Solidity: function uploadInfo(uint256 ts, string tsStr, bytes32 sign, string result, uint256 _hw) returns()
func (_Pool *PoolTransactorSession) UploadInfo(ts *big.Int, tsStr string, sign [32]byte, result string, _hw *big.Int) (*types.Transaction, error) {
	return _Pool.Contract.UploadInfo(&_Pool.TransactOpts, ts, tsStr, sign, result, _hw)
}

// PoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pool contract.
type PoolOwnershipTransferredIterator struct {
	Event *PoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolOwnershipTransferred)
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
		it.Event = new(PoolOwnershipTransferred)
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
func (it *PoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolOwnershipTransferred represents a OwnershipTransferred event raised by the Pool contract.
type PoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pool *PoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PoolOwnershipTransferredIterator{contract: _Pool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pool *PoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolOwnershipTransferred)
				if err := _Pool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pool *PoolFilterer) ParseOwnershipTransferred(log types.Log) (*PoolOwnershipTransferred, error) {
	event := new(PoolOwnershipTransferred)
	if err := _Pool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Pool contract.
type PoolPausedIterator struct {
	Event *PoolPaused // Event containing the contract specifics and raw log

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
func (it *PoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolPaused)
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
		it.Event = new(PoolPaused)
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
func (it *PoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolPaused represents a Paused event raised by the Pool contract.
type PoolPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pool *PoolFilterer) FilterPaused(opts *bind.FilterOpts) (*PoolPausedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PoolPausedIterator{contract: _Pool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pool *PoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PoolPaused) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolPaused)
				if err := _Pool.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Pool *PoolFilterer) ParsePaused(log types.Log) (*PoolPaused, error) {
	event := new(PoolPaused)
	if err := _Pool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	return event, nil
}

// PoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Pool contract.
type PoolUnpausedIterator struct {
	Event *PoolUnpaused // Event containing the contract specifics and raw log

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
func (it *PoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolUnpaused)
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
		it.Event = new(PoolUnpaused)
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
func (it *PoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolUnpaused represents a Unpaused event raised by the Pool contract.
type PoolUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pool *PoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PoolUnpausedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PoolUnpausedIterator{contract: _Pool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pool *PoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolUnpaused)
				if err := _Pool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Pool *PoolFilterer) ParseUnpaused(log types.Log) (*PoolUnpaused, error) {
	event := new(PoolUnpaused)
	if err := _Pool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	return event, nil
}
