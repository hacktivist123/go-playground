package main

type RequestError struct {
	HTTPCode int
	Body     string
	Err      string
}

func (r RequestError) Error() string {
	return r.Err
}
