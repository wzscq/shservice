docker run --name shservice -p9800:80 -v /root/shservice/font:/services/shservice/font -v /root/shservice/templetes:/services/shservice/templetes -v /root/shservice/conf:/services/shservice/conf  wangzhsh/shservice:0.1.1


2023-06-22 第一次创建

2023-10-22 更新  未更新
1、在班级人员查询中，对于学生的查询补充学号字段，并可以按学号字段查询和排序。
   1）、创建数据库视图，view_class_participant,关联学生的学号字段。
   2）、补充view_class_participant相关的模型配置信息。
   3）、班级人员菜单对应打开的模型改为view_class_participant。
   4）、原来lms_class_participant模型中的save、import操作的刷新页面改为view_class_participant。
2、增加教师课程表查询功能。
   1）、创建数据库视图，view_teacher_curriculum，关联班级教师和班级课程表。
   2）、补充view_teacher_curriculum相关模型配置。
3、增加升年级功能
   1）、增加流程配置逻辑，upgrade_class_grade
   2）、lms_class模型中增加升年级按钮配置
4、增加升学期功能
   1）、 增加流程配置逻辑，upgrade_class_semester
   2）、lms_class模型中增加升学期按钮配置

2023-11-04
1、学生信息补充以下字段：户口省份、民族、联系电话、联系人1、联系人1单位、联系人1关系、联系人1电话、联系人2、联系人2单位、联系人2关系、联系人2电话
   1）、数据库表lms_person增加以下字段
       alter table lms_person
         add column census_register varchar(25) NULL ,
         add column ethnic varchar(25) NULL ,
         add column phone varchar(25) NULL ,
         add column contact1 varchar(25) NULL ,
         add column contact1_workplace varchar(255) NULL ,
         add column contact1_relationship varchar(25) NULL ,
         add column contact1_phone varchar(25) NULL ,
         add column contact2 varchar(25) NULL ,
         add column contact2_workplace varchar(255) NULL ,
         add column contact2_relationship varchar(25) NULL ,
         add column contact2_phone varchar(25) NULL 
   2）、修改lms_person模型配置
   3）、修改学生信息导入模板文件 
   4）、修改导入流程配置 import_student
   5）、这里对dataflow中的部分逻辑做了优化，需要同步更新
   6）、修改学生信息导出流程配置 export_student

2、成长报告中增加分项分数和总分   总分的处理需要再确认一下？
   1）、数据库相关表格增加分数字段
         alter table lms_gr_ps_chinese add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ps_mathematics add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ps_english add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ps_ethics add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ps_nature add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ps_pe add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ps_music add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ps_art add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ps_explore add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ms_chinese add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ms_mathematics add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ms_english add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ms_ethics add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ms_geography add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ms_science add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ms_tech add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ms_pe add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ms_work add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ms_art add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ms_music add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ms_psychology add COLUMN score decimal(18,2) NULL;
         alter table lms_gr_ms_writing add COLUMN score decimal(18,2) NULL;
   2）、数据库视图中增加分数字段和总分字段配置
         view_gr_ms
         view_gr_ps
   3）、修改相关数据模型配置，增加对应字段信息
   4）、各科导入模板中增加成绩这一列
   5）、修改各科导入流程配置，增加对成绩导入的处理

3、导出各学科成长报告详表
   1）、配置导出处理流程
   2）、各科页面配置导出按钮

2023-12-23 成长报告更新
1、小学语文成长报告增加期末考查项目
   a、修改导入模板，增加对应成绩的导入列
   b、修改数据库表增加对应数据列
      alter table lms_gr_ps_chinese
         add column chinese_final_basics varchar(2) NULL ,
         add column chinese_final_reading varchar(2) NULL ,
         add column chinese_final_expression varchar(2) NULL ,
         add column chinese_final_writing varchar(2) NULL;
   c、模型配置修改lms_gr_ps_chinese
   d、导入导出流程修改，增加对应字段处理
      export_gr_ps_chinese.json
      import_gr_ps_chinese.json
2、小学数学成长报告增加期末考查项目
   a、修改导入模板，增加对应期末成绩导入列
   b、修改数据库表增加对应数据列
      alter table lms_gr_ps_mathematics
         add column math_final_calc varchar(2) NULL ,
         add column math_final_conception varchar(2) NULL ,
         add column math_final_app varchar(2) NULL;
   c、模型配置修改lms_gr_ps_mathematics
   d、导入导出流程修改，增加对应字段处理
      import_gr_ps_mathematics.json
      export_gr_ps_mathematics.json
3、小学英语成长报告增加期末考查项目
   a、修改导入模板，增加对应期末成绩导入列
   b、修改数据库表增加对应数据列
      alter table lms_gr_ps_english
         add column english_final_voice varchar(2) NULL ,
         add column english_final_vocabulary varchar(2) NULL ,
         add column english_final_discourse varchar(2) NULL;
   c、模型配置修改lms_gr_ps_english
   d、导入导出流程修改，增加对应字段处理
      import_gr_ps_english.json
      export_gr_ps_english.json
4、小学成长报告增加口才与展示科目
   a、创建对应数据表数据库表
   create table lms_gr_ps_ep (
      id varchar(20) NOT NULL ,
      class varchar(6) NOT NULL ,
      year varchar(4) NOT NULL ,
      semester varchar(2) NOT NULL ,
      student varchar(20) NOT NULL ,
      ep_teacher varchar(20) NULL ,
      ep_engagement varchar(2) NULL ,
      ep_listen varchar(2) NULL ,
      ep_homework varchar(2) NULL ,
      ep_comprehensive varchar(2) NULL ,
      score decimal(18,2) NULL ,
      ep_comments varchar(500) NULL ,
      sheet_name varchar(255) NOT NULL ,
      import_file_name varchar(260) NULL ,
      import_batch_number varchar(25) NULL ,
      create_time datetime(0) NOT NULL ,
      create_user varchar(25) NOT NULL ,
      update_time datetime(0) NOT NULL ,
      update_user varchar(25) NOT NULL ,
      version int(11) NOT NULL DEFAULT 0,
      PRIMARY KEY (`id`) USING BTREE)

   b、配置对应模型 lms_gr_ps_ep
   c、增加对应菜单配置
   d、添加角色
      ps_ep
   e、添加导入模板
   f、需要上传导入模板
   g、增加相应的导入、导出处理流程
      export_gr_ps_ep.json
      import_gr_ps_ep.json
5、小学成长报告增加心里科目
   a、创建对应数据表数据库表
      create table lms_gr_ps_psychology (
         id varchar(20) NOT NULL ,
         class varchar(6) NOT NULL ,
         year varchar(4) NOT NULL ,
         semester varchar(2) NOT NULL ,
         student varchar(20) NOT NULL ,
         psychology_teacher varchar(20) NULL ,
         psychology_comprehensive varchar(2) NULL ,
         psychology_comments varchar(500) NULL ,
         score decimal(18,2) NULL ,
         sheet_name varchar(255) NOT NULL ,
         import_file_name varchar(260) NULL ,
         import_batch_number varchar(25) NULL ,
         create_time datetime(0) NOT NULL ,
         create_user varchar(25) NOT NULL ,
         update_time datetime(0) NOT NULL ,
         update_user varchar(25) NOT NULL ,
         version int(11) NOT NULL DEFAULT 0,
         PRIMARY KEY (`id`) USING BTREE)
   b、配置对应模型 lms_gr_ps_psychology
   c、增加对应菜单配置
   d、添加角色
      ps_psychology
   e、添加导入模板
   f、需要上传导入模板
   g、增加相应的导入、导出处理流程
      export_gr_ps_psychology.json
      import_gr_ps_psychology.json
5、小学成长报告增加信息科目
   a、创建对应数据表数据库表
      create table lms_gr_ps_it (
         id varchar(20) NOT NULL ,
         class varchar(6) NOT NULL ,
         year varchar(4) NOT NULL ,
         semester varchar(2) NOT NULL ,
         student varchar(20) NOT NULL ,
         it_teacher varchar(20) NULL ,
         it_comprehensive varchar(2) NULL ,
         it_comments varchar(500) NULL ,
         score decimal(18,2) NULL ,
         sheet_name varchar(255) NOT NULL ,
         import_file_name varchar(260) NULL ,
         import_batch_number varchar(25) NULL ,
         create_time datetime(0) NOT NULL ,
         create_user varchar(25) NOT NULL ,
         update_time datetime(0) NOT NULL ,
         update_user varchar(25) NOT NULL ,
         version int(11) NOT NULL DEFAULT 0,
         PRIMARY KEY (`id`) USING BTREE)
   b、配置对应模型 lms_gr_ps_it
   c、增加对应菜单配置
   d、添加角色
      ps_it
   e、添加导入模板
   f、需要上传导入模板
   g、增加相应的导入、导出处理流程
      export_gr_ps_it.json
      import_gr_ps_it.json
6、初中成长报告增加物理科目
   a、创建对应数据表数据库表
   create table lms_gr_ms_physic (
      id varchar(20) NOT NULL ,
      class varchar(6) NOT NULL ,
      year varchar(4) NOT NULL ,
      semester varchar(2) NOT NULL ,
      student varchar(20) NOT NULL ,
      teacher varchar(20) NULL ,
      normal varchar(2) NULL ,
      final varchar(2) NULL ,
      comprehensive varchar(2) NULL ,
      annual_synthesis varchar(2) NULL ,
      score decimal(18,2) NULL ,
      sheet_name varchar(255) NOT NULL ,
      import_file_name varchar(260) NULL ,
      import_batch_number varchar(25) NULL ,
      create_time datetime(0) NOT NULL ,
      create_user varchar(25) NOT NULL ,
      update_time datetime(0) NOT NULL ,
      update_user varchar(25) NOT NULL ,
      version int(11) NOT NULL DEFAULT 0,
      PRIMARY KEY (`id`) USING BTREE)
   b、配置对应模型 lms_gr_ms_physic
   c、增加对应菜单配置
   d、添加角色
      ms_physic
   e、添加导入模板
   f、需要上传导入模板
   g、增加相应的导入、导出处理流程
      export_gr_ps_physic.json
      import_gr_ps_physic.json
6、初中成长报告增加探究科目
   a、创建对应数据表数据库表 
      create table lms_gr_ms_explore (
         id varchar(20) NOT NULL ,
         class varchar(6) NOT NULL ,
         year varchar(4) NOT NULL ,
         semester varchar(2) NOT NULL ,
         student varchar(20) NOT NULL ,
         teacher varchar(20) NULL ,
         normal varchar(2) NULL ,
         final varchar(2) NULL ,
         comprehensive varchar(2) NULL ,
         annual_synthesis varchar(2) NULL ,
         score decimal(18,2) NULL ,
         sheet_name varchar(255) NOT NULL ,
         import_file_name varchar(260) NULL ,
         import_batch_number varchar(25) NULL ,
         create_time datetime(0) NOT NULL ,
         create_user varchar(25) NOT NULL ,
         update_time datetime(0) NOT NULL ,
         update_user varchar(25) NOT NULL ,
         version int(11) NOT NULL DEFAULT 0,
         PRIMARY KEY (`id`) USING BTREE)
   b、配置对应模型 lms_gr_ms_explore
   c、增加对应菜单配置
   d、添加角色
      ms_explore
   e、添加导入模板
   f、需要上传导入模板
   g、增加相应的导入、导出处理流程
      export_gr_ps_explore.json
      import_gr_ps_explore.json
7、初中成长报告增加历史科目   
   a、创建对应数据库表 
   create table lms_gr_ms_history (
      id varchar(20) NOT NULL ,
      class varchar(6) NOT NULL ,
      year varchar(4) NOT NULL ,
      semester varchar(2) NOT NULL ,
      student varchar(20) NOT NULL ,
      teacher varchar(20) NULL ,
      normal varchar(2) NULL ,
      final varchar(2) NULL ,
      comprehensive varchar(2) NULL ,
      annual_synthesis varchar(2) NULL ,
      score decimal(18,2) NULL ,
      sheet_name varchar(255) NOT NULL ,
      import_file_name varchar(260) NULL ,
      import_batch_number varchar(25) NULL ,
      create_time datetime(0) NOT NULL ,
      create_user varchar(25) NOT NULL ,
      update_time datetime(0) NOT NULL ,
      update_user varchar(25) NOT NULL ,
      version int(11) NOT NULL DEFAULT 0,
      PRIMARY KEY (`id`) USING BTREE)
   b、配置对应模型 lms_gr_ms_history
   c、增加对应菜单配置
   d、添加角色
      ms_history
   e、添加导入模板
   f、需要上传导入模板
   g、增加相应的导入、导出处理流程
      export_gr_ps_history.json
      import_gr_ps_history.json
8、初中成长报告导出格式修改
   a、修改导出数据视图 view_gr_ms 添加物理、探究、历史相关字段
   b、修改模型 view_gr_ms 配置，增加相应科目字段 
   c、修改报告导出模板配置
   d、修改导出逻辑，增加对应字段的处理
9、小学成长报告导出格式修改
   a、修改导出数据视图 view_gr_ps 添加相关字段
   b、吸怪模型 view_gr_ps 配置，增加相应科目字段
   c、修改导出逻辑，增加对应字段的处理
   

