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
	"os"
	"time"

	_ "github.com/go-spring/demo-config/example"
	"github.com/go-spring/go-spring/boot-starter"
	"github.com/go-spring/go-spring/spring-boot"
)

func init() {
	os.Setenv("spring.profile", "test")
}

func main() {
	go func() {
		time.Sleep(time.Millisecond * 50)
		BootStarter.Exit()
	}()
	//SpringBoot.RunApplication("./")
	SpringBoot.RunApplication("config/")
}
