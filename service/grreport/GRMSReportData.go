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
      {"field": "enrollment_year"},
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
  {"field":"chinese_normal_score"},
  {"field":"chinese_midterm_score"},
  {"field":"chinese_final_score"},
  {"field":"chinese_comprehensive_score"},
  {"field":"chinese_score"},

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
  {"field":"mathematics_normal_score"},
  {"field":"mathematics_midterm_score"},
  {"field":"mathematics_final_score"},
  {"field":"mathematics_comprehensive_score"},
  {"field":"mathematics_score"},
  
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
  {"field":"english_normal_score"},
  {"field":"english_midterm_score"},
  {"field":"english_final_score"},
  {"field":"english_comprehensive_score"},
  {"field":"english_score"},

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
  {"field":"ethics_normal_score"},
  {"field":"ethics_final_score"},
  {"field":"ethics_comprehensive_score"},
  {"field":"ethics_score"},  
  
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
  {"field":"geography_normal_score"},
  {"field":"geography_final_score"},
  {"field":"geography_comprehensive_score"},
  {"field":"geography_score"},

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
  {"field":"science_normal_score"},
  {"field":"science_final_score"},
  {"field":"science_comprehensive_score"},
  {"field":"science_score"},
  
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
  {"field":"tech_normal_score"},
  {"field":"tech_final_score"},
  {"field":"tech_comprehensive_score"},
  {"field":"tech_score"},

  {
    "field":"physic_teacher",
    "fieldType":"many2one",
    "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
    },
  },
  {"field":"physic_normal"},
  {"field":"physic_final"},
  {"field":"physic_comprehensive"},
  {"field":"physic_annual_synthesis"},
  {"field":"physic_normal_score"},
  {"field":"physic_final_score"},
  {"field":"physic_comprehensive_score"},
  {"field":"physic_score"},

  {
    "field":"explore_teacher",
    "fieldType":"many2one",
    "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
    },
  },
  {"field":"explore_normal"},
  {"field":"explore_final"},
  {"field":"explore_comprehensive"},
  {"field":"explore_annual_synthesis"},
  {"field":"explore_normal_score"},
  {"field":"explore_final_score"},
  {"field":"explore_comprehensive_score"},
  {"field":"explore_score"},

  {
    "field":"history_teacher",
    "fieldType":"many2one",
    "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
    },
  },
  {"field":"history_normal"},
  {"field":"history_final"},
  {"field":"history_comprehensive"},
  {"field":"history_annual_synthesis"},
  {"field":"history_normal_score"},
  {"field":"history_final_score"},
  {"field":"history_comprehensive_score"},
  {"field":"history_score"},

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
  {"field":"work_normal_score"},
  {"field":"work_final_score"},
  {"field":"work_comprehensive_score"},
  {"field":"work_score"},

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
  {"field":"psychology_comprehensive_score"},
  {"field":"psychology_score"},

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