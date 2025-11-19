package conn

type Conn struct{}

func (c *Conn) Close() error {
	return nil
}
