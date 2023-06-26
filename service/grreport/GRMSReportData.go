package grreport

import (
	"shservice/common"
	"shservice/crv"
	"log"
)

var gpmsFields=[]map[string]interface{}{
	{"field":"id"},
  {
    "field":"class",
    "fieldType":"many2one",
    "relatedModelID":"lms_class",
    "fields": []map[string]interface{}{
      {"field": "id"},
      {"field": "current_grade"},
      {"field": "class_number"},
    },
  },
  {"field":"year"},
  {
    "field":"semester",
	  "fieldType":"many2one",
	  "relatedModelID":"lms_semester",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
	  },
  },
  {
    "field":"student",
	  "fieldType":"many2one",
	  "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
        {"field":"id"},
        {"field":"name"},
        {"field":"sn"},
	  },
  },
  {
    "field":"teacher",
	  "fieldType":"many2one",
	  "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
        {"field":"id"},
        {"field":"name"},
	  },
  },
  {"field":"class_comments"},
  {
    "field":"chinese_teacher",
	  "fieldType":"many2one",
	  "relatedModelID":"lms_person",
    "fields":[]map[string]interface{}{
        {"field":"id"},
        {"field":"name"},
	  },
  },
  {"field":"chinese_normal"},
  {"field":"chinese_midterm"},
  {"field":"chinese_final"},
  {"field":"chinese_comprehensive"},
  {"field":"chinese_annual_synthesis"},

  {
    "field":"mathematics_teacher",
    "fieldType":"many2one",
    "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
    },
  },
  {"field":"mathematics_normal"},
  {"field":"mathematics_midterm"},
  {"field":"mathematics_final"},
  {"field":"mathematics_comprehensive"},
  {"field":"mathematics_annual_synthesis"},

  {
    "field":"english_teacher",
	  "fieldType":"many2one",
	  "relatedModelID":"lms_person",
    "fields":[]map[string]interface{}{
        {"field":"id"},
        {"field":"name"},
	  },
  },
  {"field":"english_normal"},
  {"field":"english_midterm"},
  {"field":"english_final"},
  {"field":"english_comprehensive"},
  {"field":"english_annual_synthesis"},

  {
    "field":"ethics_teacher",
    "fieldType":"many2one",
    "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
    },
  },
  {"field":"ethics_normal"},
  {"field":"ethics_final"},
  {"field":"ethics_comprehensive"},  
  {"field":"ethics_annual_synthesis"},

  {
    "field":"geography_teacher",
    "fieldType":"many2one",
    "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
    },
  },
  {"field":"geography_normal"},
  {"field":"geography_final"},
  {"field":"geography_comprehensive"},  
  {"field":"geography_annual_synthesis"},

  {
    "field":"science_teacher",
    "fieldType":"many2one",
    "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
    },
  },
  {"field":"science_normal"},
  {"field":"science_final"},
  {"field":"science_comprehensive"},  
  {"field":"science_annual_synthesis"},

  {
    "field":"tech_teacher",
    "fieldType":"many2one",
    "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
    },
  },
  {"field":"tech_normal"},
  {"field":"tech_final"},
  {"field":"tech_comprehensive"},
  {"field":"tech_annual_synthesis"},

  {
    "field":"work_teacher",
    "fieldType":"many2one",
    "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
    },
  },
  {"field":"work_normal"},
  {"field":"work_final"},
  {"field":"work_comprehensive"},
  {"field":"work_annual_synthesis"},

  {
    "field":"art_teacher",
    "fieldType":"many2one",
    "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
    },
  },
  {"field":"art_normal"},
  {"field":"art_final"},
  {"field":"art_comprehensive"},
  {"field":"art_annual_synthesis"},

  {
    "field":"psychology_teacher",
    "fieldType":"many2one",
    "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
    },
  },
  {"field":"psychology_comprehensive"},
  {"field":"psychology_annual_synthesis"},

  {
    "field":"writing_teacher",
    "fieldType":"many2one",
    "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
    },
  },
  {"field":"writing_comprehensive"},
  {"field":"writing_annual_synthesis"},

  {
    "field":"pe_teacher",
    "fieldType":"many2one",
    "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
    },
  },
  {"field":"pe_engagement"},
  {"field":"pe_regulation"},
  {"field":"pe_team"},
  {"field":"pe_hard"},
  {"field":"pe_comprehensive"},
  {"field":"pe_annual_synthesis"},

  {
    "field":"music_teacher",
    "fieldType":"many2one",
    "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
    },
  },
  {"field":"music_play"},
  {"field":"music_act"},
  {"field":"music_singing"},
  {"field":"music_knowledge"},
  {"field":"music_engagement"},
  {"field":"music_achievement"},
  {"field":"music_comprehensive"},
  {"field":"music_annual_synthesis"},
}

func GetGRMSReportData(ids *[]string,crvClinet *crv.CRVClient,token string)([]interface{}){
	//查询数据
	commonRep:=crv.CommonReq{
		ModelID:"view_gr_ms",
		Filter:&map[string]interface{}{
			"id":map[string]interface{}{
				"Op.in":ids,
			},
		},
		Fields:&gpmsFields,
	}

	req,commonErr:=crvClinet.Query(&commonRep,token)
	if commonErr!=common.ResultSuccess {
		return nil
	}

	if req.Error == true {
		log.Println("GetGRMSReportData error:",req.ErrorCode,req.Message)
		return nil
	}

	return req.Result["list"].([]interface{})
}