package model

type Matrix struct {
	Rows []Row `xml:"row"`
}

type Row struct {
	Cols []int `xml:"col"`
}

type MatrixArgs struct {
	Matrix Matrix `xml:"matrix"`
}

type MatrixReply struct {
	Original  Matrix `xml:"original"`
	MinValue  int    `xml:"minValue"`
	Processed Matrix `xml:"processed"`
}
