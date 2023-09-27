// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const SystemConfigStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1003,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1005,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"overhead\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"scalar\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"batcherHash\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_bytes32\"},{\"astId\":1008,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"gasLimit\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_uint64\"},{\"astId\":1009,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_resourceConfig\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_struct(ResourceConfig)1013_storage\"},{\"astId\":1010,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"startBlock\",\"offset\":0,\"slot\":\"106\",\"type\":\"t_uint256\"},{\"astId\":1011,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"superchainConfig\",\"offset\":0,\"slot\":\"107\",\"type\":\"t_contract(SuperchainConfig)1012\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_contract(SuperchainConfig)1012\":{\"encoding\":\"inplace\",\"label\":\"contract SuperchainConfig\",\"numberOfBytes\":\"20\"},\"t_struct(ResourceConfig)1013_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ResourceMetering.ResourceConfig\",\"numberOfBytes\":\"32\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var SystemConfigStorageLayout = new(solc.StorageLayout)

var SystemConfigDeployedBin = "0x608060405234801561001057600080fd5b50600436106102415760003560e01c8063935f029e11610145578063dac6e63a116100bd578063f45e65d81161008c578063f8c68de011610071578063f8c68de0146105a5578063fd32aa0f146105ad578063ffa1ad74146105b557600080fd5b8063f45e65d814610588578063f68016b71461059157600080fd5b8063dac6e63a1461055c578063dc8daf9314610564578063e81b2c6d1461056c578063f2fde38b1461057557600080fd5b8063bc49ce5f11610114578063c71973f6116100f9578063c71973f614610402578063cc731b0214610415578063d908445b1461054957600080fd5b8063bc49ce5f146103f2578063c4e8ddfa146103fa57600080fd5b8063935f029e146103bc5780639b7d7f0a146103cf578063a7119869146103d7578063b40a817c146103df57600080fd5b806348cd4cb1116101d857806354fd4d50116101a757806361d157681161018c57806361d157681461038e578063715018a6146103965780638da5cb5b1461039e57600080fd5b806354fd4d501461033d5780635d73369c1461038657600080fd5b806348cd4cb1146102e45780634add321d146102ed5780634d9f15591461030e5780634f16540b1461031657600080fd5b806319f5cea81161021457806319f5cea81461029f5780631fd19ee1146102a75780632a534ab3146102af57806335e80ab3146102c457600080fd5b806306c9265714610246578063078f29cf146102615780630a49cb031461028e5780630c18c16214610296575b600080fd5b61024e6105bd565b6040519081526020015b60405180910390f35b6102696105eb565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610258565b610269610624565b61024e60655481565b61024e610654565b61026961067f565b6102c26102bd366004611a63565b6106a9565b005b606b546102699073ffffffffffffffffffffffffffffffffffffffff1681565b61024e606a5481565b6102f5610822565b60405167ffffffffffffffff9091168152602001610258565b610269610848565b61024e7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0881565b6103796040518060400160405280600581526020017f322e302e3000000000000000000000000000000000000000000000000000000081525081565b6040516102589190611afe565b61024e610878565b61024e6108a3565b6102c26108ce565b60335473ffffffffffffffffffffffffffffffffffffffff16610269565b6102c26103ca366004611b18565b6108e2565b6102696108f8565b610269610928565b6102c26103ed366004611b57565b610958565b61024e61096c565b610269610997565b6102c2610410366004611cc8565b6109c7565b6104d96040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c08101825260695463ffffffff8082168352640100000000820460ff9081166020850152650100000000008304169383019390935266010000000000008104831660608301526a0100000000000000000000810490921660808201526e0100000000000000000000000000009091046fffffffffffffffffffffffffffffffff1660a082015290565b6040516102589190600060c08201905063ffffffff80845116835260ff602085015116602084015260ff6040850151166040840152806060850151166060840152806080850151166080840152506fffffffffffffffffffffffffffffffff60a08401511660a083015292915050565b6102c2610557366004611ce4565b6109d8565b610269610d97565b6102c2610dc7565b61024e60675481565b6102c2610583366004611e43565b610f3e565b61024e60665481565b6068546102f59067ffffffffffffffff1681565b61024e610ff2565b61024e61101d565b61024e600081565b6105e860017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d611e8f565b81565b600061061f61061b60017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad6377611e8f565b5490565b905090565b600061061f61061b60017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad611e8f565b6105e860017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a8611e8f565b600061061f7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c085490565b6106b1611048565b6106ba816110c9565b6106c382611185565b6040805180820190915282815273ffffffffffffffffffffffffffffffffffffffff8216602082015260006106f7826111ad565b606b546040517fd92a09bc0000000000000000000000000000000000000000000000000000000081526004810183905291925073ffffffffffffffffffffffffffffffffffffffff169063d92a09bc90602401602060405180830381865afa158015610767573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061078b9190611ea6565b61081c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f53797374656d436f6e6669673a2053657175656e6365722068617368206e6f7460448201527f20666f756e6420696e205375706572636861696e20616c6c6f77206c6973740060648201526084015b60405180910390fd5b50505050565b60695460009061061f9063ffffffff6a0100000000000000000000820481169116611ec8565b600061061f61061b60017fe52a667f71ec761b9b381c7b76ca9b852adf7e8905da0e0ad49986a0a6871816611e8f565b6105e860017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce9580637611e8f565b6105e860017fe52a667f71ec761b9b381c7b76ca9b852adf7e8905da0e0ad49986a0a6871816611e8f565b6108d6611048565b6108e06000611206565b565b6108ea611048565b6108f4828261127d565b5050565b600061061f61061b60017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d611e8f565b600061061f61061b60017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce9580637611e8f565b610960611048565b6109698161130e565b50565b6105e860017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc598611e8f565b600061061f61061b60017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a8611e8f565b6109cf611048565b610969816113ec565b600054600290610100900460ff161580156109fa575060005460ff8083169116105b610a86576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610813565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff831617610100179055610abf611860565b610ac88c610f3e565b606b80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8d16179055610b1188611185565b610b1b8a8a61127d565b610b248761130e565b610b2d866110c9565b610b5f83610b5c60017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc598611e8f565b55565b8151610b9090610b5c60017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce9580637611e8f565b6020820151610bc490610b5c60017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a8611e8f565b6040820151610bf890610b5c60017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad6377611e8f565b6060820151610c2c90610b5c60017fe52a667f71ec761b9b381c7b76ca9b852adf7e8905da0e0ad49986a0a6871816611e8f565b6080820151610c6090610b5c60017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad611e8f565b60a0820151610c9490610b5c60017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d611e8f565b610c9d846118ff565b610ca6856113ec565b610cae610822565b67ffffffffffffffff168767ffffffffffffffff161015610d2b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f77006044820152606401610813565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050505050505050505050565b600061061f61061b60017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc598611e8f565b600060405180604001604052806067548152602001610de461067f565b73ffffffffffffffffffffffffffffffffffffffff16905290506000610e09826111ad565b606b546040517fd92a09bc0000000000000000000000000000000000000000000000000000000081526004810183905291925073ffffffffffffffffffffffffffffffffffffffff169063d92a09bc90602401602060405180830381865afa158015610e79573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e9d9190611ea6565b15610f2a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f53797374656d436f6e6669673a2063616e6e6f742072656d6f766520616c6c6f60448201527f7765642073657175656e6365722e0000000000000000000000000000000000006064820152608401610813565b610f3460006110c9565b6108f46000611185565b610f46611048565b73ffffffffffffffffffffffffffffffffffffffff8116610fe9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610813565b61096981611206565b6105e860017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad6377611e8f565b6105e860017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad611e8f565b60335473ffffffffffffffffffffffffffffffffffffffff1633146108e0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610813565b6110f1817f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0855565b6040805173ffffffffffffffffffffffffffffffffffffffff8316602082015260009101604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052905060035b60007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be836040516111799190611afe565b60405180910390a35050565b6067819055604080516020808201849052825180830390910181529082019091526000611148565b6000816000015182602001516040516020016111e992919091825273ffffffffffffffffffffffffffffffffffffffff16602082015260400190565b604051602081830303815290604052805190602001209050919050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b606582905560668190556040805160208101849052908101829052600090606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529050600160007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be836040516113019190611afe565b60405180910390a3505050565b611316610822565b67ffffffffffffffff168167ffffffffffffffff161015611393576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f77006044820152606401610813565b606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff83169081179091556040805160208082019390935281518082039093018352810190526002611148565b8060a001516fffffffffffffffffffffffffffffffff16816060015163ffffffff16111561149c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f53797374656d436f6e6669673a206d696e206261736520666565206d7573742060448201527f6265206c657373207468616e206d6178206261736500000000000000000000006064820152608401610813565b6001816040015160ff1611611533576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a2064656e6f6d696e61746f72206d757374206260448201527f65206c6172676572207468616e203100000000000000000000000000000000006064820152608401610813565b6068546080820151825167ffffffffffffffff909216916115549190611ef4565b63ffffffff1611156115c2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f77006044820152606401610813565b6000816020015160ff1611611659576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a20656c6173746963697479206d756c7469706c60448201527f6965722063616e6e6f74206265203000000000000000000000000000000000006064820152608401610813565b8051602082015163ffffffff82169160ff90911690611679908290611f13565b6116839190611f5d565b63ffffffff1614611716576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f53797374656d436f6e6669673a20707265636973696f6e206c6f73732077697460448201527f6820746172676574207265736f75726365206c696d69740000000000000000006064820152608401610813565b805160698054602084015160408501516060860151608087015160a09097015163ffffffff9687167fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000009095169490941764010000000060ff94851602177fffffffffffffffffffffffffffffffffffffffffffff0000000000ffffffffff166501000000000093909216929092027fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff1617660100000000000091851691909102177fffff0000000000000000000000000000000000000000ffffffffffffffffffff166a010000000000000000000093909416929092027fffff00000000000000000000000000000000ffffffffffffffffffffffffffff16929092176e0100000000000000000000000000006fffffffffffffffffffffffffffffffff90921691909102179055565b600054610100900460ff166118f7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610813565b6108e06119a1565b606a541561198f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603860248201527f53797374656d436f6e6669673a2063616e6e6f74206f7665727269646520616e60448201527f20616c72656164792073657420737461727420626c6f636b00000000000000006064820152608401610813565b801561199a57606a55565b43606a5550565b600054610100900460ff16611a38576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610813565b6108e033611206565b73ffffffffffffffffffffffffffffffffffffffff8116811461096957600080fd5b60008060408385031215611a7657600080fd5b823591506020830135611a8881611a41565b809150509250929050565b6000815180845260005b81811015611ab957602081850181015186830182015201611a9d565b81811115611acb576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611b116020830184611a93565b9392505050565b60008060408385031215611b2b57600080fd5b50508035926020909101359150565b803567ffffffffffffffff81168114611b5257600080fd5b919050565b600060208284031215611b6957600080fd5b611b1182611b3a565b60405160c0810167ffffffffffffffff81118282101715611bbc577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405290565b803563ffffffff81168114611b5257600080fd5b803560ff81168114611b5257600080fd5b600060c08284031215611bf957600080fd5b60405160c0810181811067ffffffffffffffff82111715611c43577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604052905080611c5283611bc2565b8152611c6060208401611bd6565b6020820152611c7160408401611bd6565b6040820152611c8260608401611bc2565b6060820152611c9360808401611bc2565b608082015260a08301356fffffffffffffffffffffffffffffffff81168114611cbb57600080fd5b60a0919091015292915050565b600060c08284031215611cda57600080fd5b611b118383611be7565b60008060008060008060008060008060008b8d036102a0811215611d0757600080fd5b8c35611d1281611a41565b9b5060208d0135611d2281611a41565b9a5060408d0135995060608d0135985060808d01359750611d4560a08e01611b3a565b965060c08d0135611d5581611a41565b9550611d648e60e08f01611be7565b94506101a08d013593506101c08d0135611d7d81611a41565b925060c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe2082011215611daf57600080fd5b50611db8611b72565b6101e08d0135611dc781611a41565b81526102008d0135611dd881611a41565b60208201526102208d0135611dec81611a41565b60408201526102408d0135611e0081611a41565b60608201526102608d0135611e1481611a41565b60808201526102808d0135611e2881611a41565b8060a083015250809150509295989b509295989b9093969950565b600060208284031215611e5557600080fd5b8135611b1181611a41565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015611ea157611ea1611e60565b500390565b600060208284031215611eb857600080fd5b81518015158114611b1157600080fd5b600067ffffffffffffffff808316818516808303821115611eeb57611eeb611e60565b01949350505050565b600063ffffffff808316818516808303821115611eeb57611eeb611e60565b600063ffffffff80841680611f51577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600063ffffffff80831681851681830481118215151615611f8057611f80611e60565b0294935050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(SystemConfigStorageLayoutJSON), SystemConfigStorageLayout); err != nil {
		panic(err)
	}

	layouts["SystemConfig"] = SystemConfigStorageLayout
	deployedBytecodes["SystemConfig"] = SystemConfigDeployedBin
}
