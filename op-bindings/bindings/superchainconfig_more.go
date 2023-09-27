// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const SuperchainConfigStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/SuperchainConfig.sol:SuperchainConfig\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/SuperchainConfig.sol:SuperchainConfig\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/SuperchainConfig.sol:SuperchainConfig\",\"label\":\"allowedSequencers\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_bytes32,t_bool)\"}],\"types\":{\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var SuperchainConfigStorageLayout = new(solc.StorageLayout)

var SuperchainConfigDeployedBin = "0x608060405234801561001057600080fd5b50600436106101775760003560e01c80636a42b8f8116100d8578063a2f9c4081161008c578063d92a09bc11610066578063d92a09bc146102ba578063f1e8cf06146102dd578063fdd07046146102f057600080fd5b8063a2f9c408146102a0578063c23a451a146102aa578063d8bff440146102b257600080fd5b80637fbf7b6a116100bd5780637fbf7b6a1461027d5780639eb17d4b14610285578063a06549561461028d57600080fd5b80636a42b8f81461026257806376ea31a41461026a57600080fd5b80634886eb9c1161012f57806354fd4d501161011457806354fd4d50146101f95780635c39fcc1146102425780635c975abb1461024a57600080fd5b80634886eb9c146101e95780634b5b189f146101f157600080fd5b8063337792541161016057806333779254146101ac5780633f4ba83a146101d9578063452a9320146101e157600080fd5b8063136439dd1461017c5780633323920214610191575b600080fd5b61018f61018a366004610fd0565b610303565b005b6101996104c5565b6040519081526020015b60405180910390f35b6101b46104f3565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a3565b61018f61052c565b6101b461064d565b61019961067d565b6101996106a8565b6102356040518060400160405280600581526020017f312e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516101a39190611054565b6101b46106d3565b610252610703565b60405190151581526020016101a3565b61019961073a565b61025261027836600461116c565b610785565b6101996107aa565b6101996107d5565b61018f61029b36600461116c565b610800565b61019962093a8081565b6101996109aa565b6101b46109d5565b6102526102c8366004610fd0565b60016020526000908152604090205460ff1681565b61018f6102eb36600461116c565b610a05565b61018f6102fe366004611188565b610b0f565b61030b61064d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103ca576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f5375706572636861696e436f6e6669673a206f6e6c7920677561726469616e2060448201527f63616e207061757365000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b62093a8081111561045d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f5375706572636861696e436f6e6669673a206475726174696f6e20657863656560448201527f6473206d6178506175736500000000000000000000000000000000000000000060648201526084016103c1565b61049961048b60017f54176ff9944c4784e5857ec4e5ef560a462c483bf534eda43f91bb01a470b1b76112b6565b61049583426112cd565b9055565b6040517f9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e75290600090a150565b6104f060017fe5134cb7d217efbc8c357a6644e3c656a6235651a8f25717e410cbf378e577536112b6565b81565b600061052761052360017fe5134cb7d217efbc8c357a6644e3c656a6235651a8f25717e410cbf378e577536112b6565b5490565b905090565b61053461064d565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105ee576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f5375706572636861696e436f6e6669673a206f6e6c7920677561726469616e2060448201527f63616e20756e706175736500000000000000000000000000000000000000000060648201526084016103c1565b61062261061c60017f54176ff9944c4784e5857ec4e5ef560a462c483bf534eda43f91bb01a470b1b76112b6565b60009055565b6040517fa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d1693390600090a1565b600061052761052360017fd30e835d3f35624761057ff5b27d558f97bd5be034621e62240e5c0b784abe696112b6565b6104f060017f704ae3ec629461681409737f623e0cebb30122362e8cb04e0a0d3581d958db7d6112b6565b6104f060017f12c56161f16f492fd4016a16e534c3a2bcceceb7f70ec9bb75867affe33703166112b6565b600061052761052360017f12c56161f16f492fd4016a16e534c3a2bcceceb7f70ec9bb75867affe33703166112b6565b60004261073461052360017f54176ff9944c4784e5857ec4e5ef560a462c483bf534eda43f91bb01a470b1b76112b6565b11905090565b600061076a61052360017f0e2f5ebd54326cdea9bf943c0fc37413dccba70cdeb76374557a8f757e8983906112b6565b73ffffffffffffffffffffffffffffffffffffffff16905090565b60008061079183610db7565b60009081526001602052604090205460ff169392505050565b6104f060017f54176ff9944c4784e5857ec4e5ef560a462c483bf534eda43f91bb01a470b1b76112b6565b6104f060017f0e2f5ebd54326cdea9bf943c0fc37413dccba70cdeb76374557a8f757e8983906112b6565b6108086106d3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108c2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f5375706572636861696e436f6e6669673a206f6e6c7920696e69746961746f7260448201527f2063616e206164642073657175656e636572000000000000000000000000000060648201526084016103c1565b60006108cd82610db7565b600081815260016020819052604090912080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169091179055905060055b7f7b743789cff01dafdeae47739925425aab5dfd02d0c8229e4a508bcd2b9f42bb8360405160200161096691908151815260209182015173ffffffffffffffffffffffffffffffffffffffff169181019190915260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529082905261099e91611054565b60405180910390a25050565b6104f060017fd30e835d3f35624761057ff5b27d558f97bd5be034621e62240e5c0b784abe696112b6565b600061052761052360017f704ae3ec629461681409737f623e0cebb30122362e8cb04e0a0d3581d958db7d6112b6565b610a0d6104f3565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ac7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603960248201527f5375706572636861696e436f6e6669673a206f6e6c792073797374656d4f776e60448201527f65722063616e2072656d6f766520612073657175656e6365720000000000000060648201526084016103c1565b6000610ad282610db7565b600081815260016020526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690559050600661090c565b600054600290610100900460ff16158015610b31575060005460ff8083169116105b610bbd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016103c1565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff831617610100179055610bf787610e10565b610c0086610ecd565b610c0985610f02565b610c1284610f37565b610c1b83610f6c565b60005b8251811015610d4f576000610c4b848381518110610c3e57610c3e6112e5565b6020026020010151610db7565b600081815260016020819052604090912080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169091179055905060057f7b743789cff01dafdeae47739925425aab5dfd02d0c8229e4a508bcd2b9f42bb858481518110610cbc57610cbc6112e5565b6020026020010151604051602001610cfc91908151815260209182015173ffffffffffffffffffffffffffffffffffffffff169181019190915260400190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052610d3491611054565b60405180910390a25080610d4781611314565b915050610c1e565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150505050505050565b600081600001518260200151604051602001610df392919091825273ffffffffffffffffffffffffffffffffffffffff16602082015260400190565b604051602081830303815290604052805190602001209050919050565b610e43610e3e60017fe5134cb7d217efbc8c357a6644e3c656a6235651a8f25717e410cbf378e577536112b6565b829055565b60005b6040805173ffffffffffffffffffffffffffffffffffffffff841660208201527f7b743789cff01dafdeae47739925425aab5dfd02d0c8229e4a508bcd2b9f42bb91015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815290829052610ec291611054565b60405180910390a250565b610efb610e3e60017f12c56161f16f492fd4016a16e534c3a2bcceceb7f70ec9bb75867affe33703166112b6565b6001610e46565b610f30610e3e60017f704ae3ec629461681409737f623e0cebb30122362e8cb04e0a0d3581d958db7d6112b6565b6002610e46565b610f65610e3e60017fd30e835d3f35624761057ff5b27d558f97bd5be034621e62240e5c0b784abe696112b6565b6003610e46565b610f9a610e3e60017f0e2f5ebd54326cdea9bf943c0fc37413dccba70cdeb76374557a8f757e8983906112b6565b60047f7b743789cff01dafdeae47739925425aab5dfd02d0c8229e4a508bcd2b9f42bb82604051602001610e8a91815260200190565b600060208284031215610fe257600080fd5b5035919050565b6000815180845260005b8181101561100f57602081850181015186830182015201610ff3565b81811115611021576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006110676020830184610fe9565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156110e4576110e461106e565b604052919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461111057600080fd5b919050565b60006040828403121561112757600080fd5b6040516040810181811067ffffffffffffffff8211171561114a5761114a61106e565b60405282358152905080611160602084016110ec565b60208201525092915050565b60006040828403121561117e57600080fd5b6110678383611115565b60008060008060008060c087890312156111a157600080fd5b6111aa876110ec565b955060206111b98189016110ec565b955060406111c8818a016110ec565b95506111d660608a016110ec565b94506080890135935060a089013567ffffffffffffffff808211156111fa57600080fd5b818b0191508b601f83011261120e57600080fd5b8135818111156112205761122061106e565b61122e858260051b0161109d565b818152858101925060069190911b83018501908d82111561124e57600080fd5b928501925b81841015611274576112658e85611115565b83529284019291850191611253565b8096505050505050509295509295509295565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000828210156112c8576112c8611287565b500390565b600082198211156112e0576112e0611287565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361134557611345611287565b506001019056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(SuperchainConfigStorageLayoutJSON), SuperchainConfigStorageLayout); err != nil {
		panic(err)
	}

	layouts["SuperchainConfig"] = SuperchainConfigStorageLayout
	deployedBytecodes["SuperchainConfig"] = SuperchainConfigDeployedBin
}
