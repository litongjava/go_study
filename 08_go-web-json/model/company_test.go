package model

import "testing"

func TestCompany_GetCompnayType(t *testing.T) {
  c := Company{ID: 12456, Name: "ABCD.LTD", Country: "China"}
  companyType := c.GetCompnayType()
  if companyType != "Limited Liability Company1" {
    t.Errorf("Error")
  }
}
