pragma solidity ^0.4.0;

contract Bid {
    // bid initialiser
    address public owner;
    uint public bidIncrement;
    uint public startBlock;
    uint public endBlock;
    string public ipfsHash;

    //bid contract state
    bool public canceled;
    uint public highestBindingBid;
    uint public ownercomission;
    address public highestBidder;
    mapping(address => uint256) public fundsByBidder;
    bool ownerHasWithdrawn;

    //Bidding game events
    //Event on placing Bid 
    event LogBid(address bidder, uint bid, address highestBidder, uint highestBid, uint highestBindingBid);
    //Event on withdrawl + bidding rewards distribution 
    event LogWithdrawal(address withdrawer, address withdrawalAccount, uint amount);
    //Bidding Game canceled 
    event LogCanceled();

  
    //Bidding Game Initialisation
    //Deciding block range for bidding game transactions
    //BLockchain merkle Root hash at which bidding game started
    //Bidding game deployer on chain, Bidding game owner/Master 
    function Bid(address _owner, uint _bidIncrement, uint _startBlock, uint _endBlock, string _ipfsHash) {
        if (_startBlock >= _endBlock) throw;

        if (_owner == 0) throw;

        owner = _owner;
        bidIncrement = _bidIncrement;
        startBlock = _startBlock;
        endBlock = _endBlock;
        ipfsHash = _ipfsHash;
    }

    //Returns highest bid placed so far
    function getHighestBid()
        constant
        returns (uint)
    {
        return fundsByBidder[highestBidder];
    }

    function placeBid()
        payable
        onlyNotCanceled
        returns (bool success)
    {
        // reject bids of 0 ETH
        if (msg.value == 0) throw;

        // the user's total bid = the current amount they've sent to the contract + whatever has been sent with this transaction
        uint newBid = fundsByBidder[msg.sender] + msg.value;

        // store highest bid so far
        uint highestBid = fundsByBidder[highestBidder];

        // Maintain a map of user and bids placed so far by user
        fundsByBidder[msg.sender] = newBid;

        if (newBid <= highestBid) {
            // if the user has overbid the maxBindingBid but not the highestBid, we simply
            // increase the highestBindingBid and leave highestBidder as it.

            highestBindingBid = min(newBid + bidIncrement, highestBid);
        } else {
            // if msg.sender is already the highest bidder, they must simply be wanting to raise
            // their maximum bid, in which case we shouldn't increase the highestBindingBid.

            // if the user is NOT highestBidder, and has overbid highestBid completely, we set them
            // as the new highestBidder and recalculate highestBindingBid.

            if (msg.sender != highestBidder) {
                highestBidder = msg.sender;
                highestBindingBid = min(newBid, highestBid + bidIncrement);
            }
            highestBid = newBid;
        }
        //Bid placed event generated
        LogBid(msg.sender, newBid, highestBidder, highestBid, highestBindingBid);
        return true;
    }

    function min(uint a, uint b)
        private
        constant
        returns (uint)
    {
        if (a < b) return a;
        return b;
    }

    function cancelBid()
        onlyOwner
        //onlyBeforeEnd
        onlyNotCanceled
        returns (bool success)
    {
        canceled = true;
     //Bid canceled event generated
        LogCanceled();
        return true;
    }

    function withdraw()
        returns (bool success)
    {
        address withdrawalAccount;
        uint withdrawalAmount;

        if (canceled) {
            // if the Bid Game was canceled, everyone should simply be allowed to withdraw their bidding game funds
            withdrawalAccount = msg.sender;
            withdrawalAmount = fundsByBidder[withdrawalAccount];

        } else {
            // the Bid finished without being canceled
            // Only game owner (commission withdrawl) and bid winner (game amount withdrawl) can withdraw
            if (msg.sender == owner) {
                // the Bid's owner should be allowed to withdraw the highestBindingBid (Game commission)
                withdrawalAccount = highestBidder;
                withdrawalAmount = highestBindingBid;
                ownerHasWithdrawn = true;

            } else if (msg.sender == highestBidder) {
                // the highest bidder should only be allowed to withdraw the difference between their
                // highest bid and the highestBindingBid (game commission)
                withdrawalAccount = highestBidder;
                if (ownerHasWithdrawn) {
                    withdrawalAmount = fundsByBidder[highestBidder];
                } else {
                    withdrawalAmount = fundsByBidder[highestBidder] - highestBindingBid;
                }

            } else {
                // anyone who participated but did not win the Bid should not be allowed to withdraw
                withdrawalAccount = msg.sender;
                withdrawalAmount = 0;
            }
        }

        fundsByBidder[withdrawalAccount] -= withdrawalAmount;

        //Bidding event is over and transfer details event is generated
        LogWithdrawal(msg.sender, withdrawalAccount, withdrawalAmount);

        return true;
    }

    modifier onlyOwner {
        if (msg.sender != owner) throw;
        _;
    }

    modifier onlyNotOwner {
        if (msg.sender == owner) throw;
        _;
    }

    modifier onlyNotCanceled {
        if (canceled) throw;
        _;
    }

    modifier onlyEndedOrCanceled {
        if (block.number < endBlock && !canceled) throw;
        _;
    }
}


