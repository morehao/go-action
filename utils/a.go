package utils

//
// import (
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// )
//
// // Menu model definition
// type Menu struct {
// 	Name             string `gorm:"not null"`
// 	VisibilityPolicy int    `gorm:"not null"`
// }
//
// // Constants for visibility policies
// const (
// 	VisibilityNormalUser    = 1
// 	VisibilityPlatformAdmin = 2
// 	VisibilityCompanyAdmin  = 3
// )
//
// // SetupDatabase initializes the database connection and migrates the Menu model
// func SetupDatabase() *gorm.DB {
// 	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect to database")
// 	}
//
// 	// Auto Migrate the Menu model
// 	db.AutoMigrate(&Menu{})
//
// 	return db
// }
//
// // CreateMenu handles the creation of a new menu
// func CreateMenu(c *gin.Context) {
// 	var menu Menu
// 	if err := c.ShouldBindJSON(&menu); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
//
// 	// Convert visibility policy array to a single integer using bitwise OR
// 	for _, policy := range menu.VisibilityPolicyList {
// 		switch policy {
// 		case VisibilityNormalUser:
// 			menu.VisibilityPolicy |= VisibilityNormalUser
// 		case VisibilityPlatformAdmin:
// 			menu.VisibilityPolicy |= VisibilityPlatformAdmin
// 		case VisibilityCompanyAdmin:
// 			menu.VisibilityPolicy |= VisibilityCompanyAdmin
// 			// Add more cases if needed for additional policies
// 		}
// 	}
//
// 	// Save the menu to the database
// 	if err := db.Create(&menu).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create menu"})
// 		return
// 	}
//
// 	c.JSON(http.StatusCreated, menu)
// }
//
// // GetMenus retrieves menus based on visibility policies
// func GetMenus(c *gin.Context) {
// 	var menus []Menu
// 	// Extract visibility policies from the request, assuming it's an array of integers
// 	var visibilityPolicyList []int
// 	if err := c.ShouldBindJSON(&visibilityPolicyList); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	// Calculate visibilityMask by summing up the values in visibilityPolicyList
// 	var visibilityMask int
// 	for _, policy := range visibilityPolicyList {
// 		visibilityMask |= policy
// 	}
// 	// Query menus with matching visibility policies
// 	if err := db.Where("(visibility_policy & ?) > 0", visibilityMask).Find(&menus).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menus"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, menus)
// }
