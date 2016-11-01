package ring
type Ring struct {
	data [][]byte
	pos int
}

func New(capcity int) *Ring {
	d := make([][]byte, capcity)
	return &Ring{
		d,
		0,
	}
}

func (r *Ring) Set(str []byte)  {
	r.data[r.pos] = str
	r.pos = (r.pos + 1) % len(r.data)
}

func (r *Ring) Get(index int) []byte  {
	return r.data[index % len(r.data)]
}
