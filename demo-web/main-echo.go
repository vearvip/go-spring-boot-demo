/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"

	_ "github.com/go-spring/demo-web/example"
	_ "github.com/go-spring/demo-web/filter"
	_ "github.com/go-spring/go-spring-boot-starter/starter-echo"
	_ "github.com/go-spring/go-spring-boot-starter/starter-web"
	"github.com/go-spring/go-spring-web/spring-echo"
	"github.com/go-spring/go-spring/spring-boot"
	"github.com/go-spring/go-spring/spring-core"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//
// 注册 echo 容器
//
func registerEchoContainer() {
	SpringBoot.RegisterModule(func(ctx SpringCore.SpringContext) {

		e := echo.New()
		e.HideBanner = true
		e.Use(middleware.Recover())

		e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				fmt.Println("use registerEchoContainer()")
				return next(c)
			}
		})

		c := &SpringEcho.Container{
			EchoServer: e,
		}

		ctx.RegisterNameBean("WebContainer", c)
	})
}

func main() {
	if false {
		registerEchoContainer()
	}
	SpringBoot.RunApplication("config/")
}
