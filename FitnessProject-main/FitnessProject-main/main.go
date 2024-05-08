package main

import (
	//"fmt"
	db "FitnessProject/db/sqlc"
	api "FitnessProject/api"
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	//"github.com/gin-gonic/gin"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/fitnessdb?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

// type lifter struct {
// 	Name 			string 	`json:"name"`
// 	Weight 			float64 `json:"weight"`
// 	Height 			float64 `json:"height"`
// 	Age    			float64 `json:"age"`
// 	WeightLifted 	float64 `json:"weightLifted"`
// 	Reps         	float64 `json:"reps"`
// 	WeightGain		bool	`json:"weightgain"`
// 	Intensity		int 	`json:"intensity"`
// 	Bmr 		   	float64 `json:"bmr"`
// 	NewCalories		float64 `json:"newcalories"`
// 	Orm           	float64 `json:"orm"`
// }

// var lifters = []lifter{
//     {Name: "Charlie", Weight: 168.0, Height: 72.0, Age: 23.0, WeightLifted: 185.0, Reps: 5.0, Intensity: 3.0},
//     {Name: "Bob", Weight: 275, Height: 80, Age: 40, WeightLifted: 300, Reps: 7, Intensity: 2.0},
//     {Name: "Joe", Weight: 140, Height: 66, Age: 18, WeightLifted: 145, Reps: 3, Intensity: 1.0},
// }

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
	
    // router := gin.Default()
    // router.GET("/lifters", getlifters)
	// router.GET("/lifters/:name", getLifterByName)
	// router.GET("/lifters/getNewCalories/:name", getNewCalories)
	// router.GET("/lifters/getORM/:name", getORM)
	// router.POST("/lifters", postLifters)

    // router.Run("localhost:8080")
}

// func getNewCalories(c *gin.Context) {
// 	name := c.Param("name")

// 	for _, a := range lifters {
// 		if a.Name == name {
// 			if a.Height == 0 || a.Weight == 0 || a.Age == 0 {
// 				a.Bmr = 0
// 			} else {
// 				a.Bmr = 10.0*a.Weight + 6.25*a.Height - 5.0*a.Age + 5.0
// 				if !a.WeightGain {
// 					if a.Intensity == 1 {
// 						a.NewCalories = a.Bmr - 250
// 					}
// 					if a.Intensity == 2 {
// 						a.NewCalories = a.Bmr - 500
// 					}
// 					if a.Intensity == 3 {
// 						a.NewCalories = a.Bmr - 750
// 					}
// 				} else {
// 					if a.Intensity == 1 {
// 						a.NewCalories = a.Bmr + 250
// 					}
// 					if a.Intensity == 2 {
// 						a.NewCalories = a.Bmr + 500
// 					}
// 					if a.Intensity == 3 {
// 						a.NewCalories = a.Bmr + 750
// 					}
// 				}
// 			}
// 			c.IndentedJSON(http.StatusOK, "Calorie Goal: " + fmt.Sprintf("%v",a.NewCalories))
// 			return
// 		}
// 	}
// }

// func Sprintf(f float64) {
// 	panic("unimplemented")
// }

// func getORM(c *gin.Context) {
// 	name := c.Param("name")

// 	for _, a := range lifters {
// 		if a.Name == name {
// 			a.Orm = a.WeightLifted / (1.0278 - 0.0278*a.Reps)
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// }

// func getlifters(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, lifters)
// }

// func postLifters(c *gin.Context) {
// 	var newLifter lifter

// 	if err := c.BindJSON(&newLifter); err != nil {
// 		return
// 	}

// 	lifters = append(lifters, newLifter)
// 	c.IndentedJSON(http.StatusCreated, newLifter)
// }

// func getLifterByName(c *gin.Context) {
// 	name := c.Param("name")

// 	for _, a := range lifters {
// 		if a.Name == name {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "lifter not found"})
// }