// Copyright 2013 bee authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.
package main

import (
	"flag"
	"os"

	"github.com/zssky/codgen/generate"
	"github.com/zssky/log"
)

func init() {
	flag.Var(&generate.Tables, "Tables", "List of table names separated by a comma.")
	flag.Var(&generate.SQLDriver, "SQLDriver", "Database SQLDriver. Either mysql, postgres or sqlite.")
	flag.Var(&generate.SQLConn, "SQLConn", "Connection string used by the SQLDriver to connect to a database instance.")
	flag.Var(&generate.Level, "Level", "Either 1, 2 or 3. i.e. 1=models; 2=models and controllers; 3=models, controllers and routers.")
	flag.Var(&generate.Fields, "Fields", "List of table Fields.")
}

func main() {

	flag.Parse()

	currpath, _ := os.Getwd()

	appCode(currpath)
}

func appCode(currpath string) {
	if generate.SQLDriver == "" {
		log.Errorf("SQLDriver can not be null")
		return
	}

	if generate.SQLConn == "" {
		log.Errorf("SQLConn can not be null")
		return
	}

	if generate.Level == "" {
		generate.Level = "3"
	}

	log.Info("Using '%s' as 'SQLDriver'", generate.SQLDriver)
	log.Info("Using '%s' as 'SQLConn'", generate.SQLConn)
	log.Info("Using '%s' as 'Tables'", generate.Tables)
	log.Info("Using '%s' as 'Level'", generate.Level)

	generate.GenerateAppcode(generate.SQLDriver.String(), generate.SQLConn.String(), generate.Level.String(), generate.Tables.String(), currpath)
}
