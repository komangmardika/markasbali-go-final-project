package controllers

import (
	"final-project/kelas-beta-golang/model"
	"final-project/kelas-beta-golang/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func Route(app *fiber.App) {
	Group := app.Group("/autores")
	Group.Get("/", GetDBLatestBackupList)
	Group.Get("/:db_name", GetAllByDBName)
	// carsGroup.Post("/")
}

func GetDBLatestBackupList(c *fiber.Ctx) error {
	autoResData, err := utils.GetDistinct()
	if err != nil {
		logrus.Error("Error on get database distinct list: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}

	responseData := make([]map[string]interface{}, len(autoResData))

	for i, row := range autoResData {

		dataFetch, err := utils.GetLatestByDBName(row)

		if err != nil {
			logrus.Error("Error on get database latest backup name list: ", err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(
				map[string]any{
					"message": "Server Error",
				},
			)
		}

		responseData[i] = map[string]interface{}{
			"database_name": dataFetch.Nama_Database,
			"latest_backup": map[string]interface{}{
				"id":         dataFetch.Id,
				"file_name":  dataFetch.Nama_File_Backup,
				"timestamp":  dataFetch.CreatedAt,
			},
		}
    }

	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    responseData,
			"message": "Success",
		},
	)
}

func GetAllByDBName(c *fiber.Ctx) error {
	nama_database := c.Params("db_name")

	historiesData, err := utils.GetAllByDBName(nama_database)
	if err != nil {
		logrus.Error("Error on get database distinct list: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "Server Error",
		})
	}
	
	responseHistoriesData := make([]map[string]interface{},0)
	for _, row := range historiesData {
		responseHistoriesData = append(responseHistoriesData, map[string]interface{}{
			"id":         row.Id,
			"file_name":  row.Nama_Database,
			"timestamp":  row.CreatedAt,
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"data": map[string]interface{}{
			"database_name": nama_database,
			"histories":     responseHistoriesData,
		},
		"message": "Success",
	})
	
}


func GetDBList(c *fiber.Ctx) error {
	autoResData, err := utils.GetList()
	if err != nil {
		logrus.Error("Error on get database list: ", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(
			map[string]any{
				"message": "Server Error",
			},
		)
	}
	return c.Status(fiber.StatusOK).JSON(
		map[string]any{
			"data":    autoResData,
			"message": "Success",
		},
	)
}

func InsertData(c *fiber.Ctx) error{
	type AddRequest struct{
		Nama_Database  string `json:"nama_database" valid:"required"`
		Nama_File_Backup string `json:"nama_file_backup" valid:"required"`
	}

	req := new(AddRequest)

    if err := c.BodyParser(req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{
            "message": "Body Not Valid",
        })
    }

	// isValid, err := govalidator.ValidateStruct(req)
	// if !isValid{
	// 	return c.Status(fiber.StatusBadRequest).
	// 		JSON(map[string] any{
	// 			"error" : err.Error(),
	// 			"message" : "One Or Other field is required",
	// 		})
	// }

	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).
	// 		JSON(map[string] any{
	// 			"message" : "Error Occured In Validation",
	// 		})
	// }

    cardata,err := utils.InsertData(model.AutoRes{
        Nama_Database:  req.Nama_Database,
        Nama_File_Backup:  req.Nama_File_Backup,
    })
	
	if err != nil {
        logrus.Printf("Terjadi Error: %s\n", err)
        return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
            "message": "Server Error",
        })
    }

    return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
		"data" : cardata,
        "message": "Success Insert Data",
    })
}


