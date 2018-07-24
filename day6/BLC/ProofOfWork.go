package BLC
import "math/big"

const targetBit = 16 //256 位HASH前必须有16个0  难度
type ProoFOfWork struct {
	Block *Block
	target *big.Int // 大数据存储
}



func (prooFOfWork * ProoFOfWork) Run() ([]byte, int64) {

	return nil,0
}

func NewProofOfWork(block  *Block) *ProoFOfWork{
	target := big.NewInt(1)

	target = target.Lsh(target,256 -targetBit)

	return &ProoFOfWork{block, target}
}