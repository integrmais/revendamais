package revendamais

func NewClient(baseUrl, storeId string) *Client {
	c := &Client{
		StoreId: storeId,
		BaseUrl: baseUrl,
	}

	c.Posts = (*PostService)(c)

	return c
}