import(
	"github.com/stretchr/testify"
	"testing"
)

var testStoreService = &StoreService{}

func init(
	testStoreService = IntializeStore()
)

func testingStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func testingInsertationAndRetrieval(t *testing.T) {
	Shorturl := "codimow"
	originalURL := "https://github.com/Codimow"
	userID := 1

	SaveURLMapping(Shorturl, originalURL, userID)

	assert.Equal(t, initalLink, retrievalUrl)
}