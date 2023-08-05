package helper

import (
	"deneme/security"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	entitiesGroup := r.Group("/entities")
	{
		entitiesGroup.GET("/", GetAllEntities)
		entitiesGroup.GET("/database", GetAllDatabaseEntities)
		entitiesGroup.GET("/database/blockchain", GetAllBlockchainEntities)
		entitiesGroup.GET("/database/blockchain/avalanche", GetAllAvalancheEntities)
		// Diğer iç içe geçmiş istemci istekleri için işleyicileri benzer şekilde uygulayın
	}

	servicesGroup := r.Group("/services")
	{
		servicesGroup.GET("/", GetAllServices)
		servicesGroup.GET("/system", GetAllSystemEntities)
		servicesGroup.GET("/web", GetAllWebEntities)
		// Diğer iç içe geçmiş istemci istekleri için işleyicileri benzer şekilde uygulayın
	}

	logicsGroup := r.Group("/logics")
	{
		logicsGroup.GET("/", GetAllLogics)
		logicsGroup.GET("/blockchain", GetAllBlockchainLogics)
		logicsGroup.GET("/service", GetAllServiceLogics)
		// Diğer iç içe geçmiş istemci istekleri için işleyicileri benzer şekilde uygulayın
	}

	databaseGroup := r.Group("/database")
	{
		databaseGroup.GET("/", GetAllDatabases)
		databaseGroup.GET("/blockchain", GetAllBlockchainDatabases)
		// Diğer iç içe geçmiş istemci istekleri için işleyicileri benzer şekilde uygulayın
	}

	securityGroup := r.Group("/security")
	{
		securityGroup.GET("/", GetAllSecurityEntities)
		securityGroup.GET("/network", GetAllNetworkSecurityEntities)
		securityGroup.GET("/user", GetAllUserSecurityEntities)
		// Diğer iç içe geçmiş istemci istekleri için işleyicileri benzer şekilde uygulayın
	}
}

// "entities" istemci isteği için işleyici
func GetAllEntities(c *gin.Context) {
	// Tüm varlıkları almak için kodu uygulayın ve yanıt olarak döndürün
}

// "services" istemci isteği için işleyici
func GetAllServices(c *gin.Context) {
	// Tüm servisleri almak için kodu uygulayın ve yanıt olarak döndürün
}

// "logics" istemci isteği için işleyici
func GetAllLogics(c *gin.Context) {
	// Tüm mantıkları almak için kodu uygulayın ve yanıt olarak döndürün
}

// "database" istemci isteği için işleyici
func GetAllDatabases(c *gin.Context) {
	// Tüm veritabanlarını almak için kodu uygulayın ve yanıt olarak döndürün
}

// "security" istemci isteği için işleyici
func GetAllSecurityEntities(c *gin.Context) {
	// Tüm güvenlik varlıklarını almak için kodu uygulayın ve yanıt olarak döndürün

	security.HandleLogin(c)
	security.HandleCallback(c)
}

// "entities" istemci isteği için alt varlıklarla birlikte işleyici
func GetAllDatabaseEntities(c *gin.Context) {
	// Tüm veritabanı varlıklarını almak için kodu uygulayın ve yanıt olarak döndürün
}

func GetAllBlockchainEntities(c *gin.Context) {
	// Tüm blok zinciri varlıklarını almak için kodu uygulayın ve yanıt olarak döndürün
}

func GetAllAvalancheEntities(c *gin.Context) {
	// Tüm Avalanche varlıklarını almak için kodu uygulayın ve yanıt olarak döndürün
}

// Diğer iç içe geçmiş istemci istekleri için işleyicileri benzer şekilde uygulayın
// ...

// "services" istemci isteği için alt varlıklarla birlikte işleyici
func GetAllSystemEntities(c *gin.Context) {
	// Tüm sistem varlıklarını almak için kodu uygulayın ve yanıt olarak döndürün
}

func GetAllWebEntities(c *gin.Context) {
	// Tüm web varlıklarını almak için kodu uygulayın ve yanıt olarak döndürün
}

// Diğer iç içe geçmiş istemci istekleri için işleyicileri benzer şekilde uygulayın
// ...

// "logics" istemci isteği için alt varlıklarla birlikte işleyici
func GetAllBlockchainLogics(c *gin.Context) {
	// Tüm blok zinciri mantıklarını almak için kodu uygulayın ve yanıt olarak döndürün
}

func GetAllServiceLogics(c *gin.Context) {
	// Tüm servis mantıklarını almak için kodu uygulayın ve yanıt olarak döndürün
}

// Diğer iç içe geçmiş istemci istekleri için işleyicileri benzer şekilde uygulayın
// ...

// "database" istemci isteği için alt varlıklarla birlikte işleyici
func GetAllBlockchainDatabases(c *gin.Context) {
	// Tüm blok zinciri veritabanlarını almak için kodu uygulayın ve yanıt olarak döndürün
}

// Diğer iç içe geçmiş istemci istekleri için işleyicileri benzer şekilde uygulayın
// ...

// "security" istemci isteği için alt varlıklarla birlikte işleyici
func GetAllNetworkSecurityEntities(c *gin.Context) {
	// Tüm ağ güvenlik varlıklarını almak için kodu uygulayın ve yanıt olarak döndürün
}

func GetAllUserSecurityEntities(c *gin.Context) {
	// Tüm kullanıcı güvenlik varlıklarını almak için kodu uygulayın ve yanıt olarak döndürün
}

// Diğer iç içe geçmiş istemci istekleri için işleyicileri benzer şekilde uygulayın
// ...
