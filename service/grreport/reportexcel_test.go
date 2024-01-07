package grreport

import (
	"testing"
	"fmt"
	"bytes"
	"shservice/common"
	"github.com/xuri/excelize/v2"
)

func TestGetTemplateFileName(t *testing.T){
	data:=map[string]interface{}{
		"semester":map[string]interface{}{
			"value":"01",
			"list":[]interface{}{
				map[string]interface{}{
					"id":"01",
					"name":"第一学期",
				},
			},
		},
		"year":"2018",
		"class":map[string]interface{}{
			"list":[]interface{}{
				map[string]interface{}{
					"enrollment_year":"2017",
				},
			},
		},
	}
	fileName:=getTemplateFileName("grps",data)
	fmt.Println(fileName)
	if fileName!="grps_2grade_01semester.xlsx" {
		t.Error("getTemplateFileName error")
	}
}

func TestCreateExcelReports(t *testing.T){
	data:=map[string]interface{}{
		"semester":map[string]interface{}{
			"value":"01",
			"list":[]interface{}{
				map[string]interface{}{
					"id":"01",
					"name":"第一学期",
				},
			},
		},
		"year":"2018",
		"class":map[string]interface{}{
			"list":[]interface{}{
				map[string]interface{}{
					"enrollment_year":"2017",
				},
			},
		},
	}
	//构造一个可以写入数据的io.Writer对象
	//这里使用bytes.Buffer，也可以使用其他的io.Writer
	bytes:=bytes.Buffer{}
	errorCode:=CreateExcelReports("../templetes","grps",[]interface{}{data},&bytes)
	if errorCode!=common.ResultSuccess {
		t.Error("CreateExcelReports error",errorCode)
		return
	}
	
	f,err:=excelize.OpenReader(&bytes)
	if err!=nil {
		t.Error("CreateExcelReports error",err)
		return
	}
	defer f.Close()

	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err!=nil {
		t.Error("CreateExcelReports error",err)
		return
	}
	//遍历读取每行数据
	for rowIndex, row := range rows {
		for colIndex, colCell := range row {
			fmt.Print(colCell, "\t")
			if rowIndex==0 && colIndex==0 {
				if colCell!="上汇实验学校2018学年第一学期学生成长报告" {
					t.Error("CreateExcelReports error")
					return
				}
			}
		}
		fmt.Println()
	}
}

