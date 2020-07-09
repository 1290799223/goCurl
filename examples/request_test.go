package goCurl

import (
	"fmt"
	"github.com/qifengzhang007/goCurl"
	"io/ioutil"
	"log"
	"net/http"
)

func ExampleRequest_GetCookies() {
	cli := goCurl.NewClient()
	resp, err := cli.Get("http://www.iwencai.com/diag/block-detail?pid=10751&codes=600422&codeType=stock&info={\"view\":{\"nolazy\":1}}")

	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Printf("%#+v\n", resp.GetCookies())
	fmt.Printf("%T", resp.GetCookie("vvvv"))
	// Output: *http.Cookie

}

func ExampleRequest_Get() {
	cli := goCurl.NewClient()

	//resp, err := cli.Get("http://127.0.0.1:8091/get")
	resp, err := cli.Get("https://finance.sina.com.cn/realstock/company/sz002614/nc.shtml")
	//resp, err := cli.Get("http://www.10jqka.com.cn/")
	//resp, err := cli.Get("http://www.zhenai.com/")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%T", resp)
	// Output: *goCurl.Response

}

func ExampleRequest_Down() {
	cli := goCurl.NewClient()

	res := cli.Down("http://139.196.101.31:2080/GinSkeleton.jpg", "F:/2020_project/go/goz/examples/", "", goCurl.Options{
		Timeout: 5.0,
	})
	fmt.Printf("%t", res)
	// Output: true
}

func ExampleRequest_Get_withQuery_arr() {
	cli := goCurl.NewClient()

	resp, err := cli.Get("http://127.0.0.1:8091/get-with-query", goCurl.Options{
		Query: map[string]interface{}{
			"key1": 123,
			"key2": []string{"value21", "value22"},
			"key3": "abc456",
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s", resp.GetRequest().URL.RawQuery)
	// Output: key1=123&key2=value21&key2=value22&key3=abc456
}

func ExampleRequest_Get_withQuery_str() {
	cli := goCurl.NewClient()

	resp, err := cli.Get("http://127.0.0.1:8091/get-with-query?key0=value0", goCurl.Options{
		Query: "key1=value1&key2=value21&key2=value22&key3=333",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s", resp.GetRequest().URL.RawQuery)
	// Output: key1=value1&key2=value21&key2=value22&key3=333
}

func ExampleRequest_Get_withProxy() {
	cli := goCurl.NewClient()

	resp, err := cli.Get("https://www.fbisb.com/ip.php", goCurl.Options{
		Timeout: 5.0,
		Proxy:   "http://127.0.0.1:1087",
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.GetStatusCode())
	// Output: 200
	fmt.Println(resp.GetContents())
	// Output: 116.153.43.128
}

func ExampleRequest_Post() {
	cli := goCurl.NewClient()

	resp, err := cli.Post("http://127.0.0.1:8091/post")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T", resp)
	// Output: *goCurl.Response
}

func ExampleRequest_Post_withHeaders() {
	cli := goCurl.NewClient()

	resp, err := cli.Post("http://127.0.0.1:8091/post-with-headers", goCurl.Options{
		Headers: map[string]interface{}{
			"User-Agent": "testing/1.0",
			"Accept":     "application/json",
			"X-Foo":      []string{"Bar", "Baz"},
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	headers := resp.GetRequest().Header["X-Foo"]
	fmt.Println(headers)
	// Output: [Bar Baz]
}

func ExampleRequest_Post_withCookies_str() {
	cli := goCurl.NewClient()

	resp, err := cli.Post("http://127.0.0.1:8091/post-with-cookies", goCurl.Options{
		Cookies: "cookie1=value1;cookie2=value2",
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%d", resp.GetContentLength())
	//Output: 385
}

func ExampleRequest_Post_withCookies_map() {
	cli := goCurl.NewClient()

	//resp, err := cli.Post("http://127.0.0.1:8091/post-with-cookies", goCurl.Options{
	resp, err := cli.Post("http://101.132.69.236/api/v2/test_network", goCurl.Options{
		Cookies: map[string]string{
			"cookie1": "value1",
			"cookie2": "value2",
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	body := resp.GetBody()
	defer body.Close()
	bytes, _ := ioutil.ReadAll(body)
	fmt.Printf("%s", bytes)
	// Output: {"code":200,"msg":"OK","data":""}
}

func ExampleRequest_Post_withCookies_obj() {
	cli := goCurl.NewClient()

	cookies := make([]*http.Cookie, 0, 2)
	cookies = append(cookies, &http.Cookie{
		Name:     "cookie133",
		Value:    "value1",
		Domain:   "httpbin.org",
		Path:     "/cookies",
		HttpOnly: true,
	})
	cookies = append(cookies, &http.Cookie{
		Name:   "cookie2",
		Value:  "value2",
		Domain: "httpbin.org",
		Path:   "/cookies",
	})

	resp, err := cli.Post("http://127.0.0.1:8091/post-with-cookies", goCurl.Options{
		Cookies: cookies,
	})
	if err != nil {
		log.Fatalln(err)
	}

	body := resp.GetBody()
	fmt.Printf("%T", body)
	//Output: *http.cancelTimerBody
}
func ExampleRequest_SimplePost() {
	cli := goCurl.NewClient()

	resp, err := cli.Post("http://101.132.69.236/api/v2/test_network", goCurl.Options{
		Headers: map[string]interface{}{
			"Content-Type": "application/x-www-form-urlencoded",
		},
		FormParams: map[string]interface{}{
			"key1": "value1",
			"key2": []string{"value21", "value22"},
			"key3": "333",
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	contents, _ := resp.GetContents()
	fmt.Printf("%s", contents)
	// Output:  {"code":200,"msg":"OK","data":""}
}

func ExampleRequest_Post_withFormParams() {
	cli := goCurl.NewClient()

	resp, err := cli.Post("http://127.0.0.1:8091/post-with-form-params", goCurl.Options{
		Headers: map[string]interface{}{
			"Content-Type": "application/x-www-form-urlencoded",
		},
		FormParams: map[string]interface{}{
			"key1": 2020,
			"key2": []string{"value21", "value22"},
			"key3": "abcd张",
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	body, err := resp.GetContents()

	fmt.Printf("%v", body)
	// Output:  form params:{"key1":["2020"],"key2":["value21","value22"],"key3":["abcd张"]}
}

func ExampleRequest_Post_withJSON() {
	cli := goCurl.NewClient()

	resp, err := cli.Post("http://127.0.0.1:8091/post-with-json", goCurl.Options{
		Headers: map[string]interface{}{
			"Content-Type": "application/json",
		},
		JSON: struct {
			Key1 string   `json:"key1"`
			Key2 []string `json:"key2"`
			Key3 int      `json:"key3"`
		}{"value1", []string{"value21", "value22"}, 333},
	})
	if err != nil {
		log.Fatalln(err)
	}

	body := resp.GetBody()
	defer body.Close()
	fmt.Printf("%T", body)
	// Output:  *http.cancelTimerBody
}

func ExampleRequest_Put() {
	cli := goCurl.NewClient()

	resp, err := cli.Put("http://127.0.0.1:8091/put")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T", resp)
	// Output: *goCurl.Response
}

func ExampleRequest_Patch() {
	cli := goCurl.NewClient()

	resp, err := cli.Patch("http://127.0.0.1:8091/patch")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T", resp)
	// Output: *goCurl.Response
}

func ExampleRequest_Delete() {
	cli := goCurl.NewClient()

	resp, err := cli.Delete("http://127.0.0.1:8091/delete")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T", resp)
	// Output: *goCurl.Response
}

func ExampleRequest_Options() {
	cli := goCurl.NewClient()

	resp, err := cli.Options("http://127.0.0.1:8091/options")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T", resp)
	// Output: *goCurl.Response
}
