package router

import (
	"hrms/handler"
	accounthandler "hrms/handler/account"
	staffhandler "hrms/handler/staff"

	"github.com/gin-gonic/gin"
)

func Init(server *gin.Engine) {
	// 测试
	server.GET("/ping", accounthandler.Ping)
	// 权限重定向
	server.GET("/authority_render/:modelName", accounthandler.RenderAuthority)
	// 首页重定向
	server.GET("/index", accounthandler.Index)
	// 账户相关
	accountGroup := server.Group("/account")
	accountGroup.POST("/login", accounthandler.Login)
	accountGroup.POST("/quit", accounthandler.Quit)
	// 部门相关
	departGroup := server.Group("/depart")
	departGroup.POST("/create", handler.DepartCreate)
	departGroup.DELETE("/del/:dep_id", handler.DepartDel)
	departGroup.POST("/edit", handler.DepartEdit)
	departGroup.GET("/query/:dep_id", handler.DepartQuery)
	// 职级相关
	rankGroup := server.Group("/rank")
	rankGroup.POST("/create", handler.RankCreate)
	rankGroup.DELETE("/del/:rank_id", handler.RankDel)
	rankGroup.POST("/edit", handler.RankEdit)
	rankGroup.GET("/query/:rank_id", handler.RankQuery)
	// 员工信息相关
	staffGroup := server.Group("/staff")
	staffGroup.POST("/create", staffhandler.StaffCreate)
	staffGroup.DELETE("/del/:staff_id", staffhandler.StaffDel)
	staffGroup.POST("/edit", staffhandler.StaffEdit)
	staffGroup.GET("/query/:staff_id", staffhandler.StaffQuery)
	staffGroup.GET("/query_by_name/:staff_name", staffhandler.StaffQueryByName)
	staffGroup.GET("/query_by_dep/:dep_name", staffhandler.StaffQueryByDep)
	staffGroup.GET("/query_by_staff_id/:staff_id", staffhandler.StaffQueryByStaffId)
	// 密码管理信息相关
	passwordGroup := server.Group("/password")
	passwordGroup.GET("/query/:staff_id", handler.PasswordQuery)
	passwordGroup.POST("/edit", handler.PasswordEdit)
	// 授权信息相关
	authorityGroup := server.Group("/authority")
	authorityGroup.POST("/create", handler.AddAuthorityDetail)
	authorityGroup.POST("/edit", handler.UpdateAuthorityDetailById)
	authorityGroup.GET("/query_by_user_type/:user_type", handler.GetAuthorityDetailListByUserType)
	authorityGroup.POST("/query_by_user_type_and_model", handler.GetAuthorityDetailByUserTypeAndModel)
	authorityGroup.POST("/set_admin/:staff_id", handler.SetAdminByStaffId)
	authorityGroup.POST("/set_normal/:staff_id", handler.SetNormalByStaffId)
	// 通知相关
	notificationGroup := server.Group("/notification")
	notificationGroup.POST("/create", handler.CreateNotification)
	notificationGroup.DELETE("/delete/:notice_id", handler.DeleteNotificationById)
	notificationGroup.POST("/edit", handler.UpdateNotificationById)
	notificationGroup.GET("/query/:notice_title", handler.GetNotificationByTitle)
	// 分公司相关
	companyGroup := server.Group("/company")
	companyGroup.GET("/query", handler.BranchCompanyQuery)
	// 薪资相关
	salaryGroup := server.Group("/salary")
	salaryGroup.POST("/create", handler.CreateSalary)
	salaryGroup.DELETE("/delete/:salary_id", handler.DelSalary)
	salaryGroup.POST("/edit", handler.UpdateSalaryById)
	salaryGroup.GET("/query/:staff_id", handler.GetSalaryByStaffId)
	// 薪资发放相关
	salaryRecordGroup := server.Group("/salary_record")
	//salaryRecordGroup.POST("/create", handler.CreateSalaryRecord)
	//salaryRecordGroup.DELETE("/delete/:salary_record_id", handler.DelSalaryRecord)
	//salaryRecordGroup.POST("/edit", handler.UpdateSalaryRecordById)
	salaryRecordGroup.GET("/query/:staff_id", handler.GetSalaryRecordByStaffId)
	salaryRecordGroup.GET("/get_salary_record_is_pay_by_id/:id", handler.GetSalaryRecordIsPayById)
	salaryRecordGroup.GET("/pay_salary_record_by_id/:id", handler.PaySalaryRecordById)
	salaryRecordGroup.GET("/query_history/:staff_id", handler.GetHadPaySalaryRecordByStaffId)
	// 考勤相关
	attendGroup := server.Group("/attendance_record")
	attendGroup.POST("/create", handler.CreateAttendRecord)
	attendGroup.DELETE("/delete/:attendance_id", handler.DelAttendRecordByAttendId)
	attendGroup.POST("/edit", handler.UpdateAttendRecordById)
	attendGroup.GET("/query/:staff_id", handler.GetAttendRecordByStaffId)
	attendGroup.GET("/query_history/:staff_id", handler.GetAttendRecordHistoryByStaffId)
	attendGroup.GET("/get_attend_record_is_pay/:staff_id/:date", handler.GetAttendRecordIsPayByStaffIdAndDate)
	attendGroup.GET("/approve/query/:leader_staff_id", handler.GetAttendRecordApproveByLeaderStaffId)
	attendGroup.GET("/approve_accept/:attendId", handler.ApproveAccept)
	attendGroup.GET("/approve_reject/:attendId", handler.ApproveReject)
	// 招聘信息相关
	recruitmentGroup := server.Group("/recruitment")
	recruitmentGroup.POST("/create", handler.CreateRecruitment)
	recruitmentGroup.DELETE("/delete/:recruitment_id", handler.DelRecruitmentByRecruitmentId)
	recruitmentGroup.POST("/edit", handler.UpdateRecruitmentById)
	recruitmentGroup.GET("/query/:job_name", handler.GetRecruitmentByJobName)
	// 候选人管理相关
	candidateGroup := server.Group("/candidate")
	candidateGroup.POST("/create", handler.CreateCandidate)
	candidateGroup.DELETE("/delete/:candidate_id", handler.DelCandidateByCandidateId)
	candidateGroup.POST("/edit", handler.UpdateCandidateById)
	candidateGroup.GET("/query_by_name/:name", handler.GetCandidateByName)
	candidateGroup.GET("/query_by_staff_id/:staff_id", handler.GetCandidateByStaffId)
	candidateGroup.GET("/reject/:id", handler.SetCandidateRejectById)
	candidateGroup.GET("/accept/:id", handler.SetCandidateAcceptById)
	// 考试管理相关
	exampleGroup := server.Group("/example")
	exampleGroup.POST("/create", handler.CreateExample)
	exampleGroup.POST("/parse_example_content", handler.ParseExampleContent)
	exampleGroup.DELETE("/delete/:example_id", handler.DelExample)
	exampleGroup.POST("/edit", handler.UpdateExampleById)
	exampleGroup.GET("/query/:name", handler.GetExampleByName)
	exampleGroup.GET("/render_example/:id", handler.RenderExample)
	// 考试成绩相关
	exampleScoreGroup := server.Group("/example_score")
	exampleScoreGroup.POST("/create", handler.CreateExampleScore)
	exampleScoreGroup.GET("/query_by_name/:name", handler.GetExampleHistoryByName)
	exampleScoreGroup.GET("/query_by_staff_id/:staff_id", handler.GetExampleHistoryByStafId)
	// 题目管理相关
	questionGroup := server.Group("/question")
	questionGroup.POST("/create", handler.AddQuestion)
	questionGroup.GET("/search/:id", handler.SearchQuestion)
	questionGroup.POST("/edit", handler.EditQuestion)
	questionGroup.DELETE("/delete/:id", handler.DelQuestion)

	// API相关
	apiGroup := server.Group("/api")
	// apiGroup.GET("wxlogin", handler.WxHandler.Login)
	apiGroup.GET("wxRegister", handler.WxHandler.Register)
}
