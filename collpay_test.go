package collpay_go_sdk

import "testing"

func panicTest(t *testing.T){
	if recover() != nil {
		t.Log("Failed here")
		t.Fail()
	}
}

func Test_If_Nil_Sent_To_ConfigureEnv_Function_Will_Return_Error (t *testing.T) {
	defer panicTest(t)
	err :=ConfigureEnv(nil)
	if err == nil {
		t.Log("Failed here")
        t.Fail()
	}
}

func Test_If_Empty_Config_Or_Empty_Public_Key_Sent_To_ConfigureEnv_Function_Will_Return_Error (t *testing.T) {
	defer panicTest(t)
	err :=ConfigureEnv(&Config{})
	if err == nil {
		t.Log("Failed here")
        t.Fail()
	}
	err =ConfigureEnv(&Config{Env: ENV_SANDBOX, Version: V1, PublicKey: ""})
	if err == nil {
		t.Log("Failed here")
        t.Fail()
	}
}

func Test_If_Empty_Config_Fields_Except_Public_Key_Sent_To_ConfigureEnv_Function_Env_Will_Be_Production_And_Version_Will_Be_V1 (t *testing.T) {
	defer panicTest(t)
	err :=ConfigureEnv(&Config{PublicKey: "xxx"})
	if err != nil {
		t.Log("Failed here")
        t.Fail()
	}
	if configData.baseUrl != productionBaseUrl+configData.Version {
		t.Log("Failed here")
        t.Fail()
	}
}

func Test_ExchangeRate_Basic_Error (t *testing.T) {
	defer panicTest(t)
	_ = ConfigureEnv(&Config{Env: ENV_SANDBOX, Version: V1, PublicKey: "xxxx"})
	exch, err  := GetExchangeRate("USD", "BTC")
	if err == nil && exch.Success {
		t.Log("Failed here")
		t.Fail()
	}
}

func Test_CreatTransaction_Basic_Error (t *testing.T) {
	defer panicTest(t)
	tr, err  := CreateTransaction(nil)
	if err == nil && tr.Success {
		t.Log("Failed here")
		t.Fail()
	}
}

func Test_GetTransaction_Basic_Error (t *testing.T) {
	defer panicTest(t)
	tr, err  := GetTransaction("")
	if err == nil && tr.Success {
		t.Log("Failed here")
		t.Fail()
	}

	tr, err  = GetTransaction("xxxxx")
	if err == nil && tr.Success {
		t.Log("Failed here")
		t.Fail()
	}
}

func Test_doRequestAndGetResponse_Basic_Error(t *testing.T) {
	defer panicTest(t)
	_, err  := doRequestAndGetResponse(nil)
	if err == nil {
		t.Log("Failed here")
		t.Fail()
	}
}