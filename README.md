## Energi Core Bidding Game Assignment

1)https://github.com/suhasagg/energirepository/tree/main/contracts/biddinggame

2)https://github.com/suhasagg/energirepository/tree/main/Utilscripts

3)https://github.com/suhasagg/hardhateventparsing

Module is tested on 

1)ganache-cli

2)ETHERSCAN_API_KEY=ABC123ABC123ABC123ABC123ABC123ABC1

KOVAN_RPC_URL=https://eth-ropsten.alchemyapi.io/v2/<YOUR ALCHEMY KEY>
  
PRIVATE_KEY=0xabc123abc123abc123abc123abc123abc123abc123abc123abc123abc123abc1

A simple game sequence 

2022/08/04 11:44:22 Placed bid - Player 1

2022/08/04 11:44:23 Placed bid - PLayer 2 

2022/08/04 11:44:24 Placed bid - Player 1

2022/08/04 11:44:25 Placed bid - Player 2

2022/08/04 11:44:26 Placed bid - Player 1

2022/08/04 11:44:27 Placed bid - Player 2

2022/08/04 11:44:28 Placed bid - Player 1

2022/08/04 11:44:29 Placed bid - Player 2

2022/08/04 11:44:30 Placed bid - Player 1

2022/08/04 11:44:31 Placed bid - Player 2

2022/08/04 11:44:46 bidding round has completed

2022/08/04 11:44:46 Bidding is over

2022/08/04 11:44:46 Player 1 tries to withdraw

2022/08/04 11:44:46 Player 2 tries to withdraw


Game logs


Bid placement logs with Bidding Game state at that timestmap


******LogBid Event Emitted******

Bidder: 0x5194b63F10691E46635B27925100CfC0a5cECa62

Bid: 470

HighestBidder: 0x5ece25D5a9F2f82066fc97ec3B9d162599b0d826

Highestbid: 470

HighestBindingBid: 470

Bidder: 0x5194b63F10691E46635B27925100CfC0a5cECa62

**************

******LogBid Event Emitted******

Bidder: 0x5ece25D5a9F2f82066fc97ec3B9d162599b0d826

Bid: 520

HighestBidder: 0x5ece25D5a9F2f82066fc97ec3B9d162599b0d826

Highestbid: 520

HighestBindingBid: 470

Bidder: 0x5ece25D5a9F2f82066fc97ec3B9d162599b0d826

**************

Detailed Game event Logs  

{
  transactionIndex: 0,
  
  blockNumber: 2,
  
  transactionHash: '0x286d4f49d3ea89094763dcfac52034f86db04ea07e92229be99267b3b12143e4',
  
  address: '0x5FbDB2315678afecb367f032d93F642f64180aa3',
  
  topics: [
    '0xf152f4ff5e488c55370a2d53925a55055228ebd8ec95bd0251bbb299e48786b0'
  ],
  
  data: '0x000000000000000000000000f39fd6e51aad88f6f4ce6ab8827279cfffb92266000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000f39fd6e51aad88f6f4ce6ab8827279cfffb92266000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000000000000000000000000000000000000000000a',
  
  logIndex: 0,
  
  blockHash: '0xf9ab08e70bce82dad8dcc2c9f451d6df6f1d22f6389fd34529c6f8000b50b6ea',
  
  args: [
    '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    BigNumber { value: "10" },
    '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    BigNumber { value: "10" },
    BigNumber { value: "10" },
    bidder: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    bid: BigNumber { value: "10" },
    highestBidder: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    highestBid: BigNumber { value: "10" },
    highestBindingBid: BigNumber { value: "10" }
  ],
  
  decode: [Function (anonymous)],
  
  event: 'LogBid',
  
  eventSignature: 'LogBid(address,uint256,address,uint256,uint256)',
  
  removeListener: [Function (anonymous)],
  
  getBlock: [Function (anonymous)],
  
  getTransaction: [Function (anonymous)],
  
  getTransactionReceipt: [Function (anonymous)]

}
{
  transactionIndex: 0,
  
  blockNumber: 3,
  
  transactionHash: '0x5aff457572999f14c567922bc57516d05ceb30d68bb42b96f9cf353a79d78fe6',
  
  address: '0x5FbDB2315678afecb367f032d93F642f64180aa3',
  
  topics: [
    '0xf152f4ff5e488c55370a2d53925a55055228ebd8ec95bd0251bbb299e48786b0'
  ],
  
  data: '0x000000000000000000000000f39fd6e51aad88f6f4ce6ab8827279cfffb92266000000000000000000000000000000000000000000000000000000000000001e000000000000000000000000f39fd6e51aad88f6f4ce6ab8827279cfffb92266000000000000000000000000000000000000000000000000000000000000001e000000000000000000000000000000000000000000000000000000000000000a',
  
  logIndex: 0,
  
  blockHash: '0x06ee1eabe09c00221bca4db2b69581173671ede550fe0cb97946efbe180f4a46',
  
  args: [
    '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    BigNumber { value: "30" },
    '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    BigNumber { value: "30" },
    BigNumber { value: "10" },
    bidder: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    bid: BigNumber { value: "30" },
    highestBidder: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    highestBid: BigNumber { value: "30" },
    highestBindingBid: BigNumber { value: "10" }
  ],
  
  decode: [Function (anonymous)],
  
  event: 'LogBid',
  
  eventSignature: 'LogBid(address,uint256,address,uint256,uint256)',
  
  removeListener: [Function (anonymous)],
  
  getBlock: [Function (anonymous)],
  
  getTransaction: [Function (anonymous)],
  
  getTransactionReceipt: [Function (anonymous)]

}
{
  
  transactionIndex: 0,
  
  blockNumber: 4,
  
  transactionHash: '0x67009ca0a16fa2d1196f144c8eb251cf4517993b720264627e97efc5145bce10',
  
  address: '0x5FbDB2315678afecb367f032d93F642f64180aa3',
  
  topics: [
    '0x462b6ca7f632601af1295aeb320851f50e8e630a309173f23535845ea4bfb3b9'
  ],
  
  data: '0x',
  
  logIndex: 0,
  
  blockHash: '0x3e4c02e4ead849f38a3cbd81a5e5e8d9e9876b9104efa67753d017ee77cbe8f5',
  
  args: [],
  
  decode: [Function (anonymous)],
  
  event: 'LogCanceled',
  
  eventSignature: 'LogCanceled()',
  
  removeListener: [Function (anonymous)],
  
  getBlock: [Function (anonymous)],
  
  getTransaction: [Function (anonymous)],
  
  getTransactionReceipt: [Function (anonymous)]

}
{
  transactionIndex: 0,
  
  blockNumber: 5,
  
  transactionHash: '0x48fe50c988ae365100fa4d784d391a8d346d29eb9a6ebe0cef97706672cac14e',
  
  address: '0x5FbDB2315678afecb367f032d93F642f64180aa3',
  
  topics: [
    '0x0ec497a8ae5b1ba29c60470ef651def995fac3deebbdcc56c47a4e5f51a4c2bd'
  ],
  
  data: '0x000000000000000000000000f39fd6e51aad88f6f4ce6ab8827279cfffb92266000000000000000000000000f39fd6e51aad88f6f4ce6ab8827279cfffb92266000000000000000000000000000000000000000000000000000000000000001e',
  
  logIndex: 0,
  
  blockHash: '0x20a3baef22dda77fc8eba15217a737d272c07c416715bcb1c5f8ba7d1f06b84d',
  
  args: [
    '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    BigNumber { value: "30" },
    withdrawer: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    withdrawalAccount: '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266',
    amount: BigNumber { value: "30" }
  ],
  
  decode: [Function (anonymous)],
  
  event: 'LogWithdrawal',
  
  eventSignature: 'LogWithdrawal(address,address,uint256)',
  
  removeListener: [Function (anonymous)],
  
  getBlock: [Function (anonymous)],
  
  getTransaction: [Function (anonymous)],
  
  getTransactionReceipt: [Function (anonymous)]

}
