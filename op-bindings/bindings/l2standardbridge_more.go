// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2StandardBridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"spacer_0_2_30\",\"offset\":2,\"slot\":\"0\",\"type\":\"t_bytes30\"},{\"astId\":1003,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"spacer_1_0_20\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1005,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"messenger\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_contract(CrossDomainMessenger)1008\"},{\"astId\":1006,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"otherBridge\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_contract(StandardBridge)1009\"},{\"astId\":1007,\"contract\":\"src/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_array(t_uint256)45_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)45_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[45]\",\"numberOfBytes\":\"1440\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes30\":{\"encoding\":\"inplace\",\"label\":\"bytes30\",\"numberOfBytes\":\"30\"},\"t_contract(CrossDomainMessenger)1008\":{\"encoding\":\"inplace\",\"label\":\"contract CrossDomainMessenger\",\"numberOfBytes\":\"20\"},\"t_contract(StandardBridge)1009\":{\"encoding\":\"inplace\",\"label\":\"contract StandardBridge\",\"numberOfBytes\":\"20\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L2StandardBridgeStorageLayout = new(solc.StorageLayout)

var L2StandardBridgeDeployedBin = "0x6080604052600436106101635760003560e01c80635c975abb116100c0578063927ede2d11610074578063c4d66de811610059578063c4d66de8146104b7578063c89701a2146104d7578063e11013dd1461050457600080fd5b8063927ede2d14610479578063a3a79548146104a457600080fd5b80637f46ddb2116100a55780637f46ddb2146102ba57806387087623146104135780638f601f661461043357600080fd5b80635c975abb146103ec578063662a633a1461040057600080fd5b806336c717c1116101175780634397dfef116100fc5780634397dfef14610333578063540abf731461037657806354fd4d501461039657600080fd5b806336c717c1146102ba5780633cb747bf1461030657600080fd5b80631635f5fd116101485780631635f5fd1461026a578063213268491461027d57806332b7006d146102a757600080fd5b80630166a07a1461023757806309fc88431461025757600080fd5b3661023257333b156101fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084015b60405180910390fd5b61023073deaddeaddeaddeaddeaddeaddeaddeaddead000033333462030d4060405180602001604052806000815250610533565b005b600080fd5b34801561024357600080fd5b50610230610252366004612703565b61060e565b6102306102653660046127b4565b6109b0565b610230610278366004612807565b610a87565b34801561028957600080fd5b50610292610f6e565b60405190151581526020015b60405180910390f35b6102306102b536600461287a565b610fad565b3480156102c657600080fd5b5060045473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161029e565b34801561031257600080fd5b506003546102e19073ffffffffffffffffffffffffffffffffffffffff1681565b34801561033f57600080fd5b50610348611087565b6040805173ffffffffffffffffffffffffffffffffffffffff909316835260ff90911660208301520161029e565b34801561038257600080fd5b506102306103913660046128ce565b611115565b3480156103a257600080fd5b506103df6040518060400160405280600581526020017f312e392e3000000000000000000000000000000000000000000000000000000081525081565b60405161029e91906129bb565b3480156103f857600080fd5b506000610292565b61023061040e366004612703565b61115a565b34801561041f57600080fd5b5061023061042e3660046129ce565b6111cd565b34801561043f57600080fd5b5061046b61044e366004612a51565b600260209081526000928352604080842090915290825290205481565b60405190815260200161029e565b34801561048557600080fd5b5060035473ffffffffffffffffffffffffffffffffffffffff166102e1565b6102306104b23660046129ce565b6112a1565b3480156104c357600080fd5b506102306104d2366004612a8a565b6112e5565b3480156104e357600080fd5b506004546102e19073ffffffffffffffffffffffffffffffffffffffff1681565b610230610512366004612aa7565b61148e565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b7fffffffffffffffffffffffff215221522152215221522152215221522153000073ffffffffffffffffffffffffffffffffffffffff8716016105825761057d85858585856114d7565b610606565b60008673ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105f39190612b0a565b905061060487828888888888611736565b505b505050505050565b60035473ffffffffffffffffffffffffffffffffffffffff16331480156106e1575060048054600354604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff938416949390921692636e296e459282820192602092908290030181865afa1580156106a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106c99190612b0a565b73ffffffffffffffffffffffffffffffffffffffff16145b610793576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4016101f3565b61079c87611a61565b156108ea576107ab8787611ac3565b61085d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a4016101f3565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590528816906340c10f1990604401600060405180830381600087803b1580156108cd57600080fd5b505af11580156108e1573d6000803e3d6000fd5b5050505061096c565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a1683529290522054610928908490612b56565b73ffffffffffffffffffffffffffffffffffffffff8089166000818152600260209081526040808320948c168352939052919091209190915561096c908585611be3565b610604878787878787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611cb792505050565b333b15610a3f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101f3565b610a823333348686868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506114d792505050565b505050565b60035473ffffffffffffffffffffffffffffffffffffffff1633148015610b5a575060048054600354604080517f6e296e45000000000000000000000000000000000000000000000000000000008152905173ffffffffffffffffffffffffffffffffffffffff938416949390921692636e296e459282820192602092908290030181865afa158015610b1e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b429190612b0a565b73ffffffffffffffffffffffffffffffffffffffff16145b610c0c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4016101f3565b610c14610f6e565b15610ca1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2063616e6e6f742062726964676520455460448201527f48207769746820637573746f6d2067617320746f6b656e00000000000000000060648201526084016101f3565b823414610d30576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f5374616e646172644272696467653a20616d6f756e742073656e7420646f657360448201527f206e6f74206d6174636820616d6f756e7420726571756972656400000000000060648201526084016101f3565b3073ffffffffffffffffffffffffffffffffffffffff851603610dd5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f207360448201527f656c66000000000000000000000000000000000000000000000000000000000060648201526084016101f3565b60035473ffffffffffffffffffffffffffffffffffffffff90811690851603610e80576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f206d60448201527f657373656e67657200000000000000000000000000000000000000000000000060648201526084016101f3565b610ec285858585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611d4592505050565b6000610edf855a8660405180602001604052806000815250611de6565b905080610606576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a20455448207472616e736665722066616960448201527f6c6564000000000000000000000000000000000000000000000000000000000060648201526084016101f3565b600080610f79611087565b5073ffffffffffffffffffffffffffffffffffffffff1673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee141592915050565b333b1561103c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101f3565b611080853333878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061053392505050565b5050505050565b60008073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16634397dfef6040518163ffffffff1660e01b81526004016040805180830381865afa1580156110e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061110c9190612b6d565b90939092509050565b61060487873388888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061173692505050565b73ffffffffffffffffffffffffffffffffffffffff87161580156111a7575073ffffffffffffffffffffffffffffffffffffffff861673deaddeaddeaddeaddeaddeaddeaddeaddead0000145b156111be576111b98585858585610a87565b610604565b6106048688878787878761060e565b333b1561125c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101f3565b61060686863333888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061173692505050565b610606863387878787878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061053392505050565b600054610100900460ff16158080156113055750600054600160ff909116105b8061131f5750303b15801561131f575060005460ff166001145b6113ab576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084016101f3565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561140957600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61142773420000000000000000000000000000000000000783611e00565b801561148a57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b6114d13385348686868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506114d792505050565b50505050565b6114df610f6e565b1561156c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2063616e6e6f742062726964676520455460448201527f48207769746820637573746f6d2067617320746f6b656e00000000000000000060648201526084016101f3565b8234146115fb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f5374616e646172644272696467653a206272696467696e6720455448206d757360448201527f7420696e636c7564652073756666696369656e74204554482076616c7565000060648201526084016101f3565b61160785858584611eea565b60035460045460405173ffffffffffffffffffffffffffffffffffffffff92831692633dbb202b9287929116907f1635f5fd000000000000000000000000000000000000000000000000000000009061166a908b908b9086908a90602401612ba2565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e086901b90921682526116fd92918890600401612beb565b6000604051808303818588803b15801561171657600080fd5b505af115801561172a573d6000803e3d6000fd5b50505050505050505050565b61173f87611a61565b1561188d5761174e8787611ac3565b611800576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a4016101f3565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015260248201859052881690639dc29fac90604401600060405180830381600087803b15801561187057600080fd5b505af1158015611884573d6000803e3d6000fd5b50505050611921565b6118af73ffffffffffffffffffffffffffffffffffffffff8816863086611f8b565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600260209081526040808320938a16835292905220546118ed908490612c30565b73ffffffffffffffffffffffffffffffffffffffff8089166000908152600260209081526040808320938b16835292905220555b61192f878787878786611fe9565b60035460045460405173ffffffffffffffffffffffffffffffffffffffff92831692633dbb202b9216907f0166a07a0000000000000000000000000000000000000000000000000000000090611993908b908d908c908c908c908b90602401612c48565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e085901b9092168252611a2692918790600401612beb565b600060405180830381600087803b158015611a4057600080fd5b505af1158015611a54573d6000803e3d6000fd5b5050505050505050505050565b6000611a8d827f1d1d8b6300000000000000000000000000000000000000000000000000000000612077565b80611abd5750611abd827fec4fc8e300000000000000000000000000000000000000000000000000000000612077565b92915050565b6000611aef837f1d1d8b6300000000000000000000000000000000000000000000000000000000612077565b15611b98578273ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa158015611b3f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b639190612b0a565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16149050611abd565b8273ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa158015611b3f573d6000803e3d6000fd5b60405173ffffffffffffffffffffffffffffffffffffffff8316602482015260448101829052610a829084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009093169290921790915261209a565b8373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89868686604051611d2f93929190612ca3565b60405180910390a46106068686868686866121a6565b8373ffffffffffffffffffffffffffffffffffffffff1673deaddeaddeaddeaddeaddeaddeaddeaddead000073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89868686604051611dd293929190612ca3565b60405180910390a46114d18484848461222e565b600080600080845160208601878a8af19695505050505050565b600054610100900460ff16611e97576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016101f3565b6003805473ffffffffffffffffffffffffffffffffffffffff9384167fffffffffffffffffffffffff00000000000000000000000000000000000000009182161790915560048054929093169116179055565b8373ffffffffffffffffffffffffffffffffffffffff1673deaddeaddeaddeaddeaddeaddeaddeaddead000073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e868686604051611f7793929190612ca3565b60405180910390a46114d18484848461229b565b60405173ffffffffffffffffffffffffffffffffffffffff808516602483015283166044820152606481018290526114d19085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401611c35565b8373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e86868660405161206193929190612ca3565b60405180910390a46106068686868686866122fa565b600061208283612372565b8015612093575061209383836123d6565b9392505050565b60006120fc826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166124a59092919063ffffffff16565b805190915015610a82578080602001905181019061211a9190612ce1565b610a82576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101f3565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167fd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd86868660405161221e93929190612ca3565b60405180910390a4505050505050565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d848460405161228d929190612d03565b60405180910390a350505050565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af5848460405161228d929190612d03565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf86868660405161221e93929190612ca3565b600061239e827f01ffc9a7000000000000000000000000000000000000000000000000000000006123d6565b8015611abd57506123cf827fffffffff000000000000000000000000000000000000000000000000000000006123d6565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d9150600051905082801561248e575060208210155b801561249a5750600081115b979650505050505050565b60606124b484846000856124bc565b949350505050565b60608247101561254e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101f3565b73ffffffffffffffffffffffffffffffffffffffff85163b6125cc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101f3565b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040516125f59190612d1c565b60006040518083038185875af1925050503d8060008114612632576040519150601f19603f3d011682016040523d82523d6000602084013e612637565b606091505b509150915061249a82828660608315612651575081612093565b8251156126615782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f391906129bb565b73ffffffffffffffffffffffffffffffffffffffff811681146126b757600080fd5b50565b60008083601f8401126126cc57600080fd5b50813567ffffffffffffffff8111156126e457600080fd5b6020830191508360208285010111156126fc57600080fd5b9250929050565b600080600080600080600060c0888a03121561271e57600080fd5b873561272981612695565b9650602088013561273981612695565b9550604088013561274981612695565b9450606088013561275981612695565b93506080880135925060a088013567ffffffffffffffff81111561277c57600080fd5b6127888a828b016126ba565b989b979a50959850939692959293505050565b803563ffffffff811681146127af57600080fd5b919050565b6000806000604084860312156127c957600080fd5b6127d28461279b565b9250602084013567ffffffffffffffff8111156127ee57600080fd5b6127fa868287016126ba565b9497909650939450505050565b60008060008060006080868803121561281f57600080fd5b853561282a81612695565b9450602086013561283a81612695565b935060408601359250606086013567ffffffffffffffff81111561285d57600080fd5b612869888289016126ba565b969995985093965092949392505050565b60008060008060006080868803121561289257600080fd5b853561289d81612695565b9450602086013593506128b26040870161279b565b9250606086013567ffffffffffffffff81111561285d57600080fd5b600080600080600080600060c0888a0312156128e957600080fd5b87356128f481612695565b9650602088013561290481612695565b9550604088013561291481612695565b9450606088013593506129296080890161279b565b925060a088013567ffffffffffffffff81111561277c57600080fd5b60005b83811015612960578181015183820152602001612948565b838111156114d15750506000910152565b60008151808452612989816020860160208601612945565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006120936020830184612971565b60008060008060008060a087890312156129e757600080fd5b86356129f281612695565b95506020870135612a0281612695565b945060408701359350612a176060880161279b565b9250608087013567ffffffffffffffff811115612a3357600080fd5b612a3f89828a016126ba565b979a9699509497509295939492505050565b60008060408385031215612a6457600080fd5b8235612a6f81612695565b91506020830135612a7f81612695565b809150509250929050565b600060208284031215612a9c57600080fd5b813561209381612695565b60008060008060608587031215612abd57600080fd5b8435612ac881612695565b9350612ad66020860161279b565b9250604085013567ffffffffffffffff811115612af257600080fd5b612afe878288016126ba565b95989497509550505050565b600060208284031215612b1c57600080fd5b815161209381612695565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015612b6857612b68612b27565b500390565b60008060408385031215612b8057600080fd5b8251612b8b81612695565b602084015190925060ff81168114612a7f57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff808716835280861660208401525083604083015260806060830152612be16080830184612971565b9695505050505050565b73ffffffffffffffffffffffffffffffffffffffff84168152606060208201526000612c1a6060830185612971565b905063ffffffff83166040830152949350505050565b60008219821115612c4357612c43612b27565b500190565b600073ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525083608083015260c060a0830152612c9760c0830184612971565b98975050505050505050565b73ffffffffffffffffffffffffffffffffffffffff84168152826020820152606060408201526000612cd86060830184612971565b95945050505050565b600060208284031215612cf357600080fd5b8151801515811461209357600080fd5b8281526040602082015260006124b46040830184612971565b60008251612d2e818460208701612945565b919091019291505056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(L2StandardBridgeStorageLayoutJSON), L2StandardBridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2StandardBridge"] = L2StandardBridgeStorageLayout
	deployedBytecodes["L2StandardBridge"] = L2StandardBridgeDeployedBin
	immutableReferences["L2StandardBridge"] = false
}
