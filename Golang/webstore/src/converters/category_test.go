package converters

import (
	"testing"
	"models"
)


func Test_ConvertsCategoryToViewModel(t *testing.T) {
	c := models.Category{}
	c.SetImageUrl("the image URL")
	c.SetTitle("the title")
	c.SetDescription("the description")
	c.SetIsOrientRight(true)
	c.SetId(42)

	r := ConvertCategoryToViewModel(c)

	if r.ImageUrl != c.ImageUrl() {
		t.Fail()
	}
	if r.Title != c.Title() {
		t.Fail()
	}
	if r.IsOrientRight != c.IsOrientRight() {
		t.Fail()
	}
	if r.Description != c.Description() {
		t.Fail()
	}

	if r.Id != c.Id() {
		t.Fail()
	}

}
