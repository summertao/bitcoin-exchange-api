package test

/**
 *   about subscribe message: see https://www.bitmex.com/app/wsAPI
 *
 */

import (
	"fmt"

	"testing"

	bmrestfulapi   "github.com/summertao/bitcoin-exchange-api/bitmex/restfulApi"
	"github.com/summertao/bitcoin-exchange-api/bitmex"
	"github.com/summertao/bitcoin-exchange-api/utils"
	"datamesh.com/common/utils/jsonutils"
)


var (
	apikey = "apikey"
	apisecret = "apisecret"
)

func init() {
	apikey = "EF8Jzeb_-li0kqc5ExsOP3H2"
	apisecret = "4Kin9X74rt8lZv8XYqq6LmGU5mJACxjlNvr4X2GtAR_hYqcu"
}


var (
	orderapi *bmrestfulapi.OrderApi
	positionapi *bmrestfulapi.PositionApi

	configuration *bitmex.Configuration
	account utils.Platform = utils.Platform{}
)


func Test_restfulapi_position(t *testing.T) {
	var (
		po *bitmex.Position
		err error
	)

	configuration = bitmex.NewConfiguration(bmrestfulapi.APIClientImpl{"http://proxy:BTMM@gzhw.o2o.ac:58900"})
	positionapi = bmrestfulapi.NewPositionApi(configuration)

	account.Apikey = apikey
	account.Secretkey = apisecret
	//orderapi.Configuration.
	configuration.Host = "https://www.bitmex.com"
	configuration.BasePath = "/api/v1"
	configuration.Account = &account
	configuration.ExpireTime = 5

	po, _, err = positionapi.PositionUpdateLeverage(bitmex.XBTUSD, 2)
	if nil != err {
		fmt.Println(err)
	}
	fmt.Println(jsonutils.JsonEncode(po, true, false, true))

}

