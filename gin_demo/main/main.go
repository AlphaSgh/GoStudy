package main

import "gin_demo/gin_middleware"

func main() {
	/*	routers.Include(shop.Routers, blog.Routers)
		r := routers.Init()
		if err := r.Run(":8080"); err != nil {
			fmt.Printf("startup service failed, err:%v\n", err)
		}
	*/
	//gin_date.GinJson()
	//gin_date.GinForm()
	//gin_date.SomeDataStruct()
	//gin_date.AsyncSync()
	//gin_middleware.AllTest()
	//gin_middleware.TestMiddleWare()
	//gin_middleware.CookieUse()
	gin_middleware.CookieTest()
}
