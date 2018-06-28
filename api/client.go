package api

type Client struct {
	//conn Connection

	Apps
}

func NewClient(url, token string) *Client {
	conn := newConn(url, token)

	return &Client{
		Apps: Apps{conn: conn},
	}
}
