// Copyright 2018 Antoine CHABERT, toHero.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"fmt"
	"github.com/chainHero/resource-manager/app/web/controllers"
	"net/http"
)

// Serve allow to serve a web app on port 3000
func Serve(app *controllers.Controller) error {
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/home", app.HomeHandler())
	http.HandleFunc("/resources", app.ResourcesHandler())
	http.HandleFunc("/resource", app.ResourceHandler())
	http.HandleFunc("/add-resource", app.AddResourceHandler())
	http.HandleFunc("/delete-resource", app.DeleteResourceHandler())
	http.HandleFunc("/acquire-resource", app.AcquireResourceHandler())
	http.HandleFunc("/release-resource", app.ReleaseResourceHandler())
	http.HandleFunc("/logout", app.LogoutHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
	})

	fmt.Println("Listening (http://localhost:3000/) ...")

	return http.ListenAndServe(":3000", nil)
}
