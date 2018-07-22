package controllers

import (
	"net/http"
	"MVC/utils"
	"MVC/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var test = models.Model{
		Entry:"Hello",
		OtherEntry:"World",
	}


	var data = map[string]interface{}{
		"model": test,
	}
	utils.ParseSite(w, "index/index.html", data)
}
