package BLC

type ProoFOfWork struct {
	Block *Block

}

func (prooFOfWork * ProoFOfWork) Run() ([]byte, int64) {

	return nil,0
}

func NewProofOfWork(block  *Block) *ProoFOfWork{

	return &ProoFOfWork{block}
}