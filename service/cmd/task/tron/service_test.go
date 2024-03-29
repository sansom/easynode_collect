package tron

import (
	"github.com/sirupsen/logrus"
	"github.com/sunjiangjun/xlog"
	"github.com/uduncloud/easynode_collect/config"
	"github.com/uduncloud/easynode_collect/service"
	"log"
	"testing"
)

func Init() (service.BlockChainInterface, config.Config, *xlog.XLog) {
	cfg := config.LoadConfig("./../../../../config_tron.json")
	x := xlog.NewXLogger()
	return NewService(cfg.Chains[0], cfg.TaskDb, cfg.SourceDb, x), cfg, x
}

func TestService_GetTx(t *testing.T) {

	s, cfg, x := Init()

	tx := s.GetTx("0x72fd440ff0542c2c28db762b4268f126c57f0fdf6daf69258cb9a306e26723e8", cfg.Chains[0].TxTask, x.WithFields(logrus.Fields{}))

	log.Printf("%+v\n", tx)
}

func TestService_GetBlockByNumber(t *testing.T) {
	s, cfg, x := Init()
	b, t1 := s.GetBlockByNumber("0xF9CC56", cfg.Chains[0].BlockTask, x.WithFields(logrus.Fields{}))
	log.Println(b)
	log.Println(t1)
}

func TestService_GetReceiptByBlock(t *testing.T) {
	s, cfg, x := Init()
	r := s.GetReceiptByBlock("", "0xF9CC56", cfg.Chains[0].ReceiptTask, x.WithFields(logrus.Fields{}))
	log.Println(r)
}

func TestService_GetReceipt(t *testing.T) {
	s, cfg, x := Init()
	r := s.GetReceipt("0x72fd440ff0542c2c28db762b4268f126c57f0fdf6daf69258cb9a306e26723e8", cfg.Chains[0].ReceiptTask, x.WithFields(logrus.Fields{}))
	log.Println(r)
}
