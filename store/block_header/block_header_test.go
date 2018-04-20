package block_header

import (
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/maichain/eth-indexer/common"
	"github.com/maichain/eth-indexer/model"
	"github.com/maichain/mapi/base/test"
	"github.com/maichain/mapi/types/reflect"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func makeHeader(number int64, hashHex string) model.Header {
	return model.Header{
		Hash:        common.HexToBytes(hashHex),
		ParentHash:  common.HexToBytes("0x35b9253b70be351059982e8d6a218146a18ef9b723e560c7efc540629b4e75f2"),
		UncleHash:   common.HexToBytes("0x2d6159f94932bd669c7161e2563ea4cc0fbf848dd59adbed7df3da74072edd50"),
		Coinbase:    common.HexToBytes("0xB287a379e6caCa6732E50b88D23c290aA990A892"),
		Root:        common.HexToBytes("0x86f9a7ccb763958d0f6c01ea89b7a49eb5a3a8aff0f998ff514b97ad1c4e1fd6"),
		TxHash:      common.HexToBytes("0x3f28c6504aa57084da641571cd710e092c716979dac2664f70fc62cd9d792a4b"),
		ReceiptHash: common.HexToBytes("0xad2ad2d0fca28f18d0d9fedc7ec2ab4b97277546c212f67519314bfb30f56736"),
		Difficulty:  927399944,
		Number:      number,
		GasLimit:    810000,
		GasUsed:     809999,
		Time:        123456789,
		MixDigest:   []byte{11, 23, 45},
		Nonce:       []byte{12, 13, 56, 77},
	}
}

var _ = Describe("Block Header Database Test", func() {
	var (
		mysql *test.MySQLContainer
		db    *gorm.DB
	)
	BeforeSuite(func() {
		var err error
		mysql, err = test.NewMySQLContainer("quay.io/amis/eth-indexer-db-migration")
		Expect(mysql).ShouldNot(BeNil())
		Expect(err).Should(Succeed())
		Expect(mysql.Start()).Should(Succeed())

		db, err = gorm.Open("mysql", mysql.URL)
		Expect(err).Should(Succeed())
		Expect(db).ShouldNot(BeNil())

		db.LogMode(os.Getenv("ENABLE_DB_LOG_IN_TEST") != "")
	})

	AfterSuite(func() {
		mysql.Stop()
	})

	It("should get latest header", func() {
		store := NewWithDB(db)
		data1 := makeHeader(1000300, "0x58bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b")
		data2 := makeHeader(1000301, "0x68bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b")

		store.Insert(&data1)
		store.Insert(&data2)

		filter := model.Header{Hash: data1.Hash}
		result, err := store.Find(&filter)
		Expect(err).Should(Succeed())
		Expect(result[0].Number).Should(Equal(data1.Number))

		filter = model.Header{Number: data2.Number}
		result, err = store.Find(&filter)
		Expect(err).Should(Succeed())
		Expect(result[0].Number).Should(Equal(data2.Number))

		lastResult, err := store.Last()
		Expect(err).Should(Succeed())
		Expect(reflect.DeepEqual(*lastResult, data2)).Should(BeTrue())
	})

	It("should insert one new record in database", func() {
		By("insert new one header")
		store := NewWithDB(db)
		data := makeHeader(1000302, "0x78bb59babd8fd8299b22acb997832a75d7b6b666579f80cc281764342f2b373b")
		err := store.Insert(&data)
		Expect(err).Should(Succeed())

		By("failed to insert again")
		err = store.Insert(&data)
		Expect(err).ShouldNot(BeNil())
	})
})

func TestBlockHeader(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Block Header Database Test")
}
