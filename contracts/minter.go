package contracts

//Mint represents something that is mintable
type Minter interface {
	//Mint will try and mint
	Mint(count uint64) error
}
