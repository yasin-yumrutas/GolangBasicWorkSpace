Anakatmanlar =>  "1.entities", "2.services", “3.logics”, "4.database”, “5.security”


  
"1.entities"
	"1.1.database_entities"
		“1.1.1.blockchain_entities”
			“1.1.1.1.avalanche_entities”
			“1.1.1.2.polygon_entities”
		“1.1.2.coucbase_entities”
		“1.1.3.ipfs_entities”
		“1.1.4.redis_entities”
	“1.2.logic_entities"
		“1.2.1.blockchain_logic_entities”
		“1.2.1.service_logic_entities”
		“1.2.1.station_logic_entities”
	“1.3.security_entities"
		“1.3.1.network_entities”
			“1.3.1.1.ip_filter_entities”
			“1.3.1.2.ip_operation_entities”
		“1.3.2.user_entities”
			“1.3.2.1.authentication_entities”
			“1.3.2.2.authorization_entities”

	“1.4.gateway”
		“1.4.1.service_entities
			“1.4.1.1.system_entities”
			“1.4.1.2.web_entities”


“2.gateway”
	“2.1.services" 
		“2.1.1.system” 
 			”2.1.1.1.system_services" 
			“2.1.1.2.system_service_filters" 
			“2.1.1.3.system_service_router”
 		“2.1.2.web”
			“2.1.2.1.web_services" 
			“2.1.2.2.web_service_filters"
			“2.1.2.3.web_service_routers”


“3.logics"
	"3.1.blockchain_logics”
	“3.2.service_logics”
		“3.2.1.system_logics”
		“3.2.1.web_logics”
	“3.3.station_logics"


"4.database"
	"4.1.blockchain"
		“4.1.1.avalanche”
		“4.1.2.polygon”
	“4.2.couchbase”
	“4.3.ipfs”
	“4.4.redis”

“5.security”
	“5.1.network”
		“5.1.1.ip_filters”
		“5.1.2.ip_operations”	
	“5.2.user”
		“5.2.1.Authentication”
		“5.2.2.Authorization”