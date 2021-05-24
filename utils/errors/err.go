package errors

type Err struct {
	Errcode int
	Reason  string
	Data    []byte
}

func (err Err) Error() string {
	return err.Reason
}

func (err Err) Code() int {
	return err.Errcode
}

func (err *Err) GetData() []byte {
	return err.Data
}

func (err *Err) SetData(data []byte) error {
	err.Data = data
	return err
}

func New(errcode int, reason string) error {
	err := new(Err)
	err.Errcode = errcode
	err.Reason = reason
	return err
}
