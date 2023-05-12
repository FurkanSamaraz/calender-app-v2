package retail

import (
	controllerCalendarBasics "main/controller/calendar_basics"
	servicesCalendarBasics "main/services/calendar_basics"

	controllerEmployeePermissions "main/controller/employee_permissions"
	servicesEmployeePermissions "main/services/employee_permissions"

	controllerEmployeeRequests "main/controller/employee_requests"
	servicesEmployeeRequests "main/services/employee_requests"

	controllerExpenseTracking "main/controller/expense_tracking"
	servicesExpenseTracking "main/services/expense_tracking"

	controllerIntegration "main/controller/integration"
	servicesIntegration "main/services/integration"

	controllerNotifications "main/controller/notifications"
	servicesNotifications "main/services/notifications"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Run(app fiber.Router, db *gorm.DB) {
	{
		{
			{
				EventInstanceGroup := app.Group("/EventInstance")
				EventInstanceController := controllerCalendarBasics.EventInstanceController{Svc: servicesCalendarBasics.EventInstanceService{DB: db}}

				EventInstanceGroup.Get("/search", EventInstanceController.GetEventInstance).Name("GetEventInstance")
				EventInstanceGroup.Post("/single", EventInstanceController.CreateEventInstance).Name("CreateEventInstance")

				subEventInstanceGroup := EventInstanceGroup.Group("/with-id/:id")
				{

					subEventInstanceGroup.Put("/", EventInstanceController.UpdateEventInstance).Name("UpdateEventInstanceWithId")
					subEventInstanceGroup.Delete("/", EventInstanceController.DeleteEventInstance).Name("DeleteEventInstanceWithId")
				}

			}
			{
				EventGroup := app.Group("/Event")
				EventController := controllerCalendarBasics.EventController{Svc: servicesCalendarBasics.EventService{DB: db}}

				EventGroup.Get("/search", EventController.GetEvent).Name("GetEvent")
				EventGroup.Post("/single", EventController.CreateEvent).Name("CreateEvent")

				subEventGroup := EventGroup.Group("/with-id/:id")
				{

					subEventGroup.Put("/", EventController.UpdateEvent).Name("UpdateEventWithId")
					subEventGroup.Delete("/", EventController.DeleteEvent).Name("DeleteEventWithId")
				}

			}
			{
				ImportantDayGroup := app.Group("/ImportantDay")
				ImportantDayController := controllerCalendarBasics.ImportantDayController{Svc: servicesCalendarBasics.ImportantDayService{DB: db}}

				ImportantDayGroup.Get("/search", ImportantDayController.GetImportantDay).Name("GetImportantDay")
				ImportantDayGroup.Post("/single", ImportantDayController.CreateImportantDay).Name("CreateImportantDay")

				subImportantDayGroup := ImportantDayGroup.Group("/with-id/:id")
				{

					subImportantDayGroup.Put("/", ImportantDayController.UpdateImportantDay).Name("UpdateImportantDayWithId")
					subImportantDayGroup.Delete("/", ImportantDayController.DeleteImportantDay).Name("DeleteImportantDayWithId")
				}

			}
			{
				RecurringEventInstanceGroup := app.Group("/RecurringEventInstance")
				RecurringEventInstanceController := controllerCalendarBasics.RecurringEventInstanceController{Svc: servicesCalendarBasics.RecurringEventInstanceService{DB: db}}

				RecurringEventInstanceGroup.Get("/search", RecurringEventInstanceController.GetRecurringEventInstance).Name("GetRecurringEventInstance")
				RecurringEventInstanceGroup.Post("/single", RecurringEventInstanceController.CreateRecurringEventInstance).Name("CreateRecurringEventInstance")

				subRecurringEventInstanceGroup := RecurringEventInstanceGroup.Group("/with-id/:id")
				{

					subRecurringEventInstanceGroup.Put("/", RecurringEventInstanceController.UpdateRecurringEventInstance).Name("UpdateRecurringEventInstanceWithId")
					subRecurringEventInstanceGroup.Delete("/", RecurringEventInstanceController.DeleteRecurringEventInstance).Name("DeleteRecurringEventInstanceWithId")
				}

			}
			{
				RecurringEventGroup := app.Group("/RecurringEvent")
				RecurringEventController := controllerCalendarBasics.RecurringEventController{Svc: servicesCalendarBasics.RecurringEventService{DB: db}}

				RecurringEventGroup.Get("/search", RecurringEventController.GetRecurringEvent).Name("GetRecurringEvent")
				RecurringEventGroup.Post("/single", RecurringEventController.CreateRecurringEvent).Name("CreateRecurringEvent")

				subRecurringEventGroup := RecurringEventGroup.Group("/with-id/:id")
				{

					subRecurringEventGroup.Put("/", RecurringEventController.UpdateRecurringEvent).Name("UpdateRecurringEventWithId")
					subRecurringEventGroup.Delete("/", RecurringEventController.DeleteRecurringEvent).Name("DeleteRecurringEventWithId")
				}

			}
			{
				ShiftGroup := app.Group("/Shift")
				ShiftController := controllerCalendarBasics.ShiftController{Svc: servicesCalendarBasics.ShiftService{DB: db}}

				ShiftGroup.Get("/search", ShiftController.GetShift).Name("GetShift")
				ShiftGroup.Post("/single", ShiftController.CreateShift).Name("CreateShift")

				subShiftGroup := ShiftGroup.Group("/with-id/:id")
				{

					subShiftGroup.Put("/", ShiftController.UpdateShift).Name("UpdateShiftWithId")
					subShiftGroup.Delete("/", ShiftController.DeleteShift).Name("DeleteShiftWithId")
				}

			}

		}
		{
			{
				AdministrativeHolidayGroup := app.Group("/AdministrativeHoliday")
				AdministrativeHolidayController := controllerEmployeePermissions.AdministrativeHolidayController{Svc: servicesEmployeePermissions.AdministrativeHolidayService{DB: db}}

				AdministrativeHolidayGroup.Get("/search", AdministrativeHolidayController.GetAdministrativeHoliday).Name("GetAdministrativeHoliday")
				AdministrativeHolidayGroup.Post("/single", AdministrativeHolidayController.CreateAdministrativeHoliday).Name("CreateAdministrativeHoliday")

				subAdministrativeHolidayGroup := AdministrativeHolidayGroup.Group("/with-id/:id")
				{

					subAdministrativeHolidayGroup.Put("/", AdministrativeHolidayController.UpdateAdministrativeHoliday).Name("UpdateAdministrativeHolidayWithId")
					subAdministrativeHolidayGroup.Delete("/", AdministrativeHolidayController.DeleteAdministrativeHoliday).Name("DeleteAdministrativeHolidayWithId")
				}

			}
			{
				AnnualHolidayHolidayGroup := app.Group("/AnnualHolidayHoliday")
				AnnualHolidayHolidayController := controllerEmployeePermissions.AnnualHolidayController{Svc: servicesEmployeePermissions.AnnualHolidayService{DB: db}}

				AnnualHolidayHolidayGroup.Get("/search", AnnualHolidayHolidayController.GetAnnualHoliday).Name("GetAnnualHolidayHoliday")
				AnnualHolidayHolidayGroup.Post("/single", AnnualHolidayHolidayController.CreateAnnualHoliday).Name("CreateAnnualHolidayHoliday")

				subAnnualHolidayHolidayGroup := AnnualHolidayHolidayGroup.Group("/with-id/:id")
				{

					subAnnualHolidayHolidayGroup.Put("/", AnnualHolidayHolidayController.UpdateAnnualHoliday).Name("UpdateAnnualHolidayHolidayWithId")
					subAnnualHolidayHolidayGroup.Delete("/", AnnualHolidayHolidayController.DeleteAnnualHoliday).Name("DeleteAnnualHolidayHolidayWithId")
				}

			}
			{
				BirthHolidayGroup := app.Group("/BirthHoliday")
				BirthHolidayController := controllerEmployeePermissions.BirthHolidayController{Svc: servicesEmployeePermissions.BirthHolidayService{DB: db}}

				BirthHolidayGroup.Get("/search", BirthHolidayController.GetBirthHoliday).Name("GetBirthHoliday")
				BirthHolidayGroup.Post("/single", BirthHolidayController.CreateBirthHoliday).Name("CreateBirthHoliday")

				subBirthHolidayGroup := BirthHolidayGroup.Group("/with-id/:id")
				{

					subBirthHolidayGroup.Put("/", BirthHolidayController.UpdateBirthHoliday).Name("UpdateBirthHolidayWithId")
					subBirthHolidayGroup.Delete("/", BirthHolidayController.DeleteBirthHoliday).Name("DeleteBirthHolidayWithId")
				}

			}
			{
				ExcuseHolidayGroup := app.Group("/ExcuseHoliday")
				ExcuseHolidayController := controllerEmployeePermissions.ExcuseHolidayController{Svc: servicesEmployeePermissions.ExcuseHolidayService{DB: db}}

				ExcuseHolidayGroup.Get("/search", ExcuseHolidayController.GetExcuseHoliday).Name("GetExcuseHoliday")
				ExcuseHolidayGroup.Post("/single", ExcuseHolidayController.CreateExcuseHoliday).Name("CreateExcuseHoliday")

				subExcuseHolidayGroup := ExcuseHolidayGroup.Group("/with-id/:id")
				{

					subExcuseHolidayGroup.Put("/", ExcuseHolidayController.UpdateExcuseHoliday).Name("UpdateExcuseHolidayWithId")
					subExcuseHolidayGroup.Delete("/", ExcuseHolidayController.DeleteExcuseHoliday).Name("DeleteExcuseHolidayWithId")
				}

			}
			{
				PermissionGroup := app.Group("/Permission")
				PermissionController := controllerEmployeePermissions.PermissionController{Svc: servicesEmployeePermissions.PermissionService{DB: db}}

				PermissionGroup.Get("/search", PermissionController.GetPermission).Name("GetPermission")
				PermissionGroup.Post("/single", PermissionController.CreatePermission).Name("CreatePermission")

				subPermissionGroup := PermissionGroup.Group("/with-id/:id")
				{

					subPermissionGroup.Put("/", PermissionController.UpdatePermission).Name("UpdatePermissionWithId")
					subPermissionGroup.Delete("/", PermissionController.DeletePermission).Name("DeletePermissionWithId")
				}

			}
			{
				SpecialHolidayGroup := app.Group("/SpecialHoliday")
				SpecialHolidayController := controllerEmployeePermissions.SpecialHolidayController{Svc: servicesEmployeePermissions.SpecialHolidayService{DB: db}}

				SpecialHolidayGroup.Get("/search", SpecialHolidayController.GetSpecialHoliday).Name("GetSpecialHoliday")
				SpecialHolidayGroup.Post("/single", SpecialHolidayController.CreateSpecialHoliday).Name("CreateSpecialHoliday")

				subSpecialHolidayGroup := SpecialHolidayGroup.Group("/with-id/:id")
				{

					subSpecialHolidayGroup.Put("/", SpecialHolidayController.UpdateSpecialHoliday).Name("UpdateSpecialHolidayWithId")
					subSpecialHolidayGroup.Delete("/", SpecialHolidayController.DeleteSpecialHoliday).Name("DeleteSpecialHolidayWithId")
				}

			}
			{
				PublicHolidayGroup := app.Group("/PublicHoliday")
				PublicHolidayController := controllerEmployeePermissions.PublicHolidayController{Svc: servicesEmployeePermissions.PublicHolidayService{DB: db}}

				PublicHolidayGroup.Get("/search", PublicHolidayController.GetPublicHoliday).Name("GetPublicHoliday")
				PublicHolidayGroup.Post("/single", PublicHolidayController.CreatePublicHoliday).Name("CreatePublicHoliday")

				subPublicHolidayGroup := PublicHolidayGroup.Group("/with-id/:id")
				{

					subPublicHolidayGroup.Put("/", PublicHolidayController.UpdatePublicHoliday).Name("UpdatePublicHolidayWithId")
					subPublicHolidayGroup.Delete("/", PublicHolidayController.DeletePublicHoliday).Name("DeletePublicHolidayWithId")
				}

			}

		}
		{

			{
				EmployessDevelopmentGroup := app.Group("/EmployessDevelopment")
				EmployessDevelopmentController := controllerEmployeeRequests.EmployeeDevelopmentController{Svc: servicesEmployeeRequests.EmployeeDevelopmentService{DB: db}}

				EmployessDevelopmentGroup.Get("/search", EmployessDevelopmentController.GetEmployeeDevelopment).Name("GetEmployessDevelopment")
				EmployessDevelopmentGroup.Post("/single", EmployessDevelopmentController.CreateEmployeeDevelopment).Name("CreateEmployessDevelopment")

				subEmployessDevelopmentGroup := EmployessDevelopmentGroup.Group("/with-id/:id")
				{

					subEmployessDevelopmentGroup.Put("/", EmployessDevelopmentController.UpdateEmployeeDevelopment).Name("UpdateEmployessDevelopmentWithId")
					subEmployessDevelopmentGroup.Delete("/", EmployessDevelopmentController.DeleteEmployeeDevelopment).Name("DeleteEmployessDevelopmentWithId")
				}

			}
			{
				EmployeeEventRequestGroup := app.Group("/EmployeeEventRequest")
				EmployeeEventRequestController := controllerEmployeeRequests.EmployeeEventRequestController{Svc: servicesEmployeeRequests.EmployeeEventRequestService{DB: db}}

				EmployeeEventRequestGroup.Get("/search", EmployeeEventRequestController.GetEmployeeEventRequest).Name("GetEmployeeEventRequest")
				EmployeeEventRequestGroup.Post("/single", EmployeeEventRequestController.CreateEmployeeEventRequest).Name("CreateEmployeeEventRequest")

				subEmployeeEventRequestGroup := EmployeeEventRequestGroup.Group("/with-id/:id")
				{

					subEmployeeEventRequestGroup.Put("/", EmployeeEventRequestController.UpdateEmployeeEventRequest).Name("UpdateEmployeeEventRequestWithId")
					subEmployeeEventRequestGroup.Delete("/", EmployeeEventRequestController.DeleteEmployeeEventRequest).Name("DeleteEmployeeEventRequestWithId")
				}

			}
			{
				EmployeeTrainingGroup := app.Group("/EmployeeTraining")
				EmployeeTrainingController := controllerEmployeeRequests.EmployeeTrainingController{Svc: servicesEmployeeRequests.EmployeeTrainingService{DB: db}}

				EmployeeTrainingGroup.Get("/search", EmployeeTrainingController.GetEmployeeTraining).Name("GetEmployeeTraining")
				EmployeeTrainingGroup.Post("/single", EmployeeTrainingController.CreateEmployeeTraining).Name("CreateEmployeeTraining")

				subEmployeeTrainingGroup := EmployeeTrainingGroup.Group("/with-id/:id")
				{

					subEmployeeTrainingGroup.Put("/", EmployeeTrainingController.UpdateEmployeeTraining).Name("UpdateEmployeeTrainingWithId")
					subEmployeeTrainingGroup.Delete("/", EmployeeTrainingController.DeleteEmployeeTraining).Name("DeleteEmployeeTrainingWithId")
				}

			}
			{
				EmployeeGroup := app.Group("/Employee")
				EmployeeController := controllerEmployeeRequests.EmployeeController{Svc: servicesEmployeeRequests.EmployeeService{DB: db}}

				EmployeeGroup.Get("/search", EmployeeController.GetEmployee).Name("GetEmployee")
				EmployeeGroup.Post("/single", EmployeeController.CreateEmployee).Name("CreateEmployee")

				subEmployeeGroup := EmployeeGroup.Group("/with-id/:id")
				{

					subEmployeeGroup.Put("/", EmployeeController.UpdateEmployee).Name("UpdateEmployeeWithId")
					subEmployeeGroup.Delete("/", EmployeeController.DeleteEmployee).Name("DeleteEmployeeWithId")
				}

			}
		}
		{
			{

				ExpenseGroup := app.Group("/Expense")
				ExpenseController := controllerExpenseTracking.ExpenseController{Svc: servicesExpenseTracking.ExpenseService{DB: db}}

				ExpenseGroup.Get("/search", ExpenseController.GetExpense).Name("GetExpense")
				ExpenseGroup.Post("/single", ExpenseController.CreateExpense).Name("CreateExpense")

				subExpenseGroup := ExpenseGroup.Group("/with-id/:id")
				{

					subExpenseGroup.Put("/", ExpenseController.UpdateExpense).Name("UpdateExpenseWithId")
					subExpenseGroup.Delete("/", ExpenseController.DeleteExpense).Name("DeleteExpenseWithId")
				}

			}
		}
		{
			{

				AuthProviderGroup := app.Group("/AuthProvider")
				AuthProviderController := controllerIntegration.AuthProviderController{Svc: servicesIntegration.AuthProviderService{DB: db}}

				AuthProviderGroup.Get("/search", AuthProviderController.GetAuthProvider).Name("GetAuthProvider")
				AuthProviderGroup.Post("/single", AuthProviderController.CreateAuthProvider).Name("CreateAuthProvider")

				subAuthProviderGroup := AuthProviderGroup.Group("/with-id/:id")
				{

					subAuthProviderGroup.Put("/", AuthProviderController.UpdateAuthProvider).Name("UpdateAuthProviderWithId")
					subAuthProviderGroup.Delete("/", AuthProviderController.DeleteAuthProvider).Name("DeleteAuthProviderWithId")
				}

			}
		}

		{
			{

				NotificationsGroup := app.Group("/Notifications")
				NotificationsController := controllerNotifications.NotificationsController{Svc: servicesNotifications.NotificationsService{DB: db}}

				NotificationsGroup.Get("/search", NotificationsController.GetNotifications).Name("GetNotifications")
				NotificationsGroup.Post("/single", NotificationsController.CreateNotifications).Name("CreateNotifications")

				subNotificationsGroup := NotificationsGroup.Group("/with-id/:id")
				{

					subNotificationsGroup.Put("/", NotificationsController.UpdateNotifications).Name("UpdateNotificationsWithId")
					subNotificationsGroup.Delete("/", NotificationsController.DeleteNotifications).Name("DeleteNotificationsWithId")
				}

			}
		}
		{
			{

				Send_EmailsGroup := app.Group("/Send_Emails")
				Send_EmailsController := controllerNotifications.Send_EmailsController{Svc: servicesNotifications.Send_EmailsService{DB: db}}

				Send_EmailsGroup.Get("/search", Send_EmailsController.GetSend_Emails).Name("GetSend_Emails")
				Send_EmailsGroup.Post("/single", Send_EmailsController.CreateSend_Emails).Name("CreateSend_Emails")

				subSend_EmailsGroup := Send_EmailsGroup.Group("/with-id/:id")
				{

					subSend_EmailsGroup.Put("/", Send_EmailsController.UpdateSend_Emails).Name("UpdateSend_EmailsWithId")

					subSend_EmailsGroup.Delete("/", Send_EmailsController.DeleteSend_Emails).Name("DeleteSend_EmailsWithId")

					subSend_EmailsGroup.Get("/", Send_EmailsController.Send_Email).Name("Send_Send_EmailsWithId")
				}

			}
		}

	}

}
