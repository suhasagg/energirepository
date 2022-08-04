// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/energicryptocurrency/energi"
	"github.com/energicryptocurrency/energi/accounts/abi"
	"github.com/energicryptocurrency/energi/accounts/abi/bind"
	"github.com/energicryptocurrency/energi/common"
	"github.com/energicryptocurrency/energi/core/types"
	"github.com/energicryptocurrency/energi/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BidABI is the input ABI used to generate the binding from.
const BidABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"endBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownercomission\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canceled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getHighestBid\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"highestBidder\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cancelBid\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bidIncrement\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ipfsHash\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"fundsByBidder\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"placeBid\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"highestBindingBid\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_bidIncrement\",\"type\":\"uint256\"},{\"name\":\"_startBlock\",\"type\":\"uint256\"},{\"name\":\"_endBlock\",\"type\":\"uint256\"},{\"name\":\"_ipfsHash\",\"type\":\"string\"}],\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"highestBidder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"highestBid\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"highestBindingBid\",\"type\":\"uint256\"}],\"name\":\"LogBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"withdrawer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"withdrawalAccount\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogCanceled\",\"type\":\"event\"}]"

// BidBin is the compiled bytecode used for deploying new contracts.
const BidBin = `0x606060405260405161071a38038061071a83398101604052805160805160a05160c05160e05193949293919290910181831061004f57610002565b505050505050506105dd8061013d6000396000f35b84600160a060020a03166000141561006657610002565b60008054600160a060020a031916861781556001858155600285815560038590556004805485519482905290936020601f9483161561010002600019019092169290920483018190047f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b90810193909186019083901061010957805160ff19168380011785555b5061003a9291505b8082111561013957600081556001016100f5565b828001600101855582156100ed579182015b828111156100ed57825182600050559160200191906001019061011b565b509056606060405236156100ae5760e060020a6000350463083c632381146100b35780631a3e1fe4146100c15780633ccfd60b146100cf5780633f9942ff1461010757806348cd4cb1146101185780634979440a146101265780638da5cb5b1461014d57806391f90157146101645780639435c8871461017b578063b3cc167a146101a1578063c623674f146101af578063ce10cf8014610212578063ecfc7ecc1461022f578063f5b56c5614610249575b610002565b346100025761025760035481565b346100025761025760075481565b34610002576102696005546000908190819060ff161561037c57505033600160a060020a03811682526009602052604082205461030f565b346100025761026960055460ff1681565b346100025761025760025481565b3461000257610257600854600160a060020a03166000908152600960205260409020545b90565b346100025761027d600054600160a060020a031681565b346100025761027d600854600160a060020a031681565b346100025761026960008054600160a060020a0390811633919091161461042d57610002565b346100025761025760015481565b34610002576040805160048054602060026001831615610100026000190190921691909104601f810182900482028401820190945283835261029a93908301828280156104a65780601f1061047b576101008083540402835291602001916104a6565b346100025761025760043560096020526000908152604090205481565b6102696005546000908190819060ff16156104ae57610002565b346100025761025760065481565b60408051918252519081900360200190f35b604080519115158252519081900360200190f35b60408051600160a060020a03929092168252519081900360200190f35b60405180806020018281038252838181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f1680156102fa5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b5033905060005b600160a060020a03828116600081815260096020908152604091829020805486900390558151339094168452830191909152818101839052517f0ec497a8ae5b1ba29c60470ef651def995fac3deebbdcc56c47a4e5f51a4c2bd9181900360600190a1600192505b505090565b600054600160a060020a039081163390911614156103bb575050600854600654600a805460ff19166001179055600160a060020a03919091169061030f565b600854600160a060020a0390811633909116141561030857600854600a54600160a060020a0391909116925060ff16156104045750600081815260096020526040902054610428565b50600854600654600160a060020a0391909116600090815260096020526040902054035b61030f565b60055460ff161561043d57610002565b6005805460ff191660011790556040517f462b6ca7f632601af1295aeb320851f50e8e630a309173f23535845ea4bfb3b990600090a150600161014a565b820191906000526020600020905b81548152906001019060200180831161048957829003601f168201915b505050505081565b34600014156104bc57610002565b5050600160a060020a033381166000818152600960205260408082208054600854909516835290822054929091523492909201918290558082116105165760015461055c908301825b6000818310156105d45750816105d7565b600854600160a060020a039081163390911614610568576008805473ffffffffffffffffffffffffffffffffffffffff1916331790556001546105649083908301610505565b60065561056b565b6006555b50805b60085460065460408051600160a060020a033381168252602082018790529390931683820152606083018490526080830191909152517ff152f4ff5e488c55370a2d53925a55055228ebd8ec95bd0251bbb299e48786b09181900360a00190a160019250610377565b50805b9291505056`

// DeployBid deploys a new Ethereum contract, binding an instance of Bid to it.
func DeployBid(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _bidIncrement *big.Int, _startBlock *big.Int, _endBlock *big.Int, _ipfsHash string) (common.Address, *types.Transaction, *Bid, error) {
	parsed, err := abi.JSON(strings.NewReader(BidABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BidBin), backend, _owner, _bidIncrement, _startBlock, _endBlock, _ipfsHash)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bid{BidCaller: BidCaller{contract: contract}, BidTransactor: BidTransactor{contract: contract}, BidFilterer: BidFilterer{contract: contract}}, nil
}

// BidBin is the compiled bytecode of contract after deployment.
const BidRuntimeBin = `0x606060405236156100ae5760e060020a6000350463083c632381146100b35780631a3e1fe4146100c15780633ccfd60b146100cf5780633f9942ff1461010757806348cd4cb1146101185780634979440a146101265780638da5cb5b1461014d57806391f90157146101645780639435c8871461017b578063b3cc167a146101a1578063c623674f146101af578063ce10cf8014610212578063ecfc7ecc1461022f578063f5b56c5614610249575b610002565b346100025761025760035481565b346100025761025760075481565b34610002576102696005546000908190819060ff161561037c57505033600160a060020a03811682526009602052604082205461030f565b346100025761026960055460ff1681565b346100025761025760025481565b3461000257610257600854600160a060020a03166000908152600960205260409020545b90565b346100025761027d600054600160a060020a031681565b346100025761027d600854600160a060020a031681565b346100025761026960008054600160a060020a0390811633919091161461042d57610002565b346100025761025760015481565b34610002576040805160048054602060026001831615610100026000190190921691909104601f810182900482028401820190945283835261029a93908301828280156104a65780601f1061047b576101008083540402835291602001916104a6565b346100025761025760043560096020526000908152604090205481565b6102696005546000908190819060ff16156104ae57610002565b346100025761025760065481565b60408051918252519081900360200190f35b604080519115158252519081900360200190f35b60408051600160a060020a03929092168252519081900360200190f35b60405180806020018281038252838181518152602001915080519060200190808383829060006004602084601f0104600302600f01f150905090810190601f1680156102fa5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b5033905060005b600160a060020a03828116600081815260096020908152604091829020805486900390558151339094168452830191909152818101839052517f0ec497a8ae5b1ba29c60470ef651def995fac3deebbdcc56c47a4e5f51a4c2bd9181900360600190a1600192505b505090565b600054600160a060020a039081163390911614156103bb575050600854600654600a805460ff19166001179055600160a060020a03919091169061030f565b600854600160a060020a0390811633909116141561030857600854600a54600160a060020a0391909116925060ff16156104045750600081815260096020526040902054610428565b50600854600654600160a060020a0391909116600090815260096020526040902054035b61030f565b60055460ff161561043d57610002565b6005805460ff191660011790556040517f462b6ca7f632601af1295aeb320851f50e8e630a309173f23535845ea4bfb3b990600090a150600161014a565b820191906000526020600020905b81548152906001019060200180831161048957829003601f168201915b505050505081565b34600014156104bc57610002565b5050600160a060020a033381166000818152600960205260408082208054600854909516835290822054929091523492909201918290558082116105165760015461055c908301825b6000818310156105d45750816105d7565b600854600160a060020a039081163390911614610568576008805473ffffffffffffffffffffffffffffffffffffffff1916331790556001546105649083908301610505565b60065561056b565b6006555b50805b60085460065460408051600160a060020a033381168252602082018790529390931683820152606083018490526080830191909152517ff152f4ff5e488c55370a2d53925a55055228ebd8ec95bd0251bbb299e48786b09181900360a00190a160019250610377565b50805b9291505056`

// Bid is an auto generated Go binding around an Ethereum contract.
type Bid struct {
	BidCaller     // Read-only binding to the contract
	BidTransactor // Write-only binding to the contract
	BidFilterer   // Log filterer for contract events
}

// BidCaller is an auto generated read-only Go binding around an Ethereum contract.
type BidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BidSession struct {
	Contract     *Bid              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BidCallerSession struct {
	Contract *BidCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BidTransactorSession struct {
	Contract     *BidTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BidRaw is an auto generated low-level Go binding around an Ethereum contract.
type BidRaw struct {
	Contract *Bid // Generic contract binding to access the raw methods on
}

// BidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BidCallerRaw struct {
	Contract *BidCaller // Generic read-only contract binding to access the raw methods on
}

// BidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BidTransactorRaw struct {
	Contract *BidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBid creates a new instance of Bid, bound to a specific deployed contract.
func NewBid(address common.Address, backend bind.ContractBackend) (*Bid, error) {
	contract, err := bindBid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bid{BidCaller: BidCaller{contract: contract}, BidTransactor: BidTransactor{contract: contract}, BidFilterer: BidFilterer{contract: contract}}, nil
}

// NewBidCaller creates a new read-only instance of Bid, bound to a specific deployed contract.
func NewBidCaller(address common.Address, caller bind.ContractCaller) (*BidCaller, error) {
	contract, err := bindBid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BidCaller{contract: contract}, nil
}

// NewBidTransactor creates a new write-only instance of Bid, bound to a specific deployed contract.
func NewBidTransactor(address common.Address, transactor bind.ContractTransactor) (*BidTransactor, error) {
	contract, err := bindBid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BidTransactor{contract: contract}, nil
}

// NewBidFilterer creates a new log filterer instance of Bid, bound to a specific deployed contract.
func NewBidFilterer(address common.Address, filterer bind.ContractFilterer) (*BidFilterer, error) {
	contract, err := bindBid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BidFilterer{contract: contract}, nil
}

// bindBid binds a generic wrapper to an already deployed contract.
func bindBid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BidABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bid *BidRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bid.Contract.BidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bid *BidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bid.Contract.BidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bid *BidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bid.Contract.BidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bid *BidCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bid.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bid *BidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bid.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bid *BidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bid.Contract.contract.Transact(opts, method, params...)
}

// BidIncrement is a free data retrieval call binding the contract method 0xb3cc167a.
//
// Solidity: function bidIncrement() constant returns(uint256)
func (_Bid *BidCaller) BidIncrement(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bid.contract.Call(opts, out, "bidIncrement")
	return *ret0, err
}

// BidIncrement is a free data retrieval call binding the contract method 0xb3cc167a.
//
// Solidity: function bidIncrement() constant returns(uint256)
func (_Bid *BidSession) BidIncrement() (*big.Int, error) {
	return _Bid.Contract.BidIncrement(&_Bid.CallOpts)
}

// BidIncrement is a free data retrieval call binding the contract method 0xb3cc167a.
//
// Solidity: function bidIncrement() constant returns(uint256)
func (_Bid *BidCallerSession) BidIncrement() (*big.Int, error) {
	return _Bid.Contract.BidIncrement(&_Bid.CallOpts)
}

// Canceled is a free data retrieval call binding the contract method 0x3f9942ff.
//
// Solidity: function canceled() constant returns(bool)
func (_Bid *BidCaller) Canceled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bid.contract.Call(opts, out, "canceled")
	return *ret0, err
}

// Canceled is a free data retrieval call binding the contract method 0x3f9942ff.
//
// Solidity: function canceled() constant returns(bool)
func (_Bid *BidSession) Canceled() (bool, error) {
	return _Bid.Contract.Canceled(&_Bid.CallOpts)
}

// Canceled is a free data retrieval call binding the contract method 0x3f9942ff.
//
// Solidity: function canceled() constant returns(bool)
func (_Bid *BidCallerSession) Canceled() (bool, error) {
	return _Bid.Contract.Canceled(&_Bid.CallOpts)
}

// EndBlock is a free data retrieval call binding the contract method 0x083c6323.
//
// Solidity: function endBlock() constant returns(uint256)
func (_Bid *BidCaller) EndBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bid.contract.Call(opts, out, "endBlock")
	return *ret0, err
}

// EndBlock is a free data retrieval call binding the contract method 0x083c6323.
//
// Solidity: function endBlock() constant returns(uint256)
func (_Bid *BidSession) EndBlock() (*big.Int, error) {
	return _Bid.Contract.EndBlock(&_Bid.CallOpts)
}

// EndBlock is a free data retrieval call binding the contract method 0x083c6323.
//
// Solidity: function endBlock() constant returns(uint256)
func (_Bid *BidCallerSession) EndBlock() (*big.Int, error) {
	return _Bid.Contract.EndBlock(&_Bid.CallOpts)
}

// FundsByBidder is a free data retrieval call binding the contract method 0xce10cf80.
//
// Solidity: function fundsByBidder(address ) constant returns(uint256)
func (_Bid *BidCaller) FundsByBidder(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bid.contract.Call(opts, out, "fundsByBidder", arg0)
	return *ret0, err
}

// FundsByBidder is a free data retrieval call binding the contract method 0xce10cf80.
//
// Solidity: function fundsByBidder(address ) constant returns(uint256)
func (_Bid *BidSession) FundsByBidder(arg0 common.Address) (*big.Int, error) {
	return _Bid.Contract.FundsByBidder(&_Bid.CallOpts, arg0)
}

// FundsByBidder is a free data retrieval call binding the contract method 0xce10cf80.
//
// Solidity: function fundsByBidder(address ) constant returns(uint256)
func (_Bid *BidCallerSession) FundsByBidder(arg0 common.Address) (*big.Int, error) {
	return _Bid.Contract.FundsByBidder(&_Bid.CallOpts, arg0)
}

// GetHighestBid is a free data retrieval call binding the contract method 0x4979440a.
//
// Solidity: function getHighestBid() constant returns(uint256)
func (_Bid *BidCaller) GetHighestBid(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bid.contract.Call(opts, out, "getHighestBid")
	return *ret0, err
}

// GetHighestBid is a free data retrieval call binding the contract method 0x4979440a.
//
// Solidity: function getHighestBid() constant returns(uint256)
func (_Bid *BidSession) GetHighestBid() (*big.Int, error) {
	return _Bid.Contract.GetHighestBid(&_Bid.CallOpts)
}

// GetHighestBid is a free data retrieval call binding the contract method 0x4979440a.
//
// Solidity: function getHighestBid() constant returns(uint256)
func (_Bid *BidCallerSession) GetHighestBid() (*big.Int, error) {
	return _Bid.Contract.GetHighestBid(&_Bid.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() constant returns(address)
func (_Bid *BidCaller) HighestBidder(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bid.contract.Call(opts, out, "highestBidder")
	return *ret0, err
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() constant returns(address)
func (_Bid *BidSession) HighestBidder() (common.Address, error) {
	return _Bid.Contract.HighestBidder(&_Bid.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() constant returns(address)
func (_Bid *BidCallerSession) HighestBidder() (common.Address, error) {
	return _Bid.Contract.HighestBidder(&_Bid.CallOpts)
}

// HighestBindingBid is a free data retrieval call binding the contract method 0xf5b56c56.
//
// Solidity: function highestBindingBid() constant returns(uint256)
func (_Bid *BidCaller) HighestBindingBid(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bid.contract.Call(opts, out, "highestBindingBid")
	return *ret0, err
}

// HighestBindingBid is a free data retrieval call binding the contract method 0xf5b56c56.
//
// Solidity: function highestBindingBid() constant returns(uint256)
func (_Bid *BidSession) HighestBindingBid() (*big.Int, error) {
	return _Bid.Contract.HighestBindingBid(&_Bid.CallOpts)
}

// HighestBindingBid is a free data retrieval call binding the contract method 0xf5b56c56.
//
// Solidity: function highestBindingBid() constant returns(uint256)
func (_Bid *BidCallerSession) HighestBindingBid() (*big.Int, error) {
	return _Bid.Contract.HighestBindingBid(&_Bid.CallOpts)
}

// IpfsHash is a free data retrieval call binding the contract method 0xc623674f.
//
// Solidity: function ipfsHash() constant returns(string)
func (_Bid *BidCaller) IpfsHash(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Bid.contract.Call(opts, out, "ipfsHash")
	return *ret0, err
}

// IpfsHash is a free data retrieval call binding the contract method 0xc623674f.
//
// Solidity: function ipfsHash() constant returns(string)
func (_Bid *BidSession) IpfsHash() (string, error) {
	return _Bid.Contract.IpfsHash(&_Bid.CallOpts)
}

// IpfsHash is a free data retrieval call binding the contract method 0xc623674f.
//
// Solidity: function ipfsHash() constant returns(string)
func (_Bid *BidCallerSession) IpfsHash() (string, error) {
	return _Bid.Contract.IpfsHash(&_Bid.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bid *BidCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bid.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bid *BidSession) Owner() (common.Address, error) {
	return _Bid.Contract.Owner(&_Bid.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bid *BidCallerSession) Owner() (common.Address, error) {
	return _Bid.Contract.Owner(&_Bid.CallOpts)
}

// Ownercomission is a free data retrieval call binding the contract method 0x1a3e1fe4.
//
// Solidity: function ownercomission() constant returns(uint256)
func (_Bid *BidCaller) Ownercomission(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bid.contract.Call(opts, out, "ownercomission")
	return *ret0, err
}

// Ownercomission is a free data retrieval call binding the contract method 0x1a3e1fe4.
//
// Solidity: function ownercomission() constant returns(uint256)
func (_Bid *BidSession) Ownercomission() (*big.Int, error) {
	return _Bid.Contract.Ownercomission(&_Bid.CallOpts)
}

// Ownercomission is a free data retrieval call binding the contract method 0x1a3e1fe4.
//
// Solidity: function ownercomission() constant returns(uint256)
func (_Bid *BidCallerSession) Ownercomission() (*big.Int, error) {
	return _Bid.Contract.Ownercomission(&_Bid.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() constant returns(uint256)
func (_Bid *BidCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bid.contract.Call(opts, out, "startBlock")
	return *ret0, err
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() constant returns(uint256)
func (_Bid *BidSession) StartBlock() (*big.Int, error) {
	return _Bid.Contract.StartBlock(&_Bid.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() constant returns(uint256)
func (_Bid *BidCallerSession) StartBlock() (*big.Int, error) {
	return _Bid.Contract.StartBlock(&_Bid.CallOpts)
}

// CancelBid is a paid mutator transaction binding the contract method 0x9435c887.
//
// Solidity: function cancelBid() returns(bool success)
func (_Bid *BidTransactor) CancelBid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bid.contract.Transact(opts, "cancelBid")
}

// CancelBid is a paid mutator transaction binding the contract method 0x9435c887.
//
// Solidity: function cancelBid() returns(bool success)
func (_Bid *BidSession) CancelBid() (*types.Transaction, error) {
	return _Bid.Contract.CancelBid(&_Bid.TransactOpts)
}

// CancelBid is a paid mutator transaction binding the contract method 0x9435c887.
//
// Solidity: function cancelBid() returns(bool success)
func (_Bid *BidTransactorSession) CancelBid() (*types.Transaction, error) {
	return _Bid.Contract.CancelBid(&_Bid.TransactOpts)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xecfc7ecc.
//
// Solidity: function placeBid() returns(bool success)
func (_Bid *BidTransactor) PlaceBid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bid.contract.Transact(opts, "placeBid")
}

// PlaceBid is a paid mutator transaction binding the contract method 0xecfc7ecc.
//
// Solidity: function placeBid() returns(bool success)
func (_Bid *BidSession) PlaceBid() (*types.Transaction, error) {
	return _Bid.Contract.PlaceBid(&_Bid.TransactOpts)
}

// PlaceBid is a paid mutator transaction binding the contract method 0xecfc7ecc.
//
// Solidity: function placeBid() returns(bool success)
func (_Bid *BidTransactorSession) PlaceBid() (*types.Transaction, error) {
	return _Bid.Contract.PlaceBid(&_Bid.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool success)
func (_Bid *BidTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bid.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool success)
func (_Bid *BidSession) Withdraw() (*types.Transaction, error) {
	return _Bid.Contract.Withdraw(&_Bid.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool success)
func (_Bid *BidTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Bid.Contract.Withdraw(&_Bid.TransactOpts)
}

// BidLogBidIterator is returned from FilterLogBid and is used to iterate over the raw logs and unpacked data for LogBid events raised by the Bid contract.
type BidLogBidIterator struct {
	Event *BidLogBid // Event containing the contract specifics and raw log

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
func (it *BidLogBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BidLogBid)
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
		it.Event = new(BidLogBid)
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
func (it *BidLogBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BidLogBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BidLogBid represents a LogBid event raised by the Bid contract.
type BidLogBid struct {
	Bidder            common.Address
	Bid               *big.Int
	HighestBidder     common.Address
	HighestBid        *big.Int
	HighestBindingBid *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLogBid is a free log retrieval operation binding the contract event 0xf152f4ff5e488c55370a2d53925a55055228ebd8ec95bd0251bbb299e48786b0.
//
// Solidity: event LogBid(address bidder, uint256 bid, address highestBidder, uint256 highestBid, uint256 highestBindingBid)
func (_Bid *BidFilterer) FilterLogBid(opts *bind.FilterOpts) (*BidLogBidIterator, error) {

	logs, sub, err := _Bid.contract.FilterLogs(opts, "LogBid")
	if err != nil {
		return nil, err
	}
	return &BidLogBidIterator{contract: _Bid.contract, event: "LogBid", logs: logs, sub: sub}, nil
}

// WatchLogBid is a free log subscription operation binding the contract event 0xf152f4ff5e488c55370a2d53925a55055228ebd8ec95bd0251bbb299e48786b0.
//
// Solidity: event LogBid(address bidder, uint256 bid, address highestBidder, uint256 highestBid, uint256 highestBindingBid)
func (_Bid *BidFilterer) WatchLogBid(opts *bind.WatchOpts, sink chan<- *BidLogBid) (event.Subscription, error) {

	logs, sub, err := _Bid.contract.WatchLogs(opts, "LogBid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BidLogBid)
				if err := _Bid.contract.UnpackLog(event, "LogBid", log); err != nil {
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

// BidLogCanceledIterator is returned from FilterLogCanceled and is used to iterate over the raw logs and unpacked data for LogCanceled events raised by the Bid contract.
type BidLogCanceledIterator struct {
	Event *BidLogCanceled // Event containing the contract specifics and raw log

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
func (it *BidLogCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BidLogCanceled)
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
		it.Event = new(BidLogCanceled)
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
func (it *BidLogCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BidLogCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BidLogCanceled represents a LogCanceled event raised by the Bid contract.
type BidLogCanceled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogCanceled is a free log retrieval operation binding the contract event 0x462b6ca7f632601af1295aeb320851f50e8e630a309173f23535845ea4bfb3b9.
//
// Solidity: event LogCanceled()
func (_Bid *BidFilterer) FilterLogCanceled(opts *bind.FilterOpts) (*BidLogCanceledIterator, error) {

	logs, sub, err := _Bid.contract.FilterLogs(opts, "LogCanceled")
	if err != nil {
		return nil, err
	}
	return &BidLogCanceledIterator{contract: _Bid.contract, event: "LogCanceled", logs: logs, sub: sub}, nil
}

// WatchLogCanceled is a free log subscription operation binding the contract event 0x462b6ca7f632601af1295aeb320851f50e8e630a309173f23535845ea4bfb3b9.
//
// Solidity: event LogCanceled()
func (_Bid *BidFilterer) WatchLogCanceled(opts *bind.WatchOpts, sink chan<- *BidLogCanceled) (event.Subscription, error) {

	logs, sub, err := _Bid.contract.WatchLogs(opts, "LogCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BidLogCanceled)
				if err := _Bid.contract.UnpackLog(event, "LogCanceled", log); err != nil {
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

// BidLogWithdrawalIterator is returned from FilterLogWithdrawal and is used to iterate over the raw logs and unpacked data for LogWithdrawal events raised by the Bid contract.
type BidLogWithdrawalIterator struct {
	Event *BidLogWithdrawal // Event containing the contract specifics and raw log

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
func (it *BidLogWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BidLogWithdrawal)
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
		it.Event = new(BidLogWithdrawal)
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
func (it *BidLogWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BidLogWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BidLogWithdrawal represents a LogWithdrawal event raised by the Bid contract.
type BidLogWithdrawal struct {
	Withdrawer        common.Address
	WithdrawalAccount common.Address
	Amount            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawal is a free log retrieval operation binding the contract event 0x0ec497a8ae5b1ba29c60470ef651def995fac3deebbdcc56c47a4e5f51a4c2bd.
//
// Solidity: event LogWithdrawal(address withdrawer, address withdrawalAccount, uint256 amount)
func (_Bid *BidFilterer) FilterLogWithdrawal(opts *bind.FilterOpts) (*BidLogWithdrawalIterator, error) {

	logs, sub, err := _Bid.contract.FilterLogs(opts, "LogWithdrawal")
	if err != nil {
		return nil, err
	}
	return &BidLogWithdrawalIterator{contract: _Bid.contract, event: "LogWithdrawal", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawal is a free log subscription operation binding the contract event 0x0ec497a8ae5b1ba29c60470ef651def995fac3deebbdcc56c47a4e5f51a4c2bd.
//
// Solidity: event LogWithdrawal(address withdrawer, address withdrawalAccount, uint256 amount)
func (_Bid *BidFilterer) WatchLogWithdrawal(opts *bind.WatchOpts, sink chan<- *BidLogWithdrawal) (event.Subscription, error) {

	logs, sub, err := _Bid.contract.WatchLogs(opts, "LogWithdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BidLogWithdrawal)
				if err := _Bid.contract.UnpackLog(event, "LogWithdrawal", log); err != nil {
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
