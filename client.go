package revendamais

type Client struct {
	StoreId string
	BaseUrl string

	Posts *PostService
}