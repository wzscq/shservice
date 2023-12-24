package grreport

import (
	"shservice/common"
	"shservice/crv"
	"log"
)

var gprsFields=[]map[string]interface{}{
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
      "field":"class_teacher",
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
    {"field":"chinese_basics"},
    {"field":"chinese_expression"},
    {"field":"chinese_reading"},
    {"field":"chinese_synthesis"},
    {"field":"chinese_comprehensive"},
    {"field":"chinese_comments"},
    {"field":"chinese_final_basics"},
    {"field":"chinese_final_reading"},
    {"field":"chinese_final_expression"},
    {"field":"chinese_final_writing"},

    {
      "field":"math_teacher",
	  "fieldType":"many2one",
	  "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
	  },
    },
    {"field":"math_algebra"},
    {"field":"math_geometry"},
    {"field":"math_statistics"},
    {"field":"math_synthesis"},
    {"field":"math_comprehensive"},
    {"field":"math_comments"},
    {"field":"math_final_calc"},
    {"field":"math_final_conception"},
    {"field":"math_final_app"},

    {
      "field":"english_teacher",
	  "fieldType":"many2one",
	  "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
	  },
    },
    {"field":"english_listen"},
    {"field":"english_expression"},
    {"field":"english_reading"},
    {"field":"english_writing"},
    {"field":"english_comprehensive"},
    {"field":"english_comments"},
    {"field":"english_final_voice"},
    {"field":"english_final_vocabulary"},
    {"field":"english_final_discourse"},

    {
      "field":"ethics_teacher",
	  "fieldType":"many2one",
	  "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
	  },
    },
    {"field":"ethics_engagement"},
    {"field":"ethics_listen"},
    {"field":"ethics_collaboration"},
    {"field":"ethics_activity"},
    {"field":"ethics_knowledge"},
    {"field":"ethics_homework"},
    {"field":"ethics_behavior"},
    {"field":"ethics_comprehensive"},
    {"field":"ethics_comments"},

    {
      "field":"nature_teacher",
	  "fieldType":"many2one",
	  "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
	  },
    },
    {"field":"nature_observe"},
    {"field":"nature_discuss"},
    {"field":"nature_operate"},
    {"field":"nature_knowledge"},
    {"field":"nature_make"},
    {"field":"nature_design"},
    {"field":"nature_comprehensive"},
    {"field":"nature_comments"},

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
    {"field":"pe_comments"},

    {
      "field":"music_teacher",
	  "fieldType":"many2one",
	  "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
	  },
    },
    {"field":"music_listen"},
    {"field":"music_play"},
    {"field":"music_act"},
    {"field":"music_singing"},
    {"field":"music_knowledge"},
    {"field":"music_engagement"},
    {"field":"music_achievement"},
    {"field":"music_comprehensive"},
    {"field":"music_comments"},  

    {
      "field":"art_teacher",
	  "fieldType":"many2one",
	  "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
	  },
    },
    {"field":"art_homework1"},
    {"field":"art_homework2"},
    {"field":"art_homework3"},
    {"field":"art_homework4"},
    {"field":"art_styling"},
    {"field":"art_design"},
    {"field":"art_appreciation"},
    {"field":"art_comprehensive"},
    {"field":"art_comments"}, 

    {
      "field":"explore_teacher",
	  "fieldType":"many2one",
	  "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
	  },
    },
    {"field":"explore_comments"},  

    {
      "field":"ep_teacher",
      "fieldType":"many2one",
	    "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
      },
    },
    {"field":"ep_engagement"},
    {"field":"ep_listen"},
    {"field":"ep_homework"},
    {"field":"ep_comprehensive"},
    {"field":"ep_score"},
    {"field":"ep_comments"},

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
    {"field":"psychology_comments"},
    {"field":"psychology_score"},

    {
      "field":"it_teacher",
      "fieldType":"many2one",
	    "relatedModelID":"lms_person",
      "fields":[]map[string]interface{}{
          {"field":"id"},
          {"field":"name"},
      },
    },
    {"field":"it_comprehensive"},
    {"field":"it_comments"},
    {"field":"it_score"},
}

func GetGRPSReportData(ids *[]string,crvClinet *crv.CRVClient,token string)([]interface{}){
	//查询数据
	commonRep:=crv.CommonReq{
		ModelID:"view_gr_ps",
		Filter:&map[string]interface{}{
			"id":map[string]interface{}{
				"Op.in":ids,
			},
		},
		Fields:&gprsFields,
	}

	req,commonErr:=crvClinet.Query(&commonRep,token)
	if commonErr!=common.ResultSuccess {
		return nil
	}

	if req.Error == true {
		log.Println("GetGRPSReportData error:",req.ErrorCode,req.Message)
		return nil
	}

	return req.Result["list"].([]interface{})
}