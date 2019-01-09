// Copyright (C) 2018 Sam Olds

package handler

import (
	stdtemplate "html/template"
	"net/http"

	"github.com/samolds/port/template"
)

func Now(w http.ResponseWriter, r *http.Request) error {
	// TODO: load "now text" from a db? then it can be quickly updated in GAE db
	profileImgSrc := "/static/profile.jpg"
	nowText := `
<p>
  I'm currently traveling the world with
  <a href="https://remoteyear.com">Remote Year</a>. I've been to about 20
  countries in the last 6 months and plan to continue for at least another 6
  months.
</p>
<p>
  I've been keeping a
  <a href="https://instagram.com/samgetslost">daily public journal</a> for
  family and friends who are interested in keeping up with my travels.
</p>
<p>
  I recently quit my job to prioritize this travel opportunity; however, I am
  interested in becoming a contributing member of society again. Positions that
  are remote friendly with flexible hours are preferable.
</p>
<p>
  I am passionate about space travel/exploration and am pursuing work there.
</p>
`

	data := make(map[string]interface{})
	data["profileImgSrc"] = profileImgSrc
	data["htmlText"] = stdtemplate.HTML(nowText)
	return template.Now.Render(w, data)
}
