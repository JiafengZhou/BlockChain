package BLC

//UTXO模型

type Transaction struct {

	//1.交易Hash
	TxHash []byte

	//2.输入
	Vins []*TXInput

	//3.输出
	Vouts []*TXOutout
}
