package funding

type Fond struct {
	 balance int
}

func NewFond (initialBalance int) *Fond{
 return &Fond{
	 balance: initialBalance,
 }
}

func (f *Fond) Balance() int {
	return f.balance
}
func (f *Fond) Withdraw (amount int){
	f.balance -=amount
}